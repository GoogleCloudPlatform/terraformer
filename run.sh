#!/bin/bash

GLOBAL_GCP_SERVICES=",dns,gcs,globalAddresses,globalForwardingRules,iam,gke,backendServices,backendBuckets,bigQuery,disks,firewall,healthChecks,httpHealthChecks,instanceTemplates,networks,project,routes,targetHttpsProxies,urlMaps,cloudsql,pubsub,networkEndpointGroups,instanceGroupManagers,instanceGroups,"
GLOBAL_AWS_SERVICES=",sts,iam,cloudfront,accessanalyzer,waf,route53,"

case $CSP in
	"GCP")
		path="generated/google/${PROJECT_ID}"
		;;
	"Azure")
		path="generated/azurerm"
		;;
	"AWS")
		path="generated/aws"
		;;
	*)
		echo "What is the local path for $CSP?"
		exit 1
esac

run_terraformer(){
	case $CSP in
		"GCP")
			if [[ "$GLOBAL_GCP_SERVICES" =~ .*",$1,".* ]]; then
				#	To be inline with the above regex, GLOBAL_GCP_SERVICES must start and end with a ","
				regions="global"
			elif [[ "$1" == "cloudFunctions" ]]; then
			  regions="us-central1,us-east1,us-east4,europe-west1,europe-west2,asia-east2,asia-northeast1,asia-northeast2,europe-west3,europe-west6,us-west3,northamerica-northeast1,australia-southeast1,us-west2,us-west4,southamerica-east1,asia-south1,asia-southeast2,asia-northeast3"
			else
				regions="asia-east1,asia-east2,asia-northeast1,asia-northeast2,asia-northeast3,asia-south1,asia-southeast1,australia-southeast1,europe-north1,europe-west1,europe-west2,europe-west3,europe-west4,europe-west6,northamerica-northeast1,southamerica-east1,us-central1,us-east1,us-east4,us-west1,us-west2,us-west3,us-west4"
			fi
			./terraformer-google import google --projects ${PROJECT_ID} -r ${1} -z ${regions} -c=false
			aws s3 sync --delete ${path}/${1}/ s3://${RESULT_BUCKET}/terraformer/${CUSTOMER_NAME}/${PROJECT_ID}/${TIMESTAMP}/${1}/
			;;
		"Azure")
			./terraformer-azure import azure -r ${1} -c=false
			aws s3 sync --delete ${path}/${1}/ s3://${RESULT_BUCKET}/terraformer/${CUSTOMER_NAME}/${ARM_SUBSCRIPTION_ID}/${TIMESTAMP}/${1}/
			;;
		"AWS")
			if [[ "$GLOBAL_AWS_SERVICES" =~ .*",$1,".* ]]; then
				#	To be inline with the above regex, GLOBAL_GCP_SERVICES must start and end with a ","
				echo "running command: ./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources ${1} --regions global -c=false? || true"
  			./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources ${1} --regions global -c=false || true
			elif [[ $1 == "eks" ]]; then
			  echo "running command: ./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources ${1} --regions us-east-1,us-east-2,us-west-2,ap-south-1,ap-southeast-1,ap-southeast-2,ap-northeast-1,ap-northeast-2,ca-central-1,eu-central-1,eu-west-1,eu-west-2,eu-west-3,eu-north-1,sa-east-1 -c=false || true"
  			./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources ${1} --regions us-east-1,us-east-2,us-west-2,ap-south-1,ap-southeast-1,ap-southeast-2,ap-northeast-1,ap-northeast-2,ca-central-1,eu-central-1,eu-west-1,eu-west-2,eu-west-3,eu-north-1,sa-east-1 -c=false || true
			else
			  echo "running command: ./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources ${1}  --regions us-east-1,us-east-2,us-west-1,us-west-2,ca-central-1,sa-east-1 -c=false --retry-number=2 --retry-sleep-ms=5000 -|| true"
  			./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources ${1} --regions us-east-1,us-east-2,us-west-1,us-west-2,ca-central-1,sa-east-1 -c=false --retry-number=2 --retry-sleep-ms=5000 -|| true
			  echo "running command: ./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources ${1}  --regions ap-south-1,ap-southeast-1,ap-southeast-2,ap-northeast-1,ap-northeast-2 -c=false || true"
  			./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources ${1} --regions ap-south-1,ap-southeast-1,ap-southeast-2,ap-northeast-1,ap-northeast-2 -c=false || true
			  echo "running command: ./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources ${1}  --regions eu-central-1,eu-west-1,eu-west-2,eu-west-3,eu-north-1 -c=false || true"
  			./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources ${1} --regions eu-central-1,eu-west-1,eu-west-2,eu-west-3,eu-north-1 -c=false || true
			fi
			echo "Completed ${1}"
			;;
		*)
			echo "terraformer doesn't run on $CSP"
			exit 1
	esac
}

