package fc

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"net/url"
)

//MD5 :Encoding MD5
func MD5(b []byte) string {
	ctx := md5.New()
	ctx.Write(b)
	return hex.EncodeToString(ctx.Sum(nil))
}

// HasPrefix check endpoint prefix
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

// GetAccessPoint get correct endpoint and host
func GetAccessPoint(endpointInput string) (endpoint, host string) {
	httpPrefix := "http://"
	httpsPrefix := "https://"
	if HasPrefix(endpointInput, httpPrefix) {
		host = endpointInput[len(httpPrefix):]
		return endpointInput, host
	} else if HasPrefix(endpointInput, httpsPrefix) {
		host = endpointInput[len(httpsPrefix):]
		return endpointInput, host
	}
	return httpPrefix + endpointInput, endpointInput
}

// IsBlank :check string pointer is nil or empty
func IsBlank(s *string) bool {
	if s == nil {
		return true
	}
	if len(*s) == 0 {
		return true
	}
	return false
}

// GetRequestID from headers
func GetRequestID(header http.Header) string {
	return header.Get(HTTPHeaderRequestID)
}

// GetErrorType get error type when call invocation
func GetErrorType(header http.Header) string {
	return header.Get(HTTPHeaderFCErrorType)
}

// GetEtag get resource etag
func GetEtag(header http.Header) string {
	return header.Get(HTTPHeaderEtag)
}

func pathEscape(s string) string {
	return url.PathEscape(s)
}
