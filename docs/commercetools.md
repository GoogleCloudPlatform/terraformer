### Use with [Commercetools](https://commercetools.com/de/)

This provider use the [terraform-provider-commercetools](https://github.com/labd/terraform-provider-commercetools). The terraformer provider was build by [Dustin Deus](https://github.com/StarpTech).

Example:

Export required variables:

```bash
export CTP_PROJECT_KEY=key
export CTP_CLIENT_ID=foo
export CTP_CLIENT_SECRET=bar
export CTP_CLIENT_SCOPE=scope
```

Export optional variables in case default values are not appropriate:

```bash
export CTP_BASE_URL=base_url # default: https://api.sphere.io
export CTP_TOKEN_URL=token_url # default: https://auth.sphere.io
```

Run terraformer

```bash
./terraformer plan commercetools -r=types # Only planning
./terraformer import commercetools -r=types # Import commercetools types
```

List of supported [commercetools](https://commercetools.com/de/) resources:

- `api_extension`
  - `commercetools_api_extension`
- `channel`
  - `commercetools_channel`
- `custom_object`
  - `commercetools_custom_object`
- `product_type`
  - `commercetools_product_type`
- `shipping_method`
  - `commercetools_shipping_method`
- `shipping_zone`
  - `commercetools_shipping_zone`
- `state`
  - `commercetools_state`
- `store`
  - `commercetools_store`
- `subscription`
  - `commercetools_subscription`
- `tax_category`
  - `commercetools_tax_category`
- `types`
  - `commercetools_type`
