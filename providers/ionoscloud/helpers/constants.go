package helpers

import "time"

// Provider Constants
const (
	ProviderName = "ionoscloud"
)

// Config Constants
const (
	IonosDebug   = "IONOS_DEBUG"
	Ionos        = "ionoscloud"
	DcId         = "datacenter_id"
	ServerId     = "server_id"
	NicId        = "nic_id"
	K8sClusterId = "k8s_cluster_id"

	UsernameArg = "username"
	PasswordArg = "password"
	TokenArg    = "token"
	UrlArg      = "url"
	// MaxRetries - number of retries in case of rate-limit
	MaxRetries = 999

	// MaxWaitTime - waits 4 seconds before retry in case of rate limit
	MaxWaitTime = 4 * time.Second

	available = "Available"
)

const (
	CredentialsError = "set IONOS_USERNAME and IONOS_PASSWORD or IONOS_TOKEN env var"
)
