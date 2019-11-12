module github.com/GoogleCloudPlatform/terraformer

go 1.13

require (
	cloud.google.com/go v0.46.3
	cloud.google.com/go/logging v1.0.0
	cloud.google.com/go/storage v1.1.0
	github.com/Azure/azure-sdk-for-go v34.1.0+incompatible
	github.com/Azure/go-autorest/autorest/azure/auth v0.3.0
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20191010082856-e76f4c50e182
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.3+incompatible
	github.com/aws/aws-sdk-go v1.25.10
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cloudflare/cloudflare-go v0.10.4
	github.com/denverdino/aliyungo v0.0.0-20190924033012-e8e08fe22cb2
	github.com/digitalocean/godo v1.24.1
	github.com/dollarshaveclub/new-relic-synthetics-go v0.0.0-20170605224734-4dc3dd6ae884
	github.com/go-resty/resty v0.0.0-00010101000000-000000000000 // indirect
	github.com/gogo/protobuf v1.3.0 // indirect
	github.com/golang/groupcache v0.0.0-20191002201903-404acd9df4cc // indirect
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/google/go-github/v25 v25.1.3
	github.com/googleapis/gnostic v0.3.1 // indirect
	github.com/gophercloud/gophercloud v0.4.1-0.20190920074709-6e93a6ba3b09
	github.com/hashicorp/go-hclog v0.9.2
	github.com/hashicorp/go-plugin v1.0.1
	github.com/hashicorp/go-version v1.2.0 // indirect
	github.com/hashicorp/hcl v1.0.1-0.20190611123218-cf7d376da96d
	github.com/hashicorp/hcl2 v0.0.0-20191002203319-fb75b3253c80
	github.com/hashicorp/hil v0.0.0-20190212132231-97b3a9cdfa93 // indirect
	github.com/hashicorp/terraform v0.12.10
	github.com/hashicorp/yamux v0.0.0-20190923154419-df201c70410d // indirect
	github.com/heroku/heroku-go/v5 v5.1.0
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
	github.com/jonboydell/logzio_client v1.2.0
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	github.com/labd/commercetools-go-sdk v0.0.0-20190722144546-80b2ca71bd4d
	github.com/linode/linodego v0.12.0
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/paultyng/go-newrelic v3.1.0+incompatible
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/tomnomnom/linkheader v0.0.0-20180905144013-02ca5825eb80 // indirect
	github.com/ulikunitz/xz v0.5.6 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vultr/govultr v0.1.6
	github.com/zclconf/go-cty v1.1.0
	github.com/zorkian/go-datadog-api v2.24.0+incompatible
	go.opencensus.io v0.22.1 // indirect
	golang.org/x/crypto v0.0.0-20191010185427-af544f31c8ac // indirect
	golang.org/x/exp v0.0.0-20191002040644-a1355ae1e2c3 // indirect
	golang.org/x/lint v0.0.0-20190930215403-16217165b5de // indirect
	golang.org/x/net v0.0.0-20191009170851-d66e71096ffb // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/sys v0.0.0-20191010194322-b09406accb47 // indirect
	golang.org/x/text v0.3.2
	golang.org/x/tools v0.0.0-20191010201905-e5ffc44a6fee // indirect
	gonum.org/v1/gonum v0.6.0
	google.golang.org/api v0.11.0
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/genproto v0.0.0-20191009194640-548a555dbc03
	google.golang.org/grpc v1.24.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/resty.v0 v0.4.1 // indirect
	k8s.io/api v0.0.0-20191010143144-fbf594f18f80 // indirect
	k8s.io/apimachinery v0.0.0-20191006235458-f9f2f3f8ab02
	k8s.io/client-go v0.0.0-20191003000419-f68efa97b39e
	k8s.io/utils v0.0.0-20191010214722-8d271d903fe4 // indirect
)

// related to invalid pseudo-version: does not match version-control timestamp (2019-04-09T20:28:23Z)
replace golang.org/x/lint => golang.org/x/lint v0.0.0-20190409202823-959b441ac422

replace golang.org/x/time => golang.org/x/time v0.0.0-20190308202827-9d24e82272b4

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0

replace git.apache.org/thrift.git v0.0.0-20180902110319-2566ecd5d999 => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible

// https://stackoverflow.com/questions/55537287/unexpected-module-path-github-com-sirupsen-logrus
replace github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.0.6
