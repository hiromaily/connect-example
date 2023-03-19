package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	connect "github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	greetv1 "github.com/hiromaily/connect-example/pkg/gen/greet/v1"        // generated by protoc-gen-go
	"github.com/hiromaily/connect-example/pkg/gen/greet/v1/greetv1connect" // generated by protoc-gen-connect-go
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	mux := http.NewServeMux()
	createHandlers(mux)

	// serve
	http.ListenAndServe(
		"localhost:8080",
		// h2c protocol is the non-TLS version of HTTP/2
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

func createHandlers(mux *http.ServeMux) {
	// params: path, handler
	mux.Handle(greetv1connect.NewGreetServiceHandler(&GreetServer{}))
}
