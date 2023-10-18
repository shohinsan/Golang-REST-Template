package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/encoding/protojson"
	"grpcTemplate/internal/routes"
	pb "grpcTemplate/pkg"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	go startGRPCServer()
	startHTTPGateway()

}

func startGRPCServer() {
	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)
	pb.RegisterMyFirstApiServer(grpcServer, &routes.MyApiOneServer{})
	pb.RegisterMySecondApiServer(grpcServer, &routes.MyApiTwoServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("gRPC server failed to start: %v", err)
	}
}

func startHTTPGateway() {
	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	mux := runtime.NewServeMux(jsonOption)

	err1 := pb.RegisterMyFirstApiHandlerServer(context.Background(), mux, &routes.MyApiOneServer{})
	if err1 != nil {
		log.Fatalf("Failed to register MyFirstApiHandlerServer: %v", err1)
	}

	err2 := pb.RegisterMySecondApiHandlerServer(context.Background(), mux, &routes.MyApiTwoServer{})
	if err2 != nil {
		log.Fatalf("Failed to register MySecondApiHandlerServer: %v", err2)
	}

	prettierMux := http.NewServeMux()
	prettierMux.Handle("/", prettierMiddleware(mux))

	log.Fatal(http.ListenAndServeTLS(
		":8081",
		"./cert.pem",
		"./key.pem",
		prettierMux,
	))
}

func prettierMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.RawQuery, "pretty") {
			r.Header.Set("Accept", "application/json+pretty")
		}
		h.ServeHTTP(w, r)
	})
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("./cert.pem", "./key.pem")
	if err != nil {
		return nil, fmt.Errorf("failed to load server certificate: %v", err)
	}
	caCert, err := os.ReadFile("./ca.pem")
	if err != nil {
		return nil, fmt.Errorf("failed to read CA certificate: %v", err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to add CA certificate to the pool")
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
		RootCAs:      certPool,
		NextProtos:   []string{"h2", "http/1.1"},
	}
	return credentials.NewTLS(config), nil
}
