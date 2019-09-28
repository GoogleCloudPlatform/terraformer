package alikafka

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"eu-west-1":             "alikafka.ap-south-1.aliyuncs.com",
			"ap-northeast-1":        "alikafka.ap-south-1.aliyuncs.com",
			"cn-shenzhen-finance-1": "alikafka.aliyuncs.com",
			"me-east-1":             "alikafka.ap-south-1.aliyuncs.com",
			"cn-chengdu":            "alikafka.aliyuncs.com",
			"cn-north-2-gov-1":      "alikafka.aliyuncs.com",
			"cn-shanghai-finance-1": "alikafka.aliyuncs.com",
			"cn-hangzhou-finance":   "alikafka.aliyuncs.com",
			"ap-southeast-2":        "alikafka.ap-south-1.aliyuncs.com",
			"ap-southeast-3":        "alikafka.ap-south-1.aliyuncs.com",
			"eu-central-1":          "alikafka.ap-south-1.aliyuncs.com",
			"us-east-1":             "alikafka.ap-south-1.aliyuncs.com",
			"us-west-1":             "alikafka.ap-south-1.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
