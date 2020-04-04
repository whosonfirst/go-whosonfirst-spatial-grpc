package main

import (
	"context"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/spatial"
	"google.golang.org/grpc"
	"log"
)

func main() {

	host := "localhost"
	port := 8282

	lat := float32(37.605)
	lon := float32(-122.405)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	addr := fmt.Sprintf("%s:%d", host, port)

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := spatial.NewServerClient(conn)

	req := &spatial.PointInPolygonRequest{
		Latitude:  lat,
		Longitude: lon,
	}

	ctx := context.Background()

	stream, err := client.PointInPolygon(ctx, req)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(stream)
}
