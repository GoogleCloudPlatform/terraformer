### Use with Vault

Example:

```
 ./terraformer import vault --resources=aws_secret_backend_role --token=YOUR_VAULT_TOKEN // or VAULT_TOKEN in env --address=YOUR_VAULT_ADDRESS // or VAULT_ADDR in env
 ./terraformer import vault --resources=policy --filter=policy=id1:id2:id4 --token=YOUR_VAULT_TOKEN // or VAULT_TOKEN in env --address=YOUR_VAULT_ADDRESS // or VAULT_ADDR in env
```

List of supported Vault resources:

* `ad_secret_backend`
    * `ad_secret_backend`
* `ad_secret_backend_role`
    * `ad_secret_backend_role`
* `alicloud_auth_backend_role`
    * `alicloud_auth_backend_role`
* `approle_auth_backend_role`
    * `approle_auth_backend_role`
* `aws_auth_backend_role`
    * `aws_auth_backend_role`
* `aws_secret_backend`
    * `aws_secret_backend`
* `aws_secret_backend_role`
    * `aws_secret_backend_role`
* `azure_auth_backend_role`
    * `azure_auth_backend_role`
* `azure_secret_backend`
    * `azure_secret_backend`
* `azure_secret_backend_role`
    * `azure_secret_backend_role`
* `cert_auth_backend_role`
    * `cert_auth_backend_role`
* `consul_secret_backend`
    * `consul_secret_backend`
* `consul_secret_backend_role`
    * `consul_secret_backend_role`
* `database_secret_backend_role`
    * `database_secret_backend_role`
* `gcp_auth_backend`
    * `gcp_auth_backend`
* `gcp_auth_backend_role`
    * `gcp_auth_backend_role`
* `gcp_secret_backend`
    * `gcp_secret_backend`
* `generic_secret`
    * `generic_secret`
* `github_auth_backend`
    * `github_auth_backend`
* `jwt_auth_backend`
    * `jwt_auth_backend`
* `jwt_auth_backend_role`
    * `jwt_auth_backend_role`
* `kubernetes_auth_backend_role`
    * `kubernetes_auth_backend_role`
* `ldap_auth_backend`
    * `ldap_auth_backend`
* `ldap_auth_backend_group`
    * `ldap_auth_backend_group`
* `ldap_auth_backend_user`
    * `ldap_auth_backend_user`
* `nomad_secret_backend`
    * `nomad_secret_backend`
* `okta_auth_backend`
    * `okta_auth_backend`
* `okta_auth_backend_group`
    * `okta_auth_backend_group`
* `okta_auth_backend_user`
    * `okta_auth_backend_user`
* `pki_secret_backend`
    * `pki_secret_backend`
* `pki_secret_backend_role`
    * `pki_secret_backend_role`
* `policy`
    * `policy`
* `rabbitmq_secret_backend`
    * `rabbitmq_secret_backend`
* `rabbitmq_secret_backend_role`
    * `rabbitmq_secret_backend_role`
* `ssh_secret_backend_role`
    * `ssh_secret_backend_role`
* `terraform_cloud_secret_backend`
    * `terraform_cloud_secret_backend`
* `token_auth_backend_role`
    * `token_auth_backend_role`

[1]: https://github.com/GoogleCloudPlatform/terraformer/blob/master/README.md#filtering
