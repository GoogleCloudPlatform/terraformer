# Terraformer

CLI tool to generate `tf` and `tfstate` files from existing infrastructure
(reverse Terraform).

*   Disclaimer: This is not an official Google product.
*   Status: beta - need improve documentations, bugs etc..
*   Created by: Created by Waze SRE.

![Waze SRE logo](docs/waze-sre-logo.png)

## Capabilities

1.  Generate `tf` + `tfstate` files from existing infrastructure for all
    supported objects by resource.
2.  Remote state can be uploaded to a GCS bucket.
3.  Connect between resources with `terraform_remote_state` (local and bucket).
4.  Compatible with terraform 0.12 syntax.
5.  Save `tf` files with custom folder tree pattern.

Terraformer use terraform providers and built for easy to add new supported resources. 
For upgrade resources with new fields you need upgrade only terraform providers. 

```
Import current state to terraform configuration from google cloud

Usage:
   import google [flags]

Flags:
  -b, --bucket string        gs://terraform-state
  -c, --connect               (default true)
  -h, --help                 help for google
  -o, --path-output string    (default "generated")
  -p, --path-patter string   {output}/{provider}/custom/{service}/ (default "{output}/{provider}/{service}/")
  -r, --resources strings    firewalls,networks
  -s, --state string         local or bucket (default "local")
      --projects strings
  -z, --zone string
```

## Supported providers

1.  Google cloud
2.  AWS

#### Permissions

Readonly permissions

## How to use Terraformer

### Installation

1.  Run `git clone <terraformer repo>`
2.  Run `GO111MODULE=on go mod vendor`
3.  Run `go build -v`
4.  Copy your Terraform provider's plugin(s) to
    `~/.terraform.d/plugins/{darwin,linux}_amd64`, as appropriate.
    
Links for download terraform providers:
* google cloud provider - [here](https://releases.hashicorp.com/terraform-provider-google/)
* aws provider - [here](https://releases.hashicorp.com/terraform-provider-aws/)

Information on provider plugins:
https://www.terraform.io/docs/configuration/providers.html

### Use with GCP
[![asciicast](https://asciinema.org/a/243953.svg)](https://asciinema.org/a/243953)
Example:

```
terraformer import google --resources=gcs,forwardingRules,httpHealthChecks --connect=true --zone=europe-west1-a --projects=aaa,fff
```

List of supported GCP services:

*   `addresses`
*   `autoscalers`
*   `backendBuckets`
*   `backendServices`
*   `disks`
*   `firewalls`
*   `forwardingRules`
*   `globalAddresses`
*   `globalForwardingRules`
*   `healthChecks`
*   `httpHealthChecks`
*   `httpsHealthChecks`
*   `images` (bug: Either raw_disk or source_disk configuration is required.)
*   `instanceGroupManagers`
*   `instanceGroups`
*   `instanceTemplates`
*   `instances`
*   `networks`
*   `regionAutoscalers`
*   `regionBackendServices`
*   `regionDisks`
*   `regionInstanceGroupManagers`
*   `routers`
*   `routes`
*   `securityPolicies`
*   `sslPolicies`
*   `subnetworks`
*   `targetHttpProxies` (bug with proxy_id uint64 issue)
*   `targetHttpsProxies`
*   `targetSslProxies`
*   `targetTcpProxies`
*   `urlMaps`
*   `vpnTunnels`
*   `gcs`
*   `monitoring`
*   `dns`
*   `cloudsql` ([bug](https://github.com/terraform-providers/terraform-provider-google/issues/2716), [bug](https://github.com/GoogleCloudPlatform/magic-modules/pull/1097))

Your `tf` and `tfstate` files are written by default to
`generated/gcp/zone/service`.

### Use with AWS

Example:

```
 terraformer import aws --resources=vpc,subnet --connect=true --regions=eu-west-1
```

```
Import current State to terraform configuration from aws

Usage:
   import aws [flags]

Flags:
  -b, --bucket string        gs://terraform-state
  -c, --connect               (default true)
  -h, --help                 help for aws
  -o, --path-output string    (default "generated")
  -p, --path-patter string   {output}/{provider}/custom/{service}/ (default "{output}/{provider}/{service}/")
      --regions strings      eu-west-1,eu-west-2,us-east-1
  -r, --resources strings    vpc,subnet,nacl
  -s, --state string         local or bucket (default "local")
```

List of support AWS services:

*   `elb`
*   `alb`
*   `auto_scaling`
*   `rds`
*   `iam`
*   `igw`
*   `nacl`
*   `s3`
*   `sg`
*   `subnet`
*   `vpc`
*   `vpn_connection`
*   `vpn_gateway`
*   `route53`
*   `elasticache`

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
`gcp_terraforming/gcp_compute_code_generator`.

To regenerate code:

```
go run gcp_terraforming/gcp_compute_code_generator/*.go
```

### Similar projects

1.  https://github.com/dtan4/terraforming
