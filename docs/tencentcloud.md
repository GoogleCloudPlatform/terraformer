### Use with TencentCloud

Example:

```
$ export TENCENTCLOUD_SECRET_ID=<SECRET_ID>
$ export TENCENTCLOUD_SECRET_KEY=<SECRET_KEY>
$ terraformer import tencentcloud --resources=cvm,cbs --regions=ap-guangzhou
```

### Region Pattern support:

You can use `{region}` with -p or --path-pattern at any directory level.
Example layouts:

```bash
terraformer import tencentcloud -r vpc --regions=ap-guangzhou,ap-nanjing -p '{output}/{provider}/{service}/{region}'
generated
└── tencentcloud
    └── vpc
        ├── ap-guangzhou
        │   ├── outputs.tf
        │   ├── provider.tf
        │   ├── terraform.tfstate
        │   └── vpc.tf
        └── ap-nanjing
            ├── outputs.tf
            ├── provider.tf
            ├── terraform.tfstate
            └── vpc.tf
```
```bash
terraformer import tencentcloud -r vpc --regions=ap-guangzhou,ap-nanjing -p '{output}/{provider}/{region}/{service}'
generated
└── tencentcloud
    ├── ap-guangzhou
    │   └── vpc
    │       ├── outputs.tf
    │       ├── provider.tf
    │       ├── terraform.tfstate
    │       └── vpc.tf
    └── ap-nanjing
        └── vpc
            ├── outputs.tf
            ├── provider.tf
            ├── terraform.tfstate
            └── vpc.tf
```
```bash
terraformer import tencentcloud -r vpc,subnet --regions=ap-guangzhou,ap-nanjing -p '{output}/{region}'
generated
├── ap-guangzhou
│   ├── outputs.tf
│   ├── provider.tf
│   ├── subnet.tf
│   ├── terraform.tfstate
│   ├── variables.tf
│   └── vpc.tf
└── ap-nanjing
    ├── outputs.tf
    ├── provider.tf
    ├── subnet.tf
    ├── terraform.tfstate
    ├── variables.tf
    └── vpc.tf
```

List of supported TencentCloud services:

*    `as`
     * `tencentcloud_as_scaling_group`
     * `tencentcloud_as_scaling_config`
*    `cbs`
     * `tencentcloud_cbs_storage`
*    `cdn`
     * `tencentcloud_cdn_domain`
*    `cfs`
     * `tencentcloud_cfs_file_system`
*    `clb`
     * `tencentcloud_clb_instance`
     * `tencentcloud_clb_listener`
     * `tencentcloud_clb_listener_rule`
*    `cos`
     * `tencentcloud_cos_bucket`
*    `cvm`
     * `tencentcloud_instance`
*    `elasticsearch`
     * `tencentcloud_elasticsearch_instance`
*    `gaap`
     * `tencentcloud_gaap_proxy`
     * `tencentcloud_gaap_realserver`
*    `key_pair`
     * `tencentcloud_key_pair`
*    `mongodb`
     * `tencentcloud_mongodb_instance`
*    `mysql`
     * `tencentcloud_mysql_instance`
     * `tencentcloud_mysql_readonly_instance`
*    `nat_gateway`
     * `tencentcloud_nat_gateway`
*    `redis`
     * `tencentcloud_redis_instance`
*    `route-table`
     * `tencentcloud_route_table`
     * `tencentcloud_route_table_entry`
*    `scf`
     * `tencentcloud_scf_function`
*    `security_group`
     * `tencentcloud_security_group`
     * `tencentcloud_security_group_lite_rule`
*    `ssl`
     * `tencentcloud_ssl_certificate`
*    `subnet`
     * `tencentcloud_subnet`
*    `tcaplus`
     * `tencentcloud_tcaplus_cluster`
*    `vpc`
     * `tencentcloud_vpc`
*    `vpc`
     * `tencentcloud_vpn_gateway`
