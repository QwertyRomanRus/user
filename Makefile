include .env
BIN_FOLDER:=$(CURDIR)/bin

install-deps:
	GOBIN=$(BIN_FOLDER) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.3
	GOBIN=$(BIN_FOLDER) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	GOBIN=$(BIN_FOLDER) go install github.com/pressly/goose/v3/cmd/goose@v3.24.1

get-deps:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

#local goose
lg-migration-status:
	${BIN_FOLDER}/goose -dir ${DB_MIGRATION_FOLDER} postgres ${DB_DSN} status -v

lg-migration-up:
	${BIN_FOLDER}/goose -dir ${DB_MIGRATION_FOLDER} postgres ${DB_DSN} up -v

lg-migration-down:
	${BIN_FOLDER}/goose -dir ${DB_MIGRATION_FOLDER} postgres ${DB_DSN} down -v

# grpc
grpc-generate:
	make grpc-user-generate

grpc-user-generate:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user_v1/user.proto

# docker
docker-build:
	docker build --no-cache -t ${APP_NAME} .
docker-stop:
	docker stop ${APP_NAME}
docker-rm:
	docker rm ${APP_NAME}
docker-run:
	docker run -d -p ${GRPC_PORT}:50051 --network=${NETWORK} --name ${APP_NAME} ${APP_NAME}
docker-start:
	docker start ${APP_NAME}