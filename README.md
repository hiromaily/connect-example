# connect-example
Example for go-connect server and connect-web client.

This program follows a `Clean Architecture` as much as I can.
For now, this runs connect-go server and JSON-RPC server which call same use-case.

<img src="https://raw.githubusercontent.com/hiromaily/documents/main/images/clean-architecture3.png"  width="50%" height="50%">

## Environment Variables
| Variables    | Explanation          | Default           |
|--------------|----------------------|-------------------|
| PORT         | gRPC server port     | 8080              |
| JSONRPC-PORT | JSON-RPC server port | 8090              |
| CONF         | config file path     | ./config/dev.toml |

## Front-end
[README](./web/README.md)

## Generator
### For Golang
- [protoc-gen-go](https://pkg.go.dev/google.golang.org/protobuf)
- [protoc-gen-connect-go](https://github.com/bufbuild/connect-go/tree/main/cmd/protoc-gen-connect-go)
### For ES
- [protoc-gen-es](https://www.npmjs.com/package/@bufbuild/protoc-gen-es)
- [protoc-gen-connect-web](https://github.com/bufbuild/connect-es/tree/main/packages/protoc-gen-connect-web) [Deprecated]
- [protoc-gen-connect-es](https://github.com/bufbuild/connect-es/tree/main/packages/protoc-gen-connect-es)

## Rererences
- [Connect Docs](https://connect.build/docs/introduction)
- [connect-go](https://github.com/bufbuild/connect-go)
- [connect-web](https://www.npmjs.com/package/@bufbuild/connect-web)
  - [connect-es](https://github.com/bufbuild/connect-es)
- [connect-demo](https://github.com/bufbuild/connect-demo)
- [Buf Docs](https://docs.buf.build/installation)
- [buf](https://github.com/bufbuild/buf)
- [grpc/grpc-web](https://github.com/grpc/grpc-web)
