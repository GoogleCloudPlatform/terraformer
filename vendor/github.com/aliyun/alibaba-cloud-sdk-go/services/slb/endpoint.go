package slb

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "slb.aliyuncs.com",
			"cn-beijing":            "slb.aliyuncs.com",
			"cn-shenzhen-finance-1": "slb.aliyuncs.com",
			"cn-north-2-gov-1":      "slb.aliyuncs.com",
			"cn-qingdao":            "slb.aliyuncs.com",
			"cn-shanghai":           "slb.aliyuncs.com",
			"cn-shanghai-finance-1": "slb.aliyuncs.com",
			"cn-hongkong":           "slb.aliyuncs.com",
			"cn-hangzhou-finance":   "slb.aliyuncs.com",
			"ap-southeast-1":        "slb.aliyuncs.com",
			"us-east-1":             "slb.aliyuncs.com",
			"us-west-1":             "slb.aliyuncs.com",
			"cn-hangzhou":           "slb.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
