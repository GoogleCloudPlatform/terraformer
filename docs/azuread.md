### Use with Azure Active Directory

Example:

```
$ export ARM_TENANT_ID=<TENANT_ID>
$ export ARM_CLIENT_ID=<CLIENT_ID>
$ export ARM_CLIENT_SECRET=<CLIENT_SECRET>
$ terraformer import azuread --resources=user,application
```

List of supported AzureAD services:

*   `app_role_assignment`
    * `azuread_app_role_assignment`
*   `application`
    * `azuread_application`
*   `group`
    * `azuread_group`
*   `service_principal`
    * `azuread_service_principal`
*   `user`
    * `azuread_user`
