package main

type gcpResourceRenderble interface {
	getTerraformName() string
	getIgnoreKeys() []string
	getAdditionalFields() map[string]string
	getAllowEmptyValues() []string
	ifNeedRegion() bool
	ifNeedZone(zoneInParameters bool) bool
	ifIDWithZone(zoneInParameters bool) bool
}

type basicGCPResource struct {
	terraformName    string
	ignoreKeys       []string
	allowEmptyValues []string
	additionalFields map[string]string
}

func (b basicGCPResource) getTerraformName() string {
	return b.terraformName
}

func (b basicGCPResource) getIgnoreKeys() []string {
	return b.ignoreKeys
}

func (b basicGCPResource) getAdditionalFields() map[string]string {
	return b.additionalFields
}

func (b basicGCPResource) getAllowEmptyValues() []string {
	return b.allowEmptyValues
}
func (b basicGCPResource) ifNeedRegion() bool {
	return true
}

func (b basicGCPResource) ifNeedZone(zoneInParameters bool) bool {
	return zoneInParameters
}

func (b basicGCPResource) ifIDWithZone(zoneInParameters bool) bool {
	return zoneInParameters
}
