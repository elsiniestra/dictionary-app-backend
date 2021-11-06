run:
	go run src/cmd/api/main.go

build:
	go build -o src/cmd/api/rest-api src/cmd/api/main.go

linux-build:
	GOOS=linux GOARCH=amd64 go build -o src/cmd/api/rest-api src/cmd/api/main.go

compose-up:
	docker-compose -f deployment/docker-compose.yml up --build

compose-down:
	docker-compose -f deployment/docker-compose.yml down

lint:
	gofmt -l -w .
	gofumpt -l -w .
	golangci-lint run --enable-all --disable tagliatelle --disable exhaustivestruct
