### Use with Auth0

Example:

```
$ export AUTH0_DOMAIN=<DOMAIN>
$ export AUTH0_CLIENT_ID=<CLIENT_ID>
$ export AUTH0_CLIENT_SECRET=<CLIENT_SECRET>
$ terraformer import auth0 --resources=rule,user
```

List of supported Auth0 services:


*   `action`
    * `auth0_action`
*   `client`
    * `auth0_client`
*   `client_grant`
    * `auth0_client_grant`
*   `hook`
    * `auth0_hook`
*   `resource_server`
    * `auth0_resource_server`
*   `role`
    * `auth0_role`
*   `rule`
    * `auth0_rule`
*   `rule_config`
    * `auth0_rule_config`
*   `trigger`
    * `auth0_trigger`
*   `user`
    * `auth0_user`

