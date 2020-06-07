#!/bin/bash

GLOBAL_GCP_SERVICES=",dns,gcs,globalAddresses,globalForwardingRules,iam,"

case $CSP in
	"GCP")
		path="generated/google/${PROJECT_ID}"
		;;
	"Azure")
		path="generated/azurerm"
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
				regions="asia-east1,asia-east2,asia-northeast1,asia-northeast2,asia-northeast3,asia-south1,asia-southeast1,australia-southeast1,europe-north1,europe-west1,europe-west2,europe-west3,europe-west4,europe-west6,northamerica-northeast1,southamerica-east1,us-central1,us-east1,us-east4,us-west1,us-west2,us-west3,us-west4,global"
			fi
			./terraformer-google import google --projects ${PROJECT_ID} -r ${1} -z ${regions}
			aws s3 sync --delete ${path}/${1}/ s3://${RESULT_BUCKET}/terraformer/${CUSTOMER_NAME}/${PROJECT_ID}/${TIMESTAMP}/${1}/
			;;
		"Azure")
			./terraformer-azure import azure -r ${1}
			aws s3 sync --delete ${path}/${1}/ s3://${RESULT_BUCKET}/terraformer/${CUSTOMER_NAME}/${ARM_SUBSCRIPTION_ID}/${TIMESTAMP}/${1}/
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
		export ARM_CLIENT_SECRET=$(cat credentials.json | jq .clientSecret | sed s/\"//g)
		export ARM_TENANT_ID=$(cat credentials.json | jq .tenantId | sed s/\"//g)
		services=$(./terraformer-azure import azure list)
		;;
	*)
		echo "$CSP isn't supported"
		exit 1
esac

for service in $services; do
	run_terraformer $service &
done

wait
