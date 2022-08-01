### Use with Google Workspace

Example:

```
$ export GOOGLEWORKSPACE_CUSTOMER_ID=<CUSTOMER_ID>
$ export GOOGLEWORKSPACE_CREDENTIALS=</PATH/TO/CREDENTIALS.json>
$ export GOOGLEWORKSPACE_IMPERSONATED_USER_EMAIL=<USER_TO_IMPERSONATE@CONTOSO.COM>
$ terraformer import googleworkspace --resources=org_unit,chrome_policy
```

List of supported Google Workspace resources:

* `Directory API` - [Reference](https://developers.google.com/admin-sdk/directory/reference/rest)
     * `org_unit` - [orgunits](https://developers.google.com/admin-sdk/directory/reference/rest/v1/orgunits)
* `Chrome Policy API` - [Reference](https://developers.google.com/chrome/policy/reference/rest)
     * `chrome_policy` - [policies](https://developers.google.com/chrome/policy/reference/rest/v1/customers.policies.orgunits/batchModify)
