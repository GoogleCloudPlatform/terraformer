### Use with Myra Security

Example using a Myra Security API Key and corresponding Token:

```
export MYRASEC_API_SECRET=[MYRASEC_API_SECRET]
export MYRASEC_API_KEY=[MYRASEC_API_KEY]
./terraformer import myrasec --resources=domain
```

List of supported Myra Security services:

* `cache_setting`
  * `myrasec_cache_setting`
* `dns_record`
  * `myrasec_dns_record`
* `domain`
  * `myrasec_domain`
* `error_page`
  * `myrasec_error_page`
* `ip_filter`
  * `myrasec_ip_filter`
* `maintenance`
  * `myrasec_maintenance`
* `redirect`
  * `myrasec_redirect`
* `settings`
  * `myrasec_settings`
* `tag`
  * `myrasec_tag`
* `tag_cache_setting`
  * `myrasec_tag_cache_setting`
* `tag_settings`
  * `myrasec_tag_settings`
* `tag_waf_rule`
  * `myrasec_tag_waf_rule`
* `waf_rule`
  * `myrasec_waf_rule`
