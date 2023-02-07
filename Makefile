APP_BIN = build/app

.PHONY: run
run:
	go run ./app/cmd/app/main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./app/cmd/app/main.go

.PHONY: clean
clean:
	rm -rf ./app/build || true

.PHONY: swagger
swagger:
	swag init -g ./app/main.go -o ./app/docs

.PHONY: migrate
migrate:
	go run ./app/cmd/migrations/main.go
