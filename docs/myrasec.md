### Use with Myra Security

Example using a Myra Security API Key and corresponding Token:

```
export MYRASEC_API_SECRET=[MYRASEC_API_SECRET]
export MYRASEC_API_KEY=[MYRASEC_API_KEY]
./terraformer import myrasec --resources=domain
```

List of supported Myra Security services:
* `domain`
  * `myrasec_domain`
* `dns_record`
  * `myrasec_dns_record`
* `redirect`
  * `myrasec_redirect`
* `cache_setting`
  * `myrasec_cache_setting`
* `ratelimit`
  * `myrasec_ratelimit`
* `ip_filter`
  * `myrasec_ip_filter`
* `settings`
  * `myrasec_settings`
* `waf_rule`
  * `myrasec_waf_rule`
* `maintenance`
  * `myrasec_maintenance`
* `error_page`
  * `myrasec_error_page`
