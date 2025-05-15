### Use with Keboola

Example:

```
# Set up authentication
export KEBOOLA_TOKEN=your-storage-api-token
export KEBOOLA_HOST=https://connection.keboola.com  # optional, default is connection.keboola.com

# Import all resources
terraformer import keboola --resources=component_configuration,scheduler

# Import specific resource type
terraformer import keboola --resources=component_configuration
```

List of supported Keboola resources:

*   `component_configuration`
    * `keboola_component_configuration`
*   `scheduler`
    * `keboola_scheduler` 