## Compile the protobuf file
    ```bash
        cd <PATH>/golang-learning/pagination/proto/
        protoc --go_out=. --go-grpc_out=. pagination.proto
    ```

## go test the pagination_test.go
    ```bash
        cd <PATH>/golang-learning/pagination/
        go test -v golang-learning/pagination
    ```