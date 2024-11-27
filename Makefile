DB_DNS := "postgres://postgres:goland@localhost:5433/tasks?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DNS)

install-deps:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

get-deps:
	go get -u github.com/golang-migrate/migrate/v4

migrate-new:
	migrate create -ext sql -dir ./migrations

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

gen-openapi:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
run-server:
	go run cmd/main.go
