module github.com/GoogleCloudPlatform/terraformer

go 1.13

require (
	cloud.google.com/go v0.45.1
	cloud.google.com/go/logging v1.0.0
	github.com/aws/aws-sdk-go v1.22.0
	github.com/chanzuckerberg/terraform-provider-snowflake v0.4.1
	github.com/cloudflare/cloudflare-go v0.9.4
	github.com/coreos/go-systemd v0.0.0-20181031085051-9002847aa142 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/dollarshaveclub/new-relic-synthetics-go v0.0.0-20170605224734-4dc3dd6ae884
	github.com/golang/groupcache v0.0.0-20181024230925-c65c006176ff // indirect
	github.com/google/go-github/v25 v25.0.2
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gophercloud/gophercloud v0.0.0-20190427020117-60507118a582
	github.com/hashicorp/consul v1.4.0 // indirect
	github.com/hashicorp/go-hclog v0.9.2
	github.com/hashicorp/go-plugin v1.0.1
	github.com/hashicorp/hcl v1.0.1-0.20190611123218-cf7d376da96d
	github.com/hashicorp/hcl2 v0.0.0-20190821123243-0c888d1241f6
	github.com/hashicorp/terraform v0.12.10
	github.com/heroku/heroku-go/v5 v5.1.0
	github.com/howeyc/gopass v0.0.0-20170109162249-bf9dde6d0d2c // indirect
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/jonboydell/logzio_client v0.0.0-20190726085421-c93d6b149c1e
	github.com/joyent/triton-go v0.0.0-20180628001255-830d2b111e62 // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lusis/go-artifactory v0.0.0-20180304164534-a47f63f234b2 // indirect
	github.com/miekg/dns v1.0.15 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/packer-community/winrmcp v0.0.0-20180921211025-c76d91c1e7db // indirect
	github.com/paultyng/go-newrelic v3.1.0+incompatible
	github.com/pkg/errors v0.8.1
	github.com/snowflakedb/gosnowflake v1.3.1
	github.com/spf13/cobra v0.0.3
	github.com/spf13/pflag v1.0.3
	github.com/terraform-providers/terraform-provider-openstack v1.18.0 // indirect
	github.com/tomnomnom/linkheader v0.0.0-20180905144013-02ca5825eb80 // indirect
	github.com/ugorji/go v1.1.2-0.20180728093225-eeb0478a81ae // indirect
	github.com/unknwon/com v0.0.0-20181010210213-41959bdd855f // indirect
	github.com/xlab/treeprint v0.0.0-20181112141820-a009c3971eca // indirect
	github.com/zclconf/go-cty v1.1.0
	github.com/zorkian/go-datadog-api v2.24.0+incompatible
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.9.0
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
	gopkg.in/resty.v0 v0.4.1 // indirect
	k8s.io/api v0.0.0-20190116205037-c89978d5f86d // indirect
	k8s.io/apimachinery v0.0.0-20190116203031-d49e237a2683
	k8s.io/client-go v7.0.0+incompatible
	k8s.io/kubectl v0.0.0-20190502165022-ce8d9f55c93c
)

// related to invalid pseudo-version: does not match version-control timestamp (2019-04-09T20:28:23Z)
replace golang.org/x/lint => golang.org/x/lint v0.0.0-20190409202823-959b441ac422

replace golang.org/x/time => golang.org/x/time v0.0.0-20190308202827-9d24e82272b4

replace git.apache.org/thrift.git v0.0.0-20180902110319-2566ecd5d999 => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
