// Copyright 2018 The Terraformer Authors.
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

package aws

import (
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/aws/aws-sdk-go/service/elbv2"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var AlbAllowEmptyValues = []string{"tags.", "^condition."}

type AlbGenerator struct {
	AWSService
}

func (g *AlbGenerator) loadLB(svc *elbv2.ELBV2) error {
	err := svc.DescribeLoadBalancersPages(&elbv2.DescribeLoadBalancersInput{}, func(lbs *elbv2.DescribeLoadBalancersOutput, lastPage bool) bool {
		for _, lb := range lbs.LoadBalancers {
			resourceName := aws.StringValue(lb.LoadBalancerName)
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				aws.StringValue(lb.LoadBalancerArn),
				resourceName,
				"aws_lb",
				"aws",
				map[string]string{},
				AlbAllowEmptyValues,
				map[string]string{},
			))
			if aws.StringValue(lb.Type) != "network" {
				err := g.loadLBListener(svc, lb.LoadBalancerArn)
				if err != nil {
					log.Println(err)
				}
			}

		}
		return !lastPage
	})
	return err
}

func (g *AlbGenerator) loadLBListener(svc *elbv2.ELBV2, loadBalancerArn *string) error {
	err := svc.DescribeListenersPages(&elbv2.DescribeListenersInput{LoadBalancerArn: loadBalancerArn}, func(lcs *elbv2.DescribeListenersOutput, lastPage bool) bool {
		for _, ls := range lcs.Listeners {
			resourceName := aws.StringValue(ls.ListenerArn)
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				resourceName,
				resourceName,
				"aws_lb_listener",
				"aws",
				map[string]string{},
				AlbAllowEmptyValues,
				map[string]string{},
			))
			err := g.loadLBListenerRule(svc, ls.ListenerArn)
			if err != nil {
				log.Println(err)
			}
			err = g.loadLBListenerCertificate(svc, ls.ListenerArn)
			if err != nil {
				log.Println(err)
			}
		}
		return !lastPage
	})
	return err
}

func (g *AlbGenerator) loadLBListenerRule(svc *elbv2.ELBV2, listenerArn *string) error {
	lsrs, err := svc.DescribeRules(&elbv2.DescribeRulesInput{ListenerArn: listenerArn})
	if err != nil {
		return err
	}
	for _, lsr := range lsrs.Rules {
		resourceName := aws.StringValue(lsr.RuleArn)
		g.Resources = append(g.Resources, terraform_utils.NewResource(
			resourceName,
			resourceName,
			"aws_lb_listener_rule",
			"aws",
			map[string]string{},
			AlbAllowEmptyValues,
			map[string]string{},
		))
	}
	return err
}

func (g *AlbGenerator) loadLBListenerCertificate(svc *elbv2.ELBV2, listenerArn *string) error {
	lcs, err := svc.DescribeListenerCertificates(&elbv2.DescribeListenerCertificatesInput{ListenerArn: listenerArn})
	if err != nil {
		return err
	}
	for _, lc := range lcs.Certificates {
		resourceName := aws.StringValue(lc.CertificateArn)
		g.Resources = append(g.Resources, terraform_utils.NewResource(
			resourceName,
			resourceName,
			"aws_lb_listener_certificate",
			"aws",
			map[string]string{},
			AlbAllowEmptyValues,
			map[string]string{},
		))
	}
	return err
}

func (g *AlbGenerator) loadLBTargetGroup(svc *elbv2.ELBV2) error {
	err := svc.DescribeTargetGroupsPages(&elbv2.DescribeTargetGroupsInput{}, func(tgs *elbv2.DescribeTargetGroupsOutput, lastPage bool) bool {
		for _, tg := range tgs.TargetGroups {
			resourceName := aws.StringValue(tg.TargetGroupName)
			g.Resources = append(g.Resources, terraform_utils.NewResource(
				aws.StringValue(tg.TargetGroupArn),
				resourceName,
				"aws_lb_target_group",
				"aws",
				map[string]string{},
				AlbAllowEmptyValues,
				map[string]string{},
			))
			err := g.loadTargetGroupTargets(svc, tg.TargetGroupArn)
			if err != nil {
				log.Println(err)
			}
		}
		return !lastPage
	})
	return err
}

func (g *AlbGenerator) loadTargetGroupTargets(svc *elbv2.ELBV2, targetGroupArn *string) error {
	targetHealths, err := svc.DescribeTargetHealth(&elbv2.DescribeTargetHealthInput{TargetGroupArn: targetGroupArn})
	if err != nil {
		return err
	}
	for _, tgh := range targetHealths.TargetHealthDescriptions {
		id := resource.PrefixedUniqueId(fmt.Sprintf("%s-", aws.StringValue(targetGroupArn)))
		g.Resources = append(g.Resources, terraform_utils.NewResource(
			id,
			id,
			"aws_lb_target_group_attachment",
			"aws",
			map[string]string{
				"target_id":        aws.StringValue(tgh.Target.Id),
				"target_group_arn": aws.StringValue(targetGroupArn),
			},
			AlbAllowEmptyValues,
			map[string]string{},
		))
	}
	return nil
}

// Generate TerraformResources from AWS API,
func (g *AlbGenerator) InitResources() error {
	sess, _ := session.NewSession(&aws.Config{Region: aws.String(g.GetArgs()["region"].(string))})
	svc := elbv2.New(sess)
	if err := g.loadLB(svc); err != nil {
		return err
	}
	if err := g.loadLBTargetGroup(svc); err != nil {
		return err
	}
	g.PopulateIgnoreKeys()
	return nil
}

func (g *AlbGenerator) PostConvertHook() error {
	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_lb_listener" {
			continue
		}
		for _, lb := range g.Resources {
			if lb.InstanceInfo.Type != "aws_lb" {
				continue
			}
			if r.InstanceState.Attributes["load_balancer_arn"] == lb.InstanceState.Attributes["arn"] {
				g.Resources[i].Item["load_balancer_arn"] = "${aws_lb." + lb.ResourceName + ".arn}"
			}
		}
	}

	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_lb_listener_rule" {
			continue
		}
		for _, lb := range g.Resources {
			if lb.InstanceInfo.Type != "aws_lb_listener" {
				continue
			}
			if r.InstanceState.Attributes["listener_arn"] == lb.InstanceState.Attributes["arn"] {
				g.Resources[i].Item["listener_arn"] = "${aws_lb_listener." + lb.ResourceName + ".arn}"
			}
		}
		for _, lb := range g.Resources {
			if lb.InstanceInfo.Type != "aws_lb_listener_certificate" {
				continue
			}
			if r.InstanceState.Attributes["certificate_arn"] == lb.InstanceState.Attributes["arn"] {
				g.Resources[i].Item["certificate_arn"] = "${aws_lb_listener_certificate." + lb.ResourceName + ".arn}"
			}
		}
	}

	for i, r := range g.Resources {
		if r.InstanceInfo.Type != "aws_lb_target_group_attachment" {
			continue
		}
		for _, lb := range g.Resources {
			if lb.InstanceInfo.Type != "aws_lb_target_group" {
				continue
			}
			if r.InstanceState.Attributes["target_group_arn"] == lb.InstanceState.Attributes["arn"] {
				g.Resources[i].Item["target_group_arn"] = "${aws_lb_target_group." + lb.ResourceName + ".arn}"
			}
		}
	}
	return nil
}
