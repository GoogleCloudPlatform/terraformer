module github.com/GoogleCloudPlatform/terraformer

go 1.13

require (
	cloud.google.com/go v0.50.0
	cloud.google.com/go/logging v1.0.0
	cloud.google.com/go/storage v1.4.0
	github.com/Azure/azure-sdk-for-go v37.2.0+incompatible
	github.com/Azure/go-autorest/autorest v0.9.3
	github.com/Azure/go-autorest/autorest/adal v0.8.1 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.3.1 // indirect
	github.com/OctopusDeploy/go-octopusdeploy v1.6.0
	github.com/ajg/form v1.5.1 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.60.295
	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20190307165228-86c17b95fcd5 // indirect
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.3+incompatible
	github.com/aws/aws-sdk-go v1.26.5 // indirect

	// shouldn't be upgraded without check of fix to https://github.com/aws/aws-sdk-go-v2/issues/492
	github.com/aws/aws-sdk-go-v2 v0.19.0
	github.com/bmatcuk/doublestar v1.2.2 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cloudflare/cloudflare-go v0.11.0
	github.com/denverdino/aliyungo v0.0.0-20191217032211-d5785803c365
	github.com/dghubble/sling v1.1.0
	github.com/digitalocean/godo v1.29.0
	github.com/dollarshaveclub/new-relic-synthetics-go v0.0.0-20170605224734-4dc3dd6ae884
	github.com/fastly/go-fastly v1.3.0
	github.com/go-resty/resty v0.0.0-00010101000000-000000000000 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20191027212112-611e8accdfc9 // indirect
	github.com/google/go-github/v25 v25.1.3
	github.com/google/jsonapi v0.0.0-20181016150055-d0428f63eb51 // indirect
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/gophercloud/gophercloud v0.7.0
	github.com/hashicorp/go-azure-helpers v0.10.0
	github.com/hashicorp/go-hclog v0.10.1
	github.com/hashicorp/go-plugin v1.0.1
	github.com/hashicorp/go-retryablehttp v0.6.4 // indirect
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/hcl/v2 v2.2.0
	github.com/hashicorp/hcl2 v0.0.0-20191002203319-fb75b3253c80 // indirect
	github.com/hashicorp/terraform v0.12.18
	github.com/hashicorp/terraform-svchost v0.0.0-20191119180714-d2e4933b9136 // indirect
	github.com/hashicorp/yamux v0.0.0-20190923154419-df201c70410d // indirect
	github.com/heroku/heroku-go/v5 v5.1.0
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
	github.com/jonboydell/logzio_client v1.2.0
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	github.com/labd/commercetools-go-sdk v0.0.0-20190722144546-80b2ca71bd4d
	github.com/linode/linodego v0.12.2
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/mrparkers/terraform-provider-keycloak v0.0.0-20191218161228-a467c7185cbc
	github.com/paultyng/go-newrelic v3.1.0+incompatible // indirect
	github.com/paultyng/go-newrelic/v4 v4.10.0
	github.com/pkg/errors v0.8.1
	github.com/posener/complete v1.2.3 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/ulikunitz/xz v0.5.6 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vultr/govultr v0.1.7
	github.com/zclconf/go-cty v1.2.1
	github.com/zorkian/go-datadog-api v2.25.0+incompatible
	go.opencensus.io v0.22.2 // indirect
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6
	golang.org/x/sys v0.0.0-20191218084908-4a24b4065292 // indirect
	golang.org/x/text v0.3.2
	golang.org/x/tools v0.0.0-20191218225520-84f0c7cf60ea // indirect
	gonum.org/v1/gonum v0.6.1
	google.golang.org/api v0.15.0
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/genproto v0.0.0-20191216205247-b31c10ee225f
	google.golang.org/grpc v1.26.0 // indirect
	gopkg.in/ini.v1 v1.51.0 // indirect
	gopkg.in/resty.v0 v0.4.1 // indirect
	k8s.io/api v0.17.0 // indirect
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.0.0-20191003000419-f68efa97b39e
	k8s.io/utils v0.0.0-20191218082557-f07c713de883 // indirect
)

// related to invalid pseudo-version: does not match version-control timestamp (2019-04-09T20:28:23Z)
replace golang.org/x/lint => golang.org/x/lint v0.0.0-20190409202823-959b441ac422

replace golang.org/x/time => golang.org/x/time v0.0.0-20190308202827-9d24e82272b4

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0

replace git.apache.org/thrift.git v0.0.0-20180902110319-2566ecd5d999 => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible

// https://stackoverflow.com/questions/55537287/unexpected-module-path-github-com-sirupsen-logrus
replace github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.0.6
