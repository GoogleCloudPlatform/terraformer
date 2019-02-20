#Terraformer
Disclaimer: This is not an official Google product.

CLI tool for generate tf + tfstate files from current infrastructure(reverse terraform)

####Created by
# ![Waze SRE](docs/waze-sre-logo.png)


#Options
1. Generate tf + tfstate files from current infrastructure for all objects by resource
2. Remote state can be uploaded to bucket(only gcs support)
3. Connect between resource with terraform_remote_state(local and bucket)
4. Compatible with terraform 0.12 syntax
5. Save tf files with custom folder tree patter

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

##Providers support
1. Google cloud
2. AWS

###Install
1. git clone project OR binary
2. Run `GO111MODULE=on go mod vendor`
3. go build -v
4. Copy your terraform provider plugin to ~/.terraform.d/plugins/{darwin,linux}_amd64 => https://www.terraform.io/docs/configuration/providers.html


#####Usage:
#####For GCP:
GOOGLE_CLOUD_PROJECT=YOUR_PROJECT ./terraformer import google --resources=networks,firewalls --connect=true --zone=europe-west1-a

Examples: 

````
./terraformer import google --resources=gcs,forwardingRules,httpHealthChecks --connect=true --zone=europe-west1-a --projects=aaa,fff
````

List of support GCP services:
````
addresses
autoscalers
backendBuckets 
backendServices
disks
firewalls
forwardingRules
globalAddresses
globalForwardingRules
healthChecks
httpHealthChecks
httpsHealthChecks
images -  bug => Either raw_disk or source_disk configuration is required.
instanceGroupManagers
instanceGroups
instanceTemplates
instances
networks
regionAutoscalers
regionBackendServices
regionDisks
regionInstanceGroupManagers
routers
routes
securityPolicies 
sslPolicies
subnetworks
targetHttpProxies - bug with proxy_id uint64 issue
targetHttpsProxies
targetSslProxies
targetTcpProxies
urlMaps
vpnTunnels
gcs
monitoring
dns
cloudsql - bug(https://github.com/terraform-providers/terraform-provider-google/issues/2716), bug(https://github.com/GoogleCloudPlatform/magic-modules/pull/1097)
````



Your tf and tfstate files generate by default to `generated/gcp/zone/service`

#####For AWS:
````
 ./terraformer import aws --resources=vpc,subnet --connect=true --regions=eu-west-1
````
````
 ./terraformer import aws --resources=vpc,subnet --connect=false --regions=eu-west-1,eu-west-2
````
````
 ./terraformer import aws --resources=vpc,subnet --connect=true --regions=eu-west-1
````
````
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
````
List of support AWS services:
````
elb
alb
auto_scaling
rds
iam
igw
nacl
s3
sg
subnet
vpc
vpn_connection
vpn_gateway
route53
elasticache
````
 
#Contributing
If you have improvements or fixes, we would love to have your contributions. Please read CONTRIBUTING.md for more information on the process we would like contributors to follow.

###Developing:
Process for generate tf + tfstate files
1. Call GCP/AWS/other api and get list of resources
2. Iterate on resources and take only ID(not need mapping fields!!!)
3. Call to provider for readonly fields
4. Call to infrastructure and take tf + tfstate.


#####Infrastructure
1. Call to provider for refresh method and get all data
2. Convert refresh data to go struct
3. Generate HCL file - tf files.
4. Generate tfstate files

All mapping of resource make by providers and terraform. Any upgrade need only for providers.
 
#####GCP compute resources
For GCP compute resources use generate code from `gcp_terraforming/compute_resources/gcp_compute_code_generator`

Regenerate code 
````
go run gcp_terraforming/compute_resources/gcp_compute_code_generator/*.go
````


###Similar projects
1. https://github.com/dtan4/terraforming