
### Use with AliCloud

You can either edit your alicloud config directly, (usually it is `~/.aliyun/config.json`)
or run `aliyun configure` and enter the credentials when prompted.

Terraformer will pick up the profile name specified in the `--profile` parameter.
It defaults to the first config in the config array.

```sh
terraformer import alicloud --resources=ecs --regions=ap-southeast-3 --profile=default
```

List of supported AliCloud resources:

* `dns`
  * `alicloud_dns`
  * `alicloud_dns_record`
* `ecs`
  * `alicloud_instance`
* `keypair`
  * `alicloud_key_pair`
* `nat`
  * `alicloud_nat_gateway`
* `pvtz`
  * `alicloud_pvtz_zone`
  * `alicloud_pvtz_zone_attachment`
  * `alicloud_pvtz_zone_record`
* `ram`
  * `alicloud_ram_role`
  * `alicloud_ram_role_policy_attachment`
* `rds`
  * `alicloud_db_instance`
* `sg`
  * `alicloud_security_group`
  * `alicloud_security_group_rule`
* `slb`
  * `alicloud_slb`
  * `alicloud_slb_server_group`
  * `alicloud_slb_listener`
* `vpc`
  * `alicloud_vpc`
* `vswitch`
  * `alicloud_vswitch`
