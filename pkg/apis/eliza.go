package apis

import (
	"context"
	"errors"
	"fmt"
	"github.com/hiromaily/connect-example/pkg/logger"
	"io"
	"net/http"
	"time"

	"github.com/bufbuild/connect-go"

	elizav1 "github.com/hiromaily/connect-example/pkg/gen/eliza/v1"
	"github.com/hiromaily/connect-example/pkg/gen/eliza/v1/elizav1connect"
)

// Refer to
// https://github.com/bufbuild/connect-demo/blob/3a30d4de07d6ac42110acd4ebf64bb4bf8a62579/main.go

type ElizaServer struct {
	// The time to sleep between sending responses on a stream
	streamDelay time.Duration
	logger      logger.Logger
}

func NewElizaHandler(logger logger.Logger) (string, http.Handler) {
	return elizav1connect.NewElizaServiceHandler(&ElizaServer{
		logger: logger,
	})
}

func (e *ElizaServer) Say(
	ctx context.Context,
	req *connect.Request[elizav1.SayRequest],
) (*connect.Response[elizav1.SayResponse], error) {
	e.logger.Info("Say()")

	//reply, _ := eliza.Reply(req.Msg.Sentence) // ignore end-of-conversation detection
	reply := fmt.Sprintf("You said, %s!", req.Msg.Sentence)
	return connect.NewResponse(&elizav1.SayResponse{
		Sentence: reply,
	}), nil
}

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
		//reply, endSession := eliza.Reply(request.Sentence)
		reply := fmt.Sprintf("You said, %s!", request.Sentence)
		if err := stream.Send(&elizav1.ConverseResponse{Sentence: reply}); err != nil {
			return fmt.Errorf("send response: %w", err)
		}
	}
}

func (e *ElizaServer) Introduce(
	ctx context.Context,
	req *connect.Request[elizav1.IntroduceRequest],
	stream *connect.ServerStream[elizav1.IntroduceResponse],
) error {
	e.logger.Info("Introduce()")

	name := req.Msg.Name
	if name == "" {
		name = "Anonymous User"
	}
	//intros := eliza.GetIntroResponses(name)
	intros := []string{"Hello", "Goodbye", "How are you"}
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
