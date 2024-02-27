package main

import (
	"context"
	"log"
	"net/http"

	gw "github.com/OyePuru/grpc-proto/gen/go/proto/grpcproto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

// RegisterGrpcServiceHandlers registers gRPC service handlers.
func RegisterGrpcServiceHandlers(ctx context.Context, mux *runtime.ServeMux, grpcServerEndPoint string, opts []grpc.DialOption) {
	if err := gw.RegisterExampleGetServiceHandlerFromEndpoint(ctx, mux, grpcServerEndPoint, opts); err != nil {
		log.Fatalln("Failed to Register:", err)
	}

	if err := gw.RegisterExamplePostServiceHandlerFromEndpoint(ctx, mux, grpcServerEndPoint, opts); err != nil {
		log.Fatalln("Failed to Register:", err)
	}
}

func main() {
	// Create a background context.
	ctx := context.Background()
	// Create a context with cancel function to gracefully shutdown the server.
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create a new ServeMux.
	mux := runtime.NewServeMux()

	// Define gRPC server endpoint we can replace this with env later to handle multiple endpoints.
	grpcServerEndpoint := "localhost:9000"

	// Set gRPC dial options with insecure credentials.
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Register gRPC service handlers to the ServeMux.
	RegisterGrpcServiceHandlers(ctx, mux, grpcServerEndpoint, opts)

	log.Println("Server is up and running on localhost:9001")
	if err := http.ListenAndServe(":9001", mux); err != nil {
		grpclog.Fatal(err)
	}
}
