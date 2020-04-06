package dynamodb

import (
	"bytes"
	"hash/crc32"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
)

func init() {
	initClient = func(c *Client) {
		if c.Config.Retryer == nil {
			// Only override the retryer with a custom one if the config
			// does not already contain a retryer
			setCustomRetryer(c)
		}

		c.Handlers.Build.PushBackNamed(disableCompressionHandler)
		c.Handlers.Unmarshal.PushFrontNamed(validateCRC32Handler)
	}

	initRequest = func(c *Client, req *aws.Request) {
		if c.DisableComputeChecksums {
			// Checksum validation is off, remove the validator.
			req.Handlers.Unmarshal.Remove(validateCRC32Handler)
		}
	}
}

func setCustomRetryer(c *Client) {
	c.Retryer = retry.AddWithMaxAttempts(c.Retryer, 10)
}

func drainBody(b io.ReadCloser, length int64) (out *bytes.Buffer, err error) {
	if length < 0 {
		length = 0
	}
	buf := bytes.NewBuffer(make([]byte, 0, length))

	if _, err = buf.ReadFrom(b); err != nil {
		return nil, err
	}
	if err = b.Close(); err != nil {
		return nil, err
	}
	return buf, nil
}

var disableCompressionHandler = aws.NamedHandler{
	Name: "dynamodb.DisableCompression",
	Fn:   disableCompression,
}

func disableCompression(r *aws.Request) {
	r.HTTPRequest.Header.Set("Accept-Encoding", "identity")
}

var validateCRC32Handler = aws.NamedHandler{
	Name: "dynamodb.ValidateCRC32",
	Fn:   validateCRC32,
}

func validateCRC32(r *aws.Request) {
	if r.Error != nil {
		return // already have an error, no need to verify CRC
	}

	// Try to get CRC from response
	header := r.HTTPResponse.Header.Get("X-Amz-Crc32")
	if header == "" {
		return // No CRC32 header, skip
	}

	expected, err := strconv.ParseUint(header, 10, 32)
	if err != nil {
		return // Could not determine CRC value, skip
	}

	// TODO this drain body should use a multi-writer to write to the buffer and
	// hash at same time. Remove the need for iterating through the bytes a
	// second time to compute the hash separately.
	buf, err := drainBody(r.HTTPResponse.Body, r.HTTPResponse.ContentLength)
	if err != nil { // failed to read the response body, skip CRC32 validation.
		return
	}

	// Reset body for subsequent reads
	r.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(buf.Bytes()))

	// Compute the CRC checksum
	crc := crc32.ChecksumIEEE(buf.Bytes())

	if crc != uint32(expected) {
		// CRC does not match, set a retryable error
		r.Error = &CRC32CheckFailedError{}
	}
}

// CRC32CheckFailedError provides the error type for when a DynamoDB operation
// response's doesn't match the precomputed CRC32 value supplied by the
// service's API.
type CRC32CheckFailedError struct{}

// RetryableError signals that the error should be retried.
func (*CRC32CheckFailedError) RetryableError() bool {
	return true
}
func (*CRC32CheckFailedError) Error() string {
	return "integrity check failed for CRC32 validation"
}
