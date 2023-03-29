package apis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"

	greetv1 "github.com/hiromaily/connect-example/pkg/gen/greet/v1"
	"github.com/hiromaily/connect-example/pkg/gen/greet/v1/greetv1connect"
	"github.com/hiromaily/connect-example/pkg/logger"
)

type GreetServer struct {
	logger logger.Logger
}

func NewGreetServer(logger logger.Logger) *GreetServer {
	return &GreetServer{
		logger: logger,
	}
}

func NewGreetHandler(logger logger.Logger) (string, http.Handler) {
	return greetv1connect.NewGreetServiceHandler(NewGreetServer(logger))
}

func (g *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	g.logger.Info("Greet()")
	g.logger.Debug(fmt.Sprintf("Request headers: %v", req.Header()))

	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
