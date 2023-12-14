include .env.local
LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-user-api

generate-user-api:
	protoc --proto_path api/user/v1 --go_out=pkg/user/v1 \
	--go_opt=paths=source_relative  --go-grpc_out=pkg/user/v1 \
    --go-grpc_opt=paths=source_relative api/user/v1/user.proto


generate-user-api2:
	mkdir -p pkg/user/v1
	protoc --proto_path api/user/v1 --proto_path vendor.protogen \
	--go_out=pkg/user/v1 --go_opt=paths=source_relative \
	--go-grpc_out=pkg/user/v1 --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pkg/user/v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/user/v1/user.proto


generate-note-api:
	mkdir -p pkg/note_v1
	protoc --proto_path api/note_v1 --proto_path vendor.protogen \
	--go_out=pkg/note_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/note_v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	--grpc-gateway_out=pkg/note_v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/note_v1/note.proto
lint:
	golangci-lint run ./... --config .golangci.pipeline.yaml

LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

DBSTRING="dbname=$(POSTGRES_DATABASE_NAME) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) port=$(POSTGRES_PORT) sslmode=disable"
goose-up:
	goose -dir ${MIGRATION_DIR} postgres ${DBSTRING} up -v

gen-http-reverse-proxy:
	protoc -I . --grpc-gateway_out pkg/user/v1 --grpc-gateway_opt paths=source_relative --grpc-gateway_opt standalone=true api/user/v1/user.proto

gen-swagger:
	protoc -I . --openapiv2_out api/user/v1/swagger2 api/user/v1/user.proto

gs:
	cd api/user/v1; buf generate

gen-swag:
	protoc --go_out=. --go-grpc_out=. \
	--grpc-gateway_out=. \
	--grpc-gateway_opt generate_unbound_methods=true \
	--openapiv2_out . api/user/v1/user.proto

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi

install-static:
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@latest

static-v1:
	$(LOCAL_BIN)/statik -src=api/user/v1/swagger -include='*.css,*.html,*.js,*.json,*.png'

test-unit:
	go test -v ./...

compose-run:
	docker-compose up --build

generate-auth-api:
	mkdir -p pkg/auth/v1
	protoc --proto_path api/auth/v1 \
	--go_out=pkg/auth/v1 --go_opt=paths=source_relative \
	--go-grpc_out=pkg/auth/v1 --go-grpc_opt=paths=source_relative \
	api/auth/v1/auth.proto

generate-access-api:
	mkdir -p pkg/access/v1
	protoc --proto_path api/access/v1 \
	--go_out=pkg/access/v1 --go_opt=paths=source_relative \
	--go-grpc_out=pkg/access/v1 --go-grpc_opt=paths=source_relative \
	api/access/v1/access.proto


grpc-load-test:
	ghz \
		--proto api/user/v1/user.proto \
		--call user_v1.UserV1.Get \
		--data '{"id": 9}' \
		--rps 100 \
		--total 3000 \
		--insecure \
		localhost:50051

grpc-load-test-secure:
	ghz \
		--proto api/user/v1/user.proto \
		--call user_v1.UserV1.Get \
		--data '{"id": 9}' \
		--rps 100 \
		--total 3000 \
		--skipTLS \
		localhost:50051

grpc-load-test-fail-secure:
	ghz \
		--proto api/user/v1/user.proto \
		--call user_v1.UserV1.Get \
		--data '{"id": 0}' \
		--rps 100 \
		--total 1500 \
		--skipTLS \
		localhost:50051

gen-cert:
	openssl genrsa -out ca.key 4096
	openssl req -new -x509 -key ca.key -sha256 -subj "//C=US//ST=NJ//O=CA, Inc." -days 365 -out ca.cert
	openssl genrsa -out service.key 4096
	openssl req -new -key service.key -out service.csr -config certificate.conf
	openssl x509 -req -in service.csr -CA ca.cert -CAkey ca.key -CAcreateserial \
    		-out service.pem -days 365 -sha256 -extfile certificate.conf -extensions req_ext
