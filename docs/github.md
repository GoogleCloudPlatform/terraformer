### Use with GitHub

Example:

```
 ./terraformer import github --owner=YOUR_ORGANIZATION --resources=repositories --token=YOUR_TOKEN // or GITHUB_TOKEN in env
 ./terraformer import github --owner=YOUR_ORGANIZATION --resources=repositories --filter=repository=id1:id2:id4 --token=YOUR_TOKEN // or GITHUB_TOKEN in env

  ./terraformer import github --owner=YOUR_ORGANIZATION --resources=repositories --base-url=https://your-enterprise-github-url
```

Supports only organizational resources. List of supported resources:

*   `members`
    * `github_membership`
*   `organization_blocks`
    * `github_organization_block`
*   `organization_projects`
    * `github_organization_project`
*   `organization_webhooks`
    * `github_organization_webhook`
*   `repositories`
    * `github_repository`
    * `github_repository_webhook`
    * `github_branch_protection`
    * `github_repository_collaborator`
    * `github_repository_deploy_key`
*   `teams`
    * `github_team`
    * `github_team_membership`
    * `github_team_repository`
*   `user_ssh_keys`
    * `github_user_ssh_key`

Notes:
* Terraformer can't get webhook secrets from the GitHub API. If you use a secret token in any of your webhooks, running `terraform plan` will result in a change being detected:
=> `configuration.#: "1" => "0"` in tfstate only.
