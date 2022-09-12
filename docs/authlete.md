### Use with Authlete


#### importing services definitions
Example:

```
$ export AUTHLETE_SO_KEY=<SERVICE_OWNER_KEY>
$ export AUTHLETE_SO_SECRET=<SERVICE_OWNER_SECRET>
$ terraformer import authlete --resources=authlete_service
```

#### importing clients definitions
Example:

```
$ export AUTHLETE_API_KEY=<API_KEY>
$ export AUTHLETE_API_SECRET=<API_SECRET>
$ terraformer import authlete --resources=authlete_client
```

#### importing services and client definitions
Example:

```
$ export AUTHLETE_SO_KEY=<SERVICE_OWNER_KEY>
$ export AUTHLETE_SO_SECRET=<SERVICE_OWNER_SECRET>
$ export AUTHLETE_API_KEY=<API_KEY>
$ export AUTHLETE_API_SECRET=<API_SECRET>
$ terraformer import authlete --resources=authlete_service,authlete_client
```

#### dedicated cloud or running on premise


```
$ export AUTHLETE_API_SERVER=https://<api-server-fqdn>
$ export AUTHLETE_SO_KEY=<SERVICE_OWNER_KEY>
$ export AUTHLETE_SO_SECRET=<SERVICE_OWNER_SECRET>
$ export AUTHLETE_API_KEY=<API_KEY>
$ export AUTHLETE_API_SECRET=<API_SECRET>
$ terraformer import authlete --resources=authlete_service,authlete_client
```



#### List of supported Authlete services:


* `authlete_service`
* `authlete_client`

