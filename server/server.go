package server

import (
	"context"
	geojson_utils "github.com/whosonfirst/go-whosonfirst-geojson-v2/utils"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/spatial"
	"github.com/whosonfirst/go-whosonfirst-spatial/app"
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

func (s *SpatialServer) PointInPolygon(ctx context.Context, req *spatial.PointInPolygonRequest) (*spatial.StandardPlacesResponse, error) {

	spatial_db := s.app.SpatialDatabase

	const CordFactor float64 = 1e7

	lat := float64(req.Latitude) / CordFactor
	lon := float64(req.Latitude) / CordFactor

	coord, err := geojson_utils.NewCoordinateFromLatLons(lat, lon)

	if err != nil {
		return nil, err
	}

	_, err = spatial_db.PointInPolygon(ctx, &coord, nil)

	if err != nil {
		return nil, err
	}

	rsp := &spatial.StandardPlacesResponse{
		Wofid: 1234,
	}

	return rsp, nil
}
