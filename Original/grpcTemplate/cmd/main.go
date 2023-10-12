package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	pb "grpcTemplate/proto/gen"
	"log"
	"net"
	"net/http"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func main() {
	// REST cmd
	go func() {
		mux := runtime.NewServeMux()
		err := pb.RegisterTestApiHandlerServer(context.Background(), mux, &testApiServer{})
		if err != nil {
			return
		}

		apiPrefix := "/api/data"
		apiMux := http.NewServeMux()
		apiMux.Handle("/", http.StripPrefix(apiPrefix, mux))

		log.Fatalln(http.ListenAndServe(":8081", apiMux))
	}()
	// gRPC cmd
	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, &testApiServer{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("gRPC cmd failed to start: %v", err)
	}
}
