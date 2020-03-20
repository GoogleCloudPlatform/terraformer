package apigateway

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

func init() {
	initClient = func(c *Client) {
		c.Handlers.Build.PushBack(func(r *aws.Request) {
			r.HTTPRequest.Header.Add("Accept", "application/json")
		})
	}
}
