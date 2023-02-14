### Use with Heroku

Heroku organizes itself by teams and apps. This importer tool is designed to capture complete apps with all their dependent resources like addons, domains, etc.

#### Apps by ID, Not Name

Apps must be identified by ID (UUID). Even though some resources may import successfully when filtering by app name, apps themselves must be identified by ID. To get an app ID, use Heroku CLI to get the top-level `id` property:

```
heroku apps:info --json --app=NAME
```

#### App Config Vars

When importing apps, their settable config vars (those not from add-ons) are added to the Terraform configuration as `config_vars`. These may contain secrets, and can manually be split into `sensitive_config_vars` before the plan/apply.

#### Example

```
export HEROKU_API_KEY=<token>

./terraformer import heroku --resources=app --team=<name>
./terraformer import heroku --resources=app --filter=app=<ID>

./terraformer import heroku --resources=app,addon,addon_attachment,app_feature --filter=app=<ID>

./terraformer import heroku --resources=account_feature
```

Heroku Terraformer resources with the terraform-provider-heroku resources they import:

*   `account_feature`
    * `heroku_account_feature`
*   `addon`
    * `heroku_addon`
    * requires `app` filter
*   `addon_attachment`
    * `heroku_addon_attachment`
    * requires `app` filter
    * imports all attachments of the app's add-ons
*   `app`
    * `heroku_app`
    * requires `--team` name or `app` filter
*   `app_feature`
    * `heroku_app_feature`
    * requires `app` filter
    * imports only `enabled = true` features
*   `app_webhook`
    * `heroku_app_webhook`
*   `build`
    * `heroku_build`
*   `cert`
    * `heroku_cert`
*   `domain`
    * `heroku_domain`
*   `drain`
    * `heroku_drain`
*   `formation`
    * `heroku_formation`
*   `pipeline`
    * `heroku_pipeline`
*   `pipeline_coupling`
    * `heroku_pipeline_coupling`
*   `team_collaborator`
    * `heroku_team_collaborator`
*   `team_member`
    * `heroku_team_member`

