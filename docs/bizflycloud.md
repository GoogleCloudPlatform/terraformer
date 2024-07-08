### Use with BizflyCloud

Example:

```
export BIZFLYCLOUD_PROJECT_ID=[BIZFLYCLOUD_PROJECT_ID]
export BIZFLYCLOUD_REGION=[BIZFLYCLOUD_REGION]
export BIZFLYCLOUD_EMAIL=[BIZFLYCLOUD_EMAIL]
export BIZFLYCLOUD_PASSWORD=[BIZFLYCLOUD_PASSWORD]
./terraformer import bizflycloud -r server,database
```

List of supported BizflyCloud resources:

*   `server`
    * `bizflycloud_server`
    * `bizflycloud_network_interface_attachment`
    * `bizflycloud_network_interface`
    * `bizflycloud_volume_attachment`
    * `bizflycloud_volume`
    * `bizflycloud_vpc_network`
    * `bizflycloud_wan_ip`
    * `bizflycloud_firewall`
    * `bizflycloud_ssh_key`
*   `database`
    * `bizflycloud_cloud_database_instance`
*   `load_balancer`
    * `bizflycloud_loadbalancer`
    * `bizflycloud_loadbalancer_listener`
    * `bizflycloud_loadbalancer_l7policy`
    * `bizflycloud_loadbalancer_pool`
*   `kubernetes_engine`
    * `bizflycloud_kubernetes`