module github.com/GoogleCloudPlatform/terraformer

go 1.22

require (
	cloud.google.com/go v0.110.2 // indirect
	cloud.google.com/go/logging v1.7.0
	cloud.google.com/go/storage v1.29.0
	github.com/Azure/azure-sdk-for-go v63.4.0+incompatible
	github.com/Azure/azure-storage-blob-go v0.10.0
	github.com/Azure/go-ansiterm v0.0.0-20210617225240-d185dfc1b5a1 // indirect
	github.com/Azure/go-autorest/autorest v0.11.27
	github.com/IBM-Cloud/bluemix-go v0.0.0-20220624043500-d538cb4fd9be
	github.com/IBM/go-sdk-core/v3 v3.3.1
	github.com/IBM/go-sdk-core/v4 v4.9.0
	github.com/IBM/ibm-cos-sdk-go v1.5.0
	github.com/IBM/keyprotect-go-client v0.8.1
	github.com/IBM/networking-go-sdk v0.30.0
	github.com/IBM/platform-services-go-sdk v0.26.1
	github.com/IBM/vpc-go-sdk v0.4.1
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/OctopusDeploy/go-octopusdeploy v1.6.0
	github.com/PaloAltoNetworks/pango v0.8.0
	github.com/SAP/go-hdb v0.105.2 // indirect
	github.com/SermoDigital/jose v0.9.1 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1247
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.2+incompatible
	github.com/apache/openwhisk-client-go v0.0.0-20210106144548-17d556327cd3
	github.com/aws/aws-sdk-go-v2 v1.24.0
	github.com/aws/aws-sdk-go-v2/config v1.26.1
	github.com/aws/aws-sdk-go-v2/credentials v1.16.12
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.26.5
	github.com/aws/aws-sdk-go-v2/service/acm v1.22.5
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.21.5
	github.com/aws/aws-sdk-go-v2/service/appsync v1.26.5
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.36.5
	github.com/aws/aws-sdk-go-v2/service/batch v1.30.5
	github.com/aws/aws-sdk-go-v2/service/budgets v1.20.5
	github.com/aws/aws-sdk-go-v2/service/cloud9 v1.22.3
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.42.4
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.32.5
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.19.5
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.35.5
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.32.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.21.5
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.29.5
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.26.5
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.19.5
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.22.1
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.22.5
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.21.5
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.31.5
	github.com/aws/aws-sdk-go-v2/service/configservice v1.43.5
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.19.5
	github.com/aws/aws-sdk-go-v2/service/devicefarm v1.20.5
	github.com/aws/aws-sdk-go-v2/service/docdb v1.29.5
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.26.6
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.141.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.24.5
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.21.5
	github.com/aws/aws-sdk-go-v2/service/ecs v1.35.5
	github.com/aws/aws-sdk-go-v2/service/efs v1.26.5
	github.com/aws/aws-sdk-go-v2/service/eks v1.35.5
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.34.5
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.20.5
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.21.5
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.26.5
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.24.5
	github.com/aws/aws-sdk-go-v2/service/emr v1.35.5
	github.com/aws/aws-sdk-go-v2/service/firehose v1.22.5
	github.com/aws/aws-sdk-go-v2/service/glue v1.72.4
	github.com/aws/aws-sdk-go-v2/service/iam v1.28.5
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.21.5
	github.com/aws/aws-sdk-go-v2/service/iot v1.46.5
	github.com/aws/aws-sdk-go-v2/service/kafka v1.28.5
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.24.5
	github.com/aws/aws-sdk-go-v2/service/kms v1.27.5
	github.com/aws/aws-sdk-go-v2/service/lambda v1.49.5
	github.com/aws/aws-sdk-go-v2/service/medialive v1.43.3
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.28.5
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.18.5
	github.com/aws/aws-sdk-go-v2/service/mq v1.20.5
	github.com/aws/aws-sdk-go-v2/service/opsworks v1.19.5
	github.com/aws/aws-sdk-go-v2/service/organizations v1.23.5
	github.com/aws/aws-sdk-go-v2/service/qldb v1.19.5
	github.com/aws/aws-sdk-go-v2/service/rds v1.64.5
	github.com/aws/aws-sdk-go-v2/service/redshift v1.39.6
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.19.5
	github.com/aws/aws-sdk-go-v2/service/route53 v1.35.5
	github.com/aws/aws-sdk-go-v2/service/s3 v1.47.5
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.25.5
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.43.5
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.25.5
	github.com/aws/aws-sdk-go-v2/service/ses v1.19.5
	github.com/aws/aws-sdk-go-v2/service/sfn v1.24.5
	github.com/aws/aws-sdk-go-v2/service/sns v1.26.5
	github.com/aws/aws-sdk-go-v2/service/sqs v1.29.5
	github.com/aws/aws-sdk-go-v2/service/ssm v1.44.5
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.23.5
	github.com/aws/aws-sdk-go-v2/service/sts v1.26.5
	github.com/aws/aws-sdk-go-v2/service/swf v1.20.5
	github.com/aws/aws-sdk-go-v2/service/waf v1.18.5
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.19.5
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.43.5
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.35.5
	github.com/aws/aws-sdk-go-v2/service/xray v1.23.5
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cloudflare/cloudflare-go v0.13.6
	github.com/cloudfoundry/jibber_jabber v0.0.0-20151120183258-bcc4c8345a21 // indirect
	github.com/containerd/continuity v0.1.0 // indirect
	github.com/ddelnano/terraform-provider-mikrotik/client v0.0.0-20210401060029-7f652169b2c4
	github.com/ddelnano/terraform-provider-xenorchestra/client v0.0.0-20210401070256-0d721c6762ef
	github.com/denisenkom/go-mssqldb v0.10.0 // indirect
	github.com/denverdino/aliyungo v0.0.0-20200327235253-d59c209c7e93
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/digitalocean/godo v1.83.0
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/duosecurity/duo_api_golang v0.0.0-20201112143038-0e07e9f869e3 // indirect
	github.com/elazarl/go-bindata-assetfs v1.0.1 // indirect
	github.com/fastly/go-fastly/v7 v7.0.0
	github.com/fatih/structs v1.1.0 // indirect
	github.com/gocql/gocql v0.0.0-20210707082121-9a3953d1826d // indirect
	github.com/google/go-github/v35 v35.1.0
	github.com/gophercloud/gophercloud v1.0.0
	github.com/grafana/grafana-api-golang-client v0.0.0-20210218192924-9ccd2365d2a6
	github.com/hashicorp/go-azure-helpers v0.36.0
	github.com/hashicorp/go-cleanhttp v0.5.2
	github.com/hashicorp/go-hclog v1.2.1
	github.com/hashicorp/go-memdb v1.3.2 // indirect
	github.com/hashicorp/go-plugin v1.4.4
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/terraform v0.12.31
	github.com/hashicorp/vault v0.10.4
	github.com/heimweh/go-pagerduty v0.0.0-20210930203304-530eff2acdc6
	github.com/heroku/heroku-go/v5 v5.4.1
	github.com/hokaccha/go-prettyjson v0.0.0-20210113012101-fb4e108d2519 // indirect
	github.com/honeycombio/terraform-provider-honeycombio v0.10.0
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/ionos-cloud/sdk-go-dbaas-mongo v1.3.1
	github.com/ionos-cloud/sdk-go-dbaas-postgres v1.1.2
	github.com/ionos-cloud/sdk-go/v6 v6.1.3
	github.com/jefferai/jsonx v1.0.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0
	github.com/jonboydell/logzio_client v1.2.0
	github.com/labd/commercetools-go-sdk v0.3.1
	github.com/linode/linodego v0.24.1
	github.com/microsoft/azure-devops-go-api/azuredevops v1.0.0-b5
	github.com/mrparkers/terraform-provider-keycloak v0.0.0-20221013232944-56f37a07590d
	github.com/nicksnyder/go-i18n v1.10.1 // indirect
	github.com/okta/okta-sdk-golang/v2 v2.12.2-0.20220602195034-d7ea6917663f
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/opencontainers/runc v1.1.12 // indirect
	github.com/opsgenie/opsgenie-go-sdk-v2 v1.2.9
	github.com/ory/dockertest v3.3.5+incompatible // indirect
	github.com/packethost/packngo v0.30.0
	github.com/pkg/errors v0.9.1
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/tencentyun/cos-go-sdk-v5 v0.7.34
	github.com/vultr/govultr v1.1.1
	github.com/xanzy/go-gitlab v0.50.2
	github.com/yandex-cloud/go-genproto v0.0.0-20220314102905-1acaee8ca7eb
	github.com/yandex-cloud/go-sdk v0.0.0-20220314105123-d0c2a928feb6
	github.com/zclconf/go-cty v1.11.0
	github.com/zorkian/go-datadog-api v2.30.0+incompatible
	golang.org/x/oauth2 v0.12.0
	golang.org/x/text v0.14.0
	gonum.org/v1/gonum v0.7.0
	google.golang.org/api v0.131.0
	google.golang.org/genproto v0.0.0-20230629202037-9506855d4529
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	k8s.io/apimachinery v0.24.2
	k8s.io/client-go v0.24.2
)

