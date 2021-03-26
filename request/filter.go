package request

import (
	wof_spatial "github.com/whosonfirst/go-whosonfirst-spatial"
	"github.com/whosonfirst/go-whosonfirst-spatial-grpc/spatial"
	wof_filter "github.com/whosonfirst/go-whosonfirst-spatial/filter"
)

func SPRFilterFromPointInPolygonRequest(req *spatial.PointInPolygonRequest) (wof_spatial.Filter, error) {

	inputs := &wof_filter.SPRInputs{
		Placetypes: req.Placetypes,
		// IsCurrent:
		// IsCeased:
		// IsDeprecated:
		// IsSuperseded:
		// IsSuperseding:
		// Geometries: req.Geometries,
		AlternateGeometries: req.AlternateGeometries,
		InceptionDate:       req.InceptionDate,
		CessationDate:       req.CessationDate,
	}

	return wof_filter.NewSPRFilterFromInputs(inputs)
}
