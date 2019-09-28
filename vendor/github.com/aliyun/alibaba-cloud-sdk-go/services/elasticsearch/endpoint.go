package elasticsearch

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "elasticsearch.aliyuncs.com",
			"cn-beijing-gov-1":            "elasticsearch.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "elasticsearch.aliyuncs.com",
			"cn-shanghai-inner":           "elasticsearch.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "elasticsearch.aliyuncs.com",
			"cn-haidian-cm12-c01":         "elasticsearch.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "elasticsearch.aliyuncs.com",
			"cn-north-2-gov-1":            "elasticsearch.aliyuncs.com",
			"cn-yushanfang":               "elasticsearch.aliyuncs.com",
			"cn-hongkong-finance-pop":     "elasticsearch.aliyuncs.com",
			"cn-qingdao-nebula":           "elasticsearch.aliyuncs.com",
			"cn-beijing-finance-pop":      "elasticsearch.aliyuncs.com",
			"cn-wuhan":                    "elasticsearch.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "elasticsearch.aliyuncs.com",
			"rus-west-1-pop":              "elasticsearch.ap-northeast-1.aliyuncs.com",
			"cn-shanghai-et15-b01":        "elasticsearch.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "elasticsearch.aliyuncs.com",
			"eu-west-1":                   "elasticsearch.ap-northeast-1.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "elasticsearch.aliyuncs.com",
			"eu-west-1-oxs":               "elasticsearch.ap-northeast-1.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "elasticsearch.aliyuncs.com",
			"cn-beijing-finance-1":        "elasticsearch.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "elasticsearch.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "elasticsearch.aliyuncs.com",
			"cn-shenzhen-finance-1":       "elasticsearch.aliyuncs.com",
			"me-east-1":                   "elasticsearch.ap-northeast-1.aliyuncs.com",
			"cn-chengdu":                  "elasticsearch.aliyuncs.com",
			"cn-hangzhou-test-306":        "elasticsearch.aliyuncs.com",
			"cn-shanghai-et2-b01":         "elasticsearch.aliyuncs.com",
			"cn-beijing-nu16-b01":         "elasticsearch.aliyuncs.com",
			"cn-edge-1":                   "elasticsearch.aliyuncs.com",
			"cn-huhehaote":                "elasticsearch.aliyuncs.com",
			"cn-fujian":                   "elasticsearch.aliyuncs.com",
			"us-east-1":                   "elasticsearch.ap-northeast-1.aliyuncs.com",
			"ap-northeast-2-pop":          "elasticsearch.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen-inner":           "elasticsearch.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "elasticsearch.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
