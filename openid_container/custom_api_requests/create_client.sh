curl -k --location --request POST 'https://127.0.0.1:5445/clients' \
--header 'Content-Type: application/json' \
--data-raw '{
    "client_name": "libary",
    "client_secret": "sdjkfdshwfjkvgdne347dsbmqwuz322334guz121shybhu321",
    "scope": "mail user",
    "grant_types": ["authorization_code","refresh_token","client_credentials","implicit"],
    "response_types": ["token", "code", "id_token"],
    "grant_scopes": "mail",
    "access_token_strategy": "opaque"
}'