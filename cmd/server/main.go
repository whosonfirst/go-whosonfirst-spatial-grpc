package main

import (
	_ "github.com/whosonfirst/go-whosonfirst-spatial-rtree"
)

import (
	"context"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/server"
	"log"
)

func main() {

	ctx := context.Background()

	app, err := server.NewServerApplication()

	if err != nil {
		log.Fatalf("Failed to create new server application, %v", err)
	}

	err = app.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run server application, %v", err)
	}
}
