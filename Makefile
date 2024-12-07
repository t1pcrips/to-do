LOCAL_BIN := $(CURDIR)/bin
LOCAL_MIGRATION_DIR := $(CURDIR)/migrations

install-deps:
	GOBIN=$(LOCAL_BIN) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

get-deps:
	go get -u github.com/golang-migrate/migrate/v4

migrate-new:
	mkdir -p migrations
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) create ${NAME} sql

migrate-up:
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) up -v

migrate-down:
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) down -v

migrate-reset:
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) reset -v

gen-openapi:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package api openapi/openapi.yaml > ./internal/web/tasks/tasks.gen.go
	oapi-codegen -config openapi/.openapi -include-tags users -package api openapi/openapi.yaml > ./internal/web/users/users.gen.go

run-server:
	go run cmd/main.go

lint:
