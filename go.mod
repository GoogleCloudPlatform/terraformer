module github.com/GoogleCloudPlatform/terraformer

go 1.16

require (
	cloud.google.com/go v0.77.0
	cloud.google.com/go/logging v1.1.2
	cloud.google.com/go/storage v1.12.0
	github.com/Azure/azure-sdk-for-go v42.3.0+incompatible
	github.com/Azure/azure-storage-blob-go v0.10.0
	github.com/Azure/go-autorest/autorest v0.11.12
	github.com/DataDog/datadog-api-client-go v1.0.0-beta.18
	github.com/IBM-Cloud/bluemix-go v0.0.0-20210203095940-db28d5e07b55
	github.com/IBM/go-sdk-core/v3 v3.3.1
	github.com/IBM/go-sdk-core/v4 v4.9.0
	github.com/IBM/ibm-cos-sdk-go v1.5.0
	github.com/IBM/keyprotect-go-client v0.6.0
	github.com/IBM/networking-go-sdk v0.13.0
	github.com/IBM/vpc-go-sdk v0.4.1
	github.com/OctopusDeploy/go-octopusdeploy v1.6.0
	github.com/aliyun/alibaba-cloud-sdk-go v1.60.295
	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.2+incompatible
	github.com/apache/openwhisk-client-go v0.0.0-20210106144548-17d556327cd3
	github.com/aws/aws-sdk-go-v2 v1.3.2
	github.com/aws/aws-sdk-go-v2/config v1.1.4
	github.com/aws/aws-sdk-go-v2/credentials v1.1.4
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.2.0
	github.com/aws/aws-sdk-go-v2/service/acm v1.2.1
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.2.1
	github.com/aws/aws-sdk-go-v2/service/appsync v1.2.1
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.2.1
	github.com/aws/aws-sdk-go-v2/service/budgets v1.1.3
	github.com/aws/aws-sdk-go-v2/service/cloud9 v1.1.3
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.3.0
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.2.1
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.1.3
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.2.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.3.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.2.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.2.1
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.2.1
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.1.3
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.2.1
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.2.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.2.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.2.1
	github.com/aws/aws-sdk-go-v2/service/configservice v1.3.0
	github.com/aws/aws-sdk-go-v2/service/datapipeline v1.1.3
	github.com/aws/aws-sdk-go-v2/service/devicefarm v1.1.3
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
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.2.1
	github.com/aws/aws-sdk-go-v2/service/emr v1.2.1
	github.com/aws/aws-sdk-go-v2/service/firehose v1.2.1
	github.com/aws/aws-sdk-go-v2/service/glue v1.3.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.3.0
	github.com/aws/aws-sdk-go-v2/service/iot v1.2.0
	github.com/aws/aws-sdk-go-v2/service/kafka v1.2.1
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.2.1
	github.com/aws/aws-sdk-go-v2/service/kms v1.2.1
	github.com/aws/aws-sdk-go-v2/service/lambda v1.2.1
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.2.1
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.1.4
	github.com/aws/aws-sdk-go-v2/service/opsworks v1.2.2
	github.com/aws/aws-sdk-go-v2/service/organizations v1.2.1
	github.com/aws/aws-sdk-go-v2/service/qldb v1.1.3
	github.com/aws/aws-sdk-go-v2/service/rds v1.2.1
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
	github.com/aws/aws-sdk-go-v2/service/swf v1.2.1
	github.com/aws/aws-sdk-go-v2/service/waf v1.1.4
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.2.1
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
	github.com/dollarshaveclub/new-relic-synthetics-go v0.0.0-20170605224734-4dc3dd6ae884
	github.com/fastly/go-fastly v1.18.0
	github.com/google/go-github/v25 v25.1.3
	github.com/gophercloud/gophercloud v0.13.0
	github.com/hashicorp/go-azure-helpers v0.10.0
	github.com/hashicorp/go-hclog v0.15.0
	github.com/hashicorp/go-plugin v1.4.0
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/terraform v0.12.29
	github.com/heroku/heroku-go/v5 v5.1.0
	github.com/hokaccha/go-prettyjson v0.0.0-20210113012101-fb4e108d2519 // indirect
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/jmespath/go-jmespath v0.4.0
	github.com/jonboydell/logzio_client v1.2.0
	github.com/labd/commercetools-go-sdk v0.0.0-20200309143931-ca72e918a79d
	github.com/linode/linodego v0.24.1
	github.com/mrparkers/terraform-provider-keycloak v0.0.0-20200506151941-509881368409
	github.com/nicksnyder/go-i18n v1.10.1 // indirect
	github.com/ns1/ns1-go v2.4.0+incompatible
	github.com/packethost/packngo v0.9.0
	github.com/paultyng/go-newrelic/v4 v4.10.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/tencentcloud/tencentcloud-sdk-go v3.0.233+incompatible
	github.com/tencentyun/cos-go-sdk-v5 v0.7.19
	github.com/vultr/govultr v0.5.0
	github.com/yandex-cloud/go-genproto v0.0.0-20200722140432-762fe965ce77
	github.com/yandex-cloud/go-sdk v0.0.0-20200722140627-2194e5077f13
	github.com/zclconf/go-cty v1.7.1
	github.com/zorkian/go-datadog-api v2.30.0+incompatible
	golang.org/x/oauth2 v0.0.0-20210113205817-d3ed898aa8a3
	golang.org/x/text v0.3.5
	gonum.org/v1/gonum v0.7.0
	google.golang.org/api v0.40.0
	google.golang.org/genproto v0.0.0-20210212180131-e7f2df4ecc2d
	gopkg.in/jarcoal/httpmock.v1 v1.0.0-00010101000000-000000000000 // indirect
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
)

replace gopkg.in/jarcoal/httpmock.v1 => github.com/jarcoal/httpmock v1.0.5
