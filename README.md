# telegram-ms

A microservice for sending message to a telegram bot using token and chait id.

This microservice is created using go lang and protocol used for client-server architecture is gRPC

### Pre-Requisites 
1. Install go
2. Install below go plugins
    - go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    - go install go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
3. Add bin path of go to PATH variable
    - export PATH="$PATH:$(go env GOPATH)/bin"

Reference :- https://www.golinuxcloud.com/golang-grpc/#gRPC_vs_REST
