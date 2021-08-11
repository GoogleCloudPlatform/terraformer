### Use with Okta

Example:

```
$ export OKTA_ORG_NAME=<ORG_NAME>
$ export OKTA_BASE_URL=<BASE_URL>
$ export OKTA_API_TOKEN=<API_TOKEN>
$ terraformer import okta --resources=okta_user,okta_group
```

List of supported Okta services:

*    `user`
     * `okta_user`
*    `user_type`
     * `okta_user_type`
*    `group`
     * `okta_group`
*    `policy`
     * `okta_policy_password`
     * `okta_policy_rule_password`
     * `okta_policy_mfa`
     * `okta_policy_rule_mfa`
     * `okta_policy_signon`
     * `okta_policy_rule_signon`
*    `authorization_server`
     * `okta_auth_server`
     * `okta_auth_server_scope`
     * `okta_auth_server_claim`
     * `okta_auth_server_policy`
*    `app`
     * `okta_app_auto_login`
     * `okta_app_basic_auth`
     * `okta_app_bookmark`
     * `okta_app_oauth`
     * `okta_app_saml`
     * `okta_app_secure_password_store`
     * `okta_app_swa`
     * `okta_app_three_field`
*    `event_hook`
*    * `okta_event_hook`
*    `factor`
     * `okta_factor`
*    `inline_hook`
*    * `okta_inline_hook`
*    `idp`
     * `okta_idp_oidc`
     * `okta_idp_saml`
     * `okta_idp_social`
*    `network_zone`
     * `okta_network_zone`
*    `template_sms`
     * `okta_template_sms`
*    `trusted_origin`
     * `okta_trusted_orgin`
