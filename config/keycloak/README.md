## How-to

### Export configurations
All the new users, roles, configurations, clients... created or modified using Keycloak realm administration dashboard won't be persisted on container restart.

If you need to export the realm config for later usage, you can overwrite the `realms/` contents by spawning the following command (make sure the container is up and running):

```shell
    $ make save-keycloak-config
```