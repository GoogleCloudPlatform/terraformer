module github.com/GoogleCloudPlatform/terraformer

go 1.18

require (
	cloud.google.com/go v0.100.2 // indirect
	cloud.google.com/go/logging v1.1.2
	cloud.google.com/go/storage v1.14.0
	github.com/Azure/azure-sdk-for-go v63.1.0+incompatible
	github.com/Azure/azure-storage-blob-go v0.10.0
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Azure/go-autorest/autorest v0.11.23
	github.com/DataDog/datadog-api-client-go v1.13.0
	github.com/IBM-Cloud/bluemix-go v0.0.0-20210203095940-db28d5e07b55
	github.com/IBM/go-sdk-core/v3 v3.3.1
	github.com/IBM/go-sdk-core/v4 v4.9.0
	github.com/IBM/ibm-cos-sdk-go v1.5.0
	github.com/IBM/keyprotect-go-client v0.6.0
	github.com/IBM/networking-go-sdk v0.28.0
	github.com/IBM/platform-services-go-sdk v0.19.1
	github.com/IBM/vpc-go-sdk v0.4.1
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/OctopusDeploy/go-octopusdeploy v1.6.0
	github.com/PaloAltoNetworks/pango v0.6.0
	github.com/SAP/go-hdb v0.105.2 // indirect
	github.com/SermoDigital/jose v0.9.1 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1247
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.2+incompatible
	github.com/apache/openwhisk-client-go v0.0.0-20210106144548-17d556327cd3
	github.com/aws/aws-sdk-go-v2 v1.16.3
	github.com/aws/aws-sdk-go-v2/config v1.1.4
	github.com/aws/aws-sdk-go-v2/credentials v1.1.4
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.2.0
	github.com/aws/aws-sdk-go-v2/service/acm v1.2.1
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.2.1
	github.com/aws/aws-sdk-go-v2/service/appsync v1.2.1
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.3.1
	github.com/aws/aws-sdk-go-v2/service/batch v1.3.1
	github.com/aws/aws-sdk-go-v2/service/budgets v1.9.0
	github.com/aws/aws-sdk-go-v2/service/cloud9 v1.1.3
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.3.0
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.3.0
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.1.3
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.2.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.3.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.2.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.2.3
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.2.1
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.1.3
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.2.1
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.13.3
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.2.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.2.1
	github.com/aws/aws-sdk-go-v2/service/configservice v1.3.0
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.1.3
	github.com/aws/aws-sdk-go-v2/service/devicefarm v1.13.3
	github.com/aws/aws-sdk-go-v2/service/docdb v1.15.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.2.1
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.3.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.2.1
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.2.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.2.1
	github.com/aws/aws-sdk-go-v2/service/efs v1.2.1
	github.com/aws/aws-sdk-go-v2/service/eks v1.2.1
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.2.1
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.2.1
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.2.1
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.2.1
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.15.4
	github.com/aws/aws-sdk-go-v2/service/emr v1.2.1
	github.com/aws/aws-sdk-go-v2/service/firehose v1.2.1
	github.com/aws/aws-sdk-go-v2/service/glue v1.3.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.3.0
	github.com/aws/aws-sdk-go-v2/service/iot v1.2.0
	github.com/aws/aws-sdk-go-v2/service/kafka v1.14.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.2.1
	github.com/aws/aws-sdk-go-v2/service/kms v1.2.2
	github.com/aws/aws-sdk-go-v2/service/lambda v1.2.1
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.15.3
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.1.4
	github.com/aws/aws-sdk-go-v2/service/opsworks v1.2.2
	github.com/aws/aws-sdk-go-v2/service/organizations v1.2.1
	github.com/aws/aws-sdk-go-v2/service/qldb v1.1.3
	github.com/aws/aws-sdk-go-v2/service/rds v1.18.1
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.2.1
	github.com/aws/aws-sdk-go-v2/service/route53 v1.3.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.4.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.2.1
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.2.1
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.2.1
	github.com/aws/aws-sdk-go-v2/service/ses v1.2.1
	github.com/aws/aws-sdk-go-v2/service/sfn v1.2.1
	github.com/aws/aws-sdk-go-v2/service/sns v1.2.1
	github.com/aws/aws-sdk-go-v2/service/sqs v1.3.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.3.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.2.1
	github.com/aws/aws-sdk-go-v2/service/swf v1.10.0
	github.com/aws/aws-sdk-go-v2/service/waf v1.1.4
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.12.3
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.18.0
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.2.1
	github.com/aws/aws-sdk-go-v2/service/xray v1.2.1
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cloudflare/cloudflare-go v0.13.6
	github.com/cloudfoundry/jibber_jabber v0.0.0-20151120183258-bcc4c8345a21 // indirect
	github.com/containerd/continuity v0.1.0 // indirect
	github.com/ddelnano/terraform-provider-mikrotik/client v0.0.0-20210401060029-7f652169b2c4
	github.com/ddelnano/terraform-provider-xenorchestra/client v0.0.0-20210401070256-0d721c6762ef
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/denverdino/aliyungo v0.0.0-20200327235253-d59c209c7e93
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/digitalocean/godo v1.57.0
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/duosecurity/duo_api_golang v0.0.0-20201112143038-0e07e9f869e3 // indirect
	github.com/elazarl/go-bindata-assetfs v1.0.1 // indirect
	github.com/fastly/go-fastly/v6 v6.0.1
	github.com/fatih/structs v1.1.0 // indirect
	github.com/gocql/gocql v0.0.0-20210707082121-9a3953d1826d // indirect
	github.com/google/go-github/v35 v35.1.0
	github.com/gophercloud/gophercloud v0.24.0
	github.com/grafana/grafana-api-golang-client v0.0.0-20210218192924-9ccd2365d2a6
	github.com/hashicorp/go-azure-helpers v0.10.0
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/hashicorp/go-hclog v0.16.2
	github.com/hashicorp/go-memdb v1.3.2 // indirect
	github.com/hashicorp/go-plugin v1.4.1
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/terraform v0.12.31
	github.com/hashicorp/vault v0.10.4
	github.com/heimweh/go-pagerduty v0.0.0-20210930203304-530eff2acdc6
	github.com/heroku/heroku-go/v5 v5.4.1
	github.com/hokaccha/go-prettyjson v0.0.0-20210113012101-fb4e108d2519 // indirect
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/jefferai/jsonx v1.0.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0
	github.com/jonboydell/logzio_client v1.2.0
	github.com/labd/commercetools-go-sdk v0.3.1
	github.com/linode/linodego v0.24.1
	github.com/microsoft/azure-devops-go-api/azuredevops v1.0.0-b5
	github.com/mrparkers/terraform-provider-keycloak v0.0.0-20200506151941-509881368409
	github.com/nicksnyder/go-i18n v1.10.1 // indirect
	github.com/ns1/ns1-go v2.4.0+incompatible
	github.com/okta/okta-sdk-golang/v2 v2.9.2
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v1.0.1 // indirect
	github.com/opsgenie/opsgenie-go-sdk-v2 v1.2.9
	github.com/ory/dockertest v3.3.5+incompatible // indirect
	github.com/packethost/packngo v0.9.0
	github.com/pkg/errors v0.9.1
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/tencentcloud/tencentcloud-sdk-go v3.0.82+incompatible
	github.com/tencentyun/cos-go-sdk-v5 v0.7.34
	github.com/vultr/govultr v0.5.0
	github.com/xanzy/go-gitlab v0.50.2
	github.com/yandex-cloud/go-genproto v0.0.0-20220314102905-1acaee8ca7eb
	github.com/yandex-cloud/go-sdk v0.0.0-20220314105123-d0c2a928feb6
	github.com/zclconf/go-cty v1.8.4
	github.com/zorkian/go-datadog-api v2.30.0+incompatible
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	golang.org/x/text v0.3.7
	gonum.org/v1/gonum v0.7.0
	google.golang.org/api v0.70.0
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106
	gopkg.in/jarcoal/httpmock.v1 v1.0.0-00010101000000-000000000000 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	k8s.io/apimachinery v0.21.0
	k8s.io/client-go v0.21.0
)

