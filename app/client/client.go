package client

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/request"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/spatial"
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

	opts, err := RunOptionsFromFlagSet(ctx, fs)

	if err != nil {
		return fmt.Errorf("Failed to derive options from flagset, %w", err)
	}

	return RunWithOptions(ctx, opts, logger)
}

func RunWithOptions(ctx context.Context, opts *RunOptions, logger *log.Logger) error {
		
	req := &pip.PointInPolygonRequest{
		Latitude:            opts.Latitude,
		Longitude:           opts.Longitude,
		Placetypes:          opts.Placetypes,
		Geometries:          opts.Geometries,
		AlternateGeometries: opts.AlternateGeometries,
		IsCurrent:           opts.IsCurrent,
		IsCeased:            opts.IsCeased,
		IsDeprecated:        opts.IsDeprecated,
		IsSuperseded:        opts.IsSuperseded,
		IsSuperseding:       opts.IsSuperseding,
		InceptionDate:       opts.InceptionDate,
		CessationDate:       opts.CessationDate,
		Properties:          opts.Properties,
		Sort:                opts.Sort,
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
