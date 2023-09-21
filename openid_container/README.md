# Information
## Why is this needed?
- To provide minimal security approach to the REST-API, i used the OpenID - ory container. Only with this generated client-id/client-secret capturing accepted bearer token is possible

## How to run?
- check in advance postgres container is running `database_container/start_postgres.sh`
- `start_openid.sh` this starts neccessary container related to OpenID
- `stop_openid.sh` this stops all containers related to OpenID
- all commands within these files can be executed seperatly manual
- some interaction with OpenID Container can be found in `./custom_api_requests`