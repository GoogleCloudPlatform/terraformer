package kinesis

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

var readDuration = 5 * time.Second

func init() {
	initRequest = func(c *Client, r *aws.Request) {
		if r.Operation.Name == opGetRecords {
			r.ApplyOptions(aws.WithResponseReadTimeout(readDuration))
		}

		// Service specific error codes.
		r.RetryErrorCodes = append(r.RetryErrorCodes, ErrCodeLimitExceededException)
	}
}
