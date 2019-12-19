package s3

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"

	request "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
)

func copyMultipartStatusOKUnmarhsalError(r *request.Request) {
	b, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		r.Error = awserr.New("SerializationError", "unable to read response body", err)
		return
	}
	body := bytes.NewReader(b)
	r.HTTPResponse.Body = ioutil.NopCloser(body)
	defer body.Seek(0, io.SeekStart)

	if body.Len() == 0 {
		// If there is no body don't attempt to parse the body.
		return
	}

	unmarshalError(r)
	if err, ok := r.Error.(awserr.Error); ok && err != nil {
		if err.Code() == "SerializationError" {
			r.Error = nil
			return
		}
		r.HTTPResponse.StatusCode = http.StatusServiceUnavailable
	}
}
