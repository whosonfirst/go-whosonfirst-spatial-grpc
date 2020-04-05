package server

import (
	"context"
	geojson_utils "github.com/whosonfirst/go-whosonfirst-geojson-v2/utils"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/spatial"
	"github.com/whosonfirst/go-whosonfirst-spatial/app"
	"github.com/whosonfirst/go-whosonfirst-spatial/filter"
	"log"
)

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

	lat := float64(req.Latitude)
	lon := float64(req.Longitude)

	coord, err := geojson_utils.NewCoordinateFromLatLons(lat, lon)

	if err != nil {
		return nil, err
	}

	f, err := filter.NewSPRFilter()

	if err != nil {
		return nil, err
	}

	rsp, err := spatial_db.PointInPolygon(ctx, &coord, f)

	if err != nil {
		return nil, err
	}

	results := rsp.Results()
	count := len(results)

	log.Println("COUNT", count, lat, lon)

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

func (s *SpatialServer) PointInPolygonStream(ctx context.Context, req *spatial.Coordinate, stream *spatial.StandardPlaceResponse) error {

	spatial_db := s.app.SpatialDatabase

	lat := float64(req.Latitude)
	lon := float64(req.Longitude)

	coord, err := geojson_utils.NewCoordinateFromLatLons(lat, lon)

	if err != nil {
		return err
	}

	f, err := filter.NewSPRFilter()

	if err != nil {
		return err
	}

	// UPDATE TO USE CHANNELS

	rsp, err := spatial_db.PointInPolygon(ctx, &coord, f)

	if err != nil {
		return err
	}

	results := rsp.Results()

	for _, spr_result := range results {

		grpc_result := &spatial.StandardPlaceResponse{
			Id: spr_result.Id(),
		}

		err := stream.Send(grpc_result)

		if err != nil {
			return err
		}
	}

	return nil
}
