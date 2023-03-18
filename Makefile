
.PHONY:tools
tools:
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest

#------------------------------------------------------------------------------
# proto using buf
#------------------------------------------------------------------------------
.PHONY:init-buf
init-buf:
	buf mod init
	touch buf.gen.yaml

# Note: set directory to ignore root directory `proto` from package name in *.proto files
.PHONY:buf-lint
buf-lint:
	#buf lint
	buf lint proto

# Note: set directory to ignore root directory `proto` from package name in *.proto files
.PHONY:buf-gen
buf-gen:
	#buf generate
	buf generate proto

#------------------------------------------------------------------------------
# go server
#------------------------------------------------------------------------------
.PHONY:run-server
run-server:
	go run ./cmd/server/main.go

.PHONY:build-server
build-server:
	go build -v -o ${GOPATH}/bin/connect-server ./cmd/server/main.go

.PHONY:exec-server
exec-server:
	connect-server

#------------------------------------------------------------------------------
# go client
#------------------------------------------------------------------------------
.PHONY:build-client
build-client:
	go build -v -o ${GOPATH}/bin/connect-client ./cmd/client/main.go

#------------------------------------------------------------------------------
# Make requests
#------------------------------------------------------------------------------
#-------------------------------------
# HTTP
#-------------------------------------
.PHONY:req-greet
req-greet:
	curl \
	--header "Content-Type: application/json" \
	--data '{"name": "Jane"}' \
	http://localhost:8080/greet.v1.GreetService/Greet

#-------------------------------------
# gRPC
#-------------------------------------
.PHONY:req-grpc-greet
req-grpc-greet:
	buf build proto -o protoset
	grpcurl \
	-protoset protoset -plaintext \
	-d '{"name": "Jane"}' \
	localhost:8080 greet.v1.GreetService/Greet
	rm -rf protoset

#------------------------------------------------------------------------------
# Utils
#------------------------------------------------------------------------------
.PHONY:clean
clean:
	rm -rf pkg/gen