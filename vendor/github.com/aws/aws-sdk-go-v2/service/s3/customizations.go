package s3

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

func init() {
	initClient = defaultInitClientFn
	initRequest = defaultInitRequestFn
}

func defaultInitClientFn(c *Client) {
	// Require SSL when using SSE keys
	c.Handlers.Validate.PushBack(validateSSERequiresSSL)
	c.Handlers.Build.PushBack(computeSSEKeys)

	// S3 uses custom error unmarshaling logic
	c.Handlers.UnmarshalError.Clear()
	c.Handlers.UnmarshalError.PushBack(unmarshalError)
}

func defaultInitRequestFn(c *Client, r *aws.Request) {
	// Add reuest handlers for specific platforms.
	// e.g. 100-continue support for PUT requests using Go 1.6
	platformRequestHandlers(c, r)

	// Support building custom endpoints based on config
	r.Handlers.Build.PushFront(buildUpdateEndpointForS3Config(c))

	switch r.Operation.Name {
	case opPutBucketCors, opPutBucketLifecycle, opPutBucketPolicy,
		opPutBucketTagging, opDeleteObjects, opPutBucketLifecycleConfiguration,
		opPutObjectLegalHold, opPutObjectRetention, opPutObjectLockConfiguration,
		opPutBucketReplication:
		// These S3 operations require Content-MD5 to be set
		r.Handlers.Build.PushBack(contentMD5)
	case opGetBucketLocation:
		// GetBucketLocation has custom parsing logic
		r.Handlers.Unmarshal.PushFront(buildGetBucketLocation)
	case opCreateBucket:
		// Auto-populate LocationConstraint with current region
		r.Handlers.Validate.PushFront(populateLocationConstraint)
	case opCopyObject, opUploadPartCopy, opCompleteMultipartUpload:
		r.Handlers.Unmarshal.PushFront(copyMultipartStatusOKUnmarhsalError)
	}
}

// bucketGetter is an accessor interface to grab the "Bucket" field from
// an S3 type.
type bucketGetter interface {
	getBucket() string
}

// sseCustomerKeyGetter is an accessor interface to grab the "SSECustomerKey"
// field from an S3 type.
type sseCustomerKeyGetter interface {
	getSSECustomerKey() string
}

// copySourceSSECustomerKeyGetter is an accessor interface to grab the
// "CopySourceSSECustomerKey" field from an S3 type.
type copySourceSSECustomerKeyGetter interface {
	getCopySourceSSECustomerKey() string
}
