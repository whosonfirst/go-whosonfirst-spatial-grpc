package client

import (
	"flag"
	"fmt"
	"github.com/sfomuseum/go-flags/flagset"
	grpc_flags "github.com/whosonfirst/go-whosonfirst-spatial-grpc/flags"
	spatial_flags "github.com/whosonfirst/go-whosonfirst-spatial/flags"
)

func DefaultFlagSet() (*flag.FlagSet, error) {

	fs := flagset.NewFlagSet("client")

	err := grpc_flags.AppendGRPCClientFlags(fs)

	if err != nil {
		return nil, fmt.Errorf("Failed to append grpc client flags, %w", err)
	}

	err = spatial_flags.AppendQueryFlags(fs)

	if err != nil {
		return nil, fmt.Errorf("Failed to append query flags, %w", err)
	}

	return fs, nil
}
