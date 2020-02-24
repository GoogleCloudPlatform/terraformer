# Terraformer

[![Build Status](https://travis-ci.com/GoogleCloudPlatform/terraformer.svg?branch=master)](https://travis-ci.com/GoogleCloudPlatform/terraformer)
[![Go Report Card](https://goreportcard.com/badge/github.com/GoogleCloudPlatform/terraformer)](https://goreportcard.com/report/github.com/GoogleCloudPlatform/terraformer)
[![AUR package](https://img.shields.io/aur/version/terraformer)](https://aur.archlinux.org/packages/terraformer/)

A CLI tool that generates `tf`/`json` and `tfstate` files based on existing infrastructure
(reverse Terraform).

*   Disclaimer: This is not an official Google product
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
        * [AliCloud](#use-with-alicloud)
    * Cloud
        * [DigitalOcean](#use-with-digitalocean)
        * [Fastly](#use-with-fastly)
        * [Heroku](#use-with-heroku)
        * [Linode](#use-with-linode)
        * [OpenStack](#use-with-openstack)
        * [Vultr](#use-with-vultr)
    * Infrastructure Software
        * [Kubernetes](#use-with-kubernetes)
        * [OctopusDeploy](#use-with-octopusdeploy)
        * [RabbitMQ](#use-with-rabbitmq)
    * Network
        * [Cloudflare](#use-with-cloudflare)
    * VCS
        * [GitHub](#use-with-github)
    * Monitoring & System Management
        * [Datadog](#use-with-datadog)
        * [New Relic](#use-with-new-relic)
    * Community
        * [Keycloak](#use-with-keycloak)
        * [Logz.io](#use-with-logzio)
        * [Commercetools](#use-with-commercetools)
- [Contributing](#contributing)
- [Developing](#developing)
- [Infrastructure](#infrastructure)

## Capabilities

1.  Generate `tf`/`json` + `tfstate` files from existing infrastructure for all
    supported objects by resource.
2.  Remote state can be uploaded to a GCS bucket.
3.  Connect between resources with `terraform_remote_state` (local and bucket).
4.  Save `tf`/`json` files using a custom folder tree pattern.
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
  -С, --compact                (default false)
  -f, --filter strings        google_compute_firewall=id1:id2:id4
  -h, --help                  help for google
  -O, --output string         output format hcl or json (default "hcl")
  -o, --path-output string     (default "generated")
  -p, --path-pattern string   {output}/{provider}/ (default "{output}/{provider}/{service}/")
      --projects strings
  -z, --regions strings       europe-west1, (default [global])
  -r, --resources strings     firewalls,networks
  -s, --state string          local or bucket (default "local")
  -v, --verbose               verbose mode

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
$ terraformer plan google --resources=networks,firewalls --projects=my-project --regions=europe-west1-d
(snip)

Saving planfile to generated/google/my-project/terraformer/plan.json
```

After reviewing/customizing the planfile, begin the import by running `import plan`.

```
$ terraformer import plan generated/google/my-project/terraformer/plan.json
```

### Resource structure

Terraformer by default separates each resource into a file, which is put into a given service directory.

The default path for resource files is `{output}/{provider}/{service}/{resource}.tf` and can vary for each provider.

It's possible to adjust the generated structure by:
1. Using `--compact` parameter to group resource files within a single service into one `resources.tf` file
2. Adjusting the `--path-pattern` parameter and passing e.g. `--path-pattern {output}/{provider}/` to generate resources for all services in one directory

It's possible to combine `--compact` `--path-pattern` parameters together.

### Installation

From source:
1.  Run `git clone <terraformer repo>`
2.  Run `GO111MODULE=on go mod vendor`
3.  Run `go build -v` for all providers OR build with one provider `go run build/main.go {google,aws,azure,kubernetes and etc}`
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
export PROVIDER={all,google,aws,kubernetes}
curl -LO https://github.com/GoogleCloudPlatform/terraformer/releases/download/$(curl -s https://api.github.com/repos/GoogleCloudPlatform/terraformer/releases/latest | grep tag_name | cut -d '"' -f 4)/terraformer-${PROVIDER}-linux-amd64
chmod +x terraformer-${PROVIDER}-linux-amd64
sudo mv terraformer-${PROVIDER}-linux-amd64 /usr/local/bin/terraformer
```
* MacOS
```
export PROVIDER={all,google,aws,kubernetes}
curl -LO https://github.com/GoogleCloudPlatform/terraformer/releases/download/$(curl -s https://api.github.com/repos/GoogleCloudPlatform/terraformer/releases/latest | grep tag_name | cut -d '"' -f 4)/terraformer-${PROVIDER}-darwin-amd64
chmod +x terraformer-${PROVIDER}-darwin-amd64
sudo mv terraformer-${PROVIDER}-darwin-amd64 /usr/local/bin/terraformer
```

#### Using a package manager

If you want to use a package manager:

- [Homebrew](https://brew.sh/) users can use `brew install terraformer`.

Links to download Terraform Providers:
* Major Cloud
    * Google Cloud provider >2.11.0 - [here](https://releases.hashicorp.com/terraform-provider-google/)
    * AWS provider >2.25.0 - [here](https://releases.hashicorp.com/terraform-provider-aws/)
    * Azure provider >1.35.0 - [here](https://releases.hashicorp.com/terraform-provider-azurerm/)
    * Alicloud provider >1.57.1 - [here](https://releases.hashicorp.com/terraform-provider-alicloud/)
* Cloud
    * DigitalOcean provider >1.9.1 - [here](https://releases.hashicorp.com/terraform-provider-digitalocean/)
    * Fastly provider >0.11.0 - [here](https://releases.hashicorp.com/terraform-provider-fastly/)
    * Heroku provider >2.2.1 - [here](https://releases.hashicorp.com/terraform-provider-heroku/)
    * Linode provider >1.8.0 - [here](https://releases.hashicorp.com/terraform-provider-linode/)
    * OpenStack provider >1.21.1 - [here](https://releases.hashicorp.com/terraform-provider-openstack/)
    * Vultr provider >1.0.5 - [here](https://releases.hashicorp.com/terraform-provider-vultr/)
* Infrastructure Software
    * Kubernetes provider >=1.9.0 - [here](https://releases.hashicorp.com/terraform-provider-kubernetes/)
    * RabbitMQ provider >=1.1.0 - [here](https://releases.hashicorp.com/terraform-provider-rabbitmq/)
* Network
    * Cloudflare provider >1.16 - [here](https://releases.hashicorp.com/terraform-provider-cloudflare/)
* VCS
    * GitHub provider >=2.2.1 - [here](https://releases.hashicorp.com/terraform-provider-github/)
* Monitoring & System Management
    * Datadog provider >2.1.0 - [here](https://releases.hashicorp.com/terraform-provider-datadog/)
    * New Relic provider >1.5.0 - [here](https://releases.hashicorp.com/terraform-provider-newrelic/)
* Community
    * Keycloak provider >=1.12.0 - [here](https://github.com/mrparkers/terraform-provider-keycloak/)
    * Logz.io provider >=1.1.1 - [here](https://github.com/jonboydell/logzio_terraform_provider/)
    * Commercetools provider >= 0.19.0 - [here](https://github.com/labd/terraform-provider-commercetools)

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

AWS configuration including environmental variables, shared credentials file (\~/.aws/credentials), and shared config file (\~/.aws/config) will be loaded by the tool by default. To use a specific profile, you can use the following command:

```
terraformer import aws --resources=vpc,subnet --regions=eu-west-1 --profile=prod
```

You can also provide no regions when importing resources:
```
terraformer import aws --resources=cloudfront --profile=prod
```
In that case terraformer will not know with which region resources are associated with and will not assume any region. That scenario is useful in case of global resources (e.g. CloudFront distributions or Route 53 records) and when region is passed implicitly through environmental variables or metadata service.

#### Supported services

*   `accessanalyzer`
    * `aws_accessanalyzer_analyzer`
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
*   `cloud9`
    * `aws_cloud9_environment_ec2`
*   `cloudfront`
    * `aws_cloudfront_distribution`
*   `cloudformation`
    * `aws_cloudformation_stack`
    * `aws_cloudformation_stack_set`
    * `aws_cloudformation_stack_set_instance`
*   `cloudtrail`
    * `aws_cloudtrail`
*   `codebuild`
    * `aws_codebuild_project`
*   `codecommit`
    * `aws_codecommit_repository`
*   `codedeploy`
    * `aws_codedeploy_app`
*   `codepipeline`
    * `aws_codepipeline`
    * `aws_codepipeline_webhook`
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
*   `elastic_beanstalk`
    * `aws_elastic_beanstalk_application`
    * `aws_elastic_beanstalk_environment`
*   `ecs`
    * `aws_ecs_cluster`
    * `aws_ecs_service`
    * `aws_ecs_task_definition`
*   `eks`
    * `aws_eks_cluster`
*   `elb`
    * `aws_elb`
*   `emr`
    * `aws_emr_cluster`
    * `aws_emr_security_configuration`
*   `eni`
    * `aws_network_interface`
*   `es`
    * `aws_elasticsearch_domain`
*   `firehose`
    * `aws_kinesis_firehose_delivery_stream`
*   `glue`
    * `glue_crawler`
    * `aws_glue_catalog_database`
    * `aws_glue_catalog_table`
*   `iam`
    * `aws_iam_group`
    * `aws_iam_group_policy`
    * `aws_iam_instance_profile`
    * `aws_iam_policy`
    * `aws_iam_policy_attachment`
    * `aws_iam_role`
    * `aws_iam_role_policy`
    * `aws_iam_user`
    * `aws_iam_user_group_membership`
    * `aws_iam_user_policy`
*   `igw`
    * `aws_internet_gateway`
*   `kinesis`
    * `aws_kinesis_stream`
*   `kms`
    * `aws_kms_key`
    * `aws_kms_alias`
*   `lambda`
    * `aws_lambda_event_source_mapping`
    * `aws_lambda_function`
    * `aws_lambda_function_event_invoke_config`
    * `aws_lambda_layer_version`
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
    * `aws_main_route_table_association`
    * `aws_route_table_association`
*   `s3`
    * `aws_s3_bucket`
    * `aws_s3_bucket_policy`
*   `sg`
    * `aws_security_group`
    * `aws_security_group_rule` (if a rule cannot be inlined)
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
*   `budgets`
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

#### SQS queues retrieval

Terraformer uses AWS [ListQueues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ListQueues.html) API call to fetch available queues. The API is able to return only up to 1000 queues and an additional name prefix should be passed to filter the list results. It's possible to pass `QueueNamePrefix` parameter by environmental variable `SQS_PREFIX`.

### Use with Azure
Support [Azure CLI](https://www.terraform.io/docs/providers/azurerm/guides/azure_cli.html), [Service Principal with Client Certificate](https://www.terraform.io/docs/providers/azurerm/guides/service_principal_client_certificate.html) & [Service Principal with Client Secret](https://www.terraform.io/docs/providers/azurerm/guides/service_principal_client_secret.html)

Example:

```
# Using Azure CLI (az login)
export ARM_SUBSCRIPTION_ID=[SUBSCRIPTION_ID]

# Using Service Principal with Client Certificate
export ARM_SUBSCRIPTION_ID=[SUBSCRIPTION_ID]
export ARM_CLIENT_ID=[CLIENT_ID]
export ARM_CLIENT_CERTIFICATE_PATH="/path/to/my/client/certificate.pfx"
export ARM_CLIENT_CERTIFICATE_PASSWORD=[CLIENT_CERTIFICATE_PASSWORD]
export ARM_TENANT_ID=[TENANT_ID]

# Service Principal with Client Secret
export ARM_SUBSCRIPTION_ID=[SUBSCRIPTION_ID]
export ARM_CLIENT_ID=[CLIENT_ID]
export ARM_CLIENT_SECRET=[CLIENT_SECRET]
export ARM_TENANT_ID=[TENANT_ID]

./terraformer import azure -r resource_group
```

List of supported Azure resources:

*   `analysis`
    * `azurerm_analysis_services_server`
*   `database`
	* `azurerm_mariadb_configuration`
	* `azurerm_mariadb_database`
	* `azurerm_mariadb_firewall_rule`
	* `azurerm_mariadb_server`
	* `azurerm_mariadb_virtual_network_rule`
	* `azurerm_mysql_configuration`
	* `azurerm_mysql_database`
	* `azurerm_mysql_firewall_rule`
	* `azurerm_mysql_server`
	* `azurerm_mysql_virtual_network_rule`
	* `azurerm_postgresql_configuration`
	* `azurerm_postgresql_database`
	* `azurerm_postgresql_firewall_rule`
	* `azurerm_postgresql_server`
	* `azurerm_postgresql_virtual_network_rule`
	* `azurerm_sql_database`
	* `azurerm_sql_active_directory_administrator`
	* `azurerm_sql_elasticpool`
	* `azurerm_sql_failover_group`
	* `azurerm_sql_firewall_rule`
	* `azurerm_sql_server`
	* `azurerm_sql_virtual_network_rule`
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

### Use with AliCloud

You can either edit your alicloud config directly, (usually it is `~/.aliyun/config.json`)
or run `aliyun configure` and enter the credentials when prompted.

Terraformer will pick up the profile name specified in the `--profile` parameter.
It defaults to the first config in the config array.

```sh
terraformer import alicloud --resources=ecs --regions=ap-southeast-3 --profile=default
```

For all *supported* resources, you can do

```sh
# https://unix.stackexchange.com/a/114948/203870
export ALL_SUPPORTED_ALICLOUD_RESOURCES=$(terraformer import alicloud list | sed -e 'H;1h;$!d;x;y/\n/,/')
terraformer import alicloud --resources=$ALL_SUPPORTED_ALICLOUD_RESOURCES --regions=ap-southeast-3
```

List of supported AliCloud resources:

* `dns`
  * `alicloud_dns`
  * `alicloud_dns_record`
* `ecs`
  * `alicloud_instance`
* `keypair`
  * `alicloud_key_pair`
* `nat`
  * `alicloud_nat_gateway`
* `pvtz`
  * `alicloud_pvtz_zone`
  * `alicloud_pvtz_zone_attachment`
  * `alicloud_pvtz_zone_record`
* `ram`
  * `alicloud_ram_role`
  * `alicloud_ram_role_policy_attachment`
* `rds`
  * `alicloud_db_instance`
* `sg`
  * `alicloud_security_group`
  * `alicloud_security_group_rule`
* `slb`
  * `alicloud_slb`
  * `alicloud_slb_server_group`
  * `alicloud_slb_listener`
* `vpc`
  * `alicloud_vpc`
* `vswitch`
  * `alicloud_vswitch`

### Use with DigitalOcean

Example:

```
export DIGITALOCEAN_TOKEN=[DIGITALOCEAN_TOKEN]
./terraformer import digitalocean -r project,droplet
```

List of supported DigitalOcean resources:

*   `cdn`
    * `digitalocean_cdn`
*   `certificate`
    * `digitalocean_certificate`
*   `database_cluster`
    * `digitalocean_database_cluster`
    * `digitalocean_database_connection_pool`
    * `digitalocean_database_db`
    * `digitalocean_database_replica`
    * `digitalocean_database_user`
*   `domain`
    * `digitalocean_domain`
    * `digitalocean_record`
*   `droplet`
    * `digitalocean_droplet`
*   `droplet_snapshot`
    * `digitalocean_droplet_snapshot`
*   `firewall`
    * `digitalocean_firewall`
*   `floating_ip`
    * `digitalocean_floating_ip`
*   `kubernetes_cluster`
    * `digitalocean_kubernetes_cluster`
    * `digitalocean_kubernetes_node_pool`
*   `loadbalancer`
    * `digitalocean_loadbalancer`
*   `project`
    * `digitalocean_project`
*   `ssh_key`
    * `digitalocean_ssh_key`
*   `tag`
    * `digitalocean_tag`
*   `volume`
    * `digitalocean_volume`
*   `volume_snapshot`
    * `digitalocean_volume_snapshot`

### Use with Fastly

Example:

```
export FASTLY_API_KEY=[FASTLY_API_KEY]
./terraformer import fastly -r service_v1
```

List of supported Fastly resources:

*   `service_v1`
    * `fastly_service_acl_entries_v1`
    * `fastly_service_dictionary_items_v1`
    * `fastly_service_dynamic_snippet_content_v1`
    * `fastly_service_v1`

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
*   `cert`
    * `heroku_cert`
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

### Use with Linode

Example:

```
export LINODE_TOKEN=[LINODE_TOKEN]
./terraformer import linode -r instance
```

List of supported Linode resources:

*   `domain`
    * `linode_domain`
    * `linode_domain_record`
*   `image`
    * `linode_image`
*   `instance`
    * `linode_instance`
*   `nodebalancer`
    * `linode_nodebalancer`
    * `linode_nodebalancer_config`
    * `linode_nodebalancer_node`
*   `rdns`
    * `linode_rdns`
*   `sshkey`
    * `linode_sshkey`
*   `stackscript`
    * `linode_stackscript`
*   `token`
    * `linode_token`
*   `volume`
    * `linode_volume`

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

### Use with Vultr

Example:

```
export VULTR_API_KEY=[VULTR_API_KEY]
./terraformer import vultr -r server
```

List of supported Vultr resources:

*   `bare_metal_server`
    * `vultr_bare_metal_server`
*   `block_storage`
    * `vultr_block_storage`
*   `dns_domain`
    * `vultr_dns_domain`
    * `vultr_dns_record`
*   `firewall_group`
    * `vultr_firewall_group`
    * `vultr_firewall_rule`
*   `network`
    * `vultr_network`
*   `reserved_ip`
    * `vultr_reserved_ip`
*   `server`
    * `vultr_server`
*   `snapshot`
    * `vultr_snapshot`
*   `ssh_key`
    * `vultr_ssh_key`
*   `startup_script`
    * `vultr_startup_script`
*   `user`
    * `vultr_user`

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

### Use with OctopusDeploy

Example:

```
export OCTOPUS_CLI_SERVER=http://localhost:8081/
export OCTOPUS_CLI_API_KEY=API-CK7DQ8BMJCUUBSHAJCDIATXUO

terraformer import octopusdeploy --resources=tagsets
```

* `accounts`
  * `octopusdeploy_account`
* `certificates`
  * `octopusdeploy_certificate`
* `environments`
  * `octopusdeploy_environment`
* `feeds`
  * `octopusdeploy_feed`
* `libraryvariablesets`
  * `octopusdeploy_library_variable_set`
* `lifecycle`
  * `octopusdeploy_lifecycle`
* `project`
  * `octopusdeploy_project`
* `projectgroups`
  * `octopusdeploy_project_group`
* `projecttriggers`
  * `octopusdeploy_project_deployment_target_trigger`
* `tagsets`
  * `octopusdeploy_tag_set`

### Use with RabbitMQ

Example:

```
 export RABBITMQ_SERVER_URL=http://foo.bar.localdomain:15672
 export RABBITMQ_USERNAME=[RABBITMQ_USERNAME]
 export RABBITMQ_PASSWORD=[RABBITMQ_PASSWORD]

 terraformer import rabbitmq --resources=vhosts,queues,exchanges
 terraformer import rabbitmq --resources=vhosts,queues,exchanges --filter=rabbitmq_vhost=name1:name2:name3
```

All RabbitMQ resources that are currently supported by the RabbitMQ provider, are also supported by this module. Here is the list of resources which are currently supported by RabbitMQ provider v.1.1.0:

*   `bindings`
    * `rabbitmq_binding`
*   `exchanges`
    * `rabbitmq_exchange`
*   `permissions`
    * `rabbitmq_permissions`
*   `policies`
    * `rabbitmq_policy`
*   `queues`
    * `rabbitmq_queue`
*   `users`
    * `rabbitmq_user`
*   `vhosts`
    * `rabbitmq_vhost`

### Use with Cloudflare

Example using a Cloudflare API Key and corresponding email:
```
export CLOUDFLARE_API_KEY=[CLOUDFLARE_API_KEY]
export CLOUDFLARE_EMAIL=[CLOUDFLARE_EMAIL]
 ./terraformer import cloudflare --resources=firewall,dns
```

or using a Cloudflare API Token:

```
export CLOUDFLARE_API_TOKEN=[CLOUDFLARE_API_TOKEN]
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
*   `organization_blocks`
    * `github_organization_block`
*   `organization_projects`
    * `github_organization_project`
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
*   `user_ssh_keys`
    * `github_user_ssh_key`

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

### Use with Keycloak

Example:

```
 export KEYCLOAK_URL=https://foo.bar.localdomain
 export KEYCLOAK_CLIENT_ID=[KEYCLOAK_CLIENT_ID]
 export KEYCLOAK_CLIENT_SECRET=[KEYCLOAK_CLIENT_SECRET]

 terraformer import keycloak --resources=realms,openid_clients
 terraformer import keycloak --resources=realms --filter=keycloak_realm=name1:name2:name3
```

Here is the list of resources which are currently supported by Keycloak provider v.1.12.0:

- `groups`
  - `keycloak_group`
  - `keycloak_group_memberships`
  - `keycloak_group_roles`
  - `keycloak_default_groups`
- `openid_clients`
  - `keycloak_openid_client`
  - `keycloak_openid_client_service_account_role`
  - `keycloak_openid_user_attribute_protocol_mapper`
  - `keycloak_openid_user_property_protocol_mapper`
  - `keycloak_openid_full_name_protocol_mapper`
  - `keycloak_openid_audience_protocol_mapper`
  - `keycloak_openid_group_membership_protocol_mapper`
  - `keycloak_openid_hardcoded_claim_protocol_mapper`
  - `keycloak_openid_hardcoded_role_protocol_mapper`
- `realms`
  - `keycloak_realm`
  - `keycloak_ldap_user_federation`
  - `keycloak_ldap_full_name_mapper`
  - `keycloak_ldap_group_mapper`
  - `keycloak_ldap_msad_user_account_control_mapper`
  - `keycloak_ldap_user_attribute_mapper`
  - `keycloak_required_action`
- `roles`
  - `keycloak_role`
- `scopes`
  - `keycloak_openid_client_scope`
  - `keycloak_openid_client_default_scopes`
  - `keycloak_openid_client_optional_scopes`
- `users`
  - `keycloak_user`

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

### Use with [Commercetools](https://commercetools.com/de/)

This provider use the [terraform-provider-commercetools](https://github.com/labd/terraform-provider-commercetools). The terraformer provider was build by [Dustin Deus](https://github.com/StarpTech).

Example:

```
CTP_CLIENT_ID=foo CTP_CLIENT_SCOPE=scope CTP_CLIENT_SECRET=bar CTP_PROJECT_KEY=key ./terraformer plan commercetools -r=types // Only planning
CTP_CLIENT_ID=foo CTP_CLIENT_SCOPE=scope CTP_CLIENT_SECRET=bar CTP_PROJECT_KEY=key ./terraformer import commercetools -r=types // Import commercetools types
```

List of supported [commercetools](https://commercetools.com/de/) resources:

*   `api_extension`
    * `commercetools_api_extension`
*   `channel`
    * `commercetools_channel`
*   `product_type`
    * `commercetools_product_type`
*   `shipping_zone`
    * `commercetools_shipping_zone`
*   `state`
    * `commercetools_state`
*   `store`
    * `commercetools_store`
*   `subscription`
    * `commercetools_subscription`
*   `tax_category`
    * `commercetools_tax_category`
*   `types`
    * `commercetools_type`

## Contributing

If you have improvements or fixes, we would love to have your contributions.
Please read CONTRIBUTING.md for more information on the process we would like
contributors to follow.

## Developing

Terraformer was built so you can easily add new providers of any kind.

Process for generating `tf`/`json` + `tfstate` files:

1.  Call GCP/AWS/other api and get list of resources.
2.  Iterate over resources and take only the ID (we don't need mapping fields!).
3.  Call to provider for readonly fields.
4.  Call to infrastructure and take tf + tfstate.

## Infrastructure

1.  Call to provider using the refresh method and get all data.
2.  Convert refresh data to go struct.
3.  Generate HCL file - `tf`/`json` files.
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
