### Use with Azure Active Directory

Example:

```
$ export ARM_TENANT_ID=<TENANT_ID>
$ export ARM_CLIENT_ID=<CLIENT_ID>
$ export ARM_CLIENT_SECRET=<CLIENT_SECRET>
$ terraformer import azuread --resources=user,application
```

List of supported AzureAD services:


*   `user`
    * `azuread_user`
*   `group`
    * `azuread_group`
*   `application`
    * `azuread_application`
*   `service_principal`
    * `azuread_service_principal`
*   `app_role_assignment`
    * `azuread_app_role_assignment`
