# Information
## Why is this needed?
- To provide minimal security approach to the REST-API, i used the OpenID - ory container. Only with this generated client-id/client-secret capturing accepted bearer token is possible

## How to run?
- `start_openid.sh` this starts neccessary container related to OpenID -> greps client_id and client_secret
- `stop_openid.sh` this stops all containers related to OpenID
- all commands within these files can be executed seperatly manual