### Use with RabbitMQ

Example:

```
 export RABBITMQ_SERVER_URL=http://foo.bar.localdomain:15672
 export RABBITMQ_USERNAME=[RABBITMQ_USERNAME]
 export RABBITMQ_PASSWORD=[RABBITMQ_PASSWORD]

 terraformer import rabbitmq --resources=vhosts,queues,exchanges
 terraformer import rabbitmq --resources=vhosts,queues,exchanges --filter=vhost=name1:name2:name3
```

All RabbitMQ resources that are currently supported by the RabbitMQ provider, are also supported by this module. Here is the list of resources which are currently supported by RabbitMQ provider v.1.1.0:

*   `bindings`
    * `rabbitmq_binding`
*   `exchanges`
    * `rabbitmq_exchange`
*   `permissions`
    * `rabbitmq_permissions`
*   `policies`
    * `rabbitmq_policy`
*   `queues`
    * `rabbitmq_queue`
*   `users`
    * `rabbitmq_user`
*   `vhosts`
    * `rabbitmq_vhost`
