#KV Service 


This is an independent KV service. This service is powered by Bolt db underneath and provides an API to perform Get, Delete, and Insert operations.


## Running locally


To run locally, you can run `go run main.go`. The default ports for the kv-service is 8080 and 8081 for the test-client.

To run in a docker container, you need to run:

```
make docker-build
docker run xx
docker run xx
```

## Code Organization

There are two directories, kv-service and test-client.

`service`: This package contains the main business logic of the application. Code is organized in multiple files within the package. The service struct depends on `storage` which is the db abstraction layer.

`api`: This package contains the webserver. All methods hang of the server struct which utlizes the `net/http` package. This is decoupled from the service layer and can be changed accordingly

`storage`: This package interfaces with persistent storage. Currently, only a BoltDB implementation exists for the storage package.

`main`: Creates all dependencies and injects them into the `service` and `server`

## Postman Collection

You can import the postman collection and test the services out.
