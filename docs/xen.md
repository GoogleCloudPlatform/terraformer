### Use with [Xen Orchestra](https://xen-orchestra.com/)

This provider uses the [terraform-provider-xenorchestra](https://github.com/ddelnano/terraform-provider-xenorchestra). The terraformer provider was built by [Dom Del Nano](https://github.com/ddelnano) on behalf of [Vates SAS](https://vates.fr/) who is sponsoring Dom to work on the project.

Example:

```
## Warning! You should not expose your xenorchestra creds through your bash history. Export them to your shell in a safe way when doing this for real!

XOA_URL=ws://your-xenorchestra-domain XOA_USER=username XOA_PASSWORD=password terraformer import xenorchestra -r=acl
```

List of supported xenorchestra resources:

* `xenorchestra_acl`
* `xenorchestra_resource_set`
