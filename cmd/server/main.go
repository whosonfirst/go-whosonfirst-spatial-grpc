package main

import (
	"context"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/server"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/spatial"
	"github.com/whosonfirst/go-whosonfirst-spatial/app"
	"github.com/whosonfirst/go-whosonfirst-spatial/flags"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	fs, err := flags.CommonFlags()

	if err != nil {
		log.Fatal(err)
	}

	flags.Parse(fs)

	ctx := context.Background()

	err = flags.ValidateCommonFlags(fs)

	if err != nil {
		log.Fatal(err)
	}

	spatial_app, err := app.NewSpatialApplicationWithFlagSet(ctx, fs)

	if err != nil {
		log.Fatalf("Failed to create new spatial application, %v", err)
	}

	paths := fs.Args()

	err = spatial_app.IndexPaths(ctx, paths...)

	if err != nil {
		log.Fatalf("Failed to index paths, %v", err)
	}

	spatial_server, err := server.NewSpatialServer(spatial_app)

	if err != nil {
		log.Fatalf("Failed to create spatial server, %v", err)
	}

	grpc_server := grpc.NewServer()

	spatial.RegisterSpatialServer(grpc_server, spatial_server)

	addr := fmt.Sprintf("localhost:%d", 8282) // FLAGS, PLEASE
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpc_server.Serve(lis)
}
