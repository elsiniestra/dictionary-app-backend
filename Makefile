run:
	go run cmd/api/main.go

build:
	go build -o cmd/api/rest-api cmd/api/main.go

lint:
	gofmt -l -w .
	gofumpt -l -w .
	golangci-lint run --enable-all --disable tagliatelle --disable exhaustivestruct
