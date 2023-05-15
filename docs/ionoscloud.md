# Use with IONOS Cloud

## Authentication

Proper credential must be configured, before it can be used.

You can set the environment variables for *HTTP basic authentication*:

export IONOS_USERNAME="username"
export IONOS_PASSWORD="password"

Or you can use *token authentication*:

export IONOS_TOKEN="token"


## List of supported IONOS Cloud resources

* [`datacenter`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/datacenter)
* [`lan`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/lan)
* [`server`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/server)
* [`volume`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/volume)
* [`nic`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/nic)
* [`firewall`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/firewall)
* [`group`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/group)
* [`backup_unit`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/backup_unit)
* [`ipblock`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/ipblock)
* [`k8s_cluster`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/k8s_cluster)
* [`k8s_node_pool`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/k8s_node_pool)
* [`target_group`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/target_group)
* [`networkloadbalancer`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/networkloadbalancer)
* [`natgateway`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/natgateway)
* [`natgateway_rule`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/natgateway_rule)
* [`application_loadbalancer`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/application_loadbalancer)
* [`networkloadbalancer_forwardingrule`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/networkloadbalancer_forwardingrule)
* [`loadbalancer`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/loadbalancer)
* [`dbaas_pgsql_cluster`](https://registry.terraform.io/providers/ionos-cloud/ionoscloud/latest/docs/resources/dbaas_pgsql_cluster)

We allow only resources that provide valid terraform plans to be imported.
If you do not see your resource in the tf plan, please enable TF_LOG=debug and check logs 
for a message that will let you know why the resource was not imported.

#### Notes:
 - A server must have a `NIC` and a `volume` attached to be allowed to be imported by terraformer.
 - A server must also have a `BootVolume` set.


