IMAGE_NAME=openid-login

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

run-openid-login:
	docker run -d \
	--name openid-login \
	-p 9020:3000 \
	-e HYDRA_ADMIN_URL=https://localhost:5445 \
	-e NODE_TLS_REJECT_UNAUTHORIZED=0 \
	$(IMAGE_NAME):latest