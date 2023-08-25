## Use with Sumologic

Example:

```
./terraformer import sumologic --resources=user --access-id=SUMOLOGIC_ACCESS_ID --access-key=SUMOLOGIC_ACCESS_KEY --environment=SUMOLOGIC_ENVIRONMENT
```

You can also export env vars instead of providing them via args.
```
export SUMOLOGIC_ENVIRONMENT="us2"
export SUMOLOGIC_ACCESS_ID="US2_ACCESS_ID"
export SUMOLOGIC_ACCESS_KEY="US2_ACCESS_KEY"

./terraformer import sumologic --resources=user
```

List of supported Sumologic services:

* `connection`
    * `sumologic_connection`
* `dashboard`
    * `sumologic_dashboard`
* `field_extraction_rule`
    * `sumologic_field_extraction_rule`
* `monitor`
    * `sumologic_monitor`
* `partition`
    * `sumologic_partition`
* `role`
    * `sumologic_role`

      _Roles can be imported by name, which can be passed via `--filter` option._
      ```
      terraformer import sumologic --resources=role --filter 'Name=name;Value=App Admin'
      ```
* `user`
    * `sumologic_user`

      _Users can be imported by email, which can be passed via `--filter` option._
      ```
      terraformer import sumologic --resources=user --filter 'Name=email;Value=johndoe@sumologic.com'
      ```
