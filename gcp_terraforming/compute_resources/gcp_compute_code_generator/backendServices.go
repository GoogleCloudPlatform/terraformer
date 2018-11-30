package main

type backendServices struct {
	basicGCPResource
}

func (b backendServices) ifNeedRegion() bool {
	return false
}
