package routes

import (
	"context"
	"errors"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MySecondApi struct {
}

func (*MySecondApi) ThirdGetRpc(ctx context.Context, in *emptypb.Empty) (*httpbody.HttpBody, error) {

	var _ string

	if errors.Is(ctx.Err(), context.Canceled) {
		_ = "Request canceled"
	} else if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		_ = "Request timed out"
	} else {
		_ = "Received a request"
	}

	response := &httpbody.HttpBody{
		ContentType: "text/html",
		Data:        []byte("Hello World"),
	}

	return response, nil
}
