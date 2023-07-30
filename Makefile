.PHONY: clean critic security lint test build run

APP_NAME = apiserver
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = postgres://postgres:12345678@localhost:5432/go_fiber?sslmode=disable

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

test: clean critic security lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

generate:
	go generate ./...

# development
dev: generate
	go run github.com/cosmtrek/air

install:
	cd .. && go install github.com/golang-migrate/migrate/v4 && \
	go install -u github.com/swaggo/swag/cmd/swag && go install -u github.com/cosmtrek/air && \
	go install github.com/vektra/mockery/v2/.../ && \
	cd ${APP_NAME} && swag init