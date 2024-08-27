package honeycombio

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	hnyclient "github.com/honeycombio/terraform-provider-honeycombio/client"
)

type RecipientGenerator struct {
	HoneycombService
}

func (g *RecipientGenerator) InitResources() error {
	client, err := g.newClient()
	if err != nil {
		return fmt.Errorf("unable to initialize Honeycomb client: %v", err)
	}

	rcpts, err := client.Recipients.List(context.TODO())
	if err != nil {
		return fmt.Errorf("unable to list Honeycomb recipients: %v", err)
	}

	for _, rcpt := range rcpts {
		var rcptResourceType string
		var unsupportedRcpt bool

		switch rcpt.Type {
		case hnyclient.RecipientTypeEmail:
			rcptResourceType = "honeycombio_email_recipient"
		case hnyclient.RecipientTypePagerDuty:
			rcptResourceType = "honeycombio_pagerduty_recipient"
		case hnyclient.RecipientTypeSlack:
			rcptResourceType = "honeycombio_slack_recipient"
		case hnyclient.RecipientTypeWebhook:
			rcptResourceType = "honeycombio_webhook_recipient"
		default:
			unsupportedRcpt = true
		}
		if unsupportedRcpt {
			fmt.Printf("WARNING: unsupported recipient type: %s\n", rcpt.Type)
			continue
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			rcpt.ID,
			rcpt.ID,
			rcptResourceType,
			"honeycombio",
			map[string]string{},
			[]string{},
			map[string]interface{}{},
		))
	}

	return nil
}
