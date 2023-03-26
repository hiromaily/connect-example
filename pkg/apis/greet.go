package apis

import (
	"context"
	"fmt"
	"log"

	"github.com/bufbuild/connect-go"

	greetv1 "github.com/hiromaily/connect-example/pkg/gen/greet/v1"
)

type GreetServer struct{}

// TODO: move somewhere
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
