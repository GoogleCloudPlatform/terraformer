### Use with Yandex Cloud

Example:

```
export YC_TOKEN=[YANDEX_CLOUD_OAUTH_OR_IAM_TOKEN]
./terraformer import yandex -r subnet --folder_ids <comma-separated folder IDs>
```

List of supported Yandex resources:

*   `disk`
    * `yandex_compute_disk`
*   `instance`
    * `yandex_compute_instance`
*   `network`
    * `yandex_vpc_network`
*   `subnet`
    * `yandex_vpc_subnet`

Your `tf` and `tfstate` files are written by default to
`generated/yandex/service`.