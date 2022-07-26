module github.com/GoogleCloudPlatform/terraformer

go 1.18

require (
	cloud.google.com/go v0.100.2 // indirect
	cloud.google.com/go/logging v1.1.2
	cloud.google.com/go/storage v1.14.0
	github.com/Azure/azure-sdk-for-go v63.4.0+incompatible
	github.com/Azure/azure-storage-blob-go v0.14.0
	github.com/Azure/go-autorest/autorest v0.11.27
	github.com/DataDog/datadog-api-client-go v1.14.0
	github.com/IBM-Cloud/bluemix-go v0.0.0-20210203095940-db28d5e07b55
	github.com/IBM/go-sdk-core/v3 v3.3.1
	github.com/IBM/go-sdk-core/v4 v4.9.0
	github.com/IBM/ibm-cos-sdk-go v1.5.0
	github.com/IBM/keyprotect-go-client v0.6.0
	github.com/IBM/networking-go-sdk v0.30.0
	github.com/IBM/platform-services-go-sdk v0.26.1
	github.com/IBM/vpc-go-sdk v0.4.1
	github.com/OctopusDeploy/go-octopusdeploy v1.6.0
	github.com/PaloAltoNetworks/pango v0.6.0
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1247
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.2+incompatible
	github.com/apache/openwhisk-client-go v0.0.0-20210106144548-17d556327cd3
	github.com/aws/aws-sdk-go-v2 v1.16.7
	github.com/aws/aws-sdk-go-v2/config v1.6.0
	github.com/aws/aws-sdk-go-v2/credentials v1.3.2
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.2.0
	github.com/aws/aws-sdk-go-v2/service/acm v1.2.1
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.2.1
	github.com/aws/aws-sdk-go-v2/service/appsync v1.14.4
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.3.1
	github.com/aws/aws-sdk-go-v2/service/batch v1.3.1
	github.com/aws/aws-sdk-go-v2/service/budgets v1.9.0
	github.com/aws/aws-sdk-go-v2/service/cloud9 v1.1.3
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.3.0
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.3.0
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.13.8
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.2.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.18.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.2.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.2.3
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.2.1
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.1.3
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.2.1
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.13.3
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.2.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.2.1
	github.com/aws/aws-sdk-go-v2/service/configservice v1.3.0
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.13.6
	github.com/aws/aws-sdk-go-v2/service/devicefarm v1.13.3
	github.com/aws/aws-sdk-go-v2/service/docdb v1.18.1
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.2.1
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.3.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.17.6
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.13.8
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
	github.com/aws/aws-sdk-go-v2/service/iot v1.24.1
	github.com/aws/aws-sdk-go-v2/service/kafka v1.14.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.2.1
	github.com/aws/aws-sdk-go-v2/service/kms v1.2.2
	github.com/aws/aws-sdk-go-v2/service/lambda v1.2.1
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.15.3
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.12.5
	github.com/aws/aws-sdk-go-v2/service/opsworks v1.2.2
	github.com/aws/aws-sdk-go-v2/service/organizations v1.2.1
	github.com/aws/aws-sdk-go-v2/service/qldb v1.1.3
	github.com/aws/aws-sdk-go-v2/service/rds v1.18.1
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.2.1
	github.com/aws/aws-sdk-go-v2/service/route53 v1.3.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.12.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.2.1
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.2.1
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.2.1
	github.com/aws/aws-sdk-go-v2/service/ses v1.14.8
	github.com/aws/aws-sdk-go-v2/service/sfn v1.2.1
	github.com/aws/aws-sdk-go-v2/service/sns v1.2.1
	github.com/aws/aws-sdk-go-v2/service/sqs v1.3.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.3.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.6.1
	github.com/aws/aws-sdk-go-v2/service/swf v1.13.8
	github.com/aws/aws-sdk-go-v2/service/waf v1.1.4
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.12.3
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.18.0
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.2.1
	github.com/aws/aws-sdk-go-v2/service/xray v1.2.1
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cloudflare/cloudflare-go v0.13.6
	github.com/cloudfoundry/jibber_jabber v0.0.0-20151120183258-bcc4c8345a21 // indirect
	github.com/ddelnano/terraform-provider-mikrotik/client v0.0.0-20210401060029-7f652169b2c4
	github.com/ddelnano/terraform-provider-xenorchestra/client v0.0.0-20210401070256-0d721c6762ef
	github.com/denverdino/aliyungo v0.0.0-20200327235253-d59c209c7e93
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/digitalocean/godo v1.57.0
	github.com/fastly/go-fastly/v6 v6.3.2
	github.com/google/go-github/v35 v35.1.0
	github.com/gophercloud/gophercloud v0.24.0
	github.com/grafana/grafana-api-golang-client v0.0.0-20210218192924-9ccd2365d2a6
	github.com/hashicorp/go-azure-helpers v0.36.0
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/hashicorp/go-hclog v1.1.0
	github.com/hashicorp/go-plugin v1.4.4
	github.com/hashicorp/hcl v1.0.1-vault-3
	github.com/hashicorp/terraform v0.12.31
	github.com/heimweh/go-pagerduty v0.0.0-20210930203304-530eff2acdc6
	github.com/heroku/heroku-go/v5 v5.4.1
	github.com/hokaccha/go-prettyjson v0.0.0-20210113012101-fb4e108d2519 // indirect
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/jmespath/go-jmespath v0.4.0
	github.com/jonboydell/logzio_client v1.2.0
	github.com/labd/commercetools-go-sdk v0.3.1
	github.com/linode/linodego v0.24.1
	github.com/microsoft/azure-devops-go-api/azuredevops v1.0.0-b5
	github.com/mrparkers/terraform-provider-keycloak v0.0.0-20200506151941-509881368409
	github.com/nicksnyder/go-i18n v1.10.1 // indirect
	github.com/ns1/ns1-go v2.4.0+incompatible
	github.com/okta/okta-sdk-golang/v2 v2.12.2-0.20220602195034-d7ea6917663f
	github.com/opsgenie/opsgenie-go-sdk-v2 v1.2.9
	github.com/packethost/packngo v0.9.0
	github.com/pkg/errors v0.9.1
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/tencentcloud/tencentcloud-sdk-go v3.0.233+incompatible
	github.com/tencentyun/cos-go-sdk-v5 v0.7.34
	github.com/vultr/govultr v1.1.1
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
	k8s.io/apimachinery v0.24.2
	k8s.io/client-go v0.24.2
)

