### Use with Heroku

This utilizes [terraform-provider-heroku](https://registry.terraform.io/providers/heroku/heroku/latest).

Heroku organizes itself by apps. This importer tool is designed to capture complete apps with all their dependent resources like addons, domains, etc.

#### Apps by ID, Not Name

Apps must be identified by ID (UUID). Even though some resources may import successfully when filtering by app name, apps themselves must be identified by ID. To get an app ID, use Heroku CLI to get the top-level `id` property:

```
heroku apps:info --json --app=<NAME>
```

#### App Config Vars

When importing apps, their settable config vars (those not from add-ons) are added to the Terraform configuration as `config_vars`. These may contain secrets, and can manually be split into `sensitive_config_vars` before the plan/apply.

#### Builds

The imported configuration cannot build & launch apps in a new place. To launch apps that have been imported with Terraformer, one of the following is required:
* source pushed to the new Heroku apps, `git push heroku master` from each app's repo
* new apps added to an existing Heroku pipelines and promoted to, via the web dashbord or CLI
* new apps connected for GitHub deployments, via the web dashboard
* a [`heroku_build` resource](https://registry.terraform.io/providers/heroku/heroku/latest/docs/resources/build) added to the Terraform configuration.

#### Example

✏️  *Please replace angle-bracketed* `<VALUES>` *with your specific values.*

```
export HEROKU_API_KEY=<token>

# All team's apps
./terraformer import heroku --resources=app --team=<NAME>

# Specific app(s), by UUID
./terraformer import heroku --resources=app --filter=app=<ID>
./terraformer import heroku --resources=app --filter=app=<ID>:<ID2>:<ID3>

# Output directory
./terraformer import heroku --resources=app --filter=app=<ID> --path-pattern='{output}/{provider}/<DIRECTORY NAME>'

# All enabled features of HEROKU_API_KEY's Heroku account
./terraformer import heroku --resources=account_feature
```

Heroku Terraformer resources with the terraform-provider-heroku resources they import:

*   `account_feature`
    * `heroku_account_feature`
*   `app`
    * `heroku_app`
    * `heroku_app_feature`
    * `heroku_app_webhook`
    * `heroku_addon`
    * `heroku_addon_attachment` (includes attachments to other apps)
    * `heroku_domain`
    * `heroku_drain`
    * `heroku_formation`
    * `heroku_ssl`
*   `pipeline`
    * `heroku_pipeline`
*   `pipeline_coupling`
    * `heroku_pipeline_coupling`
*   `team_collaborator`
    * `heroku_team_collaborator`
*   `team_member`
    * `heroku_team_member`

