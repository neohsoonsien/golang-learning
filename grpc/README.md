## Prerequisite
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
        protoc --go_out=. --go-grpc_out=. chat.proto
    ```

