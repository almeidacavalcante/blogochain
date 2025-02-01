build:
	@go build -o bin/bolder

run: build
	@./bin/bolder

test:
	@go test -v ./...