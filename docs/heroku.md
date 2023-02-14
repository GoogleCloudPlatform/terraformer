### Use with Heroku

Heroku organizes itself by teams and apps. This importer tool is designed to capture complete apps with all their dependent resources like addons, domains, etc.

Apps must be identified by ID (UUID). Even though some resources may import successfully when filtering by app name, apps themselves must be identified by ID. To get an app ID, use Heroku CLI to get the top-level `id` property:

```
heroku apps:info --json --app=NAME
```

Example:

```
export HEROKU_API_KEY=<token>

./terraformer import heroku --resources=app --team=<name>
./terraformer import heroku --resources=app --filter=app=<ID>

./terraformer import heroku --resources=app,addon --filter=app=<ID>

./terraformer import heroku --resources=account_feature
```

List of supported Heroku resources:

*   `account_feature`
    * `heroku_account_feature`
*   `addon`
    * `heroku_addon`
*   `addon_attachment`
    * `heroku_addon_attachment`
*   `app`
    * `heroku_app`
*   `app_feature`
    * `heroku_app_feature`
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

