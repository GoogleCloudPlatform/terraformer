package keycloak

import (
	"fmt"
	"strings"
)

type Key struct {
	Algorithm        *string `json:"algorithm,omitempty"`
	Certificate      *string `json:"certificate,omitempty"`
	ProviderId       *string `json:"providerId,omitempty"`
	ProviderPriority *int    `json:"providerPriority,omitempty"`
	PublicKey        *string `json:"publicKey,omitempty"`
	Kid              *string `json:"kid,omitempty"`
	Status           *string `json:"status,omitempty"`
	Type             *string `json:"type,omitempty"`
}

type Keys struct {
	Keys []Key `json:"keys"`
}

type Realm struct {
	Id              string `json:"id"`
	Realm           string `json:"realm"`
	Enabled         bool   `json:"enabled"`
	DisplayName     string `json:"displayName"`
	DisplayNameHtml string `json:"displayNameHtml"`

	// Login Config
	RegistrationAllowed         bool   `json:"registrationAllowed"`
	RegistrationEmailAsUsername bool   `json:"registrationEmailAsUsername"`
	EditUsernameAllowed         bool   `json:"editUsernameAllowed"`
	ResetPasswordAllowed        bool   `json:"resetPasswordAllowed"`
	RememberMe                  bool   `json:"rememberMe"`
	VerifyEmail                 bool   `json:"verifyEmail"`
	LoginWithEmailAllowed       bool   `json:"loginWithEmailAllowed"`
	DuplicateEmailsAllowed      bool   `json:"duplicateEmailsAllowed"`
	SslRequired                 string `json:"sslRequired,omitempty"`

	//SMTP Server
	SmtpServer SmtpServer `json:"smtpServer"`

	// Themes
	LoginTheme   string `json:"loginTheme,omitempty"`
	AccountTheme string `json:"accountTheme,omitempty"`
	AdminTheme   string `json:"adminTheme,omitempty"`
	EmailTheme   string `json:"emailTheme,omitempty"`

	// Tokens
	RevokeRefreshToken                  bool `json:"revokeRefreshToken"`
	RefreshTokenMaxReuse                int  `json:"refreshTokenMaxReuse"`
	SsoSessionIdleTimeout               int  `json:"ssoSessionIdleTimeout,omitempty"`
	SsoSessionMaxLifespan               int  `json:"ssoSessionMaxLifespan,omitempty"`
	OfflineSessionIdleTimeout           int  `json:"offlineSessionIdleTimeout,omitempty"`
	OfflineSessionMaxLifespan           int  `json:"offlineSessionMaxLifespan,omitempty"`
	AccessTokenLifespan                 int  `json:"accessTokenLifespan,omitempty"`
	AccessTokenLifespanForImplicitFlow  int  `json:"accessTokenLifespanForImplicitFlow,omitempty"`
	AccessCodeLifespan                  int  `json:"accessCodeLifespan,omitempty"`
	AccessCodeLifespanLogin             int  `json:"accessCodeLifespanLogin,omitempty"`
	AccessCodeLifespanUserAction        int  `json:"accessCodeLifespanUserAction,omitempty"`
	ActionTokenGeneratedByUserLifespan  int  `json:"actionTokenGeneratedByUserLifespan,omitempty"`
	ActionTokenGeneratedByAdminLifespan int  `json:"actionTokenGeneratedByAdminLifespan,omitempty"`

	//internationalization
	InternationalizationEnabled bool     `json:"internationalizationEnabled"`
	SupportLocales              []string `json:"supportedLocales"`
	DefaultLocale               string   `json:"defaultLocale"`

	//extra attributes of a realm
	Attributes map[string]interface{} `json:"attributes"`

	BrowserSecurityHeaders BrowserSecurityHeaders `json:"browserSecurityHeaders"`

	BruteForceProtected          bool `json:"bruteForceProtected"`
	PermanentLockout             bool `json:"permanentLockout"`
	FailureFactor                int  `json:"failureFactor"` //Max Login Failures
	WaitIncrementSeconds         int  `json:"waitIncrementSeconds"`
	QuickLoginCheckMilliSeconds  int  `json:"quickLoginCheckMilliSeconds"`
	MinimumQuickLoginWaitSeconds int  `json:"minimumQuickLoginWaitSeconds"`
	MaxFailureWaitSeconds        int  `json:"maxFailureWaitSeconds"` //Max Wait
	MaxDeltaTimeSeconds          int  `json:"maxDeltaTimeSeconds"`   //Failure Reset Time

	PasswordPolicy string `json:"passwordPolicy"`

	//flow bindings
	BrowserFlow              string `json:"browserFlow,omitempty"`
	RegistrationFlow         string `json:"registrationFlow,omitempty"`
	DirectGrantFlow          string `json:"directGrantFlow,omitempty"`
	ResetCredentialsFlow     string `json:"resetCredentialsFlow,omitempty"`
	ClientAuthenticationFlow string `json:"clientAuthenticationFlow,omitempty"`
	DockerAuthenticationFlow string `json:"dockerAuthenticationFlow,omitempty"`
}

type BrowserSecurityHeaders struct {
	ContentSecurityPolicy           string `json:"contentSecurityPolicy"`
	ContentSecurityPolicyReportOnly string `json:"contentSecurityPolicyReportOnly"`
	StrictTransportSecurity         string `json:"strictTransportSecurity"`
	XContentTypeOptions             string `json:"xContentTypeOptions"`
	XFrameOptions                   string `json:"xFrameOptions"`
	XRobotsTag                      string `json:"xRobotsTag"`
	XXSSProtection                  string `json:"xXSSProtection"`
}

