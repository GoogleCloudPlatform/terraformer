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

package main

// metadata for generate code for GCP compute service
var terraformResources = map[string]gcpResourceRenderable{
	"addresses": basicGCPResource{
		terraformName: "google_compute_address",
	},
	"autoscalers": basicGCPResource{
		terraformName: "google_compute_autoscaler",
	},
	"backendBuckets": basicGCPResource{
		terraformName: "google_compute_backend_bucket",
	},
	"backendServices": backendServices{
		basicGCPResource{
			terraformName: "google_compute_backend_service",
		},
	},
	"disks": basicGCPResource{
		terraformName: "google_compute_disk",
	},
	"firewalls": basicGCPResource{
		terraformName: "google_compute_firewall",
	},
	"forwardingRules": basicGCPResource{
		terraformName: "google_compute_forwarding_rule",
	},
	"globalAddresses": basicGCPResource{
		terraformName: "google_compute_global_address",
	},
	"globalForwardingRules": globalForwardingRules{
		basicGCPResource{
			terraformName: "google_compute_global_forwarding_rule",
		},
	},
	"healthChecks": basicGCPResource{
		terraformName: "google_compute_health_check",
	},
	"httpHealthChecks": basicGCPResource{
		terraformName: "google_compute_http_health_check",
	},
	"httpsHealthChecks": basicGCPResource{
		terraformName: "google_compute_https_health_check",
	},
	"images": basicGCPResource{
		terraformName: "google_compute_image",
	},
	"instanceGroupManagers": instanceGroupManagers{
		basicGCPResource{
			terraformName:    "google_compute_instance_group_manager",
			allowEmptyValues: []string{"^version.[0-9].name", "^auto_healing_policies.[0-9].health_check"},
		},
	},
	"instanceGroups": basicGCPResource{
		terraformName: "google_compute_instance_group",
	},
	"instanceTemplates": basicGCPResource{
		terraformName: "google_compute_instance_template",
	},
	"instances": instances{
		basicGCPResource{
			terraformName:    "google_compute_instance",
			allowEmptyValues: []string{"labels."},
			additionalFieldsForRefresh: map[string]string{
				"disk.#": "0",
			},
		},
	},
	"networks": basicGCPResource{
		terraformName: "google_compute_network",
	},
	"regionAutoscalers": basicGCPResource{
		terraformName: "google_compute_region_autoscaler",
	},
	"regionBackendServices": basicGCPResource{
		terraformName: "google_compute_region_backend_service",
	},
	"regionDisks": basicGCPResource{
		terraformName: "google_compute_region_disk",
	},
	"regionInstanceGroupManagers": basicGCPResource{
		terraformName:    "google_compute_region_instance_group_manager",
		allowEmptyValues: []string{"name", "health_check"},
	},
	"routers": basicGCPResource{
		terraformName: "google_compute_router",
	},
	"routes": basicGCPResource{
		terraformName: "google_compute_route",
	},
	"securityPolicies": basicGCPResource{
		terraformName: "google_compute_security_policy",
	},
	/*"snapshots": {
		terraformName: "google_compute_snapshot",
		ignoreKeys: []string{
			"snapshot_encryption_key_sha256",
			"source_disk_encryption_key_sha256",
			"source_disk_link",
		},
	},*/
	/*"sslCertificates": {
		terraformName:       "google_compute_ssl_certificate",
		ignoreKeys: []string{"certificate_id"},
	},*/
	"sslPolicies": basicGCPResource{
		terraformName: "google_compute_ssl_policy",
	},
	"subnetworks": basicGCPResource{
		terraformName: "google_compute_subnetwork",
	},
	"targetHttpProxies": basicGCPResource{
		terraformName: "google_compute_target_http_proxy",
	},
	"targetHttpsProxies": basicGCPResource{
		terraformName: "google_compute_target_https_proxy",
	},
	"targetSslProxies": basicGCPResource{
		terraformName: "google_compute_target_ssl_proxy",
	},
	"targetTcpProxies": basicGCPResource{
		terraformName: "google_compute_target_tcp_proxy",
	},
	"urlMaps": basicGCPResource{
		terraformName: "google_compute_url_map",
	},
	"vpnTunnels": basicGCPResource{
		terraformName: "google_compute_vpn_tunnel",
	},
}
