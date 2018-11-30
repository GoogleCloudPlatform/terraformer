package main

type instances struct {
	basicGCPResource
}

func (b instances) ifNeedZone(zoneInParameters bool) bool {
	return true
}

func (b instances) ifIDWithZone(zoneInParameters bool) bool {
	return false
}
func (b instances) ifNeedRegion() bool {
	return false
}
