package rds

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "rds.aliyuncs.com",
			"cn-beijing-gov-1":            "rds.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "rds.aliyuncs.com",
			"cn-beijing":                  "rds.aliyuncs.com",
			"cn-shanghai-inner":           "rds.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "rds.aliyuncs.com",
			"cn-haidian-cm12-c01":         "rds.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "rds.aliyuncs.com",
			"cn-north-2-gov-1":            "rds.aliyuncs.com",
			"cn-yushanfang":               "rds.aliyuncs.com",
			"cn-qingdao":                  "rds.aliyuncs.com",
			"cn-hongkong-finance-pop":     "rds.aliyuncs.com",
			"cn-qingdao-nebula":           "rds.aliyuncs.com",
			"cn-shanghai":                 "rds.aliyuncs.com",
			"cn-shanghai-finance-1":       "rds.aliyuncs.com",
			"cn-hongkong":                 "rds.aliyuncs.com",
			"cn-beijing-finance-pop":      "rds.aliyuncs.com",
			"cn-wuhan":                    "rds.aliyuncs.com",
			"us-west-1":                   "rds.aliyuncs.com",
			"cn-shenzhen":                 "rds.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "rds.aliyuncs.com",
			"rus-west-1-pop":              "rds.ap-northeast-1.aliyuncs.com",
			"cn-shanghai-et15-b01":        "rds.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "rds.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "rds.aliyuncs.com",
			"eu-west-1-oxs":               "rds.ap-northeast-1.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "rds.aliyuncs.com",
			"cn-beijing-finance-1":        "rds.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "rds.aliyuncs.com",
			"cn-shenzhen-finance-1":       "rds.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "rds.aliyuncs.com",
			"cn-hangzhou-test-306":        "rds.aliyuncs.com",
			"cn-shanghai-et2-b01":         "rds.aliyuncs.com",
			"cn-hangzhou-finance":         "rds.aliyuncs.com",
			"ap-southeast-1":              "rds.aliyuncs.com",
			"cn-beijing-nu16-b01":         "rds.aliyuncs.com",
			"cn-edge-1":                   "rds.aliyuncs.com",
			"us-east-1":                   "rds.aliyuncs.com",
			"cn-fujian":                   "rds.aliyuncs.com",
			"ap-northeast-2-pop":          "rds.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen-inner":           "rds.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "rds.aliyuncs.com",
			"cn-hangzhou":                 "rds.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
