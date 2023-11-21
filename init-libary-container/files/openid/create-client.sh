# Description: Create a new client in OpenID Connect
while [ $(curl -sw '%{http_code}' $OPENID_URL -k -o /dev/null) -ne 200 ]; do
  sleep 2;
  echo "Waiting for OpenID Connect to be ready..."
done

wget --header="Content-Type: application/json" --post-data='{
    "client_name": "'"$CLIENT_NAME"'",
    "client_id": "'"$CLIENT_ID"'",
    "client_secret": "'"$CLIENT_SECRET"'",
    "scope": "mail user",
    "grant_types": ["authorization_code", "refresh_token", "client_credentials", "implicit"],
    "response_types": ["token", "code", "id_token"],
    "grant_scopes": "mail",
    "access_token_strategy": "opaque"
}' --header="Authorization: Bearer $ACCESS_TOKEN" --no-check-certificate -O output.json $OPENID_URL

if [ $? -ne 0 || $? -ne 409 ]; then
  echo "Error creating client"
  exit 1
fi