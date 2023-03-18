# connect-example
Example for go-connect and connect-web

- [Connect Docs](https://connect.build/docs/introduction)
- [connect-go](https://github.com/bufbuild/connect-go)
- [connect-web](https://www.npmjs.com/package/@bufbuild/connect-web)
- [Buf Docs](https://docs.buf.build/installation)
- [buf](https://github.com/bufbuild/buf)
- [grpc/grpc-web](https://github.com/grpc/grpc-web)

## protocols in connect-go 
### gRPC protocol
- gRPC エコシステム全体で使用されるプロトコルであり、`connect-go` をすぐに他の gRPC 実装と互換性を持たせることができる
- `grpc-go` clientは `connect-go` server と問題なく動作し、その逆も問題なく動作する

### gRPC-Web protocol
- [grpc/grpc-web](https://github.com/grpc/grpc-web) で使用される gRPC-Web プロトコル
- Envoyといった仲介proxyを必要とせずに、`connect-go` server が `grpc-web` front-endと相互運用できるようにする

### [Connect protocol](https://connect.build/docs/protocol/)
- HTTP/1.1 または HTTP/2 で動作するシンプルな POST 専用プロトコル
- ストリーミングを含む gRPC と gRPC-Web の最良の部分を取り、それらをブラウザー、モノリス、およびマイクロサービスで同等に機能するプロトコルにパッケージ化したもの
- デフォルトでは、JSON およびバイナリでエンコードされた Protobuf がサポートされている
