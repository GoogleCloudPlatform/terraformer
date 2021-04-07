### Use with Keycloak

Example:

```
 export KEYCLOAK_URL=https://foo.bar.localdomain
 export KEYCLOAK_CLIENT_ID=[KEYCLOAK_CLIENT_ID]
 export KEYCLOAK_CLIENT_SECRET=[KEYCLOAK_CLIENT_SECRET]

 terraformer import keycloak --resources=realms
 terraformer import keycloak --resources=realms --filter=realm=name1:name2:name3
 terraformer import keycloak --resources=realms --targets realmA,realmB
```

Here is the list of resources which are currently supported by Keycloak provider v.1.19.0:

- `realms`
  - `keycloak_default_groups`
  - `keycloak_group`
  - `keycloak_group_memberships`
  - `keycloak_group_roles`
  - `keycloak_ldap_full_name_mapper`
  - `keycloak_ldap_group_mapper`
  - `keycloak_ldap_hardcoded_group_mapper`
  - `keycloak_ldap_hardcoded_role_mapper`
  - `keycloak_ldap_msad_lds_user_account_control_mapper`
  - `keycloak_ldap_msad_user_account_control_mapper`
  - `keycloak_ldap_user_attribute_mapper`
  - `keycloak_ldap_user_federation`
  - `keycloak_openid_audience_protocol_mapper`
  - `keycloak_openid_client`
  - `keycloak_openid_client_default_scopes`
  - `keycloak_openid_client_optional_scopes`
  - `keycloak_openid_client_scope`
  - `keycloak_openid_client_service_account_role`
  - `keycloak_openid_full_name_protocol_mapper`
  - `keycloak_openid_group_membership_protocol_mapper`
  - `keycloak_openid_hardcoded_claim_protocol_mapper`
  - `keycloak_openid_hardcoded_group_protocol_mapper`
  - `keycloak_openid_hardcoded_role_protocol_mapper` (only for client roles)
  - `keycloak_openid_user_attribute_protocol_mapper`
  - `keycloak_openid_user_property_protocol_mapper`
  - `keycloak_openid_user_realm_role_protocol_mapper`
  - `keycloak_openid_user_client_role_protocol_mapper`
  - `keycloak_openid_user_session_note_protocol_mapper`
  - `keycloak_realm`
  - `keycloak_required_action`
  - `keycloak_role`
  - `keycloak_user`
