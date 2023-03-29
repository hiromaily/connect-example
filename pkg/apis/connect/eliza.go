package connect

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"

	elizav1 "github.com/hiromaily/connect-example/pkg/gen/eliza/v1"
	"github.com/hiromaily/connect-example/pkg/gen/eliza/v1/elizav1connect"
	"github.com/hiromaily/connect-example/pkg/logger"
	"github.com/hiromaily/connect-example/pkg/usecases/eliza"
)

// Refer to
// https://github.com/bufbuild/connect-demo/blob/3a30d4de07d6ac42110acd4ebf64bb4bf8a62579/main.go

type ElizaServer struct {
	// The time to sleep between sending responses on a stream
	streamDelay time.Duration
	logger      logger.Logger
	ucEliza     *eliza.Eliza
}

func NewElizaServer(logger logger.Logger, ucEliza *eliza.Eliza) *ElizaServer {
	return &ElizaServer{
		logger:  logger,
		ucEliza: ucEliza,
	}
}

func NewElizaHandler(logger logger.Logger, ucEliza *eliza.Eliza) (string, http.Handler) {
	return elizav1connect.NewElizaServiceHandler(NewElizaServer(logger, ucEliza))
}

func (e *ElizaServer) Say(
	ctx context.Context,
	req *connect.Request[elizav1.SayRequest],
) (*connect.Response[elizav1.SayResponse], error) {
	e.logger.Info("Say()")

	// usecase
	reply := e.ucEliza.Say(req.Msg.Sentence)

	// create response
	return connect.NewResponse(&elizav1.SayResponse{
		Sentence: reply,
	}), nil
}

// Converse is bi-directional streaming request
func (e *ElizaServer) Converse(
	ctx context.Context,
	stream *connect.BidiStream[elizav1.ConverseRequest, elizav1.ConverseResponse],
) error {
	e.logger.Info("Converse()")

	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		request, err := stream.Receive()
		if err != nil && errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return fmt.Errorf("receive request: %w", err)
		}

		// usecase
		reply := e.ucEliza.Say(request.Sentence)

		// response
		if err := stream.Send(&elizav1.ConverseResponse{Sentence: reply}); err != nil {
			return fmt.Errorf("send response: %w", err)
		}
	}
}

// Introduce is server-streaming request
func (e *ElizaServer) Introduce(
	ctx context.Context,
	req *connect.Request[elizav1.IntroduceRequest],
	stream *connect.ServerStream[elizav1.IntroduceResponse],
) error {
	e.logger.Info("Introduce()")

	// usecase
	//name := req.Msg.Name
	//if name == "" {
	//	name = "Anonymous User"
	//}
	//intros := eliza.GetIntroResponses(name)
	intros := e.ucEliza.GetIntros(req.Msg.Name)

	// response
	var ticker *time.Ticker
	if e.streamDelay > 0 {
		ticker = time.NewTicker(e.streamDelay)
		defer ticker.Stop()
	}
	for _, resp := range intros {
		if ticker != nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-ticker.C:
			}
		}
		if err := stream.Send(&elizav1.IntroduceResponse{Sentence: resp}); err != nil {
			return err
		}
	}
	return nil
}
