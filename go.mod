module github.com/GoogleCloudPlatform/terraformer

go 1.15

require (
	cloud.google.com/go v0.74.0
	cloud.google.com/go/logging v1.1.2
	cloud.google.com/go/storage v1.12.0
	github.com/Azure/azure-sdk-for-go v42.0.0+incompatible
	github.com/Azure/azure-storage-blob-go v0.10.0
	github.com/Azure/go-autorest/autorest v0.11.12
	github.com/DataDog/datadog-api-client-go v1.0.0-beta.9
	github.com/IBM-Cloud/bluemix-go v0.0.0-20201210085054-cdf09378fdd9
	github.com/IBM/go-sdk-core/v3 v3.3.1
	github.com/IBM/go-sdk-core/v4 v4.9.0
	github.com/IBM/ibm-cos-sdk-go v1.5.0
	github.com/IBM/keyprotect-go-client v0.5.2
	github.com/IBM/networking-go-sdk v0.12.1
	github.com/IBM/vpc-go-sdk v0.3.1
	github.com/OctopusDeploy/go-octopusdeploy v1.6.0
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.865
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.2+incompatible
	github.com/aws/aws-sdk-go v1.36.19
	github.com/aws/aws-sdk-go-v2 v0.24.0
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cloudflare/cloudflare-go v0.13.6
	github.com/ddelnano/terraform-provider-mikrotik v0.0.0-20200501162830-a217572b326c
	github.com/denverdino/aliyungo v0.0.0-20200327235253-d59c209c7e93
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/digitalocean/godo v1.35.1
	github.com/dollarshaveclub/new-relic-synthetics-go v0.0.0-20170605224734-4dc3dd6ae884
	github.com/fastly/go-fastly v1.18.0
	github.com/google/go-github/v25 v25.1.3
	github.com/gophercloud/gophercloud v0.13.0
	github.com/hashicorp/go-azure-helpers v0.10.0
	github.com/hashicorp/go-hclog v0.14.1
	github.com/hashicorp/go-plugin v1.4.0
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/terraform v0.12.29
	github.com/heroku/heroku-go/v5 v5.1.0
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/jmespath/go-jmespath v0.4.0
	github.com/jonboydell/logzio_client v1.2.0
	github.com/labd/commercetools-go-sdk v0.0.0-20200309143931-ca72e918a79d
	github.com/linode/linodego v0.24.1
	github.com/mrparkers/terraform-provider-keycloak v0.0.0-20200506151941-509881368409
	github.com/ns1/ns1-go v2.4.0+incompatible
	github.com/paultyng/go-newrelic/v4 v4.10.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/vultr/govultr v0.5.0
	github.com/yandex-cloud/go-genproto v0.0.0-20200722140432-762fe965ce77
	github.com/yandex-cloud/go-sdk v0.0.0-20200722140627-2194e5077f13
	github.com/zclconf/go-cty v1.7.0
	github.com/zorkian/go-datadog-api v2.30.0+incompatible
	golang.org/x/oauth2 v0.0.0-20201208152858-08078c50e5b5
	golang.org/x/text v0.3.4
	gonum.org/v1/gonum v0.7.0
	google.golang.org/api v0.36.0
	google.golang.org/genproto v0.0.0-20201210142538-e3217bee35cc
	gopkg.in/jarcoal/httpmock.v1 v1.0.0-00010101000000-000000000000 // indirect
	k8s.io/apimachinery v0.20.1
	k8s.io/client-go v0.20.1
)

replace gopkg.in/jarcoal/httpmock.v1 => github.com/jarcoal/httpmock v1.0.5
