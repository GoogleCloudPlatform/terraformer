# Terraformer

[![Build Status](https://travis-ci.com/GoogleCloudPlatform/terraformer.svg?branch=master)](https://travis-ci.com/GoogleCloudPlatform/terraformer)
[![Go Report Card](https://goreportcard.com/badge/github.com/GoogleCloudPlatform/terraformer)](https://goreportcard.com/report/github.com/GoogleCloudPlatform/terraformer)
[![AUR package](https://img.shields.io/aur/version/terraformer)](https://aur.archlinux.org/packages/terraformer/)

A CLI tool that generates `tf` and `tfstate` files based on existing infrastructure
(reverse Terraform).

*   Disclaimer: This is not an official Google product
*   Status: beta - we still need to improve documentation, squash some bugs, etc...
*   Created by: Waze SRE

![Waze SRE logo](docs/waze-sre-logo.png)

# Table of Contents

- [Capabilities](#capabilities)
- [Installation](#installation)
- [Supported Providers](/providers)
    * Major Cloud
        * [Google Cloud](#use-with-gcp)
        * [AWS](#use-with-aws)
        * [Azure](#use-with-azure)
    * Cloud
        * [Heroku](#use-with-heroku)
        * [OpenStack](#use-with-openstack)
    * Infrastructure Software
        * [Kubernetes](#use-with-kubernetes)
    * Network
        * [Cloudflare](#use-with-cloudflare)
    * VCS
        * [GitHub](#use-with-github)
    * Monitoring & System Management
        * [Datadog](#use-with-datadog)
        * [New Relic](#use-with-new-relic)
    * Community
        * [Logz.io](#use-with-logzio)
- [Contributing](#contributing)
- [Developing](#developing)
- [Infrastructure](#infrastructure)

## Capabilities

1.  Generate `tf` + `tfstate` files from existing infrastructure for all
    supported objects by resource.
2.  Remote state can be uploaded to a GCS bucket.
3.  Connect between resources with `terraform_remote_state` (local and bucket).
4.  Save `tf` files using a custom folder tree pattern.
5.  Import by resource name and type.
6.  Support terraform 0.12 (for terraform 0.11 use v0.7.9).

Terraformer uses Terraform providers and is designed to easily support newly added resources.
To upgrade resources with new fields, all you need to do is upgrade the relevant Terraform providers.
```
Import current state to Terraform configuration from Google Cloud

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
  -z, --regions strings       europe-west1, (default [global])
  -r, --resources strings     firewalls,networks
  -s, --state string          local or bucket (default "local")

Use " import google [command] --help" for more information about a command.
```
#### Permissions

Read-only permissions

#### Filtering

Filters are a way to choose which resources `terraformer` imports. It's possible to filter resources by its identifiers or attributes. Multiple filtering values are separated by `:`. If an identifier contains this symbol, value should be wrapped in `'` e.g. `--filter=resource=id1:'project:dataset_id'`. Identifier based filters will be executed before Terraformer will try to refresh remote state.

##### Resource ID

Filtering is based on Terraform resource ID patterns. To find valid ID patterns for your resource, check the import part of the [Terraform documentation][terraform-providers].

[terraform-providers]: https://www.terraform.io/docs/providers/

Example usage:

```
terraformer import aws --resources=vpc,subnet --filter=aws_vpc=myvpcid --regions=eu-west-1
```
Will only import the vpc with id `myvpcid`. This form of filters can help when it's necessary to select resources by its identifiers.

#### Planning

The `plan` command generates a planfile that contains all the resources set to be imported. By modifying the planfile before running the `import` command, you can rename or filter the resources you'd like to import.

The rest of subcommands and parameters are identical to the `import` command.

```
$ terraformer plan google --resources=networks,firewalls --projects=my-project --zone=europe-west1-d
(snip)

Saving planfile to generated/google/my-project/terraformer/plan.json
```

After reviewing/customizing the planfile, begin the import by running `import plan`.

```
$ terraformer import plan generated/google/my-project/terraformer/plan.json
```

### Installation

From source:
1.  Run `git clone <terraformer repo>`
2.  Run `GO111MODULE=on go mod vendor`
3.  Run `go build -v`
4.  Run ```terraform init``` against an ```init.tf``` file to install the plugins required for your platform. For example, if you need plugins for the google provider, ```init.tf``` should contain:
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

#### Using a package manager

If you want to use a package manager:

- [Homebrew](https://brew.sh/) users can use `brew install terraformer`.

Links to download Terraform Providers:
* Major Cloud
    * Google Cloud provider >2.11.0 - [here](https://releases.hashicorp.com/terraform-provider-google/)
    * AWS provider >2.25.0 - [here](https://releases.hashicorp.com/terraform-provider-aws/)
    * Azure provider >1.35.0 - [here](https://releases.hashicorp.com/terraform-provider-azurerm/)
* Cloud
    * Heroku provider >2.2.1 - [here](https://releases.hashicorp.com/terraform-provider-heroku/)
    * OpenStack provider >1.21.1 - [here](https://releases.hashicorp.com/terraform-provider-openstack/)
* Infrastructure Software
    * Kubernetes provider >=1.9.0 - [here](https://releases.hashicorp.com/terraform-provider-kubernetes/)
* Network
    * Cloudflare provider >1.16 - [here](https://releases.hashicorp.com/terraform-provider-cloudflare/)
* VCS
    * GitHub provider >=2.2.1 - [here](https://releases.hashicorp.com/terraform-provider-github/)
* Monitoring & System Management
    * Datadog provider >2.1.0 - [here](https://releases.hashicorp.com/terraform-provider-datadog/)
    * New Relic provider >1.5.0 - [here](https://releases.hashicorp.com/terraform-provider-newrelic/)
* Community
    * Logz.io provider >=1.1.1 - [here](https://github.com/jonboydell/logzio_terraform_provider/)

Information on provider plugins:
https://www.terraform.io/docs/configuration/providers.html

### Use with GCP

[![asciicast](https://asciinema.org/a/243961.svg)](https://asciinema.org/a/243961)

Example:

```
terraformer import google --resources=gcs,forwardingRules,httpHealthChecks --connect=true --regions=europe-west1,europe-west4 --projects=aaa,fff
terraformer import google --resources=gcs,forwardingRules,httpHealthChecks --filter=google_compute_firewall=rule1:rule2:rule3 --regions=europe-west1 --projects=aaa,fff
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
*   `cloudFunctions`
    * `google_cloudfunctions_function`
*   `cloudsql`
    * `google_sql_database_instance`
    * `google_sql_database`
*   `dataProc`
    * `google_dataproc_cluster`
*   `disks`
    * `google_compute_disk`
*   `dns`
    * `google_dns_managed_zone`
    * `google_dns_record_set`
*   `firewalls`
    * `google_compute_firewall`
*   `forwardingRules`
    * `google_compute_forwarding_rule`
*   `gcs`
    * `google_storage_bucket`
    * `google_storage_bucket_acl`
    * `google_storage_default_object_acl`
    * `google_storage_bucket_iam_binding`
    * `google_storage_bucket_iam_member`
    * `google_storage_bucket_iam_policy`
    * `google_storage_notification`
*   `gke`
    * `google_container_cluster`
    * `google_container_node_pool`
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
*   `kms`
    * `google_kms_key_ring`
    * `google_kms_crypto_key`
*   `logging`
    * `google_logging_metric`
*   `memoryStore`
    * `google_redis_instance`
*   `monitoring`
    * `google_monitoring_alert_policy`
    * `google_monitoring_group`
    * `google_monitoring_notification_channel`
    * `google_monitoring_uptime_check_config`
*   `networks`
    * `google_compute_network`
*   `nodeGroups`
    * `google_compute_node_group`
*   `nodeTemplates`
    * `google_compute_node_template`
*   `project`
    * `google_project`
*   `pubsub`
    * `google_pubsub_subscription`
    * `google_pubsub_topic`
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
*   `schedulerJobs`
    * `google_cloud_scheduler_job`
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

Your `tf` and `tfstate` files are written by default to
`generated/gcp/zone/service`.

### Use with AWS

Example:

```
 terraformer import aws --resources=vpc,subnet --connect=true --regions=eu-west-1 --profile=prod
 terraformer import aws --resources=vpc,subnet --filter=aws_vpc=vpc_id1:vpc_id2:vpc_id3 --regions=eu-west-1
```

#### Profiles support

To load profiles from the shared AWS configuration file (typically `~/.aws/config`), set the `AWS_SDK_LOAD_CONFIG` to `true`:

```
AWS_SDK_LOAD_CONFIG=true terraformer import aws --resources=vpc,subnet --regions=eu-west-1 --profile=prod
```

You can also provide no regions when importing resources:
```
terraformer import aws --resources=cloudfront --profile=prod
```
In that case terraformer will not know with which region resources are associated with and will not assume any region. That scenario is useful in case of global resources (e.g. CloudFront distributions or Route 53 records) and when region is passed implicitly through environmental variables or metadata service.

#### Supported services

*   `acm`
    * `aws_acm_certificate`
*   `alb` (supports ALB and NLB)
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
*   `budgets`
    * `aws_budgets_budget`
*   `cloudfront`
    * `aws_cloudfront_distribution`
*   `cloudtrail`
    * `aws_cloudtrail`
*   `dynamodb`
    * `aws_dynamodb_table`
*   `ec2_instance`
    * `aws_instance`
*   `eip`
    * `aws_eip`
*   `elasticache`
    * `aws_elasticache_cluster`
    * `aws_elasticache_parameter_group`
    * `aws_elasticache_subnet_group`
    * `aws_elasticache_replication_group`
*   `ebs`
    * `aws_ebs_volume`
    * `aws_volume_attachment`
*   `ecs`
    * `aws_ecs_cluster`
    * `aws_ecs_service`
    * `aws_ecs_task_definition`
*   `elb`
    * `aws_elb`
*   `es`
    * `aws_elasticsearch_domain`
*   `firehose`
    * `aws_kinesis_firehose_delivery_stream`
*   `glue`
    * `glue_crawler`
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
*   `kinesis`
    * `aws_kinesis_stream`
*   `msk`
    * `aws_msk_cluster`
*   `nat`
    * `aws_nat_gateway`
*   `nacl`
    * `aws_network_acl`
*   `organization`
    * `aws_organizations_account`
    * `aws_organizations_organization`
    * `aws_organizations_organizational_unit`
    * `aws_organizations_policy`
    * `aws_organizations_policy_attachment`
*   `rds`
    * `aws_db_instance`
    * `aws_db_parameter_group`
    * `aws_db_subnet_group`
    * `aws_db_option_group`
    * `aws_db_event_subscription`
*   `route53`
    * `aws_route53_zone`
    * `aws_route53_record`
*   `route_table`
    * `aws_route_table`
*   `s3`
    * `aws_s3_bucket`
    * `aws_s3_bucket_policy`
*   `sg`
    * `aws_security_group`
*   `sns`
    * `aws_sns_topic`
    * `aws_sns_topic_subscription`
*   `sqs`
    * `aws_sqs_queue`
*   `subnet`
    * `aws_subnet`
*   `vpc`
    * `aws_vpc`
*   `vpc_peering`
    * `aws_vpc_peering_connection`
*   `vpn_connection`
    * `aws_vpn_connection`
*   `vpn_gateway`
    * `aws_vpn_gateway`

#### Global services

AWS services that are global will be imported without specified region even if several regions will be passed. It is to ensure only one representation of an AWS resource is imported.

List of global AWS services:
*   `cloudfront`
*   `iam`
*   `organization`
*   `route53`

#### Attribute filters

Attribute filters allow filtering across different resource types by its attributes.

```
terraformer import aws --resources=ec2_instance,ebs --filter=Name=tags.costCenter;Value=20000:'20001:1' --regions=eu-west-1
```
Will only import AWS EC2 instances along with EBS volumes annotated with tag `costCenter` with values `20000` or `20001:1`. Attribute filters are by default applicable to all resource types although it's possible to specify to what resource type a given filter should be applicable to by providing `Type=<type>` parameter. For example:
```
terraformer import aws --resources=ec2_instance,ebs --filter=Type=ec2_instance;Name=tags.costCenter;Value=20000:'20001:1' --regions=eu-west-1
```
Will work as same as example above with a change the filter will be applicable only to `ec2_instance` resources.

### Use with Azure

Example:

```
export ARM_CLIENT_ID=[CLIENT_ID]
export ARM_CLIENT_SECRET=[CLIENT_SECRET]
export ARM_SUBSCRIPTION_ID=[SUBSCRIPTION_ID]
export ARM_TENANT_ID=[TENANT_ID]

export AZURE_CLIENT_ID=[CLIENT_ID]
export AZURE_CLIENT_SECRET=[CLIENT_SECRET]
export AZURE_TENANT_ID=[TENANT_ID]

./terraformer import azure -r resource_group
```

List of supported Azure resources:

*   `disk`
    * `azurerm_managed_disk`
*   `network_interface`
    * `azurerm_network_interface`
*   `network_security_group`
    * `azurerm_network_security_group`
*   `resource_group`
    * `azurerm_resource_group`
*   `storage_account`
    * `azurerm_storage_account`
*   `virtual_machine`
    * `azurerm_virtual_machine`
*   `virtual_network`
    * `azurerm_virtual_network`

### Use with Heroku

Example:

```
export HEROKU_EMAIL=[HEROKU_EMAIL]
export HEROKU_API_KEY=[HEROKU_API_KEY]
./terraformer import heroku -r app,addon
```

List of supported Heroku resources:

*   `account_feature`
    * `heroku_account_feature`
*   `addon`
    * `heroku_addon`
*   `addon_attachment`
    * `heroku_addon_attachment`
*   `app`
    * `heroku_app`
*   `app_config_association`
    * `heroku_app_config_association`
*   `app_feature`
    * `heroku_app_feature`
*   `app_webhook`
    * `heroku_app_webhook`
*   `build`
    * `heroku_build`
*   `domain`
    * `heroku_domain`
*   `drain`
    * `heroku_drain`
*   `formation`
    * `heroku_formation`
*   `pipeline`
    * `heroku_pipeline`
*   `pipeline_coupling`
    * `heroku_pipeline_coupling`
*   `team_collaborator`
    * `heroku_team_collaborator`
*   `team_member`
    * `heroku_team_member`

### Use with OpenStack

Example:

```
 terraformer import openstack --resources=compute,networking --regions=RegionOne
```

List of supported OpenStack services:

*   `blockstorage`
    * `openstack_blockstorage_volume_v1`
    * `openstack_blockstorage_volume_v2`
    * `openstack_blockstorage_volume_v3`
*   `compute`
    * `openstack_compute_instance_v2`
*   `networking`
    * `openstack_networking_secgroup_v2`
    * `openstack_networking_secgroup_rule_v2`

### Use with Kubernetes

Example:

```
 terraformer import kubernetes --resources=deployments,services,storageclasses
 terraformer import kubernetes --resources=deployments,services,storageclasses --filter=kubernetes_deployment=name1:name2:name3
```

All Kubernetes resources that are currently supported by the Kubernetes provider, are also supported by this module. Here is the list of resources which are currently supported by Kubernetes provider v.1.4:

*   `clusterrolebinding`
    * `kubernetes_cluster_role_binding`
*   `configmaps`
    * `kubernetes_config_map`
*   `deployments`
    * `kubernetes_deployment`
*   `horizontalpodautoscalers`
    * `kubernetes_horizontal_pod_autoscaler`
*   `limitranges`
    * `kubernetes_limit_range`
*   `namespaces`
    * `kubernetes_namespace`
*   `persistentvolumes`
    * `kubernetes_persistent_volume`
*   `persistentvolumeclaims`
    * `kubernetes_persistent_volume_claim`
*   `pods`
    * `kubernetes_pod`
*   `replicationcontrollers`
    * `kubernetes_replication_controller`
*   `resourcequotas`
    * `kubernetes_resource_quota`
*   `secrets`
    * `kubernetes_secret`
*   `services`
    * `kubernetes_service`
*   `serviceaccounts`
    * `kubernetes_service_account`
*   `statefulsets`
    * `kubernetes_stateful_set`
*   `storageclasses`
    * `kubernetes_storage_class`

#### Known issues

* Terraform Kubernetes provider is rejecting resources with ":" characters in their names (as they don't meet DNS-1123), while it's allowed for certain types in Kubernetes, e.g. ClusterRoleBinding.
* Because Terraform flatmap uses "." to detect the keys for unflattening the maps, some keys with "." in their names are being considered as the maps.
* Since the library assumes empty strings to be empty values (not "0"), there are some issues with optional integer keys that are restricted to be positive.

### Use with Cloudflare

Example:
```
CLOUDFLARE_TOKEN=[CLOUDFLARE_API_TOKEN]
CLOUDFLARE_EMAIL=[CLOUDFLARE_EMAIL]
 ./terraformer import cloudflare --resources=firewall,dns
```

List of supported Cloudflare services:

* `access`
  * `cloudflare_access_application`
* `dns`
  * `cloudflare_zone`
  * `cloudflare_record`
* `firewall`
  * `cloudflare_access_rule`
  * `cloudflare_filter`
  * `cloudflare_firewall_rule`
  * `cloudflare_zone_lockdown`

### Use with GitHub

Example:

```
 ./terraformer import github --organizations=YOUR_ORGANIZATION --resources=repositories --token=YOUR_TOKEN // or GITHUB_TOKEN in env
 ./terraformer import github --organizations=YOUR_ORGANIZATION --resources=repositories --filter=github_repository=id1:id2:id4 --token=YOUR_TOKEN // or GITHUB_TOKEN in env
```

Supports only organizational resources. List of supported resources:

*   `members`
    * `github_membership`
*   `organization_webhooks`
    * `github_organization_webhook`
*   `repositories`
    * `github_repository`
    * `github_repository_webhook`
    * `github_branch_protection`
    * `github_repository_collaborator`
    * `github_repository_deploy_key`
*   `teams`
    * `github_team`
    * `github_team_membership`
    * `github_team_repository`

Notes:
* Terraformer can't get webhook secrets from the GitHub API. If you use a secret token in any of your webhooks, running `terraform plan` will result in a change being detected:
=> `configuration.#: "1" => "0"` in tfstate only.

### Use with Datadog

Example:

```
 ./terraformer import datadog --resources=monitor --api-key=YOUR_DATADOG_API_KEY // or DATADOG_API_KEY in env --app-key=YOUR_DATADOG_APP_KEY // or DATADOG_APP_KEY in env
 ./terraformer import datadog --resources=monitor --filter=datadog_monitor=id1:id2:id4 --api-key=YOUR_DATADOG_API_KEY // or DATADOG_API_KEY in env --app-key=YOUR_DATADOG_APP_KEY // or DATADOG_APP_KEY in env
```

List of supported Datadog services:

*   `dashboard`
    * `datadog_dashboard`
*   `downtime`
    * `datadog_downtime`
*   `monitor`
    * `datadog_monitor`
*   `screenboard`
    * `datadog_screenboard`
*   `synthetics`
    * `datadog_synthetics_test`
*   `timeboard`
    * `datadog_timeboard`
*   `user`
    * `datadog_user`

### Use with New Relic

Example:

```
NEWRELIC_API_KEY=[API-KEY]
./terraformer import newrelic -r alert,dashboard,infra,synthetics
```

List of supported New Relic resources:

*   `alert`
    * `newrelic_alert_channel`
    * `newrelic_alert_condition`
    * `newrelic_alert_policy`
*   `dashboard`
    * `newrelic_dashboard`
*   `infra`
    * `newrelic_infra_alert_condition`
*   `synthetics`
    * `newrelic_synthetics_monitor`
    * `newrelic_synthetics_alert_condition`

### Use with Logz.io

Example:

```
 LOGZIO_API_TOKEN=foobar LOGZIO_BASE_URL=https://api-eu.logz.io ./terraformer import logzio -r=alerts,alert_notification_endpoints // Import Logz.io alerts and alert notification endpoints
```

List of supported Logz.io resources:

*   `alerts`
    * `logzio_alert`
*   `alert_notification_endpoints`
    * `logzio_endpoint`

## Contributing

If you have improvements or fixes, we would love to have your contributions.
Please read CONTRIBUTING.md for more information on the process we would like
contributors to follow.

## Developing

Terraformer was built so you can easily add new providers of any kind.

Process for generating `tf` + `tfstate` files:

1.  Call GCP/AWS/other api and get list of resources.
2.  Iterate over resources and take only the ID (we don't need mapping fields!).
3.  Call to provider for readonly fields.
4.  Call to infrastructure and take tf + tfstate.

## Infrastructure

1.  Call to provider using the refresh method and get all data.
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

#### [terraforming](https://github.com/dtan4/terraforming)

##### Terraformer Benefits

* Simpler to add new providers and resources - already supports AWS, GCP, GitHub, Kubernetes, and Openstack. Terraforming supports only AWS.
* Better support for HCL + tfstate, including updates for Terraform 0.12.
* If a provider adds new attributes to a resource, there is no need change Terraformer code - just update the Terraform provider on your laptop.
* Automatically supports connections between resources in HCL files.

##### Comparison

Terraforming gets all attributes from cloud APIs and creates HCL and tfstate files with templating. Each attribute in the API needs to map to attribute in Terraform. Generated files from templating can be broken with illegal syntax. When a provider adds new attributes the terraforming code needs to be updated.

Terraformer instead uses Terraform provider files for mapping attributes, HCL library from Hashicorp, and Terraform code.

Look for S3 support in terraforming here and official S3 support
Terraforming lacks full coverage for resources - as an example you can see that 70% of S3 options are not supported:

* terraforming - https://github.com/dtan4/terraforming/blob/master/lib/terraforming/template/tf/s3.erb
* official S3 support - https://www.terraform.io/docs/providers/aws/r/s3_bucket.html
