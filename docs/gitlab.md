### Use with GitLab

Example:

```shell
./terraformer import gitlab --group=GROUP_TO_IMPORT --resources=projects --token=YOUR_TOKEN # or GITLAB_TOKEN in env
./terraformer import gitlab --group=GROUP_TO_IMPORT --resources=groups --base-url=https://your-self-hosted-gitlab-domain/api/v4
```

List of supported resources:

* `projects`
  * `gitlab_project`
  * `gitlab_project_value`
  * `gitlab_project_membership`
  * `gitlab_tag_protection`
  * `gitlab_branch_protection`
* `groups`
  * `gitlab_group_membership`
  * `gitlab_group_variable`
