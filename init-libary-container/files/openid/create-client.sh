# Description: Create a new client in OpenID Connect
wget --no-check-certificate --quiet \
  --method POST \
  --timeout=0 \
  --header 'Content-Type: application/json' \
  --body-data '{
    "client_name": "libary",
    "client_secret": $CLIENT_SECRET,
    "scope": "mail user",
    "grant_types": ["authorization_code","refresh_token","client_credentials","implicit"],
    "response_types": ["token", "code", "id_token"],
    "grant_scopes": "mail",
    "access_token_strategy": "opaque"
}' \ $OPENID_URL
