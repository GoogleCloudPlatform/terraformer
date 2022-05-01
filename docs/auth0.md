### Use with Auth0

Example:

```
$ export AUTH0_DOMAIN=<DOMAIN>
$ export AUTH0_CLIENT_ID=<CLIENT_ID>
$ export AUTH0_CLIENT_SECRET=<CLIENT_SECRET>
$ terraformer import auth0 --resources=rule,user
```

List of supported Auth0 services:


* `auth0_action`
* `auth0_client`
* `auth0_client_grant`
* `auth0_hook`
* `auth0_resource_server`
* `auth0_role`
* `auth0_rule`
* `auth0_rule_config`
* `auth0_trigger_binding`
* `auth0_user`
* `auth0_branding"`
* `auth0_custom_domain`
* `auth0_email`
* `auth0_prompt"`
* `auth0_log_stream`
* `auth0_tenant`

    