type SmtpServer struct {
	StartTls           KeycloakBoolQuoted `json:"starttls,omitempty"`
	Auth               KeycloakBoolQuoted `json:"auth,omitempty"`
	Port               string             `json:"port,omitempty"`
	Host               string             `json:"host,omitempty"`
	ReplyTo            string             `json:"replyTo,omitempty"`
	ReplyToDisplayName string             `json:"replyToDisplayName,omitempty"`
	From               string             `json:"from,omitempty"`
	FromDisplayName    string             `json:"fromDisplayName,omitempty"`
	EnvelopeFrom       string             `json:"envelopeFrom,omitempty"`
	Ssl                KeycloakBoolQuoted `json:"ssl,omitempty"`
	User               string             `json:"user,omitempty"`
	Password           string             `json:"password,omitempty"`
}

func (keycloakClient *KeycloakClient) NewRealm(realm *Realm) error {
	_, _, err := keycloakClient.post("/realms", realm)

	return err
}

func (keycloakClient *KeycloakClient) GetRealm(id string) (*Realm, error) {
	var realm Realm

	err := keycloakClient.get(fmt.Sprintf("/realms/%s", id), &realm, nil)
	if err != nil {
		return nil, err
	}
	return &realm, nil
}

func (keycloakClient *KeycloakClient) GetRealms() ([]*Realm, error) {
	var realms []*Realm

	err := keycloakClient.get("/realms", &realms, nil)
	if err != nil {
		return nil, err
	}

	return realms, nil
}

func (keycloakClient *KeycloakClient) GetRealmKeys(id string) (*Keys, error) {
	var keys Keys

	err := keycloakClient.get(fmt.Sprintf("/realms/%s/keys", id), &keys, nil)
	if err != nil {
		return nil, err
	}

	return &keys, nil
}

func (keycloakClient *KeycloakClient) UpdateRealm(realm *Realm) error {
	return keycloakClient.put(fmt.Sprintf("/realms/%s", realm.Id), realm)
}

func (keycloakClient *KeycloakClient) DeleteRealm(id string) error {
	err := keycloakClient.delete(fmt.Sprintf("/realms/%s", id), nil)
	if err != nil {
		// For whatever reason, this fails sometimes with a 500 during acceptance tests. try again
		return keycloakClient.delete(fmt.Sprintf("/realms/%s", id), nil)
	}

	return nil
}

func (keycloakClient *KeycloakClient) ValidateRealm(realm *Realm) error {
	if realm.RegistrationAllowed == false && realm.RegistrationEmailAsUsername == true {
		return fmt.Errorf("validation error: RegistrationEmailAsUsername cannot be true if RegistrationAllowed is false")
	}

	if realm.DuplicateEmailsAllowed == true && realm.RegistrationEmailAsUsername == true {
		return fmt.Errorf("validation error: DuplicateEmailsAllowed cannot be true if RegistrationEmailAsUsername is true")
	}

	if realm.DuplicateEmailsAllowed == true && realm.LoginWithEmailAllowed == true {
		return fmt.Errorf("validation error: DuplicateEmailsAllowed cannot be true if LoginWithEmailAllowed is true")
	}

	if realm.SslRequired != "none" && realm.SslRequired != "external" && realm.SslRequired != "all" {
		return fmt.Errorf("validation error: SslRequired should be 'none', 'external' or 'all'")
	}

	// validate if the given theme exists on the server. the keycloak API allows you to use any random string for a theme
	serverInfo, err := keycloakClient.GetServerInfo()
	if err != nil {
		return err
	}

	if realm.LoginTheme != "" && !serverInfo.ThemeIsInstalled("login", realm.LoginTheme) {
		return fmt.Errorf("validation error: theme \"%s\" does not exist on the server", realm.LoginTheme)
	}

	if realm.AccountTheme != "" && !serverInfo.ThemeIsInstalled("account", realm.AccountTheme) {
		return fmt.Errorf("validation error: theme \"%s\" does not exist on the server", realm.AccountTheme)
	}

	if realm.AdminTheme != "" && !serverInfo.ThemeIsInstalled("admin", realm.AdminTheme) {
		return fmt.Errorf("validation error: theme \"%s\" does not exist on the server", realm.AdminTheme)
	}

	if realm.EmailTheme != "" && !serverInfo.ThemeIsInstalled("email", realm.EmailTheme) {
		return fmt.Errorf("validation error: theme \"%s\" does not exist on the server", realm.EmailTheme)
	}

	if realm.InternationalizationEnabled == true && !contains(realm.SupportLocales, realm.DefaultLocale) {
		return fmt.Errorf("validation error: DefaultLocale should be in the SupportLocales")
	}

	if realm.PasswordPolicy != "" {
		policies := strings.Split(realm.PasswordPolicy, " and ")
		for _, policyTypeRepresentation := range policies {
			policy := strings.Split(policyTypeRepresentation, "(")
			if !serverInfo.providerInstalled("password-policy", policy[0]) {
				return fmt.Errorf("validation error: password-policy \"%s\" does not exist on the server, installed providers: %s", policy[0], serverInfo.getInstalledProvidersNames("password-policy"))
			}
		}
	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
