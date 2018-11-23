// AUTO-GENERATED CODE. DO NOT EDIT.
package computeTerrforming

import (
	"waze/terraform/gcp_terraforming/gcp_generator"
)

var ComputeService = map[string]gcp_generator.Generator{

	"addresses":                   AddressesGenerator{},
	"autoscalers":                 AutoscalersGenerator{},
	"backendBuckets":              BackendBucketsGenerator{},
	"backendServices":             BackendServicesGenerator{},
	"disks":                       DisksGenerator{},
	"firewalls":                   FirewallsGenerator{},
	"forwardingRules":             ForwardingRulesGenerator{},
	"globalAddresses":             GlobalAddressesGenerator{},
	"globalForwardingRules":       GlobalForwardingRulesGenerator{},
	"healthChecks":                HealthChecksGenerator{},
	"httpHealthChecks":            HttpHealthChecksGenerator{},
	"httpsHealthChecks":           HttpsHealthChecksGenerator{},
	"images":                      ImagesGenerator{},
	"instanceGroupManagers":       InstanceGroupManagersGenerator{},
	"instanceGroups":              InstanceGroupsGenerator{},
	"instanceTemplates":           InstanceTemplatesGenerator{},
	"instances":                   InstancesGenerator{},
	"networks":                    NetworksGenerator{},
	"regionAutoscalers":           RegionAutoscalersGenerator{},
	"regionBackendServices":       RegionBackendServicesGenerator{},
	"regionDisks":                 RegionDisksGenerator{},
	"regionInstanceGroupManagers": RegionInstanceGroupManagersGenerator{},
	"routers":                     RoutersGenerator{},
	"routes":                      RoutesGenerator{},
	"securityPolicies":            SecurityPoliciesGenerator{},
	"sslPolicies":                 SslPoliciesGenerator{},
	"subnetworks":                 SubnetworksGenerator{},
	"targetHttpProxies":           TargetHttpProxiesGenerator{},
	"targetHttpsProxies":          TargetHttpsProxiesGenerator{},
	"targetSslProxies":            TargetSslProxiesGenerator{},
	"targetTcpProxies":            TargetTcpProxiesGenerator{},
	"urlMaps":                     UrlMapsGenerator{},
	"vpnTunnels":                  VpnTunnelsGenerator{},
}
