#!/bin/bash

GLOBAL_GCP_SERVICES=",dns,gcs,globalAddresses,globalForwardingRules,iam,"

if [ "$CSP" == "GCP" ]; then
    path="generated/google/${PROJECT_ID}"
fi

run_terraformer(){
	if [ "$CSP" == "GCP" ]; then
		if [[ "$GLOBAL_GCP_SERVICES" =~ .*",$1,".* ]]; then
			#	To be inline with the above regex, GLOBAL_GCP_SERVICES must start and end with a ","
			regions="global"
		else
			regions="asia-east1,asia-east2,asia-northeast1,asia-northeast2,asia-northeast3,asia-south1,asia-southeast1,australia-southeast1,europe-north1,europe-west1,europe-west2,europe-west3,europe-west4,europe-west6,northamerica-northeast1,southamerica-east1,us-central1,us-east1,us-east4,us-west1,us-west2,us-west3,us-west4,global"
		fi
		./terraformer-google import google --projects ${PROJECT_ID} -r ${1} -z ${regions}
		aws s3 sync --delete ${path}/${1}/ s3://${RESULT_BUCKET}/terraformer/${CUSTOMER_NAME}/${PROJECT_ID}/${TIMESTAMP}/${1}/
	else
		echo "./terraformer-azure import azure -r ${1}"
	fi
}

aws s3 cp s3://${RESULT_BUCKET}/terraformer/${CUSTOMER_NAME}/${PROJECT_ID}/credentials.json .

ls -la

if [ "$CSP" == "GCP" ]; then
	export GOOGLE_APPLICATION_CREDENTIALS=./credentials.json
	services=$(./terraformer-google import google list --projects ${PROJECT_ID})
else
	services=$(./terraformer-azure import azure list)
fi

for service in $services; do
	run_terraformer $service &
done

wait
