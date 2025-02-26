ifneq ("$(wildcard app.env)", "")
	include app.env
	export $(shell sed 's/=.*//' app.env)
endif

DOCKER_COMPOSE_FILE = ./.dockers/docker-compose.yml

setup:
	$(MAKE) create-envs
	$(MAKE) compose-build-detached

compose:
	docker compose -f $(DOCKER_COMPOSE_FILE) up

compose-build:
	docker compose -f $(DOCKER_COMPOSE_FILE) up --build

compose-build-detached:
	docker compose -f $(DOCKER_COMPOSE_FILE) up --build -d

compose-down:
	docker compose -f $(DOCKER_COMPOSE_FILE) down

create-envs:
	cp .env.example app.env

proto:
	protoc --experimental_allow_proto3_optional \
		-I=sa_proto/services \
		--go_out=./pkg/proto --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative \
		sa_proto/services/auth.proto sa_proto/services/user.proto sa_proto/services/class.proto sa_proto/services/member.proto sa_proto/services/shared.proto

.PHONY: setup compose compose-build compose-build-detached compose-down create-envs proto