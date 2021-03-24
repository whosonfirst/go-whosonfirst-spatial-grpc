package main

import (
	_ "github.com/whosonfirst/go-whosonfirst-spatial-rtree"
)

import (
	"context"
	"fmt"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/lookup"
	grpc_flags "github.com/whosonfirst/go-whosonfirst-spatial-grpc/flags"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/server"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/spatial"
	"github.com/whosonfirst/go-whosonfirst-spatial/app"
	spatial_flags "github.com/whosonfirst/go-whosonfirst-spatial/flags"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	fs, err := spatial_flags.CommonFlags()

	if err != nil {
		log.Fatal(err)
	}

	err = spatial_flags.AppendIndexingFlags(fs)

	if err != nil {
		log.Fatal(err)
	}

	err = grpc_flags.AppendGRPCServerFlags(fs)

	if err != nil {
		log.Fatal(err)
	}

	flagset.Parse(fs)

	ctx := context.Background()

	err = spatial_flags.ValidateCommonFlags(fs)

	if err != nil {
		log.Fatal(err)
	}

	err = spatial_flags.ValidateIndexingFlags(fs)

	if err != nil {
		log.Fatal(err)
	}

	err = grpc_flags.ValidateGRPCServerFlags(fs)

	if err != nil {
		log.Fatal(err)
	}

	host, _ := lookup.StringVar(fs, grpc_flags.HOST)
	port, _ := lookup.IntVar(fs, grpc_flags.PORT)

	spatial_app, err := app.NewSpatialApplicationWithFlagSet(ctx, fs)

	if err != nil {
		log.Fatalf("Failed to create new spatial application, %v", err)
	}

	uris := fs.Args()

	err = spatial_app.IndexPaths(ctx, uris...)

	if err != nil {
		log.Fatalf("Failed to index paths, %v", err)
	}

	spatial_server, err := server.NewSpatialServer(spatial_app)

	if err != nil {
		log.Fatalf("Failed to create spatial server, %v", err)
	}

	grpc_server := grpc.NewServer()

	spatial.RegisterSpatialServer(grpc_server, spatial_server)

	addr := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Listening on %s\n", addr)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpc_server.Serve(lis)
}
