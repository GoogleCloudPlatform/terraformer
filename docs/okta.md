### Use with Okta

Example:

```
$ export OKTA_ORG_NAME=<ORG_NAME>
$ export OKTA_BASE_URL=<BASE_URL>
$ export OKTA_API_TOKEN=<API_TOKEN>
$ terraformer import okta --resources=user,group
```

List of supported Okta services:

*    `user`
     * `okta_user`
*    `user_type`     
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
*    `event_hook`
*    `inline_hook`
*    `template_sms`
*    `trusted_origin`






