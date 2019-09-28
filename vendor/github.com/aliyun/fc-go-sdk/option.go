package fc

import (
	"net/http"
	"time"

	"gopkg.in/resty.v1"
)

// ClientOption : defines client options type
type ClientOption func(*Client)

// WithTimeout : set request timeout in second
func WithTimeout(t uint) ClientOption {
	return func(c *Client) {
		c.Connect.Timeout = t
		resty.SetTimeout(time.Duration(t) * time.Second)
	}
}

// WithTransport : overrides default http.Transport with customized transport
func WithTransport(ts *http.Transport) ClientOption {
	return func(c *Client) {
		if ts != nil {
			resty.SetTransport(ts)
		}
	}
}

// WithSecurityToken : sets the STS security token
func WithSecurityToken(token string) ClientOption {
	return func(c *Client) { c.Config.SecurityToken = token }
}

// WithAccountID sets the account id in header, this enables accessing
// FC using IP address:
//
// client, _ := fc.NewClient("127.0.0.1", "api-version", "id", "key",
//	fc.WithAccountID("1234776887"))
func WithAccountID(aid string) ClientOption {
	return func(c *Client) { c.Config.AccountID = aid }
}

// WithRetryCount : config the retry count for resty
func WithRetryCount(count int) ClientOption {
	return func(c *Client) {
		resty.SetRetryCount(count)
	}
}
