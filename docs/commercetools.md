### Use with [Commercetools](https://commercetools.com/de/)

This provider use the [terraform-provider-commercetools](https://github.com/labd/terraform-provider-commercetools). The terraformer provider was build by [Dustin Deus](https://github.com/StarpTech).

Example:

```
CTP_CLIENT_ID=foo CTP_CLIENT_SCOPE=scope CTP_CLIENT_SECRET=bar CTP_PROJECT_KEY=key ./terraformer plan commercetools -r=types // Only planning
CTP_CLIENT_ID=foo CTP_CLIENT_SCOPE=scope CTP_CLIENT_SECRET=bar CTP_PROJECT_KEY=key ./terraformer import commercetools -r=types // Import commercetools types
```

List of supported [commercetools](https://commercetools.com/de/) resources:

*   `api_extension`
    * `commercetools_api_extension`
*   `channel`
    * `commercetools_channel`
*   `product_type`
    * `commercetools_product_type`
*   `shipping_method`
    * `commercetools_shipping_method`
*   `shipping_zone`
    * `commercetools_shipping_zone`
*   `state`
    * `commercetools_state`
*   `store`
    * `commercetools_store`
*   `subscription`
    * `commercetools_subscription`
*   `tax_category`
    * `commercetools_tax_category`
*   `types`
    * `commercetools_type`
