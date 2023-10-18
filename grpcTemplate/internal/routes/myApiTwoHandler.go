package routes

import (
	"context"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "grpcTemplate/pkg"
)

type MyApiTwoServer struct {
	pb.UnimplementedMySecondApiServer
}

func (*MyApiTwoServer) ThirdGetRpc(ctx context.Context, in *emptypb.Empty) (*httpbody.HttpBody, error) {
	result := "This is a hardcoded response."
	responseData := "Response: " + result

	pretty := ctx.Value("pretty")
	if pretty != nil {
		responseData = "Pretty Response: " + result
	}

	response := &httpbody.HttpBody{
		ContentType: "text/html",
		Data:        []byte(responseData),
	}
	return response, nil
}

func (*MyApiTwoServer) FourthGetRpc(ctx context.Context, in *emptypb.Empty) (*httpbody.HttpBody, error) {
	result := "This is a hardcoded response."
	responseData := "Response: " + result

	pretty := ctx.Value("pretty")
	if pretty != nil {
		responseData = "Pretty Response: " + result
	}

	response := &httpbody.HttpBody{
		ContentType: "text/html",
		Data:        []byte(responseData),
	}
	return response, nil
}
