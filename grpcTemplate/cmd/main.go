package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	pb "grpcTemplate/pkg" // Correct the import path
	"log"
	"net"
	"net/http"
	"strings"
)

type myApiOneServer struct {
	pb.UnimplementedMyFirstApiServer
}

type myApiTwoServer struct {
	pb.UnimplementedMySecondApiServer
}

func main() {
	// REST cmd
	go func() {
		mux := runtime.NewServeMux()

		err1 := pb.RegisterMyFirstApiHandlerServer(context.Background(), mux, &myApiOneServer{})
		if err1 != nil {
			log.Fatalf("Failed to register MyFirstApiHandlerServer: %v", err1)
		}

		err2 := pb.RegisterMySecondApiHandlerServer(context.Background(), mux, &myApiTwoServer{})
		if err2 != nil {
			log.Fatalf("Failed to register MySecondApiHandlerServer: %v", err2)
		}

		prettierMux := http.NewServeMux()
		prettierMux.Handle("/", prettierMiddleware(mux))

		// Listen and serve on port 8081
		log.Fatalln(http.ListenAndServe(":8081", prettierMux))
	}()

	// gRPC cmd
	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMyFirstApiServer(grpcServer, &myApiOneServer{})
	pb.RegisterMySecondApiServer(grpcServer, &myApiTwoServer{})
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("gRPC cmd failed to start: %v", err)
	}
}

func prettierMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.RawQuery, "pretty") {
			r.Header.Set("Accept", "application/json+pretty")
		}
		h.ServeHTTP(w, r)
	})
}
