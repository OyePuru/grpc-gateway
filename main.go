package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	gw "github.com/amanjain-cb/helloworld/gen/go/proto/helloworld"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	grpcServerEndpoint := "localhost:9000"
	if err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		log.Fatalln("Failed to Register:", err)
	}

	if err := gw.RegisterGreeter2HandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		log.Fatalln("Failed to Register:", err)
	}

	log.Println("Server is up and running on localhost:9001")
	if err := http.ListenAndServe(":9001", mux); err != nil {
		grpclog.Fatal(err)
	}
}
