// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ibm

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/controllerv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"

	ibmaws "github.com/IBM/ibm-cos-sdk-go/aws"
	cossession "github.com/IBM/ibm-cos-sdk-go/aws/session"
	coss3 "github.com/IBM/ibm-cos-sdk-go/service/s3"
)

type COSGenerator struct {
	IBMService
}

func (g COSGenerator) loadCOS(cosID string, cosName string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		cosID,
		cosName,
		"ibm_resource_instance",
		"ibm",
		[]string{})
	return resources
}

func (g COSGenerator) loadCOSBuckets(bucketID, bucketName string) terraformutils.Resource {
	var resources terraformutils.Resource
	resources = terraformutils.NewSimpleResource(
		bucketID,
		bucketName,
		"ibm_cos_bucket",
		"ibm",
		[]string{})
	return resources
}

func (g *COSGenerator) InitResources() error {
	bmxConfig := &bluemix.Config{
		BluemixAPIKey: os.Getenv("IC_API_KEY"),
	}
	sess, err := session.New(bmxConfig)
	if err != nil {
		return err
	}

	catalogClient, err := catalog.New(sess)
	if err != nil {
		return err
	}

	controllerClient, err := controllerv2.New(sess)
	if err != nil {
		return err
	}

	serviceID, err := catalogClient.ResourceCatalog().FindByName("cloud-object-storage", true)
	if err != nil {
		return err
	}
	query := controllerv2.ServiceInstanceQuery{
		ServiceID: serviceID[0].ID,
	}
	cosInstances, err := controllerClient.ResourceServiceInstanceV2().ListInstances(query)
	if err != nil {
		return err
	}
	authEndpoint := "https://iam.cloud.ibm.com/identity/token"
	for _, cs := range cosInstances {
		g.Resources = append(g.Resources, g.loadCOS(cs.ID, cs.Name))
		s3Conf := ibmaws.NewConfig().WithCredentials(ibmiam.NewStaticCredentials(ibmaws.NewConfig(), authEndpoint, os.Getenv("IC_API_KEY"), cs.ID)).WithS3ForcePathStyle(true).WithEndpoint("s3.us-south.cloud-object-storage.appdomain.cloud")
		s3Sess := cossession.Must(cossession.NewSession())
		s3Client := coss3.New(s3Sess, s3Conf)
		d, _ := s3Client.ListBucketsExtended(&coss3.ListBucketsExtendedInput{})
		for _, b := range d.Buckets {
			var apiType, location string
			singleSiteLocationRegex, _ := regexp.Compile("^[a-z]{3}[0-9][0-9]-[a-z]{4,8}$")
			regionLocationRegex, _ := regexp.Compile("^[a-z]{2}-[a-z]{2,5}-[a-z]{4,8}$")
			crossRegionLocationRegex, _ := regexp.Compile("^[a-z]{2}-[a-z]{4,8}$")
			bLocationConstraint := *b.LocationConstraint
			if singleSiteLocationRegex.MatchString(bLocationConstraint) {
				apiType = "ss1"
				location = strings.Split(bLocationConstraint, "-")[0]
			}
			if regionLocationRegex.MatchString(bLocationConstraint) {
				apiType = "rl"
				location = fmt.Sprintf("%s-%s", strings.Split(bLocationConstraint, "-")[0], strings.Split(bLocationConstraint, "-")[1])
			}
			if crossRegionLocationRegex.MatchString(bLocationConstraint) {
				apiType = "crl"
				location = strings.Split(bLocationConstraint, "-")[0]
			}
			bucketID := fmt.Sprintf("%s:%s:%s:meta:%s:%s", strings.Replace(cs.ID, "::", "", -1), "bucket", *b.Name, apiType, location)
			g.Resources = append(g.Resources, g.loadCOSBuckets(bucketID, *b.Name))
		}
	}

	return nil
}
