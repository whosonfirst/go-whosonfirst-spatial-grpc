package main

import (
	"context"
	"fmt"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/lookup"
	grpc_flags "github.com/whosonfirst/go-whosonfirst-spatial-grpc/flags"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/spatial"
	spatial_flags "github.com/whosonfirst/go-whosonfirst-spatial/flags"
	"google.golang.org/grpc"
	"log"
)

func main() {

	fs := flagset.NewFlagSet("client")

	err := grpc_flags.AppendGRPCClientFlags(fs)

	if err != nil {
		log.Fatal(err)
	}

	err = spatial_flags.AppendQueryFlags(fs)

	if err != nil {
		log.Fatal(err)
	}

	flagset.Parse(fs)

	ctx := context.Background()

	err = grpc_flags.ValidateGRPCClientFlags(fs)

	if err != nil {
		log.Fatal(err)
	}

	err = spatial_flags.ValidateQueryFlags(fs)

	if err != nil {
		log.Fatal(err)
	}

	host, _ := lookup.StringVar(fs, grpc_flags.HOST)
	port, _ := lookup.IntVar(fs, grpc_flags.PORT)

	lat, _ := lookup.Float64Var(fs, spatial_flags.LATITUDE)
	lon, _ := lookup.Float64Var(fs, spatial_flags.LONGITUDE)

	lat32 := float32(lat)
	lon32 := float32(lon)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	addr := fmt.Sprintf("%s:%d", host, port)

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("fail to dial '%s', %v", addr, err)
	}

	defer conn.Close()

	client := spatial.NewSpatialClient(conn)

	req := &spatial.Coordinate{
		Latitude:  lat32,
		Longitude: lon32,
	}

	stream, err := client.PointInPolygon(ctx, req)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(len(stream.Results))

	for _, r := range stream.Results {
		log.Println(r, r.Id, r.ParentId)
	}

}
