// Package auth contains types and functions to manage authentication
// credentials for service hosts.
package auth

import (
	"fmt"
	"net/http"

	"github.com/zclconf/go-cty/cty"

	"github.com/hashicorp/terraform/svchost"
)

// Credentials is a list of CredentialsSource objects that can be tried in
// turn until one returns credentials for a host, or one returns an error.
//
// A Credentials is itself a CredentialsSource, wrapping its members.
// In principle one CredentialsSource can be nested inside another, though
// there is no good reason to do so.
//
// The write operations on a Credentials are tried only on the first object,
// under the assumption that it is the primary store.
type Credentials []CredentialsSource

// NoCredentials is an empty CredentialsSource that always returns nil
// when asked for credentials.
var NoCredentials CredentialsSource = Credentials{}

// A CredentialsSource is an object that may be able to provide credentials
// for a given host.
//
// Credentials lookups are not guaranteed to be concurrency-safe. Callers
// using these facilities in concurrent code must use external concurrency
// primitives to prevent race conditions.
type CredentialsSource interface {
	// ForHost returns a non-nil HostCredentials if the source has credentials
	// available for the host, and a nil HostCredentials if it does not.
	//
	// If an error is returned, progress through a list of CredentialsSources
	// is halted and the error is returned to the user.
	ForHost(host svchost.Hostname) (HostCredentials, error)

	// StoreForHost takes a HostCredentialsWritable and saves it as the
	// credentials for the given host.
	//
	// If credentials are already stored for the given host, it will try to
	// replace those credentials but may produce an error if such replacement
	// is not possible.
	StoreForHost(host svchost.Hostname, credentials HostCredentialsWritable) error

	// ForgetForHost discards any stored credentials for the given host. It
	// does nothing and returns successfully if no credentials are saved
	// for that host.
	ForgetForHost(host svchost.Hostname) error
}

// HostCredentials represents a single set of credentials for a particular
// host.
type HostCredentials interface {
	// PrepareRequest modifies the given request in-place to apply the
	// receiving credentials. The usual behavior of this method is to
	// add some sort of Authorization header to the request.
	PrepareRequest(req *http.Request)

	// Token returns the authentication token.
	Token() string
}

// HostCredentialsWritable is an extension of HostCredentials for credentials
// objects that can be serialized as a JSON-compatible object value for
// storage.
type HostCredentialsWritable interface {
	HostCredentials

	// ToStore returns a cty.Value, always of an object type,
	// representing data that can be serialized to represent this object
	// in persistent storage.
	//
	// The resulting value may uses only cty values that can be accepted
	// by the cty JSON encoder, though the caller may elect to instead store
	// it in some other format that has a JSON-compatible type system.
	ToStore() cty.Value
}

// ForHost iterates over the contained CredentialsSource objects and
// tries to obtain credentials for the given host from each one in turn.
//
// If any source returns either a non-nil HostCredentials or a non-nil error
// then this result is returned. Otherwise, the result is nil, nil.
func (c Credentials) ForHost(host svchost.Hostname) (HostCredentials, error) {
	for _, source := range c {
		creds, err := source.ForHost(host)
		if creds != nil || err != nil {
			return creds, err
		}
	}
	return nil, nil
}

// StoreForHost passes the given arguments to the same operation on the
// first CredentialsSource in the receiver.
func (c Credentials) StoreForHost(host svchost.Hostname, credentials HostCredentialsWritable) error {
	if len(c) == 0 {
		return fmt.Errorf("no credentials store is available")
	}

	return c[0].StoreForHost(host, credentials)
}

// ForgetForHost passes the given arguments to the same operation on the
// first CredentialsSource in the receiver.
func (c Credentials) ForgetForHost(host svchost.Hostname) error {
	if len(c) == 0 {
		return fmt.Errorf("no credentials store is available")
	}

	return c[0].ForgetForHost(host)
}
