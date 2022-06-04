### Use with Snowflake
Example:

```
export SNOWFLAKE_USER=[SNOWFLAKE_EMAIL]
export SNOWFLAKE_PASSWORD=[SNOWFLAKE_PASSWORD]
export SNOWFLAKE_ROLE=[SNOWFLAKE_ROLE]
export SNOWFLAKE_ACCOUNT=[SNOWFLAKE_ACCOUNT]
export SNOWFLAKE_REGION=[SNOWFLAKE_REGION]
./terraformer import snowflake -r database
```

List of supported Snowflake resources:

* `database`
  * `snowflake_database`
* `database_grant`
  * `snowflake_database_grant`
* `role`
  * `snowflake_role`
* `role_grant`
  * `snowflake_role_grants`
* `schema`
  * `snowflake_schema`
* `schema_grant`
  * `snowflake_schema_grant`
* `user`
  * `snowflake_user`
* `view`
  * `snowflake_view`
* `warehouse`
  * `snowflake_warehouse`
* `warehouse_grant`
  * `snowflake_warehouse_grant`
