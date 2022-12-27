package cmd

import (
	confluent_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/confluent"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdConfluentImporter(options ImportOptions) *cobra.Command {
	var cloudApiKey, cloudApiSecret, endpoint, kafkaApiKey, kafkaApiSecret, kafkaRestEndpoint, maxRetries string
	cmd := &cobra.Command{
		Use:   "confluent",
		Short: "Import current state to Terraform configuration from Confluent",
		Long:  "Import current state to Terraform configuration from Confluent",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newConfluentProvider()
			err := Import(provider, options, []string{cloudApiKey, cloudApiSecret, endpoint, kafkaApiKey, kafkaApiSecret, kafkaRestEndpoint, maxRetries})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newConfluentProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "kafka_topic,kafka_cluster", "kafka_cluster=id1:id2:id4")
	cmd.PersistentFlags().StringVarP(&cloudApiKey, "cloud-api-key", "", "", "YOUR_CONFLUENT_API_KEY or env param CLOUD_API_KEY")
	cmd.PersistentFlags().StringVarP(&cloudApiSecret, "cloud-api-secret", "", "", "YOUR_CONFLUENT_API_SECRET or env param CLOUD_API_SECRET")
	cmd.PersistentFlags().StringVarP(&kafkaApiKey, "kafka-api-key", "", "", "YOUR_KAFKA_API_KEY or env param KAFKA_API_KEY")
	cmd.PersistentFlags().StringVarP(&kafkaApiSecret, "kafka-api-secret", "", "", "YOUR_KAFKA_API_SECRET or env param KAFKA_API_SECRET")
	cmd.PersistentFlags().StringVarP(&kafkaRestEndpoint, "kafka-rest-endpoint", "", "", "YOUR_KAFKA_REST_ENDPOINT or env param KAFKA_REST_ENDPOINT")
	return cmd
}

func newConfluentProvider() terraformutils.ProviderGenerator {
	return &confluent_terraforming.ConfluentProvider{}
}
