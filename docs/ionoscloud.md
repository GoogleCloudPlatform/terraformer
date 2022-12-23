# Use with IONOS Cloud

## Authentication

Proper credential must be configured, before it can be used.

You can set the environment variables for *HTTP basic authentication*:

export IONOS_USERNAME="username"
export IONOS_PASSWORD="password"

Or you can use *token authentication*:

export IONOS_TOKEN="token"


## List of supported IONOS Cloud resources

* `datacenter`
* `lan`
* `server`
* `volume`
* `nic`
* `firewall`
* `group`
* `backup_unit`
* `ipblock`
* `k8s_cluster`
* `k8s_node_pool`
* `target_group`
* `networkloadbalancer`
* `natgateway`
* `group`
* `application_loadbalancer`
* `firewall`
* `networkloadbalancer_forwardingrule`
* `loadbalancer`
* `natgateway_rule`
* `pg_cluster`

We allow only resources that provide valid terraform plans to be imported.
If you do not see your resource in the tf plan, please enable TF_LOG=debug and check logs 
for a message that will let you know why the resource was not imported.

#### Notes:
 - A server must have a `NIC` and a `volume` attached to be allowed to be imported by terraformer.
 - A server must also have a `BootVolume` set.


