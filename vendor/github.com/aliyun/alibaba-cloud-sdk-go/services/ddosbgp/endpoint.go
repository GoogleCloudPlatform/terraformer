package ddosbgp

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "ddosbgp.aliyuncs.com",
			"cn-beijing":            "ddosbgp.aliyuncs.com",
			"ap-south-1":            "ddosbgp.ap-southeast-1.aliyuncs.com",
			"eu-west-1":             "ddosbgp.ap-southeast-1.aliyuncs.com",
			"ap-northeast-1":        "ddosbgp.ap-southeast-1.aliyuncs.com",
			"cn-shenzhen-finance-1": "ddosbgp.aliyuncs.com",
			"me-east-1":             "ddosbgp.ap-southeast-1.aliyuncs.com",
			"cn-chengdu":            "ddosbgp.aliyuncs.com",
			"cn-north-2-gov-1":      "ddosbgp.aliyuncs.com",
			"cn-qingdao":            "ddosbgp.aliyuncs.com",
			"cn-shanghai":           "ddosbgp.aliyuncs.com",
			"cn-shanghai-finance-1": "ddosbgp.aliyuncs.com",
			"cn-hangzhou-finance":   "ddosbgp.aliyuncs.com",
			"ap-southeast-2":        "ddosbgp.ap-southeast-1.aliyuncs.com",
			"ap-southeast-3":        "ddosbgp.ap-southeast-1.aliyuncs.com",
			"eu-central-1":          "ddosbgp.ap-southeast-1.aliyuncs.com",
			"cn-huhehaote":          "ddosbgp.aliyuncs.com",
			"ap-southeast-5":        "ddosbgp.ap-southeast-1.aliyuncs.com",
			"cn-zhangjiakou":        "ddosbgp.aliyuncs.com",
			"cn-hangzhou":           "ddosbgp.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
