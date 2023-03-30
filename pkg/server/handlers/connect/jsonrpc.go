package connect

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"

	greetv1 "github.com/hiromaily/connect-example/pkg/gen/greet/v1"
	jsonrpcv1 "github.com/hiromaily/connect-example/pkg/gen/jsonrpc/v1"
	"github.com/hiromaily/connect-example/pkg/gen/jsonrpc/v1/jsonrpcv1connect"
	"github.com/hiromaily/connect-example/pkg/logger"
	"github.com/hiromaily/connect-example/pkg/usecases/greet"
)

type JSONRPCHandler struct {
	logger  logger.Logger
	ucGreet *greet.Greet
}

// NewJSONRPCHandler returns greet handler
// Handler is Controllers as Interface Adapters in Clean Architecture
func NewJSONRPCHandler(logger logger.Logger, ucGreet *greet.Greet) (string, http.Handler) {
	handler := &JSONRPCHandler{
		logger:  logger,
		ucGreet: ucGreet,
	}
	return jsonrpcv1connect.NewJSONRPCGreetServiceHandler(handler)
}

func (j *JSONRPCHandler) Greet(
	_ context.Context,
	req *connect.Request[jsonrpcv1.JSONRPCRequest],
) (*connect.Response[jsonrpcv1.JSONRPCResponse], error) {
	j.logger.Info("JSONRPC Greet()")
	j.logger.Debug(fmt.Sprintf("Request headers: %v", req.Header()))

	// usecase
	reply := j.ucGreet.Greet(req.Msg.Params.Name)

	// create response
	res := connect.NewResponse(&jsonrpcv1.JSONRPCResponse{
		Jsonrpc: "2.0",
		Params: &greetv1.GreetResponse{
			Greeting: reply,
		},
		Id: req.Msg.Id,
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}
