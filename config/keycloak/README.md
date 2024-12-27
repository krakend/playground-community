## How-to

### Export configurations
All the new users, roles, configurations, clients... created or modified using Keycloak realm administration dashboard won't be persisted on container restart.

If you need to export the realm config for later usage, you can overwrite the `realms/` contents by spawning the following command (make sure the container is up and running):

```
docker compose exec keycloak sh -c "cp -rp /opt/keycloak/data/h2 /tmp; \
    /opt/keycloak/bin/kc.sh export --dir /opt/keycloak/data/import \
    --realm krakend \
    --users realm_file \
    --db dev-file --db-url 'jdbc:h2:file:/tmp/h2/keycloakdb;NON_KEYWORDS=VALUE'"
```