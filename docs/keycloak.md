### Use with Keycloak

Example:

```
 export KEYCLOAK_URL=https://foo.bar.localdomain
 export KEYCLOAK_BASE_PATH=/auth # Only users of the legacy Wildfly distribution will need to set this.
 export KEYCLOAK_CLIENT_ID=[KEYCLOAK_CLIENT_ID]
 export KEYCLOAK_CLIENT_SECRET=[KEYCLOAK_CLIENT_SECRET]
 export RED_HAT_SSO=1 # Only users of the RH-SSO distribution will need to set this.

 terraformer import keycloak --resources=realms
 terraformer import keycloak --resources=realms --filter=realm=name1:name2:name3
 terraformer import keycloak --resources=realms --targets realmA,realmB
```

Here is the list of resources which are currently supported by Keycloak provider v.4.0.1:

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
  - `keycloak_openid_audience_resolve_protocol_mapper`
  - `keycloak_openid_client`
  - `keycloak_openid_client_default_scopes`
  - `keycloak_openid_client_optional_scopes`
  - `keycloak_openid_client_scope`
  - `keycloak_openid_client_service_account_role`
  - `keycloak_openid_full_name_protocol_mapper`
  - `keycloak_openid_group_membership_protocol_mapper`
  - `keycloak_openid_hardcoded_claim_protocol_mapper`
  - `keycloak_openid_hardcoded_role_protocol_mapper` (only for client roles)
  - `keycloak_openid_script_protocol_mapper` (support for this protocol mapper was removed in Keycloak 18)
  - `keycloak_openid_user_attribute_protocol_mapper`
  - `keycloak_openid_user_client_role_protocol_mapper`
  - `keycloak_openid_user_property_protocol_mapper`
  - `keycloak_openid_user_realm_role_protocol_mapper`
  - `keycloak_openid_user_session_note_protocol_mapper`
  - `keycloak_realm`
  - `keycloak_required_action`
  - `keycloak_role`
  - `keycloak_user`
