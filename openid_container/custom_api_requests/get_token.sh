curl -k --location --request POST 'https://127.0.0.1:5444/oauth2/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--header 'Authorization: user:pass' \
--data-urlencode 'grant_type=client_credentials'