package fc

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"io"
	"sort"
	"strings"
)

type headers struct {
	Keys []string
	Vals []string
}

// GetAuthStr get signature strings
func GetAuthStr(accessKeyID string, accessKeySecret string, method string, header map[string]string, resource string) string {
	return "FC " + accessKeyID + ":" + GetSignature(accessKeySecret, method, header, resource)
}

// GetSignResourceWithQueries get signature resource with queries
func GetSignResourceWithQueries(path string, queries map[string][]string) string {
	paramsList := []string{}
	for key, values := range queries {
		if len(values) == 0 {
			paramsList = append(paramsList, key)
			continue
		}
		for _, v := range values {
			paramsList = append(paramsList, fmt.Sprintf("%s=%s", key, v))
		}
	}
	sort.Strings(paramsList)
	resource := path + "\n" + strings.Join(paramsList, "\n")
	return resource
}

// GetSignature calculate user's signature
func GetSignature(key string, method string, req map[string]string, fcResource string) string {
	header := &headers{}
	for k, v := range req {
		if strings.HasPrefix(strings.ToLower(k), HTTPHeaderPrefix) {
			header.Keys = append(header.Keys, strings.ToLower(k))
			header.Vals = append(header.Vals, v)
		}
	}
	sort.Sort(header)

	fcHeaders := ""
	for i := range header.Keys {
		fcHeaders += header.Keys[i] + ":" + header.Vals[i] + "\n"
	}

	signStr := method + "\n" + req[HTTPHeaderContentMD5] + "\n" + req[HTTPHeaderContentType] + "\n" + req[HTTPHeaderDate] + "\n" + fcHeaders + fcResource

	h := hmac.New(func() hash.Hash { return sha256.New() }, []byte(key))
	io.WriteString(h, signStr)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signedStr
}

func (h *headers) Len() int {
	return len(h.Vals)
}

func (h *headers) Less(i, j int) bool {
	return bytes.Compare([]byte(h.Keys[i]), []byte(h.Keys[j])) < 0
}

func (h *headers) Swap(i, j int) {
	h.Vals[i], h.Vals[j] = h.Vals[j], h.Vals[i]
	h.Keys[i], h.Keys[j] = h.Keys[j], h.Keys[i]
}
