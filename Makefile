compose:
	docker compose -f ./.dockers/docker-compose.yml up

composebuild:
	docker compose -f ./.dockers/docker-compose.yml up --build

composebuilddetached:
	docker compose -f ./.dockers/docker-compose.yml up --build -d

proto:
	protoc --experimental_allow_proto3_optional \
		-I=sa_proto/services \
		--go_out=./pkg/proto --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative \
		sa_proto/services/auth.proto sa_proto/services/user.proto sa_proto/services/class.proto sa_proto/services/member.proto sa_proto/services/shared.proto

.PHONY: compose composebuild proto