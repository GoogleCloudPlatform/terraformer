### Use with Heroku

Example:

```
export HEROKU_EMAIL=[HEROKU_EMAIL]
export HEROKU_API_KEY=[HEROKU_API_KEY]
./terraformer import heroku -r app,addon
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
*   `app_config_association`
    * `heroku_app_config_association`
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

