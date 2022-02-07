
### Use with AWS

Example:

```
 terraformer import aws --resources=vpc,subnet --connect=true --regions=eu-west-1 --profile=prod
 terraformer import aws --resources=vpc,subnet --filter=vpc=vpc_id1:vpc_id2:vpc_id3 --regions=eu-west-1
```

#### Profiles support

AWS configuration including environmental variables, shared credentials file (\~/.aws/credentials), and shared config file (\~/.aws/config) will be loaded by the tool by default. To use a specific profile, you can use the following command:

```
terraformer import aws --resources=vpc,subnet --regions=eu-west-1 --profile=prod
```

You can also provide no regions when importing resources:
```
terraformer import aws --resources=cloudfront --profile=prod
```
In that case terraformer will not know with which region resources are associated with and will not assume any region. That scenario is useful in case of global resources (e.g. CloudFront distributions or Route 53 records) and when region is passed implicitly through environmental variables or metadata service.

Examples to import other resources-

 * Security Group-
```
terraformer import aws --resources=sg --regions=us-east-1
```



#### Supported services

*   `accessanalyzer`
    * `aws_accessanalyzer_analyzer`
*   `acm`
    * `aws_acm_certificate`
*   `alb` (supports ALB and NLB)
    * `aws_lb`
    * `aws_lb_listener`
    * `aws_lb_listener_rule`
    * `aws_lb_listener_certificate`
    * `aws_lb_target_group`
    * `aws_lb_target_group_attachment`
*   `api_gateway`
    * `aws_api_gateway_authorizer`
    * `aws_api_gateway_documentation_part`
    * `aws_api_gateway_gateway_response`
    * `aws_api_gateway_integration`
    * `aws_api_gateway_integration_response`
    * `aws_api_gateway_method`
    * `aws_api_gateway_method_response`
    * `aws_api_gateway_model`
    * `aws_api_gateway_resource`
    * `aws_api_gateway_rest_api`
    * `aws_api_gateway_stage`
    * `aws_api_gateway_usage_plan`
    * `aws_api_gateway_vpc_link`
*   `appsync`
    * `aws_appsync_graphql_api`
*   `auto_scaling`
    * `aws_autoscaling_group`
    * `aws_launch_configuration`
    * `aws_launch_template`
*   `batch`
    * `aws_batch_compute_environment`
    * `aws_batch_job_definition`
    * `aws_batch_job_queue`
*   `budgets`
    * `aws_budgets_budget`
*   `cloud9`
    * `aws_cloud9_environment_ec2`
*   `cloudformation`
    * `aws_cloudformation_stack`
    * `aws_cloudformation_stack_set`
    * `aws_cloudformation_stack_set_instance`
*   `cloudfront`
    * `aws_cloudfront_distribution`
*   `cloudhsm`
    * `aws_cloudhsm_v2_cluster`
    * `aws_cloudhsm_v2_hsm`
*   `cloudtrail`
    * `aws_cloudtrail`
*   `cloudwatch`
    * `aws_cloudwatch_dashboard`
    * `aws_cloudwatch_event_rule`
    * `aws_cloudwatch_event_target`
    * `aws_cloudwatch_metric_alarm`
*   `codebuild`
    * `aws_codebuild_project`
*   `codecommit`
    * `aws_codecommit_repository`
*   `codedeploy`
    * `aws_codedeploy_app`
*   `codepipeline`
    * `aws_codepipeline`
    * `aws_codepipeline_webhook`
*   `cognito`
    * `aws_cognito_identity_pool`
    * `aws_cognito_user_pool`
*   `config`
    * `aws_config_config_rule`
    * `aws_config_configuration_recorder`
    * `aws_config_delivery_channel`
*   `customer_gateway`
    * `aws_customer_gateway`
*   `datapipeline`
    * `aws_datapipeline_pipeline`
*   `devicefarm`
    * `aws_devicefarm_project`
*   `docdb`
    * `aws_docdb_cluster`
    * `aws_docdb_cluster_instance`
    * `aws_docdb_cluster_parameter_group`
    * `aws_docdb_subnet_group`
*   `dynamodb`
    * `aws_dynamodb_table`
*   `ebs`
    * `aws_ebs_volume`
    * `aws_volume_attachment`
*   `ec2_instance`
    * `aws_instance`
*   `ecr`
    * `aws_ecr_lifecycle_policy`
    * `aws_ecr_repository`
    * `aws_ecr_repository_policy`
*   `ecrpublic`
    * `aws_ecrpublic_repository`
*   `ecs`
    * `aws_ecs_cluster`
    * `aws_ecs_service`
    * `aws_ecs_task_definition`
