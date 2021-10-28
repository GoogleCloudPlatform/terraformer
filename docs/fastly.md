### Use with Fastly

Example:

```
export FASTLY_API_KEY=[FASTLY_API_KEY]
export FASTLY_CUSTOMER_ID=[FASTLY_CUSTOMER_ID]
./terraformer import fastly -r service_v1,user
```

List of supported Fastly resources:

*   `service_v1`
    * `fastly_service_acl_entries_v1`
    * `fastly_service_compute`
    * `fastly_service_dictionary_items_v1`
    * `fastly_service_dynamic_snippet_content_v1`
    * `fastly_service_v1`
*   `user`
    * `fastly_user_v1`
*   `tls_subscription`
    * `fastly_tls_subscription`