require (
	github.com/IBM-Cloud/container-services-go-sdk v0.0.0-20210705152127-41ca00fc9a62
	github.com/IBM/go-sdk-core v1.1.0
	github.com/aws/aws-sdk-go-v2/internal/ini v1.2.2 // indirect
	github.com/hashicorp/terraform-svchost v0.0.0-20200729002733-f050f53b9734 // indirect
	github.com/mackerelio/mackerel-client-go v0.19.0
	github.com/okta/terraform-provider-okta v0.0.0-20210924173942-a5a664459d3b
	github.com/zclconf/go-cty-yaml v1.0.2 // indirect
)

require (
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/antihax/optional v1.0.0 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/launchdarkly/api-client-go v5.3.0+incompatible
)

require github.com/newrelic/newrelic-client-go v0.71.0

require (
	github.com/Azure/azure-pipeline-go v0.2.2 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.14 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.2.0 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/IBM/go-sdk-core/v5 v5.8.0 // indirect
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/agext/levenshtein v1.2.2 // indirect
	github.com/apparentlymart/go-cidr v1.1.0 // indirect
	github.com/apparentlymart/go-textseg/v12 v12.0.0 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/appscode/go-querystring v0.0.0-20170504095604-0126cfb3f1dc // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
	github.com/aws/aws-sdk-go v1.37.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.0.5 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.0.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.2.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshift v1.10.0
	github.com/aws/aws-sdk-go-v2/service/sso v1.1.4 // indirect
	github.com/aws/smithy-go v1.11.2 // indirect
	github.com/beevik/etree v1.1.0 // indirect
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
	github.com/bgentry/speakeasy v0.1.0 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/bmatcuk/doublestar v1.1.5 // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/crewjam/saml v0.4.5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dghubble/sling v1.1.0 // indirect
	github.com/dimchansky/utfbom v1.1.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/fatih/color v1.7.0 // indirect
	github.com/form3tech-oss/jwt-go v3.2.2+incompatible // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logr/logr v0.4.0 // indirect
	github.com/go-openapi/errors v0.19.8 // indirect
	github.com/go-openapi/strfmt v0.20.2 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-resty/resty/v2 v2.1.1-0.20191201195748-d7b97669fe48 // indirect
	github.com/go-routeros/routeros v0.0.0-20210123142807-2a44d57c6730 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/jsonapi v1.0.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/googleapis/gax-go/v2 v2.1.1 // indirect
	github.com/googleapis/gnostic v0.4.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-getter v1.5.3 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.0 // indirect
	github.com/hashicorp/go-rootcerts v1.0.0 // indirect
	github.com/hashicorp/go-safetemp v1.0.0 // indirect
	github.com/hashicorp/go-sockaddr v1.0.0 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/go-version v1.3.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl/v2 v2.8.2 // indirect
	github.com/hashicorp/hil v0.0.0-20190212112733-ab17b08d6590 // indirect
	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jonboulle/clockwork v0.2.1 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/klauspost/compress v1.11.2 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattermost/xml-roundtrip-validator v0.0.0-20201213122252-bcd7e1b9601e // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-ieproxy v0.0.0-20190702010315-6dee0af9227d // indirect
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/mitchellh/cli v1.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/go-testing-interface v1.0.4 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/hashstructure v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/mozillazg/go-httpheader v0.2.1 // indirect
	github.com/oklog/run v1.0.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/patrickmn/go-cache v0.0.0-20180815053127-5633e0862627 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/peterhellberg/link v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/posener/complete v1.2.1 // indirect
	github.com/russellhaering/goxmldsig v1.1.0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/sourcegraph/jsonrpc2 v0.0.0-20210201082850-366fbb520750 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/tomnomnom/linkheader v0.0.0-20180905144013-02ca5825eb80 // indirect
	github.com/ulikunitz/xz v0.5.8 // indirect
	github.com/vmihailenco/msgpack/v4 v4.3.12 // indirect
	github.com/vmihailenco/tagparser v0.1.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.7.2 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220315194320-039c03cc5b86 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	golang.org/x/tools v0.1.8 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/auth0.v5 v5.21.1
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	k8s.io/api v0.21.0 // indirect
	k8s.io/klog/v2 v2.8.0 // indirect
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.1.0 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

require github.com/PuerkitoBio/rehttp v1.0.0 // indirect

require (
	cloud.google.com/go/cloudtasks v1.3.0
	cloud.google.com/go/iam v0.3.0
	cloud.google.com/go/monitoring v1.4.0
	github.com/manicminer/hamilton v0.43.0
)

require (
	cloud.google.com/go/compute v1.3.0 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/golang-jwt/jwt/v4 v4.3.0 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aa v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aai v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/advisor v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/af v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/afc v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ame v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ams v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/antiddos v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/apcas v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ape v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/api v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/apigateway v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/apm v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/as v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asw v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ba v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/batch v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bda v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bizlive v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bm v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bmeip v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bmlb v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bmvpc v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bri v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/bsca v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/btoe v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cam v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/captcha v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/car v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/casb v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cat v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ccc v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cds v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfg v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfs v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfw v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/chdfs v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cii v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cim v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cis v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ckafka v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudaudit v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cloudhsm v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cls v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cme v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cmq v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cms v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cpdp v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cr v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cwp v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cws v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cynosdb v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dayu v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dbbrain v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dc v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dcdb v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dlc v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/domain v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/drm v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ds v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dtf v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dts v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/eb v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecc v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecdn v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ecm v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/eiam v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/eis v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/emr v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/es v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ess v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/essbasic v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/facefusion v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/faceid v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/fmu v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ft v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gaap v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gme v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gpm v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gs v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gse v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/habo v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/hcm v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iai v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ic v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/icr v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ie v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iecp v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iir v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ims v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iot v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotcloud v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotexplorer v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iottid v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotvideo v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/iotvideoindustry v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ivld v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/kms v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lighthouse v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/live v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lowcode v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/lp v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mariadb v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/market v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/memcached v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mgobe v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mna v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mongodb v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/monitor v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mps v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mrs v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ms v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/msp v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mvj v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/nlp v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/npp v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/oceanus v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ocr v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/organization v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/partners v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/pds v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/postgres v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/privatedns v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rce v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rkp v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rp v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/rum v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/scf v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ses v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/smh v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/smpn v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/soe v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/solar v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sqlserver v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssa v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sslpod v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssm v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/taf v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tat v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tav v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tbaas v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tbm v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tbp v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcaplusdb v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcb v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcex v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tci v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcm v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcr v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcss v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdcpg v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdid v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tdmq v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tem v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/teo v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/thpc v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tia v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tic v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ticm v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tics v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tiems v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tiia v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tione v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tiw v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tkgdq v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tms v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/trtc v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tse v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tsf v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tsw v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tts v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ump v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vm v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vms v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vod v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/waf v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/wav v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/wss v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/youmall v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/yunjing v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/yunsou v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/zj v1.0.392 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.392
)

replace gopkg.in/jarcoal/httpmock.v1 => github.com/jarcoal/httpmock v1.0.5
replace github.com/tencentcloud/tencentcloud-sdk-go => github.com/tencentcloud/tencentcloud-sdk-go v1.0.392
