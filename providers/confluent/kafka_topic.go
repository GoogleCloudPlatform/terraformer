package confluent

import (
	"fmt"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	kafkarestv3 "github.com/confluentinc/ccloud-sdk-go-v2/kafkarest/v3"
)

var topicAllowEmptyValues = []string{""}
var topicAdditionalFields = map[string]interface{}{}

type KafkaTopicGenerator struct {
	ConfluentService
}

func (g KafkaTopicGenerator) createResources(topics []kafkarestv3.TopicData) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, topic := range topics {
		resources = append(resources, terraformutils.NewResource(
			createKafkaTopicId(topic.ClusterId, topic.TopicName),
			topic.TopicName,
			"confluent_kafka_topic",
			g.ProviderName,
			map[string]string{
				"topic_name":       topic.TopicName,
				"partitions_count": string(topic.PartitionsCount),
				"rest_endpoint":    g.GetArgs()["kafka_rest_endpoint"].(string),
			},
			topicAllowEmptyValues,
			topicAdditionalFields,
		))
	}
	return resources
}

func createKafkaTopicId(clusterId, topicName string) string {
	return fmt.Sprintf("%s/%s", clusterId, topicName)
}
