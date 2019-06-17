# Terraformer

[![Build Status](https://travis-ci.com/GoogleCloudPlatform/terraformer.svg?branch=master)](https://travis-ci.com/GoogleCloudPlatform/terraformer)
[![Go Report Card](https://goreportcard.com/badge/github.com/GoogleCloudPlatform/terraformer)](https://goreportcard.com/report/github.com/GoogleCloudPlatform/terraformer)

CLI tool to generate `tf` and `tfstate` files from existing infrastructure
(reverse Terraform).

*   Disclaimer: This is not an official Google product.
*   Status: beta - need improve documentations, bugs etc..
*   Created by: Waze SRE.

![Waze SRE logo](docs/waze-sre-logo.png)


# Table of Contents
- [Capabilities](#capabilities)
- [Installation](#installation)
- [Supported providers](/providers)
    * [Google Cloud](#use-with-gcp)
    * [AWS](#use-with-aws)
    * [OpenStack](#use-with-openstack)
    * [Kubernetes](#use-with-kubernetes)
    * [Github](#use-with-github)
    * [Datadog](#use-with-datadog)
- [Contributing](#contributing)
- [Developing](#developing)
- [Infrastructure](#infrastructure)

## Capabilities

1.  Generate `tf` + `tfstate` files from existing infrastructure for all
    supported objects by resource.
2.  Remote state can be uploaded to a GCS bucket.
3.  Connect between resources with `terraform_remote_state` (local and bucket).
4.  Compatible with terraform 0.12 syntax.
5.  Save `tf` files with custom folder tree pattern.
6.  Import by resource name and type.

Terraformer use terraform providers and built for easy to add new supported resources.
For upgrade resources with new fields you need upgrade only terraform providers.
```
Import current State to terraform configuration from google cloud

Usage:
   import google [flags]
   import google [command]

Available Commands:
  list        List supported resources for google provider

Flags:
  -b, --bucket string         gs://terraform-state
  -c, --connect                (default true)
  -f, --filter strings        google_compute_firewall=id1:id2:id4
  -h, --help                  help for google
  -o, --path-output string     (default "generated")
  -p, --path-pattern string   {output}/{provider}/custom/{service}/ (default "{output}/{provider}/{service}/")
      --projects strings
  -r, --resources strings     firewalls,networks
  -s, --state string          local or bucket (default "local")
  -z, --zone string
```
#### Permissions

Readonly permissions

#### Filtering

Filters are a way to choose which resources `terraformer` import.

Ex:
```
terraformer import aws --resources=vpc,subnet --filter=aws_vpc=myvpcid --regions=eu-west-1
```
will import only one VPC and not only subnets from this VPC but all subnets from all VPCs

##### Resources ID

Filtering is based on Terraform resource ID pattern. This way, it may differ from the value your provider give you.
Check the import part of [Terraform documentation][terraform-providers] for your resource for valid ID pattern.

[terraform-providers]: https://www.terraform.io/docs/providers/

#### Planning

`plan` command let you review the resources to import, rename them, or even filter only what you need.

The rest of subcommands and parameters are same to `import` command.

```
$ terraformer plan google --resources=networks,firewalls --projects=my-project --zone=europe-west1-d
(snip)

Saving planfile to generated/google/my-project/terraformer/plan.json
```

After reviewing/customizing the planfile, perform import by `import plan` command.

```
$ terraformer import plan generated/google/my-project/terraformer/plan.json
```

### Installation
From source:
1.  Run `git clone <terraformer repo>`
2.  Run `GO111MODULE=on go mod vendor`
3.  Run `go build -v`
4.  Run ```terraform init``` against an ```init.tf``` installing the platform correct required plugin(s) relative to current directory. Example ```init.tf``` to install the Google cloud plugin:
```
provider "google" {}
```
Or alternatively

4.  Copy your Terraform provider's plugin(s) to folder
    `~/.terraform.d/plugins/{darwin,linux}_amd64/`, as appropriate.

From Releases:

* Linux
```
curl -LO https://github.com/GoogleCloudPlatform/terraformer/releases/download/$(curl -s https://api.github.com/repos/GoogleCloudPlatform/terraformer/releases/latest | grep tag_name | cut -d '"' -f 4)/terraformer-linux-amd64
chmod +x terraformer-linux-amd64
sudo mv terraformer-linux-amd64 /usr/local/bin/terraformer
```
* MacOS
```
curl -LO https://github.com/GoogleCloudPlatform/terraformer/releases/download/$(curl -s https://api.github.com/repos/GoogleCloudPlatform/terraformer/releases/latest | grep tag_name | cut -d '"' -f 4)/terraformer-darwin-amd64
chmod +x terraformer-darwin-amd64
sudo mv terraformer-darwin-amd64 /usr/local/bin/terraformer
```

#### Using package manager

If you want to use a package manager:

- [Homebrew](https://brew.sh/) users can use `brew install terraformer`.

Links for download terraform providers:
* google cloud provider >2.0.0 - [here](https://releases.hashicorp.com/terraform-provider-google/)
* aws provider >1.56.0 - [here](https://releases.hashicorp.com/terraform-provider-aws/)
* openstack provider >1.17.0 - [here](https://releases.hashicorp.com/terraform-provider-openstack/)
* kubernetes provider >=1.4.0 - [here](https://releases.hashicorp.com/terraform-provider-kubernetes/)
* github provider >=2.0.0 - [here](https://releases.hashicorp.com/terraform-provider-github/)
* datadog provider >1.19.0 - [here](https://releases.hashicorp.com/terraform-provider-datadog/)

Information on provider plugins:
https://www.terraform.io/docs/configuration/providers.html

### Use with GCP
[![asciicast](https://asciinema.org/a/243961.svg)](https://asciinema.org/a/243961)
Example:

```
terraformer import google --resources=gcs,forwardingRules,httpHealthChecks --connect=true --zone=europe-west1-a --projects=aaa,fff
terraformer import google --resources=gcs,forwardingRules,httpHealthChecks --filter=google_compute_firewall=rule1:rule2:rule3 --zone=europe-west1-a --projects=aaa,fff
```

List of supported GCP services:

*   `addresses`
    * `google_compute_address`
*   `autoscalers`
    * `google_compute_autoscaler`
*   `backendBuckets`
    * `google_compute_backend_bucket`
*   `backendServices`
    * `google_compute_backend_service`
*   `bigQuery`
    * `google_bigquery_dataset`
    * `google_bigquery_table`
*   `schedulerJobs`
    * `google_cloud_scheduler_job`
*   `disks`
    * `google_compute_disk`
*   `firewalls`
    * `google_compute_firewall`
*   `forwardingRules`
    * `google_compute_forwarding_rule`
*   `globalAddresses`
    * `google_compute_global_address`
*   `globalForwardingRules`
    * `google_compute_global_forwarding_rule`
*   `healthChecks`
    * `google_compute_health_check`
*   `httpHealthChecks`
    * `google_compute_http_health_check`
*   `httpsHealthChecks`
    * `google_compute_https_health_check`
*   `images`
    * `google_compute_image`
*   `instanceGroupManagers`
    * `google_compute_instance_group_manager`
*   `instanceGroups`
    * `google_compute_instance_group`
*   `instanceTemplates`
    * `google_compute_instance_template`
*   `instances`
    * `google_compute_instance`
*   `interconnectAttachments`
    * `google_compute_interconnect_attachment`
*   `memoryStore`
    * `google_redis_instance`
*   `networks`
    * `google_compute_network`
*   `nodeGroups`
    * `google_compute_node_group`
*   `nodeTemplates`
    * `google_compute_node_template`
*   `regionAutoscalers`
    * `google_compute_region_autoscaler`
*   `regionBackendServices`
    * `google_compute_region_backend_service`
*   `regionDisks`
    * `google_compute_region_disk`
*   `regionInstanceGroupManagers`
    * `google_compute_region_instance_group_manager`
*   `routers`
    * `google_compute_router`
*   `routes`
    * `google_compute_route`
*   `securityPolicies`
    * `google_compute_security_policy`
*   `sslPolicies`
    * `google_compute_ssl_policy`
*   `subnetworks`
    * `google_compute_subnetwork`
*   `targetHttpProxies`
    * `google_compute_target_http_proxy`
*   `targetHttpsProxies`
    * `google_compute_target_https_proxy`
*   `targetInstances`
    * `google_compute_target_instance`
*   `targetPools`
    * `google_compute_target_pool`
*   `targetSslProxies`
    * `google_compute_target_ssl_proxy`
*   `targetTcpProxies`
    * `google_compute_target_tcp_proxy`
*   `targetVpnGateways`
    * `google_compute_vpn_gateway`
*   `urlMaps`
    * `google_compute_url_map`
*   `vpnTunnels`
    * `google_compute_vpn_tunnel`
*   `gke`
    * `google_container_cluster`
    * `google_container_node_pool`
*   `pubsub`
    * `google_pubsub_subscription`
    * `google_pubsub_topic`
*   `dataProc`
    * `google_dataproc_cluster`
*   `cloudFunctions`
    * `google_cloudfunctions_function`
*   `gcs`
    * `google_storage_bucket`
    * `google_storage_bucket_acl`
    * `google_storage_default_object_acl`
    * `google_storage_bucket_iam_binding`
    * `google_storage_bucket_iam_member`
    * `google_storage_bucket_iam_policy`
    * `google_storage_notification`
*   `monitoring`
    * `google_monitoring_alert_policy`
    * `google_monitoring_group`
    * `google_monitoring_notification_channel`
    * `google_monitoring_uptime_check_config`
*   `dns`
    * `google_dns_managed_zone`
    * `google_dns_record_set`
*   `cloudsql`
    * `google_sql_database_instance`
    * `google_sql_database`
*   `kms`
    * `google_kms_key_ring`
    * `google_kms_crypto_key`
*   `project`
    * `google_project`

Your `tf` and `tfstate` files are written by default to
`generated/gcp/zone/service`.

### Use with AWS

Example:

```
 terraformer import aws --resources=vpc,subnet --connect=true --regions=eu-west-1
 terraformer import aws --resources=vpc,subnet --filter=aws_vpc=vpc_id1:vpc_id2:vpc_id3 --regions=eu-west-1
```

List of support AWS services:

*   `elb`
    * `aws_elb`
*   `alb`
    * `aws_lb`
    * `aws_lb_listener`
    * `aws_lb_listener_rule`
    * `aws_lb_listener_certificate`
    * `aws_lb_target_group`
    * `aws_lb_target_group_attachment`
*   `auto_scaling`
    * `aws_autoscaling_group`
    * `aws_launch_configuration`
    * `aws_launch_template`
*   `rds`
    * `aws_db_instance`
    * `aws_db_parameter_group`
    * `aws_db_subnet_group`
    * `aws_db_option_group`
    * `aws_db_event_subscription`
*   `iam`
    * `aws_iam_role`
    * `aws_iam_role_policy`
    * `aws_iam_user`
    * `aws_iam_user_group_membership`
    * `aws_iam_user_policy`
    * `aws_iam_policy_attachment`
    * `aws_iam_policy`
    * `aws_iam_group`
    * `aws_iam_group_membership`
    * `aws_iam_group_policy`
*   `igw`
    * `aws_internet_gateway`
*   `nacl`
    * `aws_network_acl`
*   `s3`
    * `aws_s3_bucket`
    * `aws_s3_bucket_policy`
*   `sg`
    * `aws_security_group`
*   `subnet`
    * `aws_subnet`
*   `vpc`
    * `aws_vpc`
*   `vpn_connection`
    * `aws_vpn_connection`
*   `vpn_gateway`
    * `aws_vpn_gateway`
*   `route53`
    * `aws_route53_zone`
    * `aws_route53_record`
*   `acm`
    * `aws_acm_certificate`
*   `elasticache`
    * `aws_elasticache_cluster`
    * `aws_elasticache_parameter_group`
    * `aws_elasticache_subnet_group`
    * `aws_elasticache_replication_group`

### Use with OpenStack

Example:

```
 terraformer import openstack --resources=compute,networking --regions=RegionOne
```

List of support OpenStack services:

*   `compute`
    * `openstack_compute_instance_v2`
*   `networking`
    * `openstack_networking_secgroup_v2`
    * `openstack_networking_secgroup_rule_v2`
*   `blockstorage`
    * `openstack_blockstorage_volume_v1`
    * `openstack_blockstorage_volume_v2`
    * `openstack_blockstorage_volume_v3`

### Use with Kubernetes

Example:

```
 terraformer import kubernetes --resources=deployments,services,storageclasses
 terraformer import kubernetes --resources=deployments,services,storageclasses --filter=kubernetes_deployment=name1:name2:name3
```

All of the kubernetes resources that are currently being supported by kubernetes provider are supported by this module as well. Here is the list of resources which are currently supported by kubernetes provider v.1.4:

* `clusterrolebinding`
  * `kubernetes_cluster_role_binding`
* `configmaps`
  * `kubernetes_config_map`
* `deployments`
  * `kubernetes_deployment`
* `horizontalpodautoscalers`
  * `kubernetes_horizontal_pod_autoscaler`
* `limitranges`
  * `kubernetes_limit_range`
* `namespaces`
  * `kubernetes_namespace`
* `persistentvolumes`
  * `kubernetes_persistent_volume`
* `persistentvolumeclaims`
  * `kubernetes_persistent_volume_claim`
* `pods`
  * `kubernetes_pod`
* `replicationcontrollers`
  * `kubernetes_replication_controller`
* `resourcequotas`
  * `kubernetes_resource_quota`
* `secrets`
  * `kubernetes_secret`
* `services`
  * `kubernetes_service`
* `serviceaccounts`
  * `kubernetes_service_account`
* `statefulsets`
  * `kubernetes_stateful_set`
* `storageclasses`
  * `kubernetes_storage_class`

#### Known issues

* Terraform kubernetes provider is rejecting resources with ":" character in their names (As it's not meeting DNS-1123), while it's allowed for certain types in kubernetes, e.g. ClusterRoleBinding.
* As terraform flatmap is using "." to detect the keys for unflattening the maps, some keys with "." in their names are being considered as the maps.
* As the library is just assuming empty string as empty value (not "0"), there are some issues with optional integer keys that are restricted to be positive.

### Use with Github

Example:

```
 ./terraformer import github --organizations=YOUR_ORGANIZATION --resources=repositories --token=YOUR_TOKEN // or GITHUB_TOKEN in env
 ./terraformer import github --organizations=YOUR_ORGANIZATION --resources=repositories --filter=github_repository=id1:id2:id4 --token=YOUR_TOKEN // or GITHUB_TOKEN in env
```

Support only organizations resources. List of supported resources:

* `repositories`
    * `github_repository`
    * `github_repository_webhook`
    * `github_branch_protection`
    * `github_repository_collaborator`
    * `github_repository_deploy_key`
* `teams`
    * `github_team`
    * `github_team_membership`
    * `github_team_repository`
* `members`
    * `github_membership`
* `organization_webhooks`
    * `github_organization_webhook`

Notes:
* Github API don't return webhook secrets. If you have secret in webhook, you get changes on `terraform plan`
=> `configuration.#: "1" => "0"` in tfstate only.

### Use with Datadog

Example:

```
 ./terraformer import datadog --resources=monitor --api-key=YOUR_DATADOG_API_KEY // or DATADOG_API_KEY in env --app-key=YOUR_DATADOG_APP_KEY // or DATADOG_APP_KEY in env
 ./terraformer import datadog --resources=monitor --filter=datadog_monitor=id1:id2:id4 --api-key=YOUR_DATADOG_API_KEY // or DATADOG_API_KEY in env --app-key=YOUR_DATADOG_APP_KEY // or DATADOG_APP_KEY in env
```

List of support Datadog services:

* `downtime`
    * `datadog_downtime`
* `monitor`
    * `datadog_monitor`
* `screenboard`
    * `datadog_screenboard`
* `synthetics`
    * `datadog_synthetics_test`
* `timeboard`
    * `datadog_timeboard`
* `user`
    * `datadog_user`


## Contributing

If you have improvements or fixes, we would love to have your contributions.
Please read CONTRIBUTING.md for more information on the process we would like
contributors to follow.

## Developing
Terraformer built for easy to add new providers and not only cloud providers.

Process for generating `tf` + `tfstate` files:

1.  Call GCP/AWS/other api and get list of resources.
2.  Iterate over resources and take only ID (we don't need mapping fields!!!)
3.  Call to provider for readonly fields.
4.  Call to infrastructure and take tf + tfstate.

## Infrastructure

1.  Call to provider for refresh method and get all data.
2.  Convert refresh data to go struct.
3.  Generate HCL file - `tf` files.
4.  Generate `tfstate` files.

All mapping of resource is made by providers and Terraform. Upgrades are needed only
for providers.

##### GCP compute resources

For GCP compute resources, use generated code from
`providers/gcp/gcp_compute_code_generator`.

To regenerate code:

```
go run providers/gcp/gcp_compute_code_generator/*.go
```

### Similar projects

1.  https://github.com/dtan4/terraforming
