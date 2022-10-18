build: format test
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/main .

test:
	go test ./...

format:
	go fmt ./...

clean:
	rm -rf ./bin/main

lint:
	golangci-lint run ./...


docker-build:
	docker build -t test-client -f ./Dockerfile .
