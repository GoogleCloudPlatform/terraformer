package alidns

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "central"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "dns.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "dns.aliyuncs.com",
			"cn-beijing":                  "alidns.cn-beijing.aliyuncs.com",
			"cn-shanghai-inner":           "dns.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "dns.aliyuncs.com",
			"cn-north-2-gov-1":            "dns.aliyuncs.com",
			"cn-yushanfang":               "dns.aliyuncs.com",
			"cn-qingdao-nebula":           "dns.aliyuncs.com",
			"cn-beijing-finance-pop":      "dns.aliyuncs.com",
			"cn-wuhan":                    "dns.aliyuncs.com",
			"cn-zhangjiakou":              "alidns.cn-zhangjiakou.aliyuncs.com",
			"us-west-1":                   "alidns.ap-southeast-1.aliyuncs.com",
			"rus-west-1-pop":              "alidns.ap-southeast-1.aliyuncs.com",
			"cn-shanghai-et15-b01":        "dns.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "dns.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "dns.aliyuncs.com",
			"ap-northeast-1":              "alidns.ap-southeast-1.aliyuncs.com",
			"cn-shanghai-et2-b01":         "dns.aliyuncs.com",
			"ap-southeast-1":              "alidns.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2":              "alidns.ap-southeast-1.aliyuncs.com",
			"ap-southeast-3":              "alidns.ap-southeast-1.aliyuncs.com",
			"ap-southeast-5":              "alidns.ap-southeast-1.aliyuncs.com",
			"us-east-1":                   "alidns.ap-southeast-1.aliyuncs.com",
			"cn-shenzhen-inner":           "dns.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "dns.aliyuncs.com",
			"cn-beijing-gov-1":            "dns.aliyuncs.com",
			"ap-south-1":                  "alidns.ap-southeast-1.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "dns.aliyuncs.com",
			"cn-haidian-cm12-c01":         "dns.aliyuncs.com",
			"cn-qingdao":                  "dns.aliyuncs.com",
			"cn-hongkong-finance-pop":     "dns.aliyuncs.com",
			"cn-shanghai":                 "alidns.cn-shanghai.aliyuncs.com",
			"cn-shanghai-finance-1":       "dns.aliyuncs.com",
			"cn-hongkong":                 "alidns.cn-hongkong.aliyuncs.com",
			"eu-central-1":                "alidns.eu-central-1.aliyuncs.com",
			"cn-shenzhen":                 "alidns.cn-shenzhen.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "dns.aliyuncs.com",
			"eu-west-1":                   "alidns.ap-southeast-1.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "dns.aliyuncs.com",
			"eu-west-1-oxs":               "alidns.ap-southeast-1.aliyuncs.com",
			"cn-beijing-finance-1":        "dns.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "dns.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "dns.aliyuncs.com",
			"cn-shenzhen-finance-1":       "dns.aliyuncs.com",
			"me-east-1":                   "alidns.ap-southeast-1.aliyuncs.com",
			"cn-chengdu":                  "dns.aliyuncs.com",
			"cn-hangzhou-test-306":        "dns.aliyuncs.com",
			"cn-hangzhou-finance":         "dns.aliyuncs.com",
			"cn-beijing-nu16-b01":         "dns.aliyuncs.com",
			"cn-edge-1":                   "dns.aliyuncs.com",
			"cn-huhehaote":                "alidns.cn-huhehaote.aliyuncs.com",
			"cn-fujian":                   "dns.aliyuncs.com",
			"ap-northeast-2-pop":          "alidns.ap-southeast-1.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
