### Use with Yandex Cloud

Example:

```
export YC_TOKEN=[YANDEX_CLOUD_OAUTH_OR_IAM_TOKEN]
./terraformer import yandex -r subnet --folder_ids <comma-separated folder IDs>
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