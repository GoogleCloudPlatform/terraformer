package kinesis

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
)

var readDuration = 5 * time.Second

func init() {
	initClient = func(c *Client) {
		// Service specific error codes.
		c.Retryer = retry.AddWithErrorCodes(c.Retryer, ErrCodeLimitExceededException)
	}

	initRequest = func(c *Client, r *aws.Request) {
		if r.Operation.Name == opGetRecords {
			r.ApplyOptions(aws.WithResponseReadTimeout(readDuration))
		}
	}
}
