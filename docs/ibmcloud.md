### Use with IBM Cloud

If you want to run Terraformer with the IBM Cloud provider plugin on your system, complete the following steps:


1. Export IBM Cloud API key as environment variables.
    Example:

    ```
    export IC_API_KEY=<IBMCLOUD_API_KEY>
    terraformer import ibm -r ibm_cos,ibm_iam....
    ```
2. Use flag for Resource Group to classify resources accordingly.
    Example:

    ```
    export IC_API_KEY=<IBMCLOUD_API_KEY>
    terraformer import ibm --resources=ibm_is_vpc --resource_group=default
    terraformer import ibm --resources=ibm_function --region=us-south
    ```
List of supported IBM Cloud resources:

*   `ibm_kp`
    * `ibm_resource_instance`
    * `ibm_kms_key`
*   `ibm_cos`
    * `ibm_resource_instance`
    * `ibm_cos_bucket`
*   `ibm_iam`
    * `ibm_iam_user_policy`
    * `ibm_iam_access_group`
    * `ibm_iam_access_group_members`
    * `ibm_iam_access_group_policy`
    * `ibm_iam_access_group_dynamic_rule`
    * `ibm_iam_service_id`
    * `ibm_iam_authorization_policy`
    * `ibm_iam_custom_role`
    * `ibm_iam_service_policy`
*   `ibm_container_vpc_cluster`
    * `ibm_container_vpc_cluster`
    * `ibm_container_vpc_worker_pool`
*   `ibm_database_etcd`
    * `ibm_database`
*   `ibm_database_mongo`
    * `ibm_database`
*   `ibm_database_postgresql`
    * `ibm_database`
*   `ibm_database_rabbitmq`
    * `ibm_database`
*   `ibm_database_redis`
    * `ibm_database`
*   `ibm_is_instance_group`
    * `ibm_is_instance_group`
    * `ibm_is_instance_group_manager`
    * `ibm_is_instance_group_manager_policy`
*   `ibm_cis`
    * `ibm_cis`
    * `ibm_cis_dns_record`
    * `ibm_cis_firewall`
    * `ibm_cis_domain_settings`
    * `ibm_cis_global_load_balancer`
    * `ibm_cis_edge_functions_action`
    * `ibm_cis_edge_functions_trigger`
    * `ibm_cis_healthcheck`
    * `ibm_cis_rate_limit` 
    * `ibm_cis_domain`
    * `ibm_cis_origin_pool`
    * `ibm_cis_waf_package`
    * `ibm_cis_waf_group`
    * `ibm_cis_page_rule`
    * `ibm_cis_custom_page`
    * `ibm_cis_range_app`
    * `ibm_cis_certificate_order`
    * `ibm_cis_routing`
    * `ibm_cis_cache_settings`
    * `ibm_cis_tls_settings`
    * `ibm_cis_filter`
*   `ibm_is_vpc`
    * `ibm_is_vpc`
    * `ibm_is_vpc_address_prefix`
    * `ibm_is_vpc_route`
    * `ibm_is_vpc_routing_table`
    * `ibm_is_vpc_routing_table_route`
*   `ibm_is_subnet`
*   `ibm_is_instance`
* `ibm_is_security_group`
    * `ibm_is_security_group`
    * `ibm_is_security_group_rule`
*   `ibm_is_network_acl`
*   `ibm_is_public_gateway`
*   `ibm_is_volume`
* `ibm_is_vpn_gateway`
    * `ibm_is_vpn_gateway`
    * `ibm_is_vpn_gateway_connections`
*   `ibm_is_lb`
    * `ibm_is_lb_pool`
    * `ibm_is_lb_pool_member`
    * `ibm_is_lb_listener`
    * `ibm_is_lb_listener_policy`
    * `ibm_is_lb_listener_policy_rule`
*   `ibm_is_floating_ip`
*   `ibm_is_flow_log`
*   `ibm_is_ike_policy`
*   `ibm_is_image`
*   `ibm_is_instance_template`
*   `ibm_is_ipsec_policy`
*   `ibm_is_ssh_key`
*   `ibm_function`
    * `ibm_function_package`
    * `ibm_function_action`
    * `ibm_function_rule`
    * `ibm_function_trigger`
* `ibm_private_dns`
    * `ibm_resource_instance`
    * `ibm_dns_zone`
    * `ibm_dns_resource_record`
    * `ibm_dns_permitted_network`
    * `ibm_dns_glb_monitor`
    * `ibm_dns_glb_pool`
    * `ibm_dns_glb`
* `ibm_transit_gateway`
    * `ibm_tg_gateway`
    * `ibm_tg_connection`
* `ibm_direct_link`
    * `ibm_dl_gateway`
    * `ibm_dl_virtual_connection`
    * `ibm_dl_provider_gateway`
* `ibm_container_cluster`
    * `ibm_container_cluster`
    * `ibm_container_worker_pool`
    * `ibm_container_worker_pool_zone_attachment`
* `ibm_certificate_manager`
    * `ibm_resource_instance`
    * `ibm_certificate_manager_import`  
    * `ibm_certificate_manager_order`  
* `ibm_vpe_gateway`
    * `ibm_is_virtual_endpoint_gateway`
    * `ibm_is_virtual_endpoint_gateway_ip`
* `ibm_satellite`
    * `ibm_satellite_location`
    * `ibm_satellite_host`    
