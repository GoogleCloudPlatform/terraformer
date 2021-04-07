### Use with Azure
Support [Azure CLI](https://www.terraform.io/docs/providers/azurerm/guides/azure_cli.html), [Service Principal with Client Certificate](https://www.terraform.io/docs/providers/azurerm/guides/service_principal_client_certificate.html) & [Service Principal with Client Secret](https://www.terraform.io/docs/providers/azurerm/guides/service_principal_client_secret.html)

Example:

```
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
```

List of supported Azure resources:

*   `analysis`
    * `azurerm_analysis_services_server`
*   `app_service`
    * `azurerm_app_service`
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
*   `network_interface`
    * `azurerm_network_interface`
*   `network_security_group`
    * `azurerm_network_security_group`
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
*   `public_ip`
    * `azurerm_public_ip`
    * `azurerm_public_ip_prefix`
*   `redis`
    * `azurerm_redis_cache
*   `resource_group`
    * `azurerm_resource_group`
*   `scaleset`
    * `azurerm_virtual_machine_scale_set`
*   `security_center`
    * `azurerm_security_center_contact`
    * `azurerm_security_center_subscription_pricing`
*   `storage_account`
    * `azurerm_storage_account`
    * `azurerm_storage_blob`
    * `azurerm_storage_container`
*   `virtual_machine`
    * `azurerm_virtual_machine`
*   `virtual_network`
    * `azurerm_virtual_network`
