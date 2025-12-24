.PHONY: build run test clean

API_GATEWAY_DIR=services/api-gateway
NOTE_SERVICE_DIR=services/note-service
TAG_SERVICE_DIR=services/tag-service

build:
	@$(MAKE) -C $(API_GATEWAY_DIR) build
	@$(MAKE) -C $(NOTE_SERVICE_DIR) build
	@$(MAKE) -C $(TAG_SERVICE_DIR) build

run:
	@$(MAKE) -C $(API_GATEWAY_DIR) run
	@$(MAKE) -C $(NOTE_SERVICE_DIR) run
	@$(MAKE) -C $(TAG_SERVICE_DIR) run

test:
	@$(MAKE) -C $(API_GATEWAY_DIR) test
	@$(MAKE) -C $(NOTE_SERVICE_DIR) test
	@$(MAKE) -C $(TAG_SERVICE_DIR) test

clean:
	@$(MAKE) -C $(API_GATEWAY_DIR) clean
	@$(MAKE) -C $(NOTE_SERVICE_DIR) clean
	@$(MAKE) -C $(TAG_SERVICE_DIR) clean
