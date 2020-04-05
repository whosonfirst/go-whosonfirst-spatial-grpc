package server

import (
	"context"
	geojson_utils "github.com/whosonfirst/go-whosonfirst-geojson-v2/utils"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/spatial"
	"github.com/whosonfirst/go-whosonfirst-spatial/app"
)

const COORD_FACTOR float64 = 1e7

type SpatialServer struct {
	app *app.SpatialApplication
}

func NewSpatialServer(app *app.SpatialApplication) (*SpatialServer, error) {

	s := &SpatialServer{
		app: app,
	}

	return s, nil
}

func (s *SpatialServer) PointInPolygon(ctx context.Context, req *spatial.Coordinate) (*spatial.StandardPlacesResults, error) {

	spatial_db := s.app.SpatialDatabase

	lat := float64(req.Latitude) // COORD_FACTOR
	lon := float64(req.Latitude) // COORD_FACTOR

	coord, err := geojson_utils.NewCoordinateFromLatLons(lat, lon)

	if err != nil {
		return nil, err
	}

	rsp, err := spatial_db.PointInPolygon(ctx, &coord, nil)

	if err != nil {
		return nil, err
	}

	results := rsp.Results()
	count := len(results)

	grpc_results := make([]*spatial.StandardPlaceResponse, count)

	for idx, spr_result := range results {

		grpc_result := &spatial.StandardPlaceResponse{
			Id: spr_result.Id(),
		}

		grpc_results[idx] = grpc_result
	}

	grpc_rsp := &spatial.StandardPlacesResults{
		Results: grpc_results,
	}

	return grpc_rsp, nil
}