require (
	github.com/IBM-Cloud/container-services-go-sdk v0.0.0-20210705152127-41ca00fc9a62
	github.com/IBM/go-sdk-core v1.1.0
	github.com/aws/aws-sdk-go-v2/internal/ini v1.7.2 // indirect
	github.com/hashicorp/terraform-svchost v0.0.0-20200729002733-f050f53b9734 // indirect
	github.com/mackerelio/mackerel-client-go v0.21.0
	github.com/okta/terraform-provider-okta v0.0.0-20210924173942-a5a664459d3b
	github.com/zclconf/go-cty-yaml v1.0.2 // indirect
)

require (
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/antihax/optional v1.0.0 // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/launchdarkly/api-client-go v5.3.0+incompatible
)

require github.com/newrelic/newrelic-client-go v0.79.0

require (
	github.com/Azure/azure-pipeline-go v0.2.2 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.18 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.4 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/BurntSushi/toml v1.1.0 // indirect
	github.com/IBM/go-sdk-core/v5 v5.10.1 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-cidr v1.1.0 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/appscode/go-querystring v0.0.0-20170504095604-0126cfb3f1dc // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.4 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.9 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.9 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.2.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.2.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.8.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.16.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.18.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.21.5 // indirect
	github.com/aws/smithy-go v1.19.0
	github.com/beevik/etree v1.1.0 // indirect
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
	github.com/bgentry/speakeasy v0.1.0 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/bmatcuk/doublestar v1.1.5 // indirect
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/crewjam/saml v0.4.13 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dghubble/sling v1.1.0 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/form3tech-oss/jwt-go v3.2.3+incompatible // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
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
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/jsonapi v1.0.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-getter v1.7.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.0 // indirect
	github.com/hashicorp/go-safetemp v1.0.0 // indirect
	github.com/hashicorp/go-sockaddr v1.0.0 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl/v2 v2.14.0 // indirect
	github.com/hashicorp/hil v0.0.0-20190212112733-ab17b08d6590 // indirect
	github.com/hashicorp/yamux v0.0.0-20211028200310-0bc27b27de87 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kelseyhightower/envconfig v1.4.0 // indirect
	github.com/klauspost/compress v1.15.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattermost/xml-roundtrip-validator v0.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-ieproxy v0.0.0-20190702010315-6dee0af9227d // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
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
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/patrickmn/go-cache v0.0.0-20180815053127-5633e0862627 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/peterhellberg/link v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/posener/complete v1.2.1 // indirect
	github.com/russellhaering/goxmldsig v1.2.0 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/sourcegraph/jsonrpc2 v0.0.0-20210201082850-366fbb520750 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	github.com/tomnomnom/linkheader v0.0.0-20180905144013-02ca5825eb80 // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	github.com/vmihailenco/msgpack/v4 v4.3.12 // indirect
	github.com/vmihailenco/tagparser v0.1.2 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.7.5 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/term v0.15.0 // indirect
	golang.org/x/time v0.0.0-20220922220347-f3bd1da661af // indirect
	golang.org/x/tools v0.6.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.56.3 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/auth0.v5 v5.21.1
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.24.2 // indirect
	k8s.io/klog/v2 v2.60.1 // indirect
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

require github.com/PuerkitoBio/rehttp v1.0.0 // indirect

require (
	cloud.google.com/go/cloudbuild v1.9.0
	cloud.google.com/go/cloudtasks v1.10.0
	cloud.google.com/go/iam v0.13.0
	cloud.google.com/go/monitoring v1.13.0
	github.com/DataDog/datadog-api-client-go/v2 v2.11.0
	github.com/Myra-Security-GmbH/myrasec-go/v2 v2.28.0
	github.com/manicminer/hamilton v0.44.0
	github.com/opalsecurity/opal-go v1.0.19
	gopkg.in/ns1/ns1-go.v2 v2.6.5
)

require (
	cloud.google.com/go/compute v1.20.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/longrunning v0.4.1 // indirect
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/Myra-Security-GmbH/signature v1.0.0 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/aws/aws-sdk-go v1.44.122 // indirect
	github.com/clbanning/mxj v1.8.4 // indirect
	github.com/emicklei/go-restful v2.16.0+incompatible // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.3 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/s2a-go v0.1.4 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.5 // indirect
	github.com/hashicorp/terraform-plugin-log v0.7.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/manicminer/hamilton-autorest v0.2.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	golang.org/x/tools/cmd/cover v0.1.0-deprecated // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230629202037-9506855d4529 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230706204954-ccb25ca9f130 // indirect
	k8s.io/kube-openapi v0.0.0-20220328201542-3ee0da9b0b42 // indirect
	sigs.k8s.io/json v0.0.0-20211208200746-9f7c6b3444d2 // indirect
)

require (
	github.com/gofrs/uuid/v3 v3.1.2
	github.com/ionos-cloud/sdk-go-cert-manager v1.0.0
	github.com/ionos-cloud/sdk-go-container-registry v1.0.0
	github.com/ionos-cloud/sdk-go-dataplatform v1.0.1
	github.com/ionos-cloud/sdk-go-dns v1.1.1
	github.com/ionos-cloud/sdk-go-logging v1.0.1
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/as v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cbs v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdn v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cfs v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.694
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/es v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/gaap v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mongodb v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/pts v1.0.694
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/scf v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ses v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tat v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tcaplusdb v1.0.392
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc v1.0.392
)

replace gopkg.in/jarcoal/httpmock.v1 => github.com/jarcoal/httpmock v1.0.5

replace gopkg.in/ns1/ns1-go.v2 => github.com/ns1/ns1-go/v2 v2.6.5

replace github.com/tencentcloud/tencentcloud-sdk-go => github.com/tencentcloud/tencentcloud-sdk-go v1.0.392
