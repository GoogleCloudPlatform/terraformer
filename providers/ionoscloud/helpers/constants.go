package helpers

import "time"

// Provider Constants
const (
	ProviderName = "ionoscloud"
)

// Config Constants
const (
	IonosUsername = "IONOS_USERNAME"
	IonosPassword = "IONOS_PASSWORD"
	IonosToken    = "IONOS_TOKEN"
	IonosApiUrl   = "IONOS_API_URL"
	IonosDebug    = "IONOS_DEBUG"

	UsernameArg = "username"
	PasswordArg = "password"
	TokenArg    = "token"
	UrlArg      = "url"
	// MaxRetries - number of retries in case of rate-limit
	MaxRetries = 999

	// MaxWaitTime - waits 4 seconds before retry in case of rate limit
	MaxWaitTime = 4 * time.Second

	sleepInterval = 5 * time.Second
	available     = "Available"
)

// Error Messages
const (
	CredentialsError = "set IONOS_USERNAME and IONOS_PASSWORD or IONOS_TOKEN env var"
)
