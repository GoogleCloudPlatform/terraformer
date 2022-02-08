# Terraformer

[![tests](https://github.com/GoogleCloudPlatform/terraformer/actions/workflows/test.yml/badge.svg)](https://github.com/GoogleCloudPlatform/terraformer/actions/workflows/test.yml)
[![linter](https://github.com/GoogleCloudPlatform/terraformer/actions/workflows/linter.yml/badge.svg)](https://github.com/GoogleCloudPlatform/terraformer/actions/workflows/linter.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/GoogleCloudPlatform/terraformer)](https://goreportcard.com/report/github.com/GoogleCloudPlatform/terraformer)
[![AUR package](https://img.shields.io/aur/version/terraformer)](https://aur.archlinux.org/packages/terraformer/)
[![Homebrew](https://img.shields.io/badge/dynamic/json.svg?url=https://formulae.brew.sh/api/formula/terraformer.json&query=$.versions.stable&label=homebrew)](https://formulae.brew.sh/formula/terraformer)

A CLI tool that generates `tf`/`json` and `tfstate` files based on existing infrastructure
(reverse Terraform).

*   Disclaimer: This is not an official Google product
*   Created by: Waze SRE

![Waze SRE logo](assets/waze-sre-logo.png)

# Table of Contents
- [Demo GCP](#demo-gcp)
- [Capabilities](#capabilities)
- [Installation](#installation)
- [Supported Providers](/docs)
    * Major Cloud
        * [Google Cloud](/docs/gcp.md)
        * [AWS](/docs/aws.md)
        * [Azure](/docs/azure.md)
        * [AliCloud](/docs/alicloud.md)
        * [IBM Cloud](/docs/ibmcloud.md)
    * Cloud
        * [DigitalOcean](/docs/digitalocean.md)
        * [Equinix Metal](/docs/equinixmetal.md)
        * [Fastly](/docs/fastly.md)
        * [Heroku](/docs/heroku.md)
        * [LaunchDarkly](/docs/launchdarkly.md)
        * [Linode](/docs/linode.md)
        * [NS1](/docs/ns1.md)
        * [OpenStack](/docs/openstack.md)
        * [TencentCloud](/docs/tencentcloud.md)
        * [Vultr](/docs/vultr.md)
        * [Yandex.Cloud](/docs/yandex.md)
    * Infrastructure Software
        * [Kubernetes](/docs/kubernetes.md)
        * [OctopusDeploy](/docs/octopus.md)
        * [RabbitMQ](/docs/rabbitmq.md)
    * Network
        * [Cloudflare](/docs/cloudflare.md)
        * [PAN-OS](/docs/panos.md)
    * VCS
        * [Azure DevOps](/docs/azuredevops.md)
        * [GitHub](/docs/github.md)
        * [Gitlab](/docs/gitlab.md)
    * Monitoring & System Management
        * [Datadog](/docs/datadog.md)
        * [New Relic](/docs/relic.md)
        * [Mackerel](/docs/mackerel.md)
        * [PagerDuty](/docs/pagerduty.md)
        * [Opsgenie](/docs/opsgenie.md)
    * Community
        * [Keycloak](/docs/keycloak.md)
        * [Logz.io](/docs/logz.md)
        * [Commercetools](/docs/commercetools.md)
        * [Mikrotik](/docs/mikrotik.md)
        * [Xen Orchestra](/docs/xen.md)
        * [GmailFilter](/docs/gmailfilter.md)
        * [Grafana](/docs/grafana.md)
        * [Vault](/docs/vault.md)
    * Identity
        * [Okta](/docs/okta.md)
        * [Auth0](/docs/auth0.md)
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
1.  Run `git clone <terraformer repo> && cd terraformer/`
2.  Run `go mod download`
3.  Run `go build -v` for all providers OR build with one provider
`go run build/main.go {google,aws,azure,kubernetes,etc}`
4.  Run ```terraform init``` against a ```versions.tf``` file to install the plugins required for your platform. For example, if you need plugins for the google provider, ```versions.tf``` should contain:

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
* Windows
1. Install Terraform - https://www.terraform.io/downloads
2. Download exe file for required provider from here - https://github.com/GoogleCloudPlatform/terraformer/releases
3. Add the exe file path to path variable
4. Create a folder and initialize the terraform provider and run terraformer commands from there
   * For AWS -  refer https://learn.hashicorp.com/tutorials/terraform/aws-build?in=terraform/aws-get-started



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
    * Heroku provider >2.2.1 - [here](https://releases.hashicorp.com/terraform-provider-heroku/)
    * LaunchDarkly provider >=2.1.1 - [here](https://releases.hashicorp.com/terraform-provider-launchdarkly/)
    * Linode provider >1.8.0 - [here](https://releases.hashicorp.com/terraform-provider-linode/)
    * OpenStack provider >1.21.1 - [here](https://releases.hashicorp.com/terraform-provider-openstack/)
    * TencentCloud provider >1.50.0 - [here](https://releases.hashicorp.com/terraform-provider-tencentcloud/)
    * Vultr provider >1.0.5 - [here](https://releases.hashicorp.com/terraform-provider-vultr/)
    * Yandex provider >0.42.0 - [here](https://releases.hashicorp.com/terraform-provider-yandex/)
* Infrastructure Software
    * Kubernetes provider >=1.9.0 - [here](https://releases.hashicorp.com/terraform-provider-kubernetes/)
    * RabbitMQ provider >=1.1.0 - [here](https://releases.hashicorp.com/terraform-provider-rabbitmq/)
* Network
    * Cloudflare provider >1.16 - [here](https://releases.hashicorp.com/terraform-provider-cloudflare/)
    * Fastly provider >0.16.1 - [here](https://releases.hashicorp.com/terraform-provider-fastly/)
    * NS1 provider >1.8.3 - [here](https://releases.hashicorp.com/terraform-provider-ns1/)
    * PAN-OS provider >= 1.8.3 - [here](https://github.com/PaloAltoNetworks/terraform-provider-panos)
* VCS
    * GitHub provider >=2.2.1 - [here](https://releases.hashicorp.com/terraform-provider-github/)
* Monitoring & System Management
    * Datadog provider >2.1.0 - [here](https://releases.hashicorp.com/terraform-provider-datadog/)
    * New Relic provider >2.0.0 - [here](https://releases.hashicorp.com/terraform-provider-newrelic/)
    * Mackerel provider > 0.0.6 - [here](https://github.com/mackerelio-labs/terraform-provider-mackerel)
    * Pagerduty >=1.9 - [here](https://releases.hashicorp.com/terraform-provider-pagerduty/)
    * Opsgenie >= 0.6.0 [here](https://releases.hashicorp.com/terraform-provider-opsgenie/)
* Community
    * Keycloak provider >=1.19.0 - [here](https://github.com/mrparkers/terraform-provider-keycloak/)
    * Logz.io provider >=1.1.1 - [here](https://github.com/jonboydell/logzio_terraform_provider/)
    * Commercetools provider >= 0.21.0 - [here](https://github.com/labd/terraform-provider-commercetools)
    * Mikrotik provider >= 0.2.2 - [here](https://github.com/ddelnano/terraform-provider-mikrotik)
    * Xen Orchestra provider >= 0.18.0 - [here](https://github.com/ddelnano/terraform-provider-xenorchestra)
    * GmailFilter provider >= 1.0.1 - [here](https://github.com/yamamoto-febc/terraform-provider-gmailfilter)
    * Vault provider - [here](https://github.com/hashicorp/terraform-provider-vault)
    * Auth0 provider - [here](https://github.com/alexkappa/terraform-provider-auth0)

Information on provider plugins:
https://www.terraform.io/docs/configuration/providers.html


## High-Level steps to add new provider
 * Initialize provider details in cmd/root.go and create a provider initialization file in the terraformer/cmd folder
 * Create a folder under terraformer/providers/ for your provider
 * Create two files under this folder
   * <provide_name>_provider.go
   * <provide_name>_service.go
* Initialize all provider's supported services in <provide_name>_provider.go file
* Create script for each supported service in same folder

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
