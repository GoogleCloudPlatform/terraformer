#!/bin/bash

GLOBAL_GCP_SERVICES=",dns,gcs,globalAddresses,globalForwardingRules,iam,gke,backendServices,bigQuery,disks,firewall,healthChecks,httpHealthChecks,instanceTemplates,networks,project,routes,targetHttpsProxies,urlMaps,"
GLOBAL_AWS_SERVICES=",sts,iam,route53,route53domains,s3,s3control,cloudfront,organizations,"

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
			else
				regions="asia-east1,asia-east2,asia-northeast1,asia-northeast2,asia-northeast3,asia-south1,asia-southeast1,australia-southeast1,europe-north1,europe-west1,europe-west2,europe-west3,europe-west4,europe-west6,northamerica-northeast1,southamerica-east1,us-central1,us-east1,us-east4,us-west1,us-west2,us-west3,us-west4"
			fi
			./terraformer-google import google --projects ${PROJECT_ID} -r ${1} -z ${regions}
			aws s3 sync --delete ${path}/${1}/ s3://${RESULT_BUCKET}/terraformer/${CUSTOMER_NAME}/${PROJECT_ID}/${TIMESTAMP}/${1}/
			;;
		"Azure")
			./terraformer-azure import azure -r ${1}
			aws s3 sync --delete ${path}/${1}/ s3://${RESULT_BUCKET}/terraformer/${CUSTOMER_NAME}/${ARM_SUBSCRIPTION_ID}/${TIMESTAMP}/${1}/
			;;
		"AWS")
			if [[ "$GLOBAL_AWS_SERVICES" =~ .*",$1,".* ]]; then
				#	To be inline with the above regex, GLOBAL_GCP_SERVICES must start and end with a ","
				regions="global"
			elif [[ $1 == "s3" ]]; then
			  regions="us-east-1"
			elif [[ $1 == "devicefarm" ]]; then
			  regions="us-east-1"
			elif [[ $1 == "eks" ]]; then
			  regions="us-east-1"
			elif [[ $1 == "media_store" ]]; then
			  regions="us-east-1"
			else
				regions="us-east-1,us-east-2,us-west-1,us-west-2,ap-south-1,ap-southeast-1,ap-southeast-2,ap-northeast-1,ap-northeast-2,ca-central-1,eu-central-1,eu-west-1,eu-west-2,eu-west-3,eu-north-1,sa-east-1"
			fi

			AWS_ACCESS_KEY_ID=${CUSTOMER_AWS_ACCESS_KEY_ID} AWS_SECRET_ACCESS_KEY=${CUSTOMER_AWS_SECRET_ACCESS_KEY} AWS_SESSION_TOKEN=${CUSTOMER_AWS_SESSION_TOKEN} ./terraformer-aws import aws --resources ${1} --regions ${regions} || true
			aws s3 sync --delete ${path}/${1}/ s3://${RESULT_BUCKET}/terraformer/${CUSTOMER_NAME}/${ACCOUNT_ID}/${TIMESTAMP}/${1}/
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
		export CUSTOMER_AWS_ACCESS_KEY_ID=$(cat credentials.json | jq .accessKeyId | sed s/\"//g)
		export CUSTOMER_AWS_SECRET_ACCESS_KEY=$(cat credentials.json | jq .secretAccessKey | sed s/\"//g)
		export CUSTOMER_AWS_SESSION_TOKEN=$(cat credentials.json | jq .sessionToken | sed s/\"//g)
		services=$(./terraformer-aws import aws list)
		;;
	*)
		echo "$CSP isn't supported"
		exit 1
esac

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
  if [[ $service == "api_gateway" && $CSP == "AWS" ]]; then ### us-east-1
    continue
  fi
  if [[ $service == "budgets" && $CSP == "AWS" ]]; then ### >>>>>>>>>>>     default region
    continue
  fi
  if [[ $service == "cloudwatch" && $CSP == "AWS" ]]; then ### us-east-1
    continue
  fi
  if [[ $service == "codepipeline" && $CSP == "AWS" ]]; then ### us-east-1
    continue
  fi
  if [[ $service == "emr" && $CSP == "AWS" ]]; then ### us-east-1
    continue
  fi
  if [[ $service == "glue" && $CSP == "AWS" ]]; then ### us-east-1
    continue
  fi
  if [[ $service == "media_package" && $CSP == "AWS" ]]; then ### us-east-1
    continue
  fi
  if [[ $service == "msk" && $CSP == "AWS" ]]; then ### us-east-1
    continue
  fi
  if [[ $service == "qldb" && $CSP == "AWS" ]]; then ### us-east-1
    continue
  fi
  if [[ $service == "resourcegroups" && $CSP == "AWS" ]]; then ### us-east-1
    continue
  fi
  if [[ $service == "servicecatalog" && $CSP == "AWS" ]]; then ### us-east-1
    continue
  fi
  if [[ $service == "ses" && $CSP == "AWS" ]]; then ### us-east-1
    continue
  fi
  if [[ $service == "waf" && $CSP == "AWS" ]]; then ### >>>>>>>>>>>     default region
    continue
  fi
  if [[ $service == "waf_regional" && $CSP == "AWS" ]]; then ### us-east-1
    continue
  fi
  
  run_terraformer $service &
	
done

wait
