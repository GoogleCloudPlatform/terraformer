### Use with Cloudflare

Example using a Cloudflare API Key and corresponding email:
```
export CLOUDFLARE_API_KEY=[CLOUDFLARE_API_KEY]
export CLOUDFLARE_EMAIL=[CLOUDFLARE_EMAIL]
export CLOUDFLARE_ACCOUNT_ID=[CLOUDFLARE_ACCOUNT_ID]
 ./terraformer import cloudflare --resources=firewall,dns
```

or using a Cloudflare API Token:

```
export CLOUDFLARE_API_TOKEN=[CLOUDFLARE_API_TOKEN]
export CLOUDFLARE_ACCOUNT_ID=[CLOUDFLARE_ACCOUNT_ID]
 ./terraformer import cloudflare --resources=firewall,dns
```

List of supported Cloudflare services:

* `access`
  * `cloudflare_access_application`
* `dns`
  * `cloudflare_zone`
  * `cloudflare_record`
* `firewall`
  * `cloudflare_access_rule`
  * `cloudflare_filter`
  * `cloudflare_firewall_rule`
  * `cloudflare_zone_lockdown`
  * `cloudflare_rate_limit`
* `page_rule`
  * `cloudflare_page_rule`
* `account_member`
  * `cloudflare_account_member`
