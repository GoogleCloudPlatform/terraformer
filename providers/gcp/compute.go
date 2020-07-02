// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.
package gcp

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

// Map of supported GCP compute service with code generate
var ComputeServices = map[string]terraformutils.ServiceGenerator{

	"addresses":                   &AddressesGenerator{},
	"autoscalers":                 &AutoscalersGenerator{},
	"backendBuckets":              &BackendBucketsGenerator{},
	"backendServices":             &BackendServicesGenerator{},
	"disks":                       &DisksGenerator{},
	"externalVpnGateways":         &ExternalVpnGatewaysGenerator{},
	"firewall":                    &FirewallGenerator{},
	"forwardingRules":             &ForwardingRulesGenerator{},
	"globalAddresses":             &GlobalAddressesGenerator{},
	"globalForwardingRules":       &GlobalForwardingRulesGenerator{},
	"healthChecks":                &HealthChecksGenerator{},
	"httpHealthChecks":            &HttpHealthChecksGenerator{},
	"httpsHealthChecks":           &HttpsHealthChecksGenerator{},
	"images":                      &ImagesGenerator{},
	"instanceGroupManagers":       &InstanceGroupManagersGenerator{},
	"instanceGroups":              &InstanceGroupsGenerator{},
	"instanceTemplates":           &InstanceTemplatesGenerator{},
	"interconnectAttachments":     &InterconnectAttachmentsGenerator{},
	"networkEndpointGroups":       &NetworkEndpointGroupsGenerator{},
	"networks":                    &NetworksGenerator{},
	"nodeGroups":                  &NodeGroupsGenerator{},
	"nodeTemplates":               &NodeTemplatesGenerator{},
	"packetMirrorings":            &PacketMirroringsGenerator{},
	"regionAutoscalers":           &RegionAutoscalersGenerator{},
	"regionBackendServices":       &RegionBackendServicesGenerator{},
	"regionDisks":                 &RegionDisksGenerator{},
	"regionHealthChecks":          &RegionHealthChecksGenerator{},
	"regionInstanceGroupManagers": &RegionInstanceGroupManagersGenerator{},
	"regionInstanceGroups":        &RegionInstanceGroupsGenerator{},
	"regionSslCertificates":       &RegionSslCertificatesGenerator{},
	"regionTargetHttpProxies":     &RegionTargetHttpProxiesGenerator{},
	"regionTargetHttpsProxies":    &RegionTargetHttpsProxiesGenerator{},
	"regionUrlMaps":               &RegionUrlMapsGenerator{},
	"reservations":                &ReservationsGenerator{},
	"resourcePolicies":            &ResourcePoliciesGenerator{},
	"routers":                     &RoutersGenerator{},
	"routes":                      &RoutesGenerator{},
	"securityPolicies":            &SecurityPoliciesGenerator{},
	"sslCertificates":             &SslCertificatesGenerator{},
	"sslPolicies":                 &SslPoliciesGenerator{},
	"subnetworks":                 &SubnetworksGenerator{},
	"targetHttpProxies":           &TargetHttpProxiesGenerator{},
	"targetHttpsProxies":          &TargetHttpsProxiesGenerator{},
	"targetInstances":             &TargetInstancesGenerator{},
	"targetPools":                 &TargetPoolsGenerator{},
	"targetSslProxies":            &TargetSslProxiesGenerator{},
	"targetTcpProxies":            &TargetTcpProxiesGenerator{},
	"targetVpnGateways":           &TargetVpnGatewaysGenerator{},
	"urlMaps":                     &UrlMapsGenerator{},
	"vpnTunnels":                  &VpnTunnelsGenerator{},
}
