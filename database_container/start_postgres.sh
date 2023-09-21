
if (make container && make run-postgress); then
    echo "Created Postgress"
else
    exit 2
fi