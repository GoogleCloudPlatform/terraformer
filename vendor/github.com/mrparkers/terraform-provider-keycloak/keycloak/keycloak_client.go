package keycloak

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

type KeycloakClient struct {
	baseUrl           string
	realm             string
	clientCredentials *ClientCredentials
	httpClient        *http.Client
	initialLogin      bool
}

type ClientCredentials struct {
	ClientId     string
	ClientSecret string
	Username     string
	Password     string
	GrantType    string
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

const (
	apiUrl   = "/auth/admin"
	tokenUrl = "%s/auth/realms/%s/protocol/openid-connect/token"
)

func NewKeycloakClient(baseUrl, clientId, clientSecret, realm, username, password string, initialLogin bool, clientTimeout int, caCert string, tlsInsecureSkipVerify bool) (*KeycloakClient, error) {
	cookieJar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})

	if err != nil {
		return nil, err
	}
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: tlsInsecureSkipVerify},
		Proxy:           http.ProxyFromEnvironment,
	}

	httpClient := &http.Client{
		Timeout:   time.Second * time.Duration(clientTimeout),
		Transport: transport,
		Jar:       cookieJar,
	}

	if caCert != "" {
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM([]byte(caCert))
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		}
	}
	clientCredentials := &ClientCredentials{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
	if password != "" && username != "" {
		clientCredentials.Username = username
		clientCredentials.Password = password
		clientCredentials.GrantType = "password"
	} else if clientSecret != "" {
		clientCredentials.GrantType = "client_credentials"
	} else {
		return nil, fmt.Errorf("must specify client id, username and password for password grant, or client id and secret for client credentials grant")
	}

	keycloakClient := KeycloakClient{
		baseUrl:           baseUrl,
		clientCredentials: clientCredentials,
		httpClient:        httpClient,
		initialLogin:      initialLogin,
		realm:             realm,
	}

	if keycloakClient.initialLogin {
		err := keycloakClient.login()
		if err != nil {
			return nil, err
		}
	}

	return &keycloakClient, nil
}

func (keycloakClient *KeycloakClient) login() error {
	accessTokenUrl := fmt.Sprintf(tokenUrl, keycloakClient.baseUrl, keycloakClient.realm)
	accessTokenData := url.Values{}
	accessTokenData.Set("client_id", keycloakClient.clientCredentials.ClientId)
	accessTokenData.Set("grant_type", keycloakClient.clientCredentials.GrantType)

	if keycloakClient.clientCredentials.GrantType == "password" {
		accessTokenData.Set("username", keycloakClient.clientCredentials.Username)
		accessTokenData.Set("password", keycloakClient.clientCredentials.Password)

		if keycloakClient.clientCredentials.ClientSecret != "" {
			accessTokenData.Set("client_secret", keycloakClient.clientCredentials.ClientSecret)
		}

	} else if keycloakClient.clientCredentials.GrantType == "client_credentials" {
		accessTokenData.Set("client_secret", keycloakClient.clientCredentials.ClientSecret)
	}

	log.Printf("[DEBUG] Login request: %s", accessTokenData.Encode())

	accessTokenRequest, _ := http.NewRequest(http.MethodPost, accessTokenUrl, strings.NewReader(accessTokenData.Encode()))
	accessTokenRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	accessTokenResponse, err := keycloakClient.httpClient.Do(accessTokenRequest)
	if err != nil {
		return err
	}

	defer accessTokenResponse.Body.Close()

	body, _ := ioutil.ReadAll(accessTokenResponse.Body)

	log.Printf("[DEBUG] Login response: %s", body)

	var clientCredentials ClientCredentials
	err = json.Unmarshal(body, &clientCredentials)

	if err != nil {
		return err
	}

	keycloakClient.clientCredentials.AccessToken = clientCredentials.AccessToken
	keycloakClient.clientCredentials.RefreshToken = clientCredentials.RefreshToken
	keycloakClient.clientCredentials.TokenType = clientCredentials.TokenType

	return nil
}

func (keycloakClient *KeycloakClient) refresh() error {
	refreshTokenUrl := fmt.Sprintf(tokenUrl, keycloakClient.baseUrl, keycloakClient.realm)
	refreshTokenData := url.Values{}
	refreshTokenData.Set("client_id", keycloakClient.clientCredentials.ClientId)
	refreshTokenData.Set("grant_type", keycloakClient.clientCredentials.GrantType)

	if keycloakClient.clientCredentials.GrantType == "password" {
		refreshTokenData.Set("username", keycloakClient.clientCredentials.Username)
		refreshTokenData.Set("password", keycloakClient.clientCredentials.Password)
	} else if keycloakClient.clientCredentials.GrantType == "client_credentials" {
		refreshTokenData.Set("client_secret", keycloakClient.clientCredentials.ClientSecret)
	}

	log.Printf("[DEBUG] Refresh request: %s", refreshTokenData.Encode())

	accessTokenRequest, _ := http.NewRequest(http.MethodPost, refreshTokenUrl, strings.NewReader(refreshTokenData.Encode()))
	accessTokenRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	refreshTokenResponse, err := keycloakClient.httpClient.Do(accessTokenRequest)
	if err != nil {
		return err
	}

	defer refreshTokenResponse.Body.Close()

	body, _ := ioutil.ReadAll(refreshTokenResponse.Body)

	log.Printf("[DEBUG] Refresh response: %s", body)

	// Handle 401 "User or client no longer has role permissions for client key" until I better understand why that happens in the first place
	if refreshTokenResponse.StatusCode == http.StatusBadRequest {
		log.Printf("[DEBUG] Unexpected 400, attemting to log in again")

		return keycloakClient.login()
	}

	var clientCredentials ClientCredentials
	err = json.Unmarshal(body, &clientCredentials)
	if err != nil {
		return err
	}

	keycloakClient.clientCredentials.AccessToken = clientCredentials.AccessToken
	keycloakClient.clientCredentials.RefreshToken = clientCredentials.RefreshToken
	keycloakClient.clientCredentials.TokenType = clientCredentials.TokenType

	return nil
}

