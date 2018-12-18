#Terraformer

CLI tool for generate tf + tfstate files from current configuration(reverse terraform)

This is not an official Google product


###Install
1. git clone project OR binary
2. Run `GO111MODULE=on go mod vendor`
3. go build -v
4. Copy your terraform provider plugin to ~/.terraform.d/plugins/{darwin,linux}_amd64 => https://www.terraform.io/docs/configuration/providers.html


#####Usage:
#####For GCP:
GOOGLE_CLOUD_PROJECT=YOUR_PROJECT ./terraformer google YOUR_SERVICE YOUR_ZONE
Examples: 

````
./terraformer google firewalls europe-west1-c
````
````
./terraformer google gcs europe-west1-c
````

````
./terraformer google addresses europe-west1-c
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
alerts
dns
````



Your tf and tfstate file generate to `generated/gcp/zone/service`

#####For AWS:
./terraformer aws YOUR_SERVICE YOUR_REGION


````
./terraformer aws sg eu-west1
````
````
./terraformer aws s3 eu-west1
````
````
./terraformer aws subnet eu-west1
````
List of support AWS services:
````
elb
autoscaling
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
````

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
