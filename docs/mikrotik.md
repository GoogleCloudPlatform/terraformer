### Use with [Mikrotik](https://wiki.mikrotik.com/wiki/Manual:TOC)

This provider uses the [terraform-provider-mikrotik](https://github.com/ddelnano/terraform-provider-mikrotik). The terraformer provider was built by [Dom Del Nano](https://github.com/ddelnano).

Example:

```
## Warning! You should not expose your mikrotik creds through your bash history. Export them to your shell in a safe way when doing this for real!

MIKROTIK_HOST=router-hostname:8728 MIKROTIK_USER=username MIKROTIK_PASSWORD=password terraformer  import mikrotik -r=dhcp_lease

# Import only static IPs
MIKROTIK_HOST=router-hostname:8728 MIKROTIK_USER=username MIKROTIK_PASSWORD=password terraformer  import mikrotik -r=dhcp_lease --filter='Name=dynamic;Value=false'
```

List of supported mikrotik resources:

* `mikrotik_dhcp_lease`