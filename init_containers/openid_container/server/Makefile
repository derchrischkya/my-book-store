IMAGE_NAME=openid

$(VERBOSE).SILENT:
default: help

help:
	@echo "Usage:"
	@echo "- make container_m1 - Build the container needed to test lambda function local for M1 Processor"
	@echo "- make container - Build the container needed to test lambda function local"
	@echo "-------------------------------------------------------------------------"
	@echo "- make run_dev - Start the dev build container in an interactive way"
	@echo "- make run_prod - Start the prod build container in an interactive way"
	@echo "- make exec - Attach an additional shell to the running container"
	@echo "- make help - Print this help"

container:
	docker buildx build --platform linux/arm64 --pull -t $(IMAGE_NAME):latest .

migrate-openid:
	docker run -it --rm \
	--net=host \
  	$(IMAGE_NAME):latest \
  	migrate sql --yes postgres://libary:secret@host.docker.internal:5432/libary?sslmode=disable

run-openid:
	docker run -d \
	--name openid-server \
	-p 5444:4444 \
	-p 5445:4445 \
	-e DSN=postgres://libary:secret@host.docker.internal:5432/libary?sslmode=disable \
	-e URLS_SELF_ISSUER=https://localhost:5444/ \
	-e URLS_CONSENT=http://localhost:9020/consent \
	-e URLS_LOGIN=http://localhost:9020/login \
	$(IMAGE_NAME):latest serve all