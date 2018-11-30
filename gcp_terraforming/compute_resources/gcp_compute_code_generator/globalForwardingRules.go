package main


type globalForwardingRules struct {
	basicGCPResource
}

func (b globalForwardingRules) ifNeedRegion() bool {
	return false
}

