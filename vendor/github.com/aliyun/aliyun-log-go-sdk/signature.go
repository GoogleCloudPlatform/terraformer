package sls

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

// GMT location
var gmtLoc = time.FixedZone("GMT", 0)

// NowRFC1123 returns now time in RFC1123 format with GMT timezone,
// eg. "Mon, 02 Jan 2006 15:04:05 GMT".
func nowRFC1123() string {
	return time.Now().In(gmtLoc).Format(time.RFC1123)
}

// signature calculates a request's signature digest.
func signature(accessKeySecret, method, uri string,
	headers map[string]string) (digest string, err error) {
	var contentMD5, contentType, date, canoHeaders, canoResource string
	var slsHeaderKeys sort.StringSlice

	if val, ok := headers["Content-MD5"]; ok {
		contentMD5 = val
	}

	if val, ok := headers["Content-Type"]; ok {
		contentType = val
	}

	date, ok := headers["Date"]
	if !ok {
		err = fmt.Errorf("Can't find 'Date' header")
		return
	}

	// Calc CanonicalizedSLSHeaders
	slsHeaders := make(map[string]string, len(headers))
	for k, v := range headers {
		l := strings.TrimSpace(strings.ToLower(k))
		if strings.HasPrefix(l, "x-log-") || strings.HasPrefix(l, "x-acs-") {
			slsHeaders[l] = strings.TrimSpace(v)
			slsHeaderKeys = append(slsHeaderKeys, l)
		}
	}

	sort.Sort(slsHeaderKeys)
	for i, k := range slsHeaderKeys {
		canoHeaders += k + ":" + slsHeaders[k]
		if i+1 < len(slsHeaderKeys) {
			canoHeaders += "\n"
		}
	}

	// Calc CanonicalizedResource
	u, err := url.Parse(uri)
	if err != nil {
		return
	}

	canoResource += u.EscapedPath()
	if u.RawQuery != "" {
		var keys sort.StringSlice

		vals := u.Query()
		for k := range vals {
			keys = append(keys, k)
		}

		sort.Sort(keys)
		canoResource += "?"
		for i, k := range keys {
			if i > 0 {
				canoResource += "&"
			}

			for _, v := range vals[k] {
				canoResource += k + "=" + v
			}
		}
	}

	signStr := method + "\n" +
		contentMD5 + "\n" +
		contentType + "\n" +
		date + "\n" +
		canoHeaders + "\n" +
		canoResource

	// Signature = base64(hmac-sha1(UTF8-Encoding-Of(SignString)，AccessKeySecret))
	mac := hmac.New(sha1.New, []byte(accessKeySecret))
	_, err = mac.Write([]byte(signStr))
	if err != nil {
		return
	}
	digest = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return
}
