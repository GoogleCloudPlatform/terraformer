### Use with Equinix Metal

Example:

```
export METAL_AUTH_TOKEN=[METAL_AUTH_TOKEN]
export PACKET_PROJECT_ID=[PROJECT_ID]
./terraformer import metal -r volume,device
```

List of supported Equinix Metal resources:

*   `device`
    * `metal_device`
*   `spotmarketrequest`
    * `metal_spot_market_request`
*   `sshkey`
    * `metal_ssh_key`
*   `volume`
    * `metal_volume`
