package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
)

type MachinePolicyService struct {
	sling *sling.Sling
}

func NewMachinePolicyService(sling *sling.Sling) *MachinePolicyService {
	return &MachinePolicyService{
		sling: sling,
	}
}

type MachinePolicies struct {
	Items []MachinePolicy `json:"Items"`
	PagedResults
}

type MachinePolicy struct {
	ID                        string                   `json:"Id"`
	Name                      string                   `json:"Name"`
	Description               string                   `json:"Description"`
	IsDefault                 bool                     `json:"IsDefault"`
	MachineHealthCheckPolicy  MachineHealthCheckPolicy `json:"MachineHealthCheckPolicy"`
	MachineConnectivityPolicy map[string]string        `json:"MachineConnectivityPolicy"`
	MachineCleanupPolicy      map[string]string        `json:"MachineCleanupPolicy"`
	MachineUpdatePolicymap    map[string]string        `json:"MachineUpdatePolicymap"`
	LastModifiedOn            *string                  `json:"LastModifiedOn,omitempty"`
	LastModifiedBy            *string                  `json:"LastModifiedBy,omitempty"`
}

type MachineHealthCheckPolicy struct {
	TentacleEndpointHealthCheckPolicy map[string]string `json:"TentacleEndpointHealthCheckPolicy"`
	SSHEndpointHealthCheckPolicy      map[string]string `json:"SshEndpointHealthCheckPolicy"`
	HealthCheckInterval               string            `json:"HealthCheckInterval"`
}

// Get returns a single machine with a given MachineID
func (s *MachinePolicyService) Get(MachinePolicyID string) (*MachinePolicy, error) {
	path := fmt.Sprintf("machinepolicies/%s", MachinePolicyID)
	resp, err := apiGet(s.sling, new(Machine), path)

	if err != nil {
		return nil, err
	}

	return resp.(*MachinePolicy), nil
}

// GetAll returns all registered machines
func (s *MachinePolicyService) GetAll() (*[]MachinePolicy, error) {
	var p []MachinePolicy
	path := "machinepolicies"
	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(MachinePolicies), path)
		if err != nil {
			return nil, err
		}

		r := resp.(*MachinePolicies)
		for _, item := range r.Items {
			p = append(p, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}
	return &p, nil
}