func (keycloakClient *KeycloakClient) addRequestHeaders(request *http.Request) {
	tokenType := keycloakClient.clientCredentials.TokenType
	accessToken := keycloakClient.clientCredentials.AccessToken

	request.Header.Set("Authorization", fmt.Sprintf("%s %s", tokenType, accessToken))
	request.Header.Set("Accept", "application/json")

	if request.Method == http.MethodPost || request.Method == http.MethodPut || request.Method == http.MethodDelete {
		request.Header.Set("Content-type", "application/json")
	}
}

/**
Sends an HTTP request and refreshes credentials on 403 or 401 errors
*/
func (keycloakClient *KeycloakClient) sendRequest(request *http.Request) ([]byte, string, error) {
	if !keycloakClient.initialLogin {
		keycloakClient.initialLogin = true
		err := keycloakClient.login()
		if err != nil {
			return nil, "", fmt.Errorf("error logging in: %s", err)
		}
	}
	requestMethod := request.Method
	requestPath := request.URL.Path

	log.Printf("[DEBUG] Sending %s to %s", requestMethod, requestPath)
	showBody := false
	if request.Body != nil {
		showBody = true
		requestBody, err := request.GetBody()
		if err != nil {
			return nil, "", err
		}

		requestBodyBuffer := new(bytes.Buffer)
		requestBodyBuffer.ReadFrom(requestBody)

		log.Printf("[DEBUG] Request body: %s", requestBodyBuffer.String())
	}

	keycloakClient.addRequestHeaders(request)

	dump, err := httputil.DumpRequest(request, showBody)
	if err != nil {
		return nil, "", err
	}
	log.Printf("[DEBUG] %s", dump)

	response, err := keycloakClient.httpClient.Do(request)
	if err != nil {
		return nil, "", err
	}

	// Unauthorized: Token could have expired
	// Forbidden: After creating a realm, following GETs for the realm return 403 until you refresh
	if response.StatusCode == http.StatusUnauthorized || response.StatusCode == http.StatusForbidden {
		log.Printf("[DEBUG] Response: %s.  Attempting refresh", response.Status)

		err := keycloakClient.refresh()
		if err != nil {
			return nil, "", fmt.Errorf("error refreshing credentials: %s", err)
		}

		keycloakClient.addRequestHeaders(request)

		response, err = keycloakClient.httpClient.Do(request)
		if err != nil {
			return nil, "", err
		}
	}

	log.Printf("[DEBUG] Response: %s", response.Status)

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, "", err
	}

	if len(body) != 0 {
		log.Printf("[DEBUG] Response body: %s", body)
	}

	if response.StatusCode >= 400 {
		return nil, "", &ApiError{
			Code:    response.StatusCode,
			Message: fmt.Sprintf("error sending %s request to %s: %s", request.Method, request.URL.Path, response.Status),
		}
	}

	return body, response.Header.Get("Location"), nil
}

func (keycloakClient *KeycloakClient) get(path string, resource interface{}, params map[string]string) error {
	body, err := keycloakClient.getRaw(path, params)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, resource)
}

func (keycloakClient *KeycloakClient) getRaw(path string, params map[string]string) ([]byte, error) {
	resourceUrl := keycloakClient.baseUrl + apiUrl + path

	request, err := http.NewRequest(http.MethodGet, resourceUrl, nil)
	if err != nil {
		return nil, err
	}

	if params != nil {
		query := url.Values{}
		for k, v := range params {
			query.Add(k, v)
		}
		request.URL.RawQuery = query.Encode()
	}

	body, _, err := keycloakClient.sendRequest(request)
	return body, err
}

func (keycloakClient *KeycloakClient) post(path string, requestBody interface{}) ([]byte, string, error) {
	resourceUrl := keycloakClient.baseUrl + apiUrl + path

	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, "", err
	}

	request, err := http.NewRequest(http.MethodPost, resourceUrl, bytes.NewReader(payload))
	if err != nil {
		return nil, "", err
	}

	body, location, err := keycloakClient.sendRequest(request)

	return body, location, err
}

func (keycloakClient *KeycloakClient) put(path string, requestBody interface{}) error {
	resourceUrl := keycloakClient.baseUrl + apiUrl + path

	payload, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, resourceUrl, bytes.NewReader(payload))
	if err != nil {
		return err
	}

	_, _, err = keycloakClient.sendRequest(request)

	return err
}

func (keycloakClient *KeycloakClient) delete(path string, requestBody interface{}) error {
	resourceUrl := keycloakClient.baseUrl + apiUrl + path

	var body io.Reader

	if requestBody != nil {
		payload, err := json.Marshal(requestBody)
		if err != nil {
			return err
		}
		body = bytes.NewReader(payload)
	}

	request, err := http.NewRequest(http.MethodDelete, resourceUrl, body)
	if err != nil {
		return err
	}

	_, _, err = keycloakClient.sendRequest(request)

	return err
}
