package client

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/lookup"
	grpc_flags "github.com/whosonfirst/go-whosonfirst-spatial-grpc/flags"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/request"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/spatial"
	spatial_flags "github.com/whosonfirst/go-whosonfirst-spatial/flags"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
)

func Run(ctx context.Context, logger *log.Logger) error {

	fs, err := DefaultFlagSet()

	if err != nil {
		return fmt.Errorf("Failed to create default flagset, %w", err)
	}

	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *log.Logger) error {

	flagset.Parse(fs)

	err := grpc_flags.ValidateGRPCClientFlags(fs)

	if err != nil {
		return fmt.Errorf("Failed to validate grpc client flags, %w", err)
	}

	err = spatial_flags.ValidateQueryFlags(fs)

	if err != nil {
		return fmt.Errorf("Failed to validate query flags, %w", err)
	}

	req, err := request.NewPointInPolygonRequestFromFlagSet(fs)

	if err != nil {
		return fmt.Errorf("Failed to create PIP request, %v", err)
	}

	host, _ := lookup.StringVar(fs, grpc_flags.HOST)
	port, _ := lookup.IntVar(fs, grpc_flags.PORT)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	addr := fmt.Sprintf("%s:%d", host, port)

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		return fmt.Errorf("Failed to dial '%s', %v", addr, err)
	}

	defer conn.Close()

	client := spatial.NewSpatialClient(conn)

	stream, err := client.PointInPolygon(ctx, req)

	if err != nil {
		return fmt.Errorf("Failed to perform point in polygon operation, %w", err)
	}

	to_stdout, _ := lookup.BoolVar(fs, grpc_flags.TO_STDOUT)
	to_null, _ := lookup.BoolVar(fs, grpc_flags.TO_NULL)

	writers := make([]io.Writer, 0)

	if to_stdout {
		writers = append(writers, os.Stdout)
	}

	if to_null {
		writers = append(writers, io.Discard)
	}

	wr := io.MultiWriter(writers...)

	enc := json.NewEncoder(wr)
	err = enc.Encode(stream)

	if err != nil {
		return fmt.Errorf("Failed to encode stream, %v", err)
	}

	return nil
}
