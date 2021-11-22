# Use with Azure

Supports [Azure CLI](https://www.terraform.io/docs/providers/azurerm/guides/azure_cli.html), [Service Principal with Client Certificate](https://www.terraform.io/docs/providers/azurerm/guides/service_principal_client_certificate.html), and [Service Principal with Client Secret](https://www.terraform.io/docs/providers/azurerm/guides/service_principal_client_secret.html).

## Example

``` sh
# Using Azure CLI (az login)
export ARM_SUBSCRIPTION_ID=[SUBSCRIPTION_ID]

# Using Service Principal with Client Certificate
export ARM_SUBSCRIPTION_ID=[SUBSCRIPTION_ID]
export ARM_CLIENT_ID=[CLIENT_ID]
export ARM_CLIENT_CERTIFICATE_PATH="/path/to/my/client/certificate.pfx"
export ARM_CLIENT_CERTIFICATE_PASSWORD=[CLIENT_CERTIFICATE_PASSWORD]
export ARM_TENANT_ID=[TENANT_ID]

# Service Principal with Client Secret
export ARM_SUBSCRIPTION_ID=[SUBSCRIPTION_ID]
export ARM_CLIENT_ID=[CLIENT_ID]
export ARM_CLIENT_SECRET=[CLIENT_SECRET]
export ARM_TENANT_ID=[TENANT_ID]

./terraformer import azure -r resource_group
./terraformer import azure -R my_resource_group -r virtual_network,resource_group
./terraformer import azure -r resource_group --filter=resource_group=/subscriptions/<Subscription id>/resourceGroups/<RGNAME>
```


## List of supported Azure resources

*   `analysis`
    * `azurerm_analysis_services_server`
*   `app_service`
    * `azurerm_app_service`
*   `application_gateway`
    * `azurerm_application_gateway`
*   `container`
    * `azurerm_container_group`
    * `azurerm_container_registry`
    * `azurerm_container_registry_webhook`
*   `cosmosdb`
	* `azurerm_cosmosdb_account`
	* `azurerm_cosmosdb_sql_container`
	* `azurerm_cosmosdb_sql_database`
	* `azurerm_cosmosdb_table`
*   `database`
	* `azurerm_mariadb_configuration`
	* `azurerm_mariadb_database`
	* `azurerm_mariadb_firewall_rule`
	* `azurerm_mariadb_server`
	* `azurerm_mariadb_virtual_network_rule`
	* `azurerm_mysql_configuration`
	* `azurerm_mysql_database`
	* `azurerm_mysql_firewall_rule`
	* `azurerm_mysql_server`
	* `azurerm_mysql_virtual_network_rule`
	* `azurerm_postgresql_configuration`
	* `azurerm_postgresql_database`
	* `azurerm_postgresql_firewall_rule`
	* `azurerm_postgresql_server`
	* `azurerm_postgresql_virtual_network_rule`
	* `azurerm_sql_database`
	* `azurerm_sql_active_directory_administrator`
	* `azurerm_sql_elasticpool`
	* `azurerm_sql_failover_group`
	* `azurerm_sql_firewall_rule`
	* `azurerm_sql_server`
	* `azurerm_sql_virtual_network_rule`
*   `databricks`
    * `azurerm_databricks_workspace`
*   `data_factory`
    * `azurerm_data_factory`
    * `azurerm_data_factory_pipeline`
    * `azurerm_data_factory_data_flow`
    * `azurerm_data_factory_dataset_azure_blob`
    * `azurerm_data_factory_dataset_binary`
    * `azurerm_data_factory_dataset_cosmosdb_sqlapi`
    * `azurerm_data_factory_custom_dataset`
    * `azurerm_data_factory_dataset_delimited_text`
    * `azurerm_data_factory_dataset_http`
    * `azurerm_data_factory_dataset_json`
    * `azurerm_data_factory_dataset_mysql`
    * `azurerm_data_factory_dataset_parquet`
    * `azurerm_data_factory_dataset_postgresql`
    * `azurerm_data_factory_dataset_snowflake`
    * `azurerm_data_factory_dataset_sql_server_table`
    * `azurerm_data_factory_integration_runtime_azure`
    * `azurerm_data_factory_integration_runtime_managed`
    * `azurerm_data_factory_integration_runtime_azure_ssis`
    * `azurerm_data_factory_integration_runtime_self_hosted`
    * `azurerm_data_factory_linked_service_azure_blob_storage`
    * `azurerm_data_factory_linked_service_azure_databricks`
    * `azurerm_data_factory_linked_service_azure_file_storage`
    * `azurerm_data_factory_linked_service_azure_function`
    * `azurerm_data_factory_linked_service_azure_search`
    * `azurerm_data_factory_linked_service_azure_sql_database`
    * `azurerm_data_factory_linked_service_azure_table_storage`
    * `azurerm_data_factory_linked_service_cosmosdb`
    * `azurerm_data_factory_linked_custom_service`
    * `azurerm_data_factory_linked_service_data_lake_storage_gen2`
    * `azurerm_data_factory_linked_service_key_vault`
    * `azurerm_data_factory_linked_service_kusto`
    * `azurerm_data_factory_linked_service_mysql`
    * `azurerm_data_factory_linked_service_odata`
    * `azurerm_data_factory_linked_service_postgresql`
    * `azurerm_data_factory_linked_service_sftp`
    * `azurerm_data_factory_linked_service_snowflake`
    * `azurerm_data_factory_linked_service_sql_server`
    * `azurerm_data_factory_linked_service_synapse`
    * `azurerm_data_factory_linked_service_web`
    * `azurerm_data_factory_trigger_blob_event`
    * `azurerm_data_factory_trigger_schedule`
    * `azurerm_data_factory_trigger_tumbling_window`
*   `disk`
    * `azurerm_managed_disk`
*   `dns`
    * `azurerm_dns_a_record`
    * `azurerm_dns_aaaa_record`
    * `azurerm_dns_caa_record`
    * `azurerm_dns_cname_record`
    * `azurerm_dns_mx_record`
    * `azurerm_dns_ns_record`
    * `azurerm_dns_ptr_record`
    * `azurerm_dns_srv_record`
    * `azurerm_dns_txt_record`
    * `azurerm_dns_zone`
*   `load_balancer`
    * `azurerm_lb`
    * `azurerm_lb_backend_address_pool`
    * `azurerm_lb_nat_rule`
    * `azurerm_lb_probe`
*   `eventhub`
    * `azurerm_eventhub_namespace`
    * `azurerm_eventhub`
    * `azurerm_eventhub_consumer_group`
    * `azurerm_eventhub_namespace_authorization_rule`
*   `network_interface`
    * `azurerm_network_interface`
*   `network_security_group`
    * `azurerm_network_security_group`
    * `azurerm_network_security_rule`
*   `network_watcher`
    * `azurerm_network_watcher`
    * `azurerm_network_watcher_flow_log`
    * `azurerm_network_packet_capture`
*   `private_dns`
    * `azurerm_private_dns_a_record`
    * `azurerm_private_dns_aaaa_record`
    * `azurerm_private_dns_cname_record`
    * `azurerm_private_dns_mx_record`
    * `azurerm_private_dns_ptr_record`
    * `azurerm_private_dns_srv_record`
    * `azurerm_private_dns_txt_record`
    * `azurerm_private_dns_zone`
    * `azurerm_private_dns_zone_virtual_network_link`
*   `private_endpoint`
    * `azurerm_private_endpoint`
    * `azurerm_private_link_service`
*   `public_ip`
    * `azurerm_public_ip`
    * `azurerm_public_ip_prefix`
*   `redis`
    * `azurerm_redis_cache`
*   `purview`
    * `azurerm_purview_account`
*   `resource_group`
    * `azurerm_resource_group`
    * `azurerm_management_lock`
*   `route_table`
    * `azurerm_route_table`
    * `azurerm_route`
    * `azurerm_route_filter`
*   `scaleset`
    * `azurerm_virtual_machine_scale_set`
*   `security_center`
    * `azurerm_security_center_contact`
    * `azurerm_security_center_subscription_pricing`
*   `storage_account`
    * `azurerm_storage_account`
    * `azurerm_storage_blob`
    * `azurerm_storage_container`
*   `synapse`
    * `azurerm_synapse_workspace`
    * `azurerm_synapse_sql_pool`
    * `azurerm_synapse_spark_pool`
    * `azurerm_synapse_firewall_rule`
    * `azurerm_synapse_managed_private_endpoint`
    * `azurerm_synapse_private_link_hub`
*   `virtual_machine`
    * `azurerm_ssh_public_key`
    * `azurerm_virtual_machine`
*   `virtual_network`
    * `azurerm_virtual_network`
*   `subnet`
    * `azurerm_subnet`
    * `azurerm_subnet_service_endpoint_storage_policy`
    * `azurerm_subnet_nat_gateway_association`
    * `azurerm_subnet_route_table_association`
    * `azurerm_subnet_network_security_group_association`

## Notes

### Virtual networks and subnets

Terraformer will import `azurerm_virtual_network` config with inlined subnet information swipped, in order to avoid any potential circular dependencies. To import the subnet information, please also import `azurerm_subnet`.
