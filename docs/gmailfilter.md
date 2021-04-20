### Use with GmailFilter

Support [Using Service Accounts](https://github.com/yamamoto-febc/terraform-provider-gmailfilter/blob/master/README.md#using-a-service-accountg-suite-users-only) or [Using Application Default Credentials](https://github.com/yamamoto-febc/terraform-provider-gmailfilter/blob/master/README.md#using-an-application-default-credential).

Example:

```
# Using Service Accounts
export GOOGLE_CREDENTIALS=/path/to/client_secret.json
export IMPERSONATED_USER_EMAIL="foobar@example.com"

# Using Application Default Credentials
gcloud auth application-default login \
  --client-id-file=client_secret.json \
  --scopes \
https://www.googleapis.com/auth/gmail.labels,\
https://www.googleapis.com/auth/gmail.settings.basic

./terraformer import gmailfilter -r=filter,label
```

List of supported GmailFilter resources:

*   `label`
    * `gmailfilter_label`
*   `filter`
    * `gmailfilter_filter`