*   `efs`
    * `aws_efs_access_point`
    * `aws_efs_file_system`
    * `aws_efs_file_system_policy`
    * `aws_efs_mount_target`
*   `eip`
    * `aws_eip`
*   `eks`
    * `aws_eks_cluster`
*   `elasticache`
    * `aws_elasticache_cluster`
    * `aws_elasticache_parameter_group`
    * `aws_elasticache_subnet_group`
    * `aws_elasticache_replication_group`
*   `elastic_beanstalk`
    * `aws_elastic_beanstalk_application`
    * `aws_elastic_beanstalk_environment`
*   `elb`
    * `aws_elb`
*   `emr`
    * `aws_emr_cluster`
    * `aws_emr_security_configuration`
*   `eni`
    * `aws_network_interface`
*   `es`
    * `aws_elasticsearch_domain`
*   `firehose`
    * `aws_kinesis_firehose_delivery_stream`
*   `glue`
    * `glue_crawler`
    * `aws_glue_catalog_database`
    * `aws_glue_catalog_table`
    * `aws_glue_job`
    * `aws_glue_trigger`
*   `iam`
    * `aws_iam_group`
    * `aws_iam_group_policy`
    * `aws_iam_group_policy_attachment`
    * `aws_iam_instance_profile`
    * `aws_iam_policy`
    * `aws_iam_role`
    * `aws_iam_role_policy`
    * `aws_iam_role_policy_attachment`
    * `aws_iam_user`
    * `aws_iam_user_group_membership`
    * `aws_iam_user_policy`
    * `aws_iam_user_policy_attachment`
*   `igw`
    * `aws_internet_gateway`
*   `iot`
    * `aws_iot_thing`
    * `aws_iot_thing_type`
    * `aws_iot_topic_rule`
    * `aws_iot_role_alias`
*   `kinesis`
    * `aws_kinesis_stream`
*   `kms`
    * `aws_kms_key`
    * `aws_kms_alias`
    * `aws_kms_grant`
*   `lambda`
    * `aws_lambda_event_source_mapping`
    * `aws_lambda_function`
    * `aws_lambda_function_event_invoke_config`
    * `aws_lambda_layer_version`
*   `logs`
    * `aws_cloudwatch_log_group`
*   `media_package`
    * `aws_media_package_channel`
*   `media_store`
    * `aws_media_store_container`
*   `msk`
    * `aws_msk_cluster`
*   `nacl`
    * `aws_network_acl`
*   `nat`
    * `aws_nat_gateway`
*   `opsworks`
    * `aws_opsworks_application`
    * `aws_opsworks_custom_layer`
    * `aws_opsworks_instance`
    * `aws_opsworks_java_app_layer`
    * `aws_opsworks_php_app_layer`
    * `aws_opsworks_rds_db_instance`
    * `aws_opsworks_stack`
    * `aws_opsworks_static_web_layer`
    * `aws_opsworks_user_profile`
*   `organization`
    * `aws_organizations_account`
    * `aws_organizations_organization`
    * `aws_organizations_organizational_unit`
    * `aws_organizations_policy`
    * `aws_organizations_policy_attachment`
*   `qldb`
    * `aws_qldb_ledger`
*   `rds`
    * `aws_db_instance`
    * `aws_db_proxy`
    * `aws_db_cluster`
    * `aws_db_parameter_group`
    * `aws_db_subnet_group`
    * `aws_db_option_group`
    * `aws_db_event_subscription`
*   `resourcegroups`
    * `aws_resourcegroups_group`
*   `route53`
    * `aws_route53_zone`
    * `aws_route53_record`
*   `route_table`
    * `aws_route_table`
    * `aws_main_route_table_association`
    * `aws_route_table_association`
*   `s3`
    * `aws_s3_bucket`
*   `secretsmanager`
    * `aws_secretsmanager_secret`
*   `securityhub`
    * `aws_securityhub_account`
    * `aws_securityhub_member`
    * `aws_securityhub_standards_subscription`
*   `servicecatalog`
    * `aws_servicecatalog_portfolio`
*   `ses`
    * `aws_ses_configuration_set`
    * `aws_ses_domain_identity`
    * `aws_ses_email_identity`
    * `aws_ses_receipt_rule`
    * `aws_ses_receipt_rule_set`
    * `aws_ses_template`
*   `sfn`
    * `aws_sfn_activity`
    * `aws_sfn_state_machine`
*   `sg`
    * `aws_security_group`
    * `aws_security_group_rule` (if a rule cannot be inlined)
