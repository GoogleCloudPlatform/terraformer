package sls

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/golang/glog"
	"golang.org/x/net/context"
)

// timeout configs
var (
	defaultRequestTimeout = 10 * time.Second
	defaultRetryTimeout   = 30 * time.Second
	defaultHttpClient     = &http.Client{
		Timeout: defaultRequestTimeout,
	}
)

func retryReadErrorCheck(ctx context.Context, err error) (bool, error) {
	if err == nil {
		return false, nil
	}

	switch e := err.(type) {
	case *url.Error:
		return true, e
	case *Error:
		if e.HTTPCode >= 500 && e.HTTPCode <= 599 {
			return true, e
		}
	case *BadResponseError:
		if e.HTTPCode >= 500 && e.HTTPCode <= 599 {
			return true, e
		}
	default:
		return false, e
	}

	return false, err
}

func retryWriteErrorCheck(ctx context.Context, err error) (bool, error) {
	if err == nil {
		return false, nil
	}

	switch e := err.(type) {
	case *Error:
		if e.HTTPCode == 502 || e.HTTPCode == 503 {
			return true, e
		}
	case *BadResponseError:
		if e.HTTPCode == 502 || e.HTTPCode == 503 {
			return true, e
		}
	default:
		return false, e
	}

	return false, err
}

// request sends a request to SLS.
// mock param only for test, default is []
func request(project *LogProject, method, uri string, headers map[string]string,
	body []byte, mock ...interface{}) (*http.Response, error) {

	var r *http.Response
	var slsErr error
	var err error
	var mockErr *mockErrorRetry

	project.init()
	ctx, cancel := context.WithTimeout(context.Background(), project.retryTimeout)
	defer cancel()

	//fmt.Println("request ", project, method, uri, headers, body)
	// all GET method is read function
	if method == http.MethodGet {
		err = RetryWithCondition(ctx, backoff.NewExponentialBackOff(), func() (bool, error) {
			if len(mock) == 0 {
				//fmt.Println("real request", project, method, uri, headers, body)
				r, slsErr = realRequest(ctx, project, method, uri, headers, body)
				//fmt.Println("real request done")
			} else {
				r, mockErr = nil, mock[0].(*mockErrorRetry)
				mockErr.RetryCnt--
				if mockErr.RetryCnt <= 0 {
					r = &http.Response{}
					slsErr = nil
					return false, nil
				}
				slsErr = &mockErr.Err
			}
			return retryReadErrorCheck(ctx, slsErr)
		})
	} else {
		err = RetryWithCondition(ctx, backoff.NewExponentialBackOff(), func() (bool, error) {
			if len(mock) == 0 {
				r, slsErr = realRequest(ctx, project, method, uri, headers, body)
			} else {
				r, mockErr = nil, mock[0].(*mockErrorRetry)
				mockErr.RetryCnt--
				if mockErr.RetryCnt <= 0 {
					r = &http.Response{}
					slsErr = nil
					return false, nil
				}
				slsErr = &mockErr.Err
			}
			return retryWriteErrorCheck(ctx, slsErr)
		})
	}

	if err != nil {
		return r, err
	}
	return r, slsErr
}

// request sends a request to alibaba cloud Log Service.
// @note if error is nil, you must call http.Response.Body.Close() to finalize reader
func realRequest(ctx context.Context, project *LogProject, method, uri string, headers map[string]string,
	body []byte) (*http.Response, error) {

	// The caller should provide 'x-log-bodyrawsize' header
	if _, ok := headers["x-log-bodyrawsize"]; !ok {
		return nil, NewClientError(fmt.Errorf("Can't find 'x-log-bodyrawsize' header"))
	}

	// SLS public request headers
	// var hostStr string
	// if len(project.Name) == 0 {
	// 	hostStr = project.Endpoint
	// } else {
	// 	hostStr = project.Name + "." + project.Endpoint
	// }
	baseURL := project.getBaseURL()
	headers["Host"] = baseURL
	headers["Date"] = nowRFC1123()
	headers["x-log-apiversion"] = version
	headers["x-log-signaturemethod"] = signatureMethod
	if len(project.UserAgent) > 0 {
		headers["User-Agent"] = project.UserAgent
	} else {
		headers["User-Agent"] = defaultLogUserAgent
	}

	// Access with token
	if project.SecurityToken != "" {
		headers["x-acs-security-token"] = project.SecurityToken
	}

	if body != nil {
		bodyMD5 := fmt.Sprintf("%X", md5.Sum(body))
		headers["Content-MD5"] = bodyMD5
		if _, ok := headers["Content-Type"]; !ok {
			return nil, NewClientError(fmt.Errorf("Can't find 'Content-Type' header"))
		}
	}

	// Calc Authorization
	// Authorization = "SLS <AccessKeyId>:<Signature>"
	digest, err := signature(project.AccessKeySecret, method, uri, headers)
	if err != nil {
		return nil, NewClientError(err)
	}
	auth := fmt.Sprintf("SLS %v:%v", project.AccessKeyID, digest)
	headers["Authorization"] = auth

	// Initialize http request
	reader := bytes.NewReader(body)

	// Handle the endpoint
	urlStr := fmt.Sprintf("%s%s", baseURL, uri)
	req, err := http.NewRequest(method, urlStr, reader)
	if err != nil {
		return nil, NewClientError(err)
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
	resp, err := project.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Parse the sls error from body.
	if resp.StatusCode != http.StatusOK {
		err := &Error{}
		err.HTTPCode = (int32)(resp.StatusCode)
		defer resp.Body.Close()
		buf, ioErr := ioutil.ReadAll(resp.Body)
		if ioErr != nil {
			return nil, NewBadResponseError(ioErr.Error(), resp.Header, resp.StatusCode)
		}
		if jErr := json.Unmarshal(buf, err); jErr != nil {
			return nil, NewBadResponseError(string(buf), resp.Header, resp.StatusCode)
		}
		err.RequestID = resp.Header.Get(RequestIDHeader)
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
