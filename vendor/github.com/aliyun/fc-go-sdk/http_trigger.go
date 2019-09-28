package fc

const (
	// AuthAnonymous defines http trigger without authorized
	AuthAnonymous = "anonymous"

	// AuthFunction defines http trigger authorized by AK
	AuthFunction = "function"
)

// DefaultAuthType ...
var DefaultAuthType = AuthFunction

// HTTPTriggerConfig ..
type HTTPTriggerConfig struct {
	AuthType *string  `json:"authType"`
	Methods  []string `json:"methods"`
}

// NewHTTPTriggerConfig ...
func NewHTTPTriggerConfig() *HTTPTriggerConfig {
	return &HTTPTriggerConfig{}
}

// WithMethods ...
func (t *HTTPTriggerConfig) WithMethods(methods ...string) *HTTPTriggerConfig {
	t.Methods = make([]string, len(methods))
	for i := range methods {
		t.Methods[i] = methods[i]
	}
	return t
}

// WithAuthType ...
func (t *HTTPTriggerConfig) WithAuthType(authType string) *HTTPTriggerConfig {
	t.AuthType = &authType
	return t
}
