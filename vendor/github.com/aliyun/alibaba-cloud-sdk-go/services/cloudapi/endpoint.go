package cloudapi

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "apigateway.cn-shenzhen.aliyuncs.com",
			"cn-beijing":            "apigateway.cn-beijing.aliyuncs.com",
			"ap-south-1":            "apigateway.ap-south-1.aliyuncs.com",
			"eu-west-1":             "apigateway.eu-west-1.aliyuncs.com",
			"ap-northeast-1":        "apigateway.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen-finance-1": "apigateway.aliyuncs.com",
			"me-east-1":             "apigateway.me-east-1.aliyuncs.com",
			"cn-chengdu":            "apigateway.cn-chengdu.aliyuncs.com",
			"cn-north-2-gov-1":      "apigateway.cn-north-2-gov-1.aliyuncs.com",
			"cn-qingdao":            "apigateway.cn-qingdao.aliyuncs.com",
			"cn-shanghai":           "apigateway.cn-shanghai.aliyuncs.com",
			"cn-shanghai-finance-1": "apigateway.aliyuncs.com",
			"cn-hongkong":           "apigateway.cn-hongkong.aliyuncs.com",
			"cn-hangzhou-finance":   "apigateway.aliyuncs.com",
			"ap-southeast-1":        "apigateway.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2":        "apigateway.ap-southeast-2.aliyuncs.com",
			"ap-southeast-3":        "apigateway.ap-southeast-3.aliyuncs.com",
			"eu-central-1":          "apigateway.eu-central-1.aliyuncs.com",
			"cn-huhehaote":          "apigateway.cn-huhehaote.aliyuncs.com",
			"ap-southeast-5":        "apigateway.ap-southeast-5.aliyuncs.com",
			"us-east-1":             "apigateway.us-east-1.aliyuncs.com",
			"cn-zhangjiakou":        "apigateway.cn-zhangjiakou.aliyuncs.com",
			"us-west-1":             "apigateway.us-west-1.aliyuncs.com",
			"cn-hangzhou":           "apigateway.cn-hangzhou.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
