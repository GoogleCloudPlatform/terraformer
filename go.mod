module github.com/GoogleCloudPlatform/terraformer

go 1.14

require (
	cloud.google.com/go v0.56.0
	cloud.google.com/go/logging v1.0.0
	cloud.google.com/go/storage v1.6.0
	github.com/Azure/azure-sdk-for-go v42.0.0+incompatible
	github.com/Azure/azure-storage-blob-go v0.10.0
	github.com/Azure/go-autorest/autorest v0.10.0
	github.com/OctopusDeploy/go-octopusdeploy v1.6.0
	github.com/aliyun/alibaba-cloud-sdk-go v1.60.295
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.2+incompatible
	github.com/aws/aws-sdk-go v1.30.19
	github.com/aws/aws-sdk-go-v2 v0.22.0
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cloudflare/cloudflare-go v0.11.7
	github.com/ddelnano/terraform-provider-mikrotik v0.0.0-20200501162830-a217572b326c
	github.com/denverdino/aliyungo v0.0.0-20200327235253-d59c209c7e93
	github.com/digitalocean/godo v1.35.1
	github.com/dollarshaveclub/new-relic-synthetics-go v0.0.0-20170605224734-4dc3dd6ae884
	github.com/fastly/go-fastly v1.15.0
	github.com/google/go-github/v25 v25.1.3
	github.com/gophercloud/gophercloud v0.10.0
	github.com/hashicorp/go-azure-helpers v0.10.0
	github.com/hashicorp/go-hclog v0.12.2
	github.com/hashicorp/go-plugin v1.3.0
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/terraform v0.12.28
	github.com/heroku/heroku-go/v5 v5.1.0
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/jmespath/go-jmespath v0.3.0
	github.com/jonboydell/logzio_client v1.2.0
	github.com/labd/commercetools-go-sdk v0.0.0-20200309143931-ca72e918a79d
	github.com/linode/linodego v0.15.0
	github.com/mrparkers/terraform-provider-keycloak v0.0.0-20200506151941-509881368409
	github.com/ns1/ns1-go v2.4.0+incompatible
	github.com/paultyng/go-newrelic/v4 v4.10.0
	github.com/pkg/errors v0.9.1
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/vultr/govultr v0.4.0
	github.com/zclconf/go-cty v1.4.0
	github.com/zorkian/go-datadog-api v2.29.0+incompatible
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/text v0.3.2
	gonum.org/v1/gonum v0.7.0
	google.golang.org/api v0.22.0
	google.golang.org/genproto v0.0.0-20200430143042-b979b6f78d84
	gopkg.in/jarcoal/httpmock.v1 v1.0.0-00010101000000-000000000000 // indirect
	k8s.io/apimachinery v0.17.5
	k8s.io/client-go v0.17.5
	k8s.io/utils v0.0.0-20191218082557-f07c713de883 // indirect
)

replace gopkg.in/jarcoal/httpmock.v1 => github.com/jarcoal/httpmock v1.0.5
