### Use with Linode

Example:

```
export LINODE_TOKEN=[LINODE_TOKEN]
./terraformer import linode -r instance
```

List of supported Linode resources:

*   `domain`
    * `linode_domain`
    * `linode_domain_record`
*   `image`
    * `linode_image`
*   `instance`
    * `linode_instance`
*   `nodebalancer`
    * `linode_nodebalancer`
    * `linode_nodebalancer_config`
    * `linode_nodebalancer_node`
*   `rdns`
    * `linode_rdns`
*   `sshkey`
    * `linode_sshkey`
*   `stackscript`
    * `linode_stackscript`
*   `token`
    * `linode_token`
*   `volume`
    * `linode_volume`