aws s3 cp s3://${RESULT_BUCKET}/terraformer/${CUSTOMER_NAME}/${PROJECT_ID}/credentials.json .

ls -la

case $CSP in
	"GCP")
		export GOOGLE_APPLICATION_CREDENTIALS=./credentials.json
		services=$(./terraformer-google import google list --projects ${PROJECT_ID})
		;;
	"Azure")
		export ARM_SUBSCRIPTION_ID=$(cat credentials.json | jq .subscriptionId | sed s/\"//g)
		export ARM_CLIENT_ID=$(cat credentials.json | jq .clientId | sed s/\"//g)
		export ARM_TENANT_ID=$(cat credentials.json | jq .tenantId | sed s/\"//g)
		export ARM_CLIENT_SECRET=$(cat credentials.json | jq .clientSecret | sed s/\"//g)
		services=$(./terraformer-azure import azure list)
		;;
	"AWS")
		CUSTOMER_ARN_ROLE=$(cat credentials.json | jq .roleArn | sed s/\"//g)
		EXTERNAL_ID=$(cat credentials.json | jq .externalId | sed s/\"//g)
		mkdir ~/.aws
		cat << AWS_CREDS > ~/.aws/credentials
[${ACCOUNT_ID}]
credential_source = EcsContainer
role_arn = ${CUSTOMER_ARN_ROLE}
external_id = ${EXTERNAL_ID}
AWS_CREDS

		services="vpc,sg,nacl,nat,igw,subnet,vpc_peering,route_table vpn_connection,vpn_gateway,transit_gateway eip,customer_gateway alb,elb ec2_instance,ebs auto_scaling eks sts,iam,cloudfront,accessanalyzer,waf,route53 ecr,ecs,acm,kinesis,codecommit elasticache rds,sqs,cloudtrail,config kms lambda,s3,sns,dynamodb,es logs,glue,waf_regional"
		;;
	*)
		echo "$CSP isn't supported"
		exit 1
esac

services_to_run=""

for service in $services; do
  if [[ $service == "kms" && $CSP == "GCP" ]]; then
    continue
  fi
  if [[ $service == "monitoring" && $CSP == "GCP" ]]; then
    continue
  fi
  if [[ $service == "schedulerJobs" && $CSP == "GCP" ]]; then
    continue
  fi
  if [[ $service == "storage_blob" && $CSP == "Azure" ]]; then
    continue
  fi
  if [[ $service == "alb,elb" && $CSP == "AWS" && ($CUSTOMER_NAME == "nubank" || $CUSTOMER_NAME == "brex" || $CUSTOMER_NAME == "lendinghome") ]]; then
    continue
  fi
  if [[ $service == "ec2_instance,ebs" && $CSP == "AWS" && ($CUSTOMER_NAME == "robinhood" || $CUSTOMER_NAME == "nubank") ]]; then
    continue
  fi
  if [[ $service == "kms" && $CSP == "AWS" && $CUSTOMER_NAME == "nubank" ]]; then
    continue
  fi
  if [[ $service == "elasticache" && $CSP == "AWS" && $CUSTOMER_NAME == "robinhood" ]]; then
    continue
  fi
  if [[ $service == "route53" && $CSP == "AWS" && $CUSTOMER_NAME == "realgeeks" ]]; then
    continue
  fi
  if [[ $CUSTOMER_NAME != "nubank" && $CUSTOMER_NAME != "bridgecrew66" ]]; then
    run_terraformer $service &
    continue
  fi
  if [[ $services_to_run == "" ]]; then
    services_to_run="$service"
    continue
  fi
  services_to_run="$services_to_run,$service"
done

if [[ $CUSTOMER_NAME = "bridgecrew66" ]]; then
   echo "running not global ./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources=vpc,sg,nacl,nat,igw,subnet,vpc_peering,route_table,vpn_connection,vpn_gateway,transit_gateway,eip,customer_gateway,alb,elb,ec2_instance,auto_scaling,eks,ecr,ecs,acm,kinesis,codecommit,elasticache,rds,sqs,cloudtrail,conig,kms,lambda,s3,sns,dynamodb,es,logs --regions=us-east-1,us-east-2,us-west-1,us-west-2,ca-central-1,sa-east-1,ap-south-1,ap-southeast-1,ap-southeast-2,ap-northeast-1,ap-northeast-2,eu-central-1,eu-west-1,eu-west-2,eu-west-3,eu-north-1 -c=false || true"
  ./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources=vpc,sg,nacl,nat,igw,subnet,vpc_peering,route_table,vpn_connection,vpn_gateway,transit_gateway,eip,customer_gateway,alb,elb,ec2_instance,auto_scaling,eks,ecr,ecs,acm,kinesis,codecommit,elasticache,rds,sqs,cloudtrail,conig,kms,lambda,s3,sns,dynamodb,es,logs --regions=us-east-1,us-east-2,us-west-1,us-west-2,ca-central-1,sa-east-1,ap-south-1,ap-southeast-1,ap-southeast-2,ap-northeast-1,ap-northeast-2,eu-central-1,eu-west-1,eu-west-2,eu-west-3,eu-north-1 -c=false || true
  echo "running global import aws --profile ${ACCOUNT_ID} --resources=sts,iam,cloudfront,accessanalyzer --regions global -c=false || true"
  ./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources=sts,iam,cloudfront,accessanalyzer --regions global -c=false || true
fi
if [[ $CUSTOMER_NAME == "nubank" ]]; then
  echo "running not global ./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources=vpc,sg,nacl,nat,igw,subnet,vpc_peering,route_table,vpn_connection,vpn_gateway,transit_gateway,eip,customer_gateway,auto_scaling,eks,sts,ecr,ecs,acm,kinesis,codecommit,elasticache,rds,sqs,cloudtrail,config,lambda,s3,sns,dynamodb,es,logs --regions=us-east-1 -c=false || true"
  ./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources=vpc,sg,nacl,nat,igw,subnet,vpc_peering,route_table,vpn_connection,vpn_gateway,transit_gateway,eip,customer_gateway,ec2_instance,auto_scaling,eks,sts,ecr,ecs,acm,kinesis,codecommit,elasticache,rds,sqs,cloudtrail,config,lambda,s3,sns,dynamodb,es,logs --regions=us-east-1 -c=false || true
  echo "running global ./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources=sts,cloudfront,accessanalyzer --regions global -c=false || true"
  ./terraformer-aws import aws --profile ${ACCOUNT_ID} --resources=sts,cloudfront,accessanalyzer --regions global -c=false || true
fi

wait

if [[ $CSP == "AWS" ]]; then
  aws s3 sync ${path}/ s3://${RESULT_BUCKET}/terraformer/${CUSTOMER_NAME}/${ACCOUNT_ID}/${TIMESTAMP}/
fi