*   `sns`
    * `aws_sns_topic`
    * `aws_sns_topic_subscription`
*   `sqs`
    * `aws_sqs_queue`
*   `ssm`
    * `aws_ssm_parameter`
*   `subnet`
    * `aws_subnet`
*   `swf`
    * `aws_swf_domain`
*   `transit_gateway`
    * `aws_ec2_transit_gateway_route_table`
    * `aws_ec2_transit_gateway_vpc_attachment`
*   `waf`
    * `aws_waf_byte_match_set`
    * `aws_waf_geo_match_set`
    * `aws_waf_ipset`
    * `aws_waf_rate_based_rule`
    * `aws_waf_regex_match_set`
    * `aws_waf_regex_pattern_set`
    * `aws_waf_rule`
    * `aws_waf_rule_group`
    * `aws_waf_size_constraint_set`
    * `aws_waf_sql_injection_match_set`
    * `aws_waf_web_acl`
    * `aws_waf_xss_match_set`
*   `waf_regional`
    * `aws_wafregional_byte_match_set`
    * `aws_wafregional_geo_match_set`
    * `aws_wafregional_ipset`
    * `aws_wafregional_rate_based_rule`
    * `aws_wafregional_regex_match_set`
    * `aws_wafregional_regex_pattern_set`
    * `aws_wafregional_rule`
    * `aws_wafregional_rule_group`
    * `aws_wafregional_size_constraint_set`
    * `aws_wafregional_sql_injection_match_set`
    * `aws_wafregional_web_acl`
    * `aws_wafregional_xss_match_set`
*   `vpc`
    * `aws_vpc`
*   `vpc_peering`
    * `aws_vpc_peering_connection`
*   `vpn_connection`
    * `aws_vpn_connection`
*   `vpn_gateway`
    * `aws_vpn_gateway`
*   `workspaces`
    * `aws_workspaces_directory`
    * `aws_workspaces_ip_group`
    * `aws_workspaces_workspace`
*   `xray`
    * `aws_xray_sampling_rule`

#### Global services

AWS services that are global will be imported without specified region even if several regions will be passed. It is to ensure only one representation of an AWS resource is imported.

List of global AWS services:
*   `budgets`
*   `cloudfront`
*   `ecrpublic`
*   `iam`
*   `organization`
*   `route53`
*   `waf`

#### Attribute filters

Attribute filters allow filtering across different resource types by its attributes.

```
terraformer import aws --resources=ec2_instance,ebs --filter="Name=tags.costCenter;Value=20000:'20001:1'" --regions=eu-west-1
```
Will only import AWS EC2 instances along with EBS volumes annotated with tag `costCenter` with values `20000` or `20001:1`. Attribute filters are by default applicable to all resource types although it's possible to specify to what resource type a given filter should be applicable to by providing `Type=<type>` parameter. For example:
```
terraformer import aws --resources=ec2_instance,ebs --filter=Type=ec2_instance;Name=tags.costCenter;Value=20000:'20001:1' --regions=eu-west-1
```
Will work as same as example above with a change the filter will be applicable only to `ec2_instance` resources.

Few more examples - How to import ec2 instance based on instance name and id
```
terraformer import aws --resources=ec2_instance --filter="Name=tags.Name;Value=Terraformer" --regions=us-east-1
```
This command imports ec2 instance having name as Terraformer.
```
terraformer import aws --resources=ec2_instance --filter="Name=id;Value=i-0xxxxxxxxx" --regions=us-east-1
```
This command imports ec2 instance having instance-id as i-0xxxxxxxxx.

Due to fact API Gateway generates a lot of resources, it's possible to issue a filtering query to retrieve resources related to a given REST API by tags. To fetch resources related to a REST API resource with a tag `STAGE` and value `dev`, add parameter `--filter="Type=api_gateway_rest_api;Name=tags.STAGE;Value=dev"`.

#### SQS queues retrieval

Terraformer uses AWS [ListQueues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_ListQueues.html) API call to fetch available queues. The API is able to return only up to 1000 queues and an additional name prefix should be passed to filter the list results. It's possible to pass `QueueNamePrefix` parameter by environmental variable `SQS_PREFIX`.

#### Security groups and rules

Terraformer by default will try to keep rules in security groups as long as no circular dependencies are detected. This approach is implemented to keep the rules as tidy as possible but there can be cases when this behaviour is not desirable (see [GoogleCloudPlatform/terraformer#493](https://github.com/GoogleCloudPlatform/terraformer/issues/493)). To make Terraformer split rules from security groups, add `SPLIT_SG_RULES` environmental variable with any value.
