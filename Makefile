compose:
	docker compose -f ./.dockers/docker-compose.yml up

composebuild:
	docker compose -f ./.dockers/docker-compose.yml up --build

composebuilddetached:
	docker compose -f ./.dockers/docker-compose.yml up --build -d

protouser:
	protoc --experimental_allow_proto3_optional \
	  --go_out=./pkg/proto --go_opt=paths=source_relative \
	  --go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative \
	  ./sa_proto/user/service.proto && \
	mv ./pkg/proto/sa_proto/user/* ./pkg/proto/user/ && \
	rm -rf ./pkg/proto/sa_proto

protoauth:
	protoc --experimental_allow_proto3_optional \
	  --go_out=./pkg/proto --go_opt=paths=source_relative \
	  --go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative \
	  ./sa_proto/auth/service.proto && \
	mv ./pkg/proto/sa_proto/auth/* ./pkg/proto/auth/ && \
	rm -rf ./pkg/proto/sa_proto

.PHONY: compose composebuild protousers protoauth