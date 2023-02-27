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
	swag init -g ./app/cmd/app/main.go -o ./docs
