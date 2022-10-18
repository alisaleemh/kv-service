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

destroy:
	docker-compose down -v --remove-orphans

docker-build:
	docker build -t kv-service -f ./Dockerfile .
