package main

type instanceGroupManagers struct {
	basicGCPResource
}

func (b instanceGroupManagers) ifNeedZone(zoneInParameters bool) bool {
	return true
}

func (b instanceGroupManagers) ifIDWithZone(zoneInParameters bool) bool {
	return false
}
func (b instanceGroupManagers) ifNeedRegion() bool {
	return false
}
