package emr

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":    "emr.aliyuncs.com",
			"cn-beijing":     "emr.aliyuncs.com",
			"cn-shanghai":    "emr.aliyuncs.com",
			"ap-southeast-1": "emr.aliyuncs.com",
			"us-west-1":      "emr.aliyuncs.com",
			"cn-hangzhou":    "emr.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
