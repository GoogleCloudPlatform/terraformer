// +build !go1.6

package s3

import request "github.com/aws/aws-sdk-go-v2/aws"

func platformRequestHandlers(r *request.Request) {
}
