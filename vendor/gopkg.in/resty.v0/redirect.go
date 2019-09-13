// Copyright (c) 2015 Jeevanandam M (jeeva@myjeeva.com), All rights reserved.
// resty source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package resty

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
)

// RedirectPolicy to regulate the redirects in the resty client.
// Objects implementing the RedirectPolicy interface can be registered as
//
// Apply function should return nil to continue the redirect jounery, otherwise
// return error to stop the redirect.
type RedirectPolicy interface {
	Apply(req *http.Request, via []*http.Request) error
}

// The RedirectPolicyFunc type is an adapter to allow the use of ordinary functions as RedirectPolicy.
// If f is a function with the appropriate signature, RedirectPolicyFunc(f) is a RedirectPolicy object that calls f.
type RedirectPolicyFunc func(*http.Request, []*http.Request) error

// Apply calls f(req, via).
func (f RedirectPolicyFunc) Apply(req *http.Request, via []*http.Request) error {
	return f(req, via)
}

// NoRedirectPolicy is used to disable redirects in the HTTP client
// 		resty.SetRedirectPolicy(NoRedirectPolicy())
func NoRedirectPolicy() RedirectPolicy {
	return RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
		return errors.New("Auto redirect is disabled")
	})
}

// FlexibleRedirectPolicy is convenient method to create No of redirect policy for HTTP client.
// 		resty.SetRedirectPolicy(FlexibleRedirectPolicy(20))
func FlexibleRedirectPolicy(noOfRedirect int) RedirectPolicy {
	return RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
		if len(via) >= noOfRedirect {
			return fmt.Errorf("Stopped after %d redirects", noOfRedirect)
		}
		return nil
	})
}

// DomainCheckRedirectPolicy is convenient method to define domain name redirect rule in resty client.
// Redirect is allowed for only mentioned host in the policy.
// 		resty.SetRedirectPolicy(DomainCheckRedirectPolicy("host1.com", "host2.org", "host3.net"))
func DomainCheckRedirectPolicy(hostnames ...string) RedirectPolicy {
	hosts := make(map[string]bool)
	for _, h := range hostnames {
		hosts[strings.ToLower(h)] = true
	}

	fn := RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
		hostname := ""
		if strings.Index(req.URL.Host, ":") > 0 {
			host, _, _ := net.SplitHostPort(req.URL.Host)
			hostname = strings.ToLower(host)
		} else {
			hostname = strings.ToLower(req.URL.Host)
		}

		if ok := hosts[hostname]; !ok {
			return errors.New("Redirect is not allowed as per DomainCheckRedirectPolicy")
		}

		return nil
	})

	return fn
}
