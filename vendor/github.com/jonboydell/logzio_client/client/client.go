package client

import (
	"crypto/tls"
	"log"
	"net/http"
)


const (
	ERROR_CODE    = "errorCode"
	ERROR_MESSAGE = "errorMessage"
)

type Client struct {
	ApiToken string
	BaseUrl  string
	log      log.Logger
}

// Entry point into the logz.io client
func New(apiToken, baseUrl string) *Client {
	var c Client
	c.ApiToken = apiToken
	c.BaseUrl = baseUrl
	return &c
}

func GetHttpClient(req *http.Request) *http.Client {
	url, err := http.ProxyFromEnvironment(req)
	if url != nil && err == nil {
		tr := &http.Transport{
			Proxy:           http.ProxyURL(url),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		return &http.Client{Transport: tr}
	} else {
		return &http.Client{}
	}
}

func IsErrorResponse(response map[string]interface{}) (bool, string) {
	if _, ok := response[ERROR_CODE]; ok {
		return true, response[ERROR_CODE].(string)
	}
	if _, ok := response[ERROR_MESSAGE]; ok {
		return true, response[ERROR_MESSAGE].(string)
	}
	return false, ""
}