require (
	github.com/IBM-Cloud/container-services-go-sdk v0.0.0-20210705152127-41ca00fc9a62
	github.com/IBM/go-sdk-core v1.1.0
	github.com/aws/aws-sdk-go-v2/internal/ini v1.2.2 // indirect
	github.com/hashicorp/terraform-svchost v0.0.0-20200729002733-f050f53b9734 // indirect
	github.com/mackerelio/mackerel-client-go v0.21.0
	github.com/okta/terraform-provider-okta v0.0.0-20210924173942-a5a664459d3b
	github.com/zclconf/go-cty-yaml v1.0.2 // indirect
)

require (
	github.com/antihax/optional v1.0.0 // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/launchdarkly/api-client-go v5.3.0+incompatible
)

require github.com/newrelic/newrelic-client-go v0.79.0

require (
	github.com/Azure/azure-pipeline-go v0.2.3 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.18 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.5 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/IBM/go-sdk-core/v5 v5.10.1 // indirect
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
	github.com/aws/aws-sdk-go v1.37.19 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.4.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.14 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.2.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.5.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshift v1.10.0
	github.com/aws/aws-sdk-go-v2/service/sso v1.3.2 // indirect
	github.com/aws/smithy-go v1.12.0 // indirect
	github.com/beevik/etree v1.1.0 // indirect
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
	github.com/bgentry/speakeasy v0.1.0 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/bmatcuk/doublestar v1.1.5 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/crewjam/saml v0.4.5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dghubble/sling v1.1.0 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/form3tech-oss/jwt-go v3.2.5+incompatible // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/go-openapi/errors v0.19.8 // indirect
	github.com/go-openapi/strfmt v0.21.2 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-resty/resty/v2 v2.1.1-0.20191201195748-d7b97669fe48 // indirect
	github.com/go-routeros/routeros v0.0.0-20210123142807-2a44d57c6730 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/jsonapi v1.0.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/gax-go/v2 v2.1.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-getter v1.5.3 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-safetemp v1.0.0 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/go-version v1.4.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl/v2 v2.8.2 // indirect
	github.com/hashicorp/hil v0.0.0-20190212112733-ab17b08d6590 // indirect
	github.com/hashicorp/yamux v0.0.0-20211028200310-0bc27b27de87 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jonboulle/clockwork v0.2.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattermost/xml-roundtrip-validator v0.0.0-20201213122252-bcd7e1b9601e // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-ieproxy v0.0.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/cli v1.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/hashstructure v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mozillazg/go-httpheader v0.2.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/peterhellberg/link v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/posener/complete v1.2.3 // indirect
	github.com/russellhaering/goxmldsig v1.1.0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/sourcegraph/jsonrpc2 v0.0.0-20210201082850-366fbb520750 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	github.com/tomnomnom/linkheader v0.0.0-20180905144013-02ca5825eb80 // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	github.com/vmihailenco/msgpack/v4 v4.3.12 // indirect
	github.com/vmihailenco/tagparser v0.1.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.7.5 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220422013727-9388b58f7150 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	golang.org/x/tools v0.1.8 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/auth0.v5 v5.21.1
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
	k8s.io/api v0.24.2 // indirect
	k8s.io/klog/v2 v2.60.1 // indirect
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

require github.com/PuerkitoBio/rehttp v1.0.0 // indirect

require (
	cloud.google.com/go/cloudtasks v1.3.0
	cloud.google.com/go/iam v0.3.0
	cloud.google.com/go/monitoring v1.4.0
	github.com/hashicorp/vault/api v1.7.2
	github.com/manicminer/hamilton v0.44.0
)

require (
	cloud.google.com/go/compute v1.3.0 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/armon/go-metrics v0.3.10 // indirect
	github.com/benbjohnson/clock v1.1.0 // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/emicklei/go-restful v2.9.5+incompatible // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/go-test/deep v1.0.8 // indirect
	github.com/golang-jwt/jwt/v4 v4.3.0 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-secure-stdlib/mlock v0.1.2 // indirect
	github.com/hashicorp/go-secure-stdlib/parseutil v0.1.6 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/hashicorp/vault/sdk v0.5.3-0.20220621155127-c9ca5e0e239b // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/manicminer/hamilton-autorest v0.2.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/rogpeppe/go-internal v1.6.2 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	k8s.io/kube-openapi v0.0.0-20220328201542-3ee0da9b0b42 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
)

replace gopkg.in/jarcoal/httpmock.v1 => github.com/jarcoal/httpmock v1.0.5
