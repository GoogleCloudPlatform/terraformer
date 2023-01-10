### Use with GitLab

Optional parameters:
* `base-url` - Base API url for your GitLab instance. Defaults to GitLab.com if not provided
* `include-sub-groups` - Controls whether to traverse down to subgroups of the the given group, to include projects within the subgroup structure. Defaults to false if not provided
Example:

```shell
./terraformer import gitlab --group=GROUP_TO_IMPORT --resources=projects --token=YOUR_TOKEN --include-sub-groups=true # or GITLAB_TOKEN in env
./terraformer import gitlab --group=GROUP_TO_IMPORT --resources=groups --base-url=https://your-self-hosted-gitlab-domain/api/v4 --include-sub-groups=true
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
