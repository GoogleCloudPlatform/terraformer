package gcp

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"cloud.google.com/go/iam/apiv1/iampb"
	run "cloud.google.com/go/run/apiv2"
	"cloud.google.com/go/run/apiv2/runpb"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var cloudRunAllowEmptyValues = []string{""}

type CloudRunGenerator struct {
	GCPService
}

func (g *CloudRunGenerator) InitResources() error {
	ctx := context.Background()
	filename := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	project := g.GetArgs()["project"].(string)
	location := g.GetArgs()["region"].(compute.Region).Name
	// endpoint := fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)

	serviceClient, err := run.NewServicesClient(ctx, option.WithCredentialsFile(filename))
	if err != nil {
		return err
	}
	defer serviceClient.Close()

	sit := serviceClient.ListServices(ctx, &runpb.ListServicesRequest{Parent: fmt.Sprintf("projects/%s/locations/%s", project, location)})
	if err := g.createServices(serviceClient, sit); err != nil {
		return err
	}

	jobsClient, err := run.NewJobsClient(ctx, option.WithCredentialsFile(filename))
	if err != nil {
		return err
	}
	defer jobsClient.Close()

	jit := jobsClient.ListJobs(ctx, &runpb.ListJobsRequest{Parent: fmt.Sprintf("projects/%s/locations/%s", project, location)})
	if err := g.createJobs(jobsClient, jit); err != nil {
		return err
	}

	return nil
}

func (g *CloudRunGenerator) createServices(client *run.ServicesClient, it *run.ServiceIterator) error {
	for {
		svc, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				return nil
			}
			return err
		}

		project := g.GetArgs()["project"].(string)
		location := g.GetArgs()["region"].(compute.Region).Name

		policy, err := client.GetIamPolicy(context.Background(), &iampb.GetIamPolicyRequest{Resource: svc.GetName()})
		if err != nil {
			return err
		}

		if len(policy.GetBindings()) > 0 {
			if policyData, err := json.Marshal(map[string]interface{}{"bindings": policy.GetBindings()}); err == nil {
				g.Resources = append(g.Resources, terraformutils.NewResource(
					svc.GetName(),
					svc.GetName(),
					"google_cloud_run_v2_service_iam_policy",
					g.GetProviderName(),
					map[string]string{
						"name":     svc.GetName(),
						"project":  project,
						"location": location,
					},
					cloudRunAllowEmptyValues,
					map[string]interface{}{
						"policy_data": string(policyData),
					},
				))
			}
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			svc.GetName(),
			svc.GetName(),
			"google_cloud_run_v2_service",
			g.GetProviderName(),
			map[string]string{
				"name":    svc.GetName(),
				"project": project,
				"region":  location,
			},
			cloudRunAllowEmptyValues,
			map[string]interface{}{
				"template": svc.Template,
			},
		))
	}
}

func (g *CloudRunGenerator) createJobs(client *run.JobsClient, it *run.JobIterator) error {
	for {
		job, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				return nil
			}
			return err
		}

		project := g.GetArgs()["project"].(string)
		location := g.GetArgs()["region"].(compute.Region).Name

		policy, err := client.GetIamPolicy(context.Background(), &iampb.GetIamPolicyRequest{Resource: job.GetName()})
		if err != nil {
			return err
		}

		if len(policy.GetBindings()) > 0 {
			if policyData, err := json.Marshal(map[string]interface{}{"bindings": policy.GetBindings()}); err == nil {
				g.Resources = append(g.Resources, terraformutils.NewResource(
					job.GetName(),
					job.GetName(),
					"google_cloud_run_v2_job_iam_policy",
					g.GetProviderName(),
					map[string]string{
						"name":     job.GetName(),
						"project":  project,
						"location": location,
					},
					cloudRunAllowEmptyValues,
					map[string]interface{}{
						"policy_data": string(policyData),
					},
				))
			}
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			job.GetName(),
			job.GetName(),
			"google_cloud_run_v2_job",
			g.GetProviderName(),
			map[string]string{
				"name":    job.GetName(),
				"project": project,
				"region":  location,
			},
			cloudRunAllowEmptyValues,
			map[string]interface{}{
				"template": job.Template,
			},
		))
	}
}
