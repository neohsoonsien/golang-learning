## Install Go protocol buffers plugin
1.  Install the Go protocol buffers plugin:
    ```bash
        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest  # register gRPC server
        go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest # install  gRPC-Gateway, a reverse-proxy server which translates a RESTful HTTP API into gRPC
        go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
    ```

2.  The compiler plugin `protoc-gen-go` will be installed in `$GOBIN`, defaulting to `$GOPATH/bin`.
    Check the path location with:
    ```bash
        go env GOPATH
        go env GOBIN
    ```
    The path variable `GOPATH` is usually set to `~/go`, whereas `GOBIN` is usually not set.

3.  Update the PATH so that the protoc compiler can find the plugins:
    ```bash
        export PATH="$PATH:$(go env GOPATH)/bin"
    ```

## Compile the proto
1.  Compile the proto file to `helloworld.pg.go` with
    ```bash
        protoc -I=. --go_out=. --go_opt=paths=source_relative \
            --go-grpc_out=. --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=generate_unbound_methods=true \
            ./helloworld.proto
    ```

## References
[Protocol Buffers Documentation](https://protobuf.dev/getting-started/gotutorial/) <br>
[gRPC](https://grpc.io/docs/languages/go/quickstart/)

