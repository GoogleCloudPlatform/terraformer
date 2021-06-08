### Use with PAN-OS

Example:

```
 export PANOS_HOSTNAME=192.168.1.1
 export PANOS_USERNAME=[PANOS_USERNAME]
 export PANOS_PASSWORD=[PANOS_PASSWORD]

 terraformer import panos --resources=device_config,firewall_networking,firewall_objects,firewall_policy
```
The list of usable environment variables is the same as the [pango go-client](https://github.com/PaloAltoNetworks/pango):
*  `PANOS_HOSTNAME`
*  `PANOS_USERNAME`
*  `PANOS_PASSWORD`
*  `PANOS_API_KEY`
*  `PANOS_PROTOCOL`
*  `PANOS_PORT`
*  `PANOS_TIMEOUT`
*  `PANOS_TARGET`
*  `PANOS_HEADERS`
*  `PANOS_VERIFY_CERTIFICATE`
*  `PANOS_LOGGING`

Here is the list of resources which are currently supported:

*   `device_config`
    * `panos_general_settings`
    * `panos_telemetry` 
    * `panos_email_server_profile`
    * `panos_http_server_profile`
    * `panos_snmptrap_server_profile`
    * `panos_syslog_server_profile`
*   `firewall_networking`
    * `panos_aggregate_interface`
    * `panos_bfd_profile`
    * `panos_bgp`
    * `panos_bgp_aggregate`
    * `panos_bgp_aggregate_advertise_filter`
    * `panos_bgp_aggregate_suppress_filter`
    * `panos_bgp_auth_profile` # The secret argument will contain "(incorrect)"
    * `panos_bgp_conditional_adv`
    * `panos_bgp_conditional_adv_advertise_filter`
    * `panos_bgp_conditional_adv_non_exist_filter`
    * `panos_bgp_dampening_profile`
    * `panos_bgp_export_rule_group`
    * `panos_bgp_import_rule_group`
    * `panos_bgp_peer`
    * `panos_bgp_peer_group`
    * `panos_bgp_redist_rule`
    * `panos_ethernet_interface`
    * `panos_gre_tunnel`
    * `panos_ike_crypto_profile`
    * `panos_ike_gateway`
    * `panos_ipsec_crypto_profile`
    * `panos_ipsec_tunnel`
    * `panos_ipsec_tunnel_proxy_id_ipv4`
    * `panos_layer2_subinterface`
    * `panos_layer3_subinterface`
    * `panos_loopback_interface`
    * `panos_management_profile`
    * `panos_monitor_profile`
    * `panos_redistribution_profile`
    * `panos_static_route_ipv4`
    * `panos_tunnel_interface`
    * `panos_virtual_router`
    * `panos_vlan`
    * `panos_vlan_interface`
    * `panos_zone`
*   `firewall_objects`
    * `panos_address_group`
    * `panos_administrative_tag`
    * `panos_application_group`
    * `panos_application_object`
    * `panos_edl`
    * `panos_log_forwarding_profile`
    * `panos_service_group`
    * `panos_service_object`
    * `panos_address_object`
    * `panos_anti_spyware_security_profile`
    * `panos_antivirus_security_profile`
    * `panos_custom_data_pattern_object`
    * `panos_data_filtering_security_profile`
    * `panos_dos_protection_profile`
    * `panos_dynamic_user_group`
    * `panos_file_blocking_security_profile`
    * `panos_url_filtering_security_profile`
    * `panos_vulnerability_security_profile`
    * `panos_wildfire_analysis_security_profile`
*   `firewall_policy`
    * `panos_nat_rule_group`
    * `panos_pbf_rule_group`
    * `panos_security_rule_group`
