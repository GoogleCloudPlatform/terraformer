package octopusdeploy

type AutoDeployReleaseOverride struct {

	// environment Id
	// Read Only: true
	EnvironmentID string `json:"EnvironmentId,omitempty"`

	// release Id
	// Read Only: true
	ReleaseID string `json:"ReleaseId,omitempty"`

	// tenant Id
	// Read Only: true
	TenantID string `json:"TenantId,omitempty"`
}
