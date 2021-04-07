### Use with Yandex

Example:

```
export YC_TOKEN=[YANDEX_CLOUD_OAUTH_TOKEN]
export YC_FOLDER_ID=[YANDEX_FOLDER_ID]
./terraformer import yandex -r subnet
```

List of supported Yandex resources:

*   `instance`
    * `yandex_compute_instance`
*   `disk`
    * `yandex_compute_disk`
*   `subnet`
    * `yandex_vpc_subnet`
*   `network`
    * `yandex_vpc_network`

Your `tf` and `tfstate` files are written by default to
`generated/yandex/service`.