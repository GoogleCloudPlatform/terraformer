package sls

// request sends a request to SLS.
import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/golang/glog"
)

// request sends a request to alibaba cloud Log Service.
// @note if error is nil, you must call http.Response.Body.Close() to finalize reader
func (c *Client) request(project, method, uri string, headers map[string]string, body []byte) (*http.Response, error) {
	// The caller should provide 'x-log-bodyrawsize' header
	if _, ok := headers["x-log-bodyrawsize"]; !ok {
		return nil, fmt.Errorf("Can't find 'x-log-bodyrawsize' header")
	}

	var endpoint string
	var usingHTTPS bool
	if strings.HasPrefix(c.Endpoint, "https://") {
		endpoint = c.Endpoint[8:]
		usingHTTPS = true
	} else if strings.HasPrefix(c.Endpoint, "http://") {
		endpoint = c.Endpoint[7:]
	} else {
		endpoint = c.Endpoint
	}

	// SLS public request headers
	var hostStr string
	if len(project) == 0 {
		hostStr = project
	} else {
		hostStr = project + "." + endpoint
	}
	headers["Host"] = hostStr
	headers["Date"] = nowRFC1123()
	headers["x-log-apiversion"] = version
	headers["x-log-signaturemethod"] = signatureMethod

	if len(c.UserAgent) > 0 {
		headers["User-Agent"] = c.UserAgent
	} else {
		headers["User-Agent"] = defaultLogUserAgent
	}

	c.accessKeyLock.RLock()
	stsToken := c.SecurityToken
	accessKeyID := c.AccessKeyID
	accessKeySecret := c.AccessKeySecret
	c.accessKeyLock.RUnlock()

	// Access with token
	if stsToken != "" {
		headers["x-acs-security-token"] = stsToken
	}

	if body != nil {
		bodyMD5 := fmt.Sprintf("%X", md5.Sum(body))
		headers["Content-MD5"] = bodyMD5
		if _, ok := headers["Content-Type"]; !ok {
			return nil, fmt.Errorf("Can't find 'Content-Type' header")
		}
	}

	// Calc Authorization
	// Authorization = "SLS <AccessKeyId>:<Signature>"
	digest, err := signature(accessKeySecret, method, uri, headers)
	if err != nil {
		return nil, err
	}
	auth := fmt.Sprintf("SLS %v:%v", accessKeyID, digest)
	headers["Authorization"] = auth

	// Initialize http request
	reader := bytes.NewReader(body)
	var urlStr string
	// using http as default
	if !GlobalForceUsingHTTP && usingHTTPS {
		urlStr = "https://"
	} else {
		urlStr = "http://"
	}
	urlStr += hostStr + uri
	req, err := http.NewRequest(method, urlStr, reader)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if glog.V(5) {
		dump, e := httputil.DumpRequest(req, true)
		if e != nil {
			glog.Info(e)
		}
		glog.Infof("HTTP Request:\n%v", string(dump))
	}

	// Get ready to do request
	resp, err := defaultHttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Parse the sls error from body.
	if resp.StatusCode != http.StatusOK {
		err := &Error{}
		err.HTTPCode = (int32)(resp.StatusCode)
		defer resp.Body.Close()
		buf, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(buf, err)
		err.RequestID = resp.Header.Get("x-log-requestid")
		return nil, err
	}

	if glog.V(5) {
		dump, e := httputil.DumpResponse(resp, true)
		if e != nil {
			glog.Info(e)
		}
		glog.Infof("HTTP Response:\n%v", string(dump))
	}
	return resp, nil
}
