package octopusdeploy

// TenantedDeploymentMode indicates what types of deployments a resource participates in.
/*
ENUM(
	Untenanted // Untenanted resources only participate in Untenanted deployments
	TenantedOrUntenanted // TenantedOrUntenanted resources participate in any type of deployment
	Tenanted // Tenanted resources only participate in Tenanted deployments
)
*/
type TenantedDeploymentMode int
