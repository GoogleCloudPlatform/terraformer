package openstack

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/fwaas/rules"
)

// RuleCreateOpts represents the attributes used when creating a new firewall rule.
type RuleCreateOpts struct {
	rules.CreateOpts
	ValueSpecs map[string]string `json:"value_specs,omitempty"`
}

// ToRuleCreateMap casts a CreateOpts struct to a map.
// It overrides rules.ToRuleCreateMap to add the ValueSpecs field.
func (opts RuleCreateOpts) ToRuleCreateMap() (map[string]interface{}, error) {
	b, err := BuildRequest(opts, "firewall_rule")
	if err != nil {
		return nil, err
	}

	if m := b["firewall_rule"].(map[string]interface{}); m["protocol"] == "any" {
		m["protocol"] = nil
	}

	return b, nil
}

func expandFWRuleV1IPVersion(ipv int) gophercloud.IPVersion {
	// Determine the IP Version
	var ipVersion gophercloud.IPVersion
	switch ipv {
	case 4:
		ipVersion = gophercloud.IPv4
	case 6:
		ipVersion = gophercloud.IPv6
	}

	return ipVersion
}

func expandFWRuleV1Protocol(p string) rules.Protocol {
	var protocol rules.Protocol
	switch p {
	case "any":
		protocol = rules.ProtocolAny
	case "icmp":
		protocol = rules.ProtocolICMP
	case "tcp":
		protocol = rules.ProtocolTCP
	case "udp":
		protocol = rules.ProtocolUDP
	}

	return protocol
}
