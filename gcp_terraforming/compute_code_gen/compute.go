
package compute_code_gen

import (
	
	"waze/terraform/gcp_terraforming/compute_code_gen/addresses"
	"waze/terraform/gcp_terraforming/compute_code_gen/autoscalers"
	"waze/terraform/gcp_terraforming/compute_code_gen/backendBuckets"
	"waze/terraform/gcp_terraforming/compute_code_gen/backendServices"
	"waze/terraform/gcp_terraforming/compute_code_gen/disks"
	"waze/terraform/gcp_terraforming/compute_code_gen/firewalls"
	"waze/terraform/gcp_terraforming/compute_code_gen/forwardingRules"
	"waze/terraform/gcp_terraforming/compute_code_gen/globalAddresses"
	"waze/terraform/gcp_terraforming/compute_code_gen/globalForwardingRules"
	"waze/terraform/gcp_terraforming/compute_code_gen/healthChecks"
	"waze/terraform/gcp_terraforming/compute_code_gen/httpHealthChecks"
	"waze/terraform/gcp_terraforming/compute_code_gen/httpsHealthChecks"
	"waze/terraform/gcp_terraforming/compute_code_gen/images"
	"waze/terraform/gcp_terraforming/compute_code_gen/instanceGroupManagers"
	"waze/terraform/gcp_terraforming/compute_code_gen/instanceGroups"
	"waze/terraform/gcp_terraforming/compute_code_gen/instanceTemplates"
	"waze/terraform/gcp_terraforming/compute_code_gen/instances"
	"waze/terraform/gcp_terraforming/compute_code_gen/networks"
	"waze/terraform/gcp_terraforming/compute_code_gen/regionAutoscalers"
	"waze/terraform/gcp_terraforming/compute_code_gen/regionBackendServices"
	"waze/terraform/gcp_terraforming/compute_code_gen/regionDisks"
	"waze/terraform/gcp_terraforming/compute_code_gen/regionInstanceGroupManagers"
	"waze/terraform/gcp_terraforming/compute_code_gen/routers"
	"waze/terraform/gcp_terraforming/compute_code_gen/routes"
	"waze/terraform/gcp_terraforming/compute_code_gen/securityPolicies"
	"waze/terraform/gcp_terraforming/compute_code_gen/snapshots"
	"waze/terraform/gcp_terraforming/compute_code_gen/sslCertificates"
	"waze/terraform/gcp_terraforming/compute_code_gen/sslPolicies"
	"waze/terraform/gcp_terraforming/compute_code_gen/subnetworks"
	"waze/terraform/gcp_terraforming/compute_code_gen/targetHttpProxies"
	"waze/terraform/gcp_terraforming/compute_code_gen/targetHttpsProxies"
	"waze/terraform/gcp_terraforming/compute_code_gen/targetSslProxies"
	"waze/terraform/gcp_terraforming/compute_code_gen/targetTcpProxies"
	"waze/terraform/gcp_terraforming/compute_code_gen/urlMaps"
	"waze/terraform/gcp_terraforming/compute_code_gen/vpnTunnels"
	"waze/terraform/gcp_terraforming/gcp_generator"
)

var ComputeService = map[string]gcp_generator.Generator{

	"addresses":                   addresses.AddressesGenerator{},
	"autoscalers":                   autoscalers.AutoscalersGenerator{},
	"backendBuckets":                   backendBuckets.BackendBucketsGenerator{},
	"backendServices":                   backendServices.BackendServicesGenerator{},
	"disks":                   disks.DisksGenerator{},
	"firewalls":                   firewalls.FirewallsGenerator{},
	"forwardingRules":                   forwardingRules.ForwardingRulesGenerator{},
	"globalAddresses":                   globalAddresses.GlobalAddressesGenerator{},
	"globalForwardingRules":                   globalForwardingRules.GlobalForwardingRulesGenerator{},
	"healthChecks":                   healthChecks.HealthChecksGenerator{},
	"httpHealthChecks":                   httpHealthChecks.HttpHealthChecksGenerator{},
	"httpsHealthChecks":                   httpsHealthChecks.HttpsHealthChecksGenerator{},
	"images":                   images.ImagesGenerator{},
	"instanceGroupManagers":                   instanceGroupManagers.InstanceGroupManagersGenerator{},
	"instanceGroups":                   instanceGroups.InstanceGroupsGenerator{},
	"instanceTemplates":                   instanceTemplates.InstanceTemplatesGenerator{},
	"instances":                   instances.InstancesGenerator{},
	"networks":                   networks.NetworksGenerator{},
	"regionAutoscalers":                   regionAutoscalers.RegionAutoscalersGenerator{},
	"regionBackendServices":                   regionBackendServices.RegionBackendServicesGenerator{},
	"regionDisks":                   regionDisks.RegionDisksGenerator{},
	"regionInstanceGroupManagers":                   regionInstanceGroupManagers.RegionInstanceGroupManagersGenerator{},
	"routers":                   routers.RoutersGenerator{},
	"routes":                   routes.RoutesGenerator{},
	"securityPolicies":                   securityPolicies.SecurityPoliciesGenerator{},
	"snapshots":                   snapshots.SnapshotsGenerator{},
	"sslCertificates":                   sslCertificates.SslCertificatesGenerator{},
	"sslPolicies":                   sslPolicies.SslPoliciesGenerator{},
	"subnetworks":                   subnetworks.SubnetworksGenerator{},
	"targetHttpProxies":                   targetHttpProxies.TargetHttpProxiesGenerator{},
	"targetHttpsProxies":                   targetHttpsProxies.TargetHttpsProxiesGenerator{},
	"targetSslProxies":                   targetSslProxies.TargetSslProxiesGenerator{},
	"targetTcpProxies":                   targetTcpProxies.TargetTcpProxiesGenerator{},
	"urlMaps":                   urlMaps.UrlMapsGenerator{},
	"vpnTunnels":                   vpnTunnels.VpnTunnelsGenerator{},

}

