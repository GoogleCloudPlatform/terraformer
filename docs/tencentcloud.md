### Use with TencentCloud

Example:

```
$ export TENCENTCLOUD_SECRET_ID=<SECRET_ID>
$ export TENCENTCLOUD_SECRET_KEY=<SECRET_KEY>
$ terraformer import tencentcloud --resources=cvm,cbs --regions=ap-guangzhou
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
*    `redis`
     * `tencentcloud_redis_instance`
*    `scf`
     * `tencentcloud_scf_function`
*    `security_group`
     * `tencentcloud_security_group`
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
