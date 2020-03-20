package octopusdeploy

type ProjectConnectivityPolicy struct {
	AllowDeploymentsToNoTargets bool     `json:"AllowDeploymentsToNoTargets,omitempty"`
	TargetRoles                 []string `json:"TargetRoles,omitempty"`
	SkipMachineBehavior         string   `json:"SkipMachineBehavior,omitempty"`
}
