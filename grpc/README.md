## Prerequisites
1.  One should have installed the following packages:
    -   `go`: `/usr/local/go/bin`
    -   `protoc`: `/usr/local/bin`
2.  Install the protocol compiler plugins for Go using the following commands:
    ```bash
        go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```
3.  Update your PATH so that the _protoc compiler_ can find the plugins:
    ```bash
        export PATH="$PATH:$(go env GOPATH)/bin"
    ```

## Project Setup
1.  Install the _golang/grpc_ dependency package.
    ```bash
        go get -u google.golang.org/grpc
    ```
2.  Generate the _go_ specific **gRPC** code using the _protoc_ tool
    ```bash
        cd <PATH>/golang-learning/grpc/
        protoc -I ./proto --go_out=./server --go-grpc_out=require_unimplemented_servers=false:./server --go_out=./client --go-grpc_out=require_unimplemented_servers=false:./client proto/chat.proto
    ```
    or
    ```bash
        cd <PATH>/golang-learning/grpc/proto/
        protoc --go_out=. --go-grpc_out=. chat.proto
    ```
3.  The _server_ could be setup by just running the command from the project directory
    ```bash
        cd <PATH>/golang-learning/grpc/server/
        go run server.go
    ```
4.  In order to run the _client_, one has to go to the project folder and execute `client.go` from there
    ```bash
        cd <PATH>/golang-learning/grpc/client/
        go run client.go
    ```

