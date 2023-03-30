package connect

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"

	greetv1 "github.com/hiromaily/connect-example/pkg/gen/greet/v1"
	"github.com/hiromaily/connect-example/pkg/gen/greet/v1/greetv1connect"
	"github.com/hiromaily/connect-example/pkg/logger"
	"github.com/hiromaily/connect-example/pkg/usecases/greet"
)

type GreetHandler struct {
	logger  logger.Logger
	ucGreet *greet.Greet
}

// NewGreetHandler returns greet handler
// Handler is Controllers as Interface Adapters in Clean Architecture
func NewGreetHandler(logger logger.Logger, ucGreet *greet.Greet) (string, http.Handler) {
	handler := &GreetHandler{
		logger:  logger,
		ucGreet: ucGreet,
	}
	return greetv1connect.NewGreetServiceHandler(handler)
}

func (g *GreetHandler) Greet(
	_ context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	g.logger.Info("Greet()")
	g.logger.Debug(fmt.Sprintf("Request headers: %v", req.Header()))

	// usecase
	reply := g.ucGreet.Greet(req.Msg.Name)

	// create response
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: reply,
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
