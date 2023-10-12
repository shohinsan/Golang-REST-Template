package routes

import (
	"context"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MyFirstApi struct {
}

func (*MyFirstApi) FirstGetRpc(ctx context.Context, request *httpbody.HttpBody, in *emptypb.Empty) (*httpbody.HttpBody, error) {
	result := "This is a hardcoded response."

	response := &httpbody.HttpBody{
		ContentType: "text/html",
		Data:        []byte("Response: " + result),
	}

	return response, nil
}
