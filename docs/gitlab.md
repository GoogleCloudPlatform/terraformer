### Use with GitLab

Example:

```
 ./terraformer import gitlab --group=YOUR_GROUP --resources=projects --token=YOUR_TOKEN // or GITLAB_TOKEN in env
 ./terraformer import gitlab --group=YOUR_GROUP --resources=projects --filter=repository=id1:id2:id4 --token=YOUR_TOKEN // or GITLAB_TOKEN in env

  ./terraformer import gitlab --group=YOUR_GROUP --resources=projects --base-url=https://your-self-hosted-gitlab-url
```

Supports only organizational resources. List of supported resources:

*   `projects`
    * `gitlab_project`
    * `gitlab_project_value`