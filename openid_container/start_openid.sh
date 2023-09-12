
if (cd ./db && make run-postgress); then
    echo "Created Postgress"
else
    exit 2
fi

sleep 10
if (cd ./server && make container migrate-openid run-openid); then
    echo "Created OpenID Server and Migrated"
else
    exit 2
fi

sleep 10
if (cd ./login && make container run-openid-login); then
    echo "Created OpenID Login"
else
    exit 2
fi

if (sh ./custom_api_requests/create_client.sh); then
    echo $1
else
    exit 2
fi