# Terraformer

[![Build Status](https://travis-ci.com/GoogleCloudPlatform/terraformer.svg?branch=master)](https://travis-ci.com/GoogleCloudPlatform/terraformer)
[![Go Report Card](https://goreportcard.com/badge/github.com/GoogleCloudPlatform/terraformer)](https://goreportcard.com/report/github.com/GoogleCloudPlatform/terraformer)
[![AUR package](https://img.shields.io/aur/version/terraformer)](https://aur.archlinux.org/packages/terraformer/)
[![Homebrew](https://img.shields.io/badge/dynamic/json.svg?url=https://formulae.brew.sh/api/formula/terraformer.json&query=$.versions.stable&label=homebrew)](https://formulae.brew.sh/formula/terraformer)

A CLI tool that generates `tf`/`json` and `tfstate` files based on existing infrastructure
(reverse Terraform).

*   Disclaimer: This is not an official Google product
*   Created by: Waze SRE

![Waze SRE logo](docs/waze-sre-logo.png)

# Table of Contents
- [Demo GCP](#demo-gcp)
- [Capabilities](#capabilities)
- [Installation](#installation)
- [Supported Providers](/providers)
    * Major Cloud
        * [Google Cloud](/docs/gcp.md)
        * [AWS](/docs/aws.md)
        * [Azure](/docs/azure.md)
        * [AliCloud](/docs/alicloud.md)
        * [IBM Cloud](/docs/ibmcloud.md)
    * Cloud
        * [DigitalOcean](#use-with-digitalocean)
        * [Equinix Metal](#use-with-equinix-metal)
        * [Fastly](#use-with-fastly)
        * [Heroku](#use-with-heroku)
        * [Linode](#use-with-linode)
        * [NS1](#use-with-ns1)
        * [OpenStack](#use-with-openstack)
        * [TencentCloud](#use-with-tencentcloud)
        * [Vultr](#use-with-vultr)
        * [Yandex.Cloud](#use-with-yandex)
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
        * [Mikrotik](#use-with-mikrotik)
        * [Xen Orchestra](#use-with-xenorchestra)
        * [GmailFilter](#use-with-gmailfilter)
- [Contributing](#contributing)
- [Developing](#developing)
- [Infrastructure](#infrastructure)
- [Stargazers over time](#stargazers-over-time)

## Demo GCP
[![asciicast](https://asciinema.org/a/243961.svg)](https://asciinema.org/a/243961)

## Capabilities

1.  Generate `tf`/`json` + `tfstate` files from existing infrastructure for all
    supported objects by resource.
2.  Remote state can be uploaded to a GCS bucket.
3.  Connect between resources with `terraform_remote_state` (local and bucket).
4.  Save `tf`/`json` files using a custom folder tree pattern.
5.  Import by resource name and type.
6.  Support terraform 0.13 (for terraform 0.11 use v0.7.9).

Terraformer uses Terraform providers and is designed to easily support newly added resources.
To upgrade resources with new fields, all you need to do is upgrade the relevant Terraform providers.
```
Import current state to Terraform configuration from a provider

Usage:
   import [provider] [flags]
   import [provider] [command]

Available Commands:
  list        List supported resources for a provider

Flags:
  -b, --bucket string         gs://terraform-state
  -c, --connect                (default true)
  -С, --compact                (default false)
  -x, --excludes strings      firewalls,networks
  -f, --filter strings        compute_firewall=id1:id2:id4
  -h, --help                  help for google
  -O, --output string         output format hcl or json (default "hcl")
  -o, --path-output string     (default "generated")
  -p, --path-pattern string   {output}/{provider}/ (default "{output}/{provider}/{service}/")
      --projects strings
  -z, --regions strings       europe-west1, (default [global])
  -r, --resources strings     firewall,networks or * for all services
  -s, --state string          local or bucket (default "local")
  -v, --verbose               verbose mode
  -n, --retry-number          number of retries to perform if refresh fails
  -m, --retry-sleep-ms        time in ms to sleep between retries

Use " import [provider] [command] --help" for more information about a command.
```
#### Permissions

The tool requires read-only permissions to list service resources.

#### Resources

You can use `--resources` parameter to tell resources from what service you want to import.

To import resources from all services, use `--resources="*"` . If you want to exclude certain services, you can combine the parameter with `--excludes` to exclude resources from services you don't want to import e.g. `--resources="*" --excludes="iam"`.

#### Filtering

Filters are a way to choose which resources `terraformer` imports. It's possible to filter resources by its identifiers or attributes. Multiple filtering values are separated by `:`. If an identifier contains this symbol, value should be wrapped in `'` e.g. `--filter=resource=id1:'project:dataset_id'`. Identifier based filters will be executed before Terraformer will try to refresh remote state.

Use `Type` when you need to filter only one of several types of resources. Multiple filters can be combined when importing different resource types. An example would be importing all AWS security groups from a specific AWS VPC: 
```
terraformer import aws -r sg,vpc --filter Type=sg;Name=vpc_id;Value=VPC_ID --filter Type=vpc;Name=id;Value=VPC_ID 
```
Notice how the `Name` is different for `sg` than it is for `vpc`.

##### Resource ID

Filtering is based on Terraform resource ID patterns. To find valid ID patterns for your resource, check the import part of the [Terraform documentation][terraform-providers].

[terraform-providers]: https://www.terraform.io/docs/providers/

Example usage:

```
terraformer import aws --resources=vpc,subnet --filter=vpc=myvpcid --regions=eu-west-1
```
Will only import the vpc with id `myvpcid`. This form of filters can help when it's necessary to select resources by its identifiers.

##### Field name only

It is possible to filter by specific field name only. It can be used e.g. when you want to retrieve resources only with a specific tag key.

Example usage:

```
terraformer import aws --resources=s3 --filter="Name=tags.Abc" --regions=eu-west-1
```
Will only import the s3 resources that have tag `Abc`. This form of filters can help when the field values are not important from filtering perspective.

##### Field with dots

It is possible to filter by a field that contains a dot.

Example usage:

```
terraformer import aws --resources=s3 --filter="Name=tags.Abc.def" --regions=eu-west-1
```
Will only import the s3 resources that have tag `Abc.def`.

#### Planning

The `plan` command generates a planfile that contains all the resources set to be imported. By modifying the planfile before running the `import` command, you can rename or filter the resources you'd like to import.

The rest of subcommands and parameters are identical to the `import` command.

```
$ terraformer plan google --resources=networks,firewall --projects=my-project --regions=europe-west1-d
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
2.  Run `go mod download`
3.  Run `go build -v` for all providers OR build with one provider `go run build/main.go {google,aws,azure,kubernetes and etc}`
4.  Run ```terraform init``` against an ```versions.tf``` file to install the plugins required for your platform. For example, if you need plugins for the google provider, ```versions.tf``` should contain:

```
terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
  }
  required_version = ">= 0.13"
}
```
Or alternatively

*  Copy your Terraform provider's plugin(s) to folder
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
- [MacPorts](https://www.macports.org/) users can use `sudo port install terraformer`.
- [Chocolatey](https://chocolatey.org/) users can use `choco install terraformer`.

Links to download Terraform Providers:
* Major Cloud
    * Google Cloud provider >2.11.0 - [here](https://releases.hashicorp.com/terraform-provider-google/)
    * AWS provider >2.25.0 - [here](https://releases.hashicorp.com/terraform-provider-aws/)
    * Azure provider >1.35.0 - [here](https://releases.hashicorp.com/terraform-provider-azurerm/)
    * Alicloud provider >1.57.1 - [here](https://releases.hashicorp.com/terraform-provider-alicloud/)
* Cloud
    * DigitalOcean provider >1.9.1 - [here](https://releases.hashicorp.com/terraform-provider-digitalocean/)
    * Fastly provider >0.16.1 - [here](https://releases.hashicorp.com/terraform-provider-fastly/)
    * Heroku provider >2.2.1 - [here](https://releases.hashicorp.com/terraform-provider-heroku/)
    * Linode provider >1.8.0 - [here](https://releases.hashicorp.com/terraform-provider-linode/)
    * NS1 provider >1.8.3 - [here](https://releases.hashicorp.com/terraform-provider-ns1/)
    * OpenStack provider >1.21.1 - [here](https://releases.hashicorp.com/terraform-provider-openstack/)
    * TencentCloud provider >1.50.0 - [here](https://releases.hashicorp.com/terraform-provider-tencentcloud/)
    * Vultr provider >1.0.5 - [here](https://releases.hashicorp.com/terraform-provider-vultr/)
    * Yandex provider >0.42.0 - [here](https://releases.hashicorp.com/terraform-provider-yandex/)
* Infrastructure Software
    * Kubernetes provider >=1.9.0 - [here](https://releases.hashicorp.com/terraform-provider-kubernetes/)
    * RabbitMQ provider >=1.1.0 - [here](https://releases.hashicorp.com/terraform-provider-rabbitmq/)
* Network
    * Cloudflare provider >1.16 - [here](https://releases.hashicorp.com/terraform-provider-cloudflare/)
* VCS
    * GitHub provider >=2.2.1 - [here](https://releases.hashicorp.com/terraform-provider-github/)
* Monitoring & System Management
    * Datadog provider >2.1.0 - [here](https://releases.hashicorp.com/terraform-provider-datadog/)
    * New Relic provider >2.0.0 - [here](https://releases.hashicorp.com/terraform-provider-newrelic/)
* Community
    * Keycloak provider >=1.19.0 - [here](https://github.com/mrparkers/terraform-provider-keycloak/)
    * Logz.io provider >=1.1.1 - [here](https://github.com/jonboydell/logzio_terraform_provider/)
    * Commercetools provider >= 0.21.0 - [here](https://github.com/labd/terraform-provider-commercetools)
    * Mikrotik provider >= 0.2.2 - [here](https://github.com/ddelnano/terraform-provider-mikrotik)
    * Xen Orchestra provider >= 0.18.0 - [here](https://github.com/ddelnano/terraform-provider-xenorchestra)
    * GmailFilter provider >= 1.0.1 - [here](https://github.com/yamamoto-febc/terraform-provider-gmailfilter)

Information on provider plugins:
https://www.terraform.io/docs/configuration/providers.html


 
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

### Use with Equinix Metal

Example:

```
export METAL_AUTH_TOKEN=[METAL_AUTH_TOKEN]
export PACKET_PROJECT_ID=[PROJECT_ID]
./terraformer import metal -r volume,device
```

List of supported Equinix Metal resources:

*   `device`
    * `metal_device`
*   `volume`
    * `metal_volume`
*   `sshkey`
    * `metal_ssh_key`
*   `spotmarketrequest`
    * `metal_spot_market_request`

### Use with Fastly

Example:

```
export FASTLY_API_KEY=[FASTLY_API_KEY]
export FASTLY_CUSTOMER_ID=[FASTLY_CUSTOMER_ID]
./terraformer import fastly -r service_v1,user
```

List of supported Fastly resources:

*   `service_v1`
    * `fastly_service_acl_entries_v1`
    * `fastly_service_dictionary_items_v1`
    * `fastly_service_dynamic_snippet_content_v1`
    * `fastly_service_v1`
*   `user`
    * `fastly_user_v1`

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

### Use with NS1

Example:

```
$ export NS1_APIKEY=[NS1_APIKEY]
$ terraformer import ns1 -r zone,monitoringjob,team
```

List of supported NS1 resources:

*   `zone`
    * `ns1_zone`
*   `monitoringjob`
    * `ns1_monitoringjob`
*   `team`
    * `ns1_team`

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

### Use with TencentCloud

Example:

```
$ export TENCENTCLOUD_SECRET_ID=<SECRET_ID>
$ export TENCENTCLOUD_SECRET_KEY=<SECRET_KEY>
$ terraformer import tencentcloud --resources=cvm,cbs --regions=ap-guangzhou
```

List of supported TencentCloud services:

*    `as`
     * `tencentcloud_as_scaling_group`
     * `tencentcloud_as_scaling_config`
*    `cbs`
     * `tencentcloud_cbs_storage`
*    `cdn`
     * `tencentcloud_cdn_domain`
*    `cfs`
     * `tencentcloud_cfs_file_system`
*    `clb`
     * `tencentcloud_clb_instance`
*    `cos`
     * `tencentcloud_cos_bucket`
*    `cvm`
     * `tencentcloud_instance`
*    `elasticsearch`
     * `tencentcloud_elasticsearch_instance`
*    `gaap`
     * `tencentcloud_gaap_proxy`
     * `tencentcloud_gaap_realserver`
*    `key_pair`
     * `tencentcloud_key_pair`
*    `mongodb`
     * `tencentcloud_mongodb_instance`
*    `mysql`
     * `tencentcloud_mysql_instance`
     * `tencentcloud_mysql_readonly_instance`
*    `redis`
     * `tencentcloud_redis_instance`
*    `scf`
     * `tencentcloud_scf_function`
*    `security_group`
     * `tencentcloud_security_group`
*    `ssl`
     * `tencentcloud_ssl_certificate`
*    `subnet`
     * `tencentcloud_subnet`
*    `tcaplus`
     * `tencentcloud_tcaplus_cluster`
*    `vpc`
     * `tencentcloud_vpc`
*    `vpc`
     * `tencentcloud_vpn_gateway`

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

### Use with Yandex

Example:

```
export YC_TOKEN=[YANDEX_CLOUD_OAUTH_TOKEN]
export YC_FOLDER_ID=[YANDEX_FOLDER_ID]
./terraformer import yandex -r subnet
```

List of supported Yandex resources:

*   `instance`
    * `yandex_compute_instance`
*   `disk`
    * `yandex_compute_disk`
*   `subnet`
    * `yandex_vpc_subnet`
*   `network`
    * `yandex_vpc_network`

Your `tf` and `tfstate` files are written by default to
`generated/yandex/service`.

### Use with Kubernetes

Example:

```
 terraformer import kubernetes --resources=deployments,services,storageclasses
 terraformer import kubernetes --resources=deployments,services,storageclasses --filter=deployment=name1:name2:name3
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
 terraformer import rabbitmq --resources=vhosts,queues,exchanges --filter=vhost=name1:name2:name3
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
export CLOUDFLARE_ACCOUNT_ID=[CLOUDFLARE_ACCOUNT_ID]
 ./terraformer import cloudflare --resources=firewall,dns
```

or using a Cloudflare API Token:

```
export CLOUDFLARE_API_TOKEN=[CLOUDFLARE_API_TOKEN]
export CLOUDFLARE_ACCOUNT_ID=[CLOUDFLARE_ACCOUNT_ID]
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
  * `cloudflare_rate_limit`
* `page_rule`
  * `cloudflare_page_rule`
* `account_member`
  * `cloudflare_account_member`

### Use with GitHub

Example:

```
 ./terraformer import github --organizations=YOUR_ORGANIZATION --resources=repositories --token=YOUR_TOKEN // or GITHUB_TOKEN in env
 ./terraformer import github --organizations=YOUR_ORGANIZATION --resources=repositories --filter=repository=id1:id2:id4 --token=YOUR_TOKEN // or GITHUB_TOKEN in env
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
 ./terraformer import datadog --resources=monitor --api-key=YOUR_DATADOG_API_KEY // or DATADOG_API_KEY in env --app-key=YOUR_DATADOG_APP_KEY // or DATADOG_APP_KEY in env --api-url=DATADOG_API_URL // or DATADOG_HOST in env
 ./terraformer import datadog --resources=monitor --filter=monitor=id1:id2:id4 --api-key=YOUR_DATADOG_API_KEY // or DATADOG_API_KEY in env --app-key=YOUR_DATADOG_APP_KEY // or DATADOG_APP_KEY in env
```

List of supported Datadog services:

*   `dashboard`
    * `datadog_dashboard`
*   `dashboard_list`
    * `datadog_dashboard_list`
*   `downtime`
    * `datadog_downtime`
*   `logs_archive`
    * `datadog_logs_archive`
*   `logs_archive_order`
    * `datadog_logs_archive_order`
*   `logs_custom_pipeline`
    * `datadog_logs_custom_pipeline`
*   `logs_integration_pipeline`
    * `datadog_logs_integration_pipeline`
*   `logs_pipeline_order`
    * `datadog_logs_pipeline_order`
*   `logs_index`
    * `datadog_logs_index`
*   `logs_index_order`
    * `datadog_logs_index_order`
*   `integration_aws`
    * `datadog_integration_aws`
*   `integration_aws_lambda_arn`
    * `datadog_integration_aws_lambda_arn`
*   `integration_aws_log_collection`
    * `datadog_integration_aws_log_collection`
*   `integration_azure`
    * `datadog_integration_azure`
        * **_NOTE:_** Sensitive field `client_secret` is not generated and needs to be manually set
*   `integration_gcp`
    * `datadog_integration_gcp`
        * **_NOTE:_** Sensitive fields `private_key, private_key_id, client_id` is not generated and needs to be manually set
*   `integration_pagerduty`
    * `datadog_integration_pagerduty`
*   `integration_pagerduty_service_object`
    * `datadog_integration_pagerduty_service_object`
*   `metric_metadata`
    * `datadog_metric_metadata`
        * **_NOTE:_** Importing resource requires resource ID's to be passed via [Filter](#filtering) option
*   `monitor`
    * `datadog_monitor`
*   `role`
    * `datadog_role`
*   `screenboard`
    * `datadog_screenboard`
*   `security_monitoring_default_rule`
    * `datadog_security_monitoring_default_rule`
*   `security_monitoring_rule`
    * `datadog_security_monitoring_rule`
*   `service_level_objective`
    * `datadog_service_level_objective`
        * **_NOTE:_** Importing resource requires resource ID's to be passed via [Filter](#filtering) option
*   `synthetics`
    * `datadog_synthetics_test`
*   `synthetics_global_variables`
    * `datadog_synthetics_global_variables`
        * **_NOTE:_** Importing resource requires resource ID's to be passed via [Filter](#filtering) option
*   `synthetics_private_location`
    * `datadog_synthetics_private_location`
*   `timeboard`
    * `datadog_timeboard`
*   `user`
    * `datadog_user`

### Use with New Relic

Example:

```
./terraformer import newrelic -r alert,dashboard,infra,synthetics --api-key=NRAK-XXXXXXXX --account-id=XXXXX
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

 terraformer import keycloak --resources=realms
 terraformer import keycloak --resources=realms --filter=realm=name1:name2:name3
 terraformer import keycloak --resources=realms --targets realmA,realmB
```

Here is the list of resources which are currently supported by Keycloak provider v.1.19.0:

- `realms`
  - `keycloak_default_groups`
  - `keycloak_group`
  - `keycloak_group_memberships`
  - `keycloak_group_roles`
  - `keycloak_ldap_full_name_mapper`
  - `keycloak_ldap_group_mapper`
  - `keycloak_ldap_hardcoded_group_mapper`
  - `keycloak_ldap_hardcoded_role_mapper`
  - `keycloak_ldap_msad_lds_user_account_control_mapper`
  - `keycloak_ldap_msad_user_account_control_mapper`
  - `keycloak_ldap_user_attribute_mapper`
  - `keycloak_ldap_user_federation`
  - `keycloak_openid_audience_protocol_mapper`
  - `keycloak_openid_client`
  - `keycloak_openid_client_default_scopes`
  - `keycloak_openid_client_optional_scopes`
  - `keycloak_openid_client_scope`
  - `keycloak_openid_client_service_account_role`
  - `keycloak_openid_full_name_protocol_mapper`
  - `keycloak_openid_group_membership_protocol_mapper`
  - `keycloak_openid_hardcoded_claim_protocol_mapper`
  - `keycloak_openid_hardcoded_group_protocol_mapper`
  - `keycloak_openid_hardcoded_role_protocol_mapper` (only for client roles)
  - `keycloak_openid_user_attribute_protocol_mapper`
  - `keycloak_openid_user_property_protocol_mapper`
  - `keycloak_openid_user_realm_role_protocol_mapper`
  - `keycloak_openid_user_client_role_protocol_mapper`
  - `keycloak_openid_user_session_note_protocol_mapper`
  - `keycloak_realm`
  - `keycloak_required_action`
  - `keycloak_role`
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
*   `shipping_method`
    * `commercetools_shipping_method`
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

### Use with [Mikrotik](https://wiki.mikrotik.com/wiki/Manual:TOC)

This provider uses the [terraform-provider-mikrotik](https://github.com/ddelnano/terraform-provider-mikrotik). The terraformer provider was built by [Dom Del Nano](https://github.com/ddelnano).

Example:

```
## Warning! You should not expose your mikrotik creds through your bash history. Export them to your shell in a safe way when doing this for real!

MIKROTIK_HOST=router-hostname:8728 MIKROTIK_USER=username MIKROTIK_PASSWORD=password terraformer  import mikrotik -r=dhcp_lease

# Import only static IPs
MIKROTIK_HOST=router-hostname:8728 MIKROTIK_USER=username MIKROTIK_PASSWORD=password terraformer  import mikrotik -r=dhcp_lease --filter='Name=dynamic;Value=false'
```

List of supported mikrotik resources:

* `mikrotik_dhcp_lease`

### Use with [Xen Orchestra](https://xen-orchestra.com/)

This provider uses the [terraform-provider-xenorchestra](https://github.com/ddelnano/terraform-provider-xenorchestra). The terraformer provider was built by [Dom Del Nano](https://github.com/ddelnano) on behalf of [Vates SAS](https://vates.fr/) who is sponsoring Dom to work on the project.

Example:

```
## Warning! You should not expose your xenorchestra creds through your bash history. Export them to your shell in a safe way when doing this for real!

XOA_URL=ws://your-xenorchestra-domain XOA_USER=username XOA_PASSWORD=password terraformer import xenorchestra -r=acl
```

List of supported xenorchestra resources:

* `xenorchestra_acl`
* `xenorchestra_resource_set`


### Use with GmailFilter

Support [Using Service Accounts](https://github.com/yamamoto-febc/terraform-provider-gmailfilter/blob/master/README.md#using-a-service-accountg-suite-users-only) or [Using Application Default Credentials](https://github.com/yamamoto-febc/terraform-provider-gmailfilter/blob/master/README.md#using-an-application-default-credential).

Example:

```
# Using Service Accounts
export GOOGLE_CREDENTIALS=/path/to/client_secret.json
export IMPERSONATED_USER_EMAIL="foobar@example.com"

# Using Application Default Credentials
gcloud auth application-default login \
  --client-id-file=client_secret.json \
  --scopes \
https://www.googleapis.com/auth/gmail.labels,\
https://www.googleapis.com/auth/gmail.settings.basic

./terraformer import gmailfilter -r=filter,label
```

List of supported GmailFilter resources:

*   `label`
    * `gmailfilter_label`
*   `filter`
    * `gmailfilter_filter`

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

## Stargazers over time

[![Stargazers over time](https://starchart.cc/GoogleCloudPlatform/terraformer.svg)](https://starchart.cc/GoogleCloudPlatform/terraformer)
