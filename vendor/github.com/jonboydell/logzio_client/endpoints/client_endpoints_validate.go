package endpoints

func validSlackEndpoint(endpointType Endpoint) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Url) > 0 && len(endpointType.Description) > 0
}

func validCustomEndpoint(endpointType Endpoint) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Url) > 0 && len(endpointType.Description) > 0 &&
		len(endpointType.Method) > 0
}

func validPagerDutyEndpoint(endpointType Endpoint) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Description) > 0 &&
		len(endpointType.ServiceKey) > 0
}

func validBigPandaEndpoint(endpointType Endpoint) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Description) > 0 &&
		len(endpointType.ApiToken) > 0 && len(endpointType.AppKey) > 0
}

func validDataDogEndpoint(endpointType Endpoint) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Description) > 0 &&
		len(endpointType.ApiKey) > 0
}

func validVictorOpsEndpoint(endpointType Endpoint) bool {
	return len(endpointType.Title) > 0 &&
		len(endpointType.Description) > 0 &&
		len(endpointType.RoutingKey) > 0 && len(endpointType.MessageType) > 0 && len(endpointType.ServiceApiKey) > 0
}

// ValidateEndpointRequest validates an endpoint request for correctness given its type,
// returns FALSE if validation failed, true otherwise
func ValidateEndpointRequest(endpoint Endpoint) bool {
	switch endpoint.EndpointType {
	case EndpointTypeSlack:
		return validSlackEndpoint(endpoint)
	case EndpointTypeCustom:
		return validCustomEndpoint(endpoint)
	case EndpointTypePagerDuty:
		return validPagerDutyEndpoint(endpoint)
	case EndpointTypeBigPanda:
		return validBigPandaEndpoint(endpoint)
	case EndpointTypeDataDog:
		return validDataDogEndpoint(endpoint)
	case EndpointTypeVictorOps:
		return validVictorOpsEndpoint(endpoint)
	default:
		return false
	}
}
