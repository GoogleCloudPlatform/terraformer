package confluent

import (
	"errors"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
	"os"
	"strconv"
)

type ConfluentProvider struct { //nolint
	terraformutils.Provider
	cloudApiKey       string
	cloudApiSecret    string
	endpoint          string
	kafkaApiKey       string
	kafkaApiSecret    string
	kafkaRestEndpoint string
	maxRetries        int
}

const (
	confluentEndpoint = "https://api.confluent.cloud"
	maxRetry          = 3
)

func (p *ConfluentProvider) Init(args []string) error {
	if os.Getenv("CLOUD_API_KEY") == "" {
		return errors.New("set CLOUD_API_KEY env var")
	}
	p.cloudApiKey = os.Getenv("CLOUD_API_KEY")

	if os.Getenv("CLOUD_API_SECRET") == "" {
		return errors.New("set CLOUD_API_KEY env var")
	}
	p.cloudApiSecret = os.Getenv("CLOUD_API_SECRET")

	if os.Getenv("KAFKA_API_KEY") == "" {
		return errors.New("set KAFKA_API_KEY env var")
	}
	p.kafkaApiKey = os.Getenv("KAFKA_API_KEY")

	if os.Getenv("KAFKA_API_SECRET") == "" {
		return errors.New("set KAFKA_API_SECRET env var")
	}
	p.kafkaApiSecret = os.Getenv("KAFKA_API_SECRET")

	if os.Getenv("KAFKA_REST_ENDPOINT") == "" {
		return errors.New("set KAFKA_REST_ENDPOINT env var")
	}
	p.kafkaRestEndpoint = os.Getenv("KAFKA_REST_ENDPOINT")

	if os.Getenv("MAX_RETRIES") == "" {
		p.maxRetries = maxRetry
	} else {
		maxRetries, err := strconv.Atoi(os.Getenv("MAX_RETRIES"))
		if err != nil {
			return errors.New("set MAX_RETRIES env var")
		}
		p.maxRetries = maxRetries
	}

	p.endpoint = confluentEndpoint
	return nil
}

func (p *ConfluentProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"cloud_api_key":       cty.StringVal(p.cloudApiKey),
		"cloud_api_secret":    cty.StringVal(p.cloudApiSecret),
		"endpoint":            cty.StringVal(p.endpoint),
		"kafka_api_key":       cty.StringVal(p.kafkaApiKey),
		"kafka_api_secret":    cty.StringVal(p.kafkaApiSecret),
		"kafka_rest_endpoint": cty.StringVal(p.kafkaRestEndpoint),
		"max_retries":         cty.NumberIntVal(int64(p.maxRetries)),
	})
}

func (p *ConfluentProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"kafka_topic":     &KafkaTopicGenerator{},
		"kafka_acl":       &KafkaACLGenerator{},
		"kafka_cluster":   &KafkaClusterGenerator{},
		"role_binding":    &RoleBindinGenerator{},
		"service_account": &ServiceAccountGenerator{},
		"api_key":         &ApiKeyGenerator{},
	}
}

func (p *ConfluentProvider) GetName() string {
	return "confluent"
}

func (p *ConfluentProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *ConfluentProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("confluent: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"cloud_api_key":       p.cloudApiKey,
		"cloud_api_secret":    p.cloudApiSecret,
		"endpoint":            p.endpoint,
		"kafka_api_key":       p.kafkaApiKey,
		"kafka_api_secret":    p.kafkaApiSecret,
		"kafka_rest_endpoint": p.kafkaRestEndpoint,
		"max_retries":         p.maxRetries,
	})
	return nil
}

func (p ConfluentProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"confluent": map[string]interface{}{
				"cloud_api_key":       p.cloudApiKey,
				"cloud_api_secret":    p.cloudApiSecret,
				"endpoint":            p.endpoint,
				"kafka_api_key":       p.kafkaApiKey,
				"kafka_api_secret":    p.kafkaApiSecret,
				"kafka_rest_endpoint": p.kafkaRestEndpoint,
				"max_retries":         p.maxRetries,
			},
		},
	}
}
