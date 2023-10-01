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