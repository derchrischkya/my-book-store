IMAGE_NAME=openid

$(VERBOSE).SILENT:

default: help

help:
	@echo "Usage:"
	@echo "- make run-postgress - Build the container needed to test lambda function local for M1 Processor"
	@echo "- make container - Build the container needed to test lambda function local"
	@echo "-------------------------------------------------------------------------"
	@echo "- make run_dev - Start the dev build container in an interactive way"
	@echo "- make run_prod - Start the prod build container in an interactive way"
	@echo "- make exec - Attach an additional shell to the running container"
	@echo "- make help - Print this help"

run-postgress:
	docker run \
  --name postgres \
  -p 5432:5432 \
  -v /tmp/postgres:/var/lib/postgresql/data \
  -e POSTGRES_USER=hydra \
  -e POSTGRES_PASSWORD=secret \
  -e POSTGRES_DB=hydra \
  -d postgres:latest

