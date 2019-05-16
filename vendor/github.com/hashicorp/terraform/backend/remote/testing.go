package remote

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	tfe "github.com/hashicorp/go-tfe"
	"github.com/hashicorp/terraform/backend"
	"github.com/hashicorp/terraform/state/remote"
	"github.com/hashicorp/terraform/svchost"
	"github.com/hashicorp/terraform/svchost/auth"
	"github.com/hashicorp/terraform/svchost/disco"
	"github.com/mitchellh/cli"

	backendLocal "github.com/hashicorp/terraform/backend/local"
)

const (
	testCred = "test-auth-token"
)

var (
	tfeHost  = svchost.Hostname(defaultHostname)
	credsSrc = auth.StaticCredentialsSource(map[svchost.Hostname]map[string]interface{}{
		tfeHost: {"token": testCred},
	})
)

func testInput(t *testing.T, answers map[string]string) *mockInput {
	return &mockInput{answers: answers}
}

func testBackendDefault(t *testing.T) (*Remote, func()) {
	c := map[string]interface{}{
		"organization": "hashicorp",
		"workspaces": []interface{}{
			map[string]interface{}{
				"name": "prod",
			},
		},
	}
	return testBackend(t, c)
}

func testBackendNoDefault(t *testing.T) (*Remote, func()) {
	c := map[string]interface{}{
		"organization": "hashicorp",
		"workspaces": []interface{}{
			map[string]interface{}{
				"prefix": "my-app-",
			},
		},
	}
	return testBackend(t, c)
}

func testBackendNoOperations(t *testing.T) (*Remote, func()) {
	c := map[string]interface{}{
		"organization": "no-operations",
		"workspaces": []interface{}{
			map[string]interface{}{
				"name": "prod",
			},
		},
	}
	return testBackend(t, c)
}

func testRemoteClient(t *testing.T) remote.Client {
	b, bCleanup := testBackendDefault(t)
	defer bCleanup()

	raw, err := b.State(backend.DefaultStateName)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	return raw.(*remote.State).Client
}

func testBackend(t *testing.T, c map[string]interface{}) (*Remote, func()) {
	s := testServer(t)
	b := New(testDisco(s))

	// Configure the backend so the client is created.
	backend.TestBackendConfig(t, b, c)

	// Get a new mock client.
	mc := newMockClient()

	// Replace the services we use with our mock services.
	b.CLI = cli.NewMockUi()
	b.client.Applies = mc.Applies
	b.client.ConfigurationVersions = mc.ConfigurationVersions
	b.client.Organizations = mc.Organizations
	b.client.Plans = mc.Plans
	b.client.PolicyChecks = mc.PolicyChecks
	b.client.Runs = mc.Runs
	b.client.StateVersions = mc.StateVersions
	b.client.Workspaces = mc.Workspaces

	// Set local to a local test backend.
	b.local = testLocalBackend(t, b)

	ctx := context.Background()

	// Create the organization.
	_, err := b.client.Organizations.Create(ctx, tfe.OrganizationCreateOptions{
		Name: tfe.String(b.organization),
	})
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	// Create the default workspace if required.
	if b.workspace != "" {
		_, err = b.client.Workspaces.Create(ctx, b.organization, tfe.WorkspaceCreateOptions{
			Name: tfe.String(b.workspace),
		})
		if err != nil {
			t.Fatalf("error: %v", err)
		}
	}

	return b, s.Close
}

func testLocalBackend(t *testing.T, remote *Remote) backend.Enhanced {
	b := backendLocal.NewWithBackend(remote)
	b.CLI = remote.CLI

	// Add a test provider to the local backend.
	backendLocal.TestLocalProvider(t, b, "null")

	return b
}

// testServer returns a *httptest.Server used for local testing.
func testServer(t *testing.T) *httptest.Server {
	mux := http.NewServeMux()

	// Respond to service discovery calls.
	mux.HandleFunc("/well-known/terraform.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
  "tfe.v2.1": "/api/v2/",
  "versions.v1": "/v1/versions/"
}`)
	})

	// Respond to service version constraints calls.
	mux.HandleFunc("/v1/versions/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
  "service": "tfe.v2.1",
  "product": "terraform",
  "minimum": "0.1.0",
  "maximum": "10.0.0"
}`)
	})

	// Respond to the initial query to read the hashicorp org entitlements.
	mux.HandleFunc("/api/v2/organizations/hashicorp/entitlement-set", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/vnd.api+json")
		io.WriteString(w, `{
  "data": {
    "id": "org-GExadygjSbKP8hsY",
    "type": "entitlement-sets",
    "attributes": {
      "operations": true,
      "private-module-registry": true,
      "sentinel": true,
      "state-storage": true,
      "teams": true,
      "vcs-integrations": true
    }
  }
}`)
	})

	// Respond to the initial query to read the no-operations org entitlements.
	mux.HandleFunc("/api/v2/organizations/no-operations/entitlement-set", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/vnd.api+json")
		io.WriteString(w, `{
  "data": {
    "id": "org-ufxa3y8jSbKP8hsT",
    "type": "entitlement-sets",
    "attributes": {
      "operations": false,
      "private-module-registry": true,
      "sentinel": true,
      "state-storage": true,
      "teams": true,
      "vcs-integrations": true
    }
  }
}`)
	})

	// All tests that are assumed to pass will use the hashicorp organization,
	// so for all other organization requests we will return a 404.
	mux.HandleFunc("/api/v2/organizations/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		io.WriteString(w, `{
  "errors": [
    {
      "status": "404",
      "title": "not found"
    }
  ]
}`)
	})

	return httptest.NewServer(mux)
}

// testDisco returns a *disco.Disco mapping app.terraform.io and
// localhost to a local test server.
func testDisco(s *httptest.Server) *disco.Disco {
	services := map[string]interface{}{
		"tfe.v2.1":    fmt.Sprintf("%s/api/v2/", s.URL),
		"versions.v1": fmt.Sprintf("%s/v1/versions/", s.URL),
	}
	d := disco.NewWithCredentialsSource(credsSrc)

	d.ForceHostServices(svchost.Hostname(defaultHostname), services)
	d.ForceHostServices(svchost.Hostname("localhost"), services)
	return d
}
