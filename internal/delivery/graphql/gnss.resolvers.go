package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/service"
)

// Listgnss is the resolver for the listgnss field.
func (r *queryResolver) Listgnss(ctx context.Context, filter model.GNSSFilter) (*model.GNSSPagination, error) {
	panic(fmt.Errorf("not implemented: Listgnss - listgnss"))
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) ListGnss(ctx context.Context, filter model.GNSSFilter) (*model.GNSSPagination, error) {
	jsonRawData, err := r.gnssSevice.ListGnss(ctx, service.ListRequest{X: "", Y: "", Z: ""})
	if err != nil {
		return nil, err
	}
	return jsonRawData, nil

}

// func getCoords() []*model.Gnss {
// 	jsonData := []*model.Gnss{
// 		&model.Gnss{
// 			ID:            "PC06",
// 			SatelliteID:   "PC06",
// 			SatelliteName: "PC06",
// 			Coordinates: &model.CoordsResults{
// 				X: "-16806.320344",
// 				Y: "29291.120310",
// 				Z: "-25355.710938",
// 			},
// 		},
// 		&model.Gnss{
// 			ID:            "PC07",
// 			SatelliteID:   "PC07",
// 			SatelliteName: "PC07",
// 			Coordinates: &model.CoordsResults{
// 				X: "-6959.418476",
// 				Y: "39332.954409",
// 				Z: "-13000.851001",
// 			},
// 		},
// 		&model.Gnss{
// 			ID:            "PC08",
// 			SatelliteID:   "PC08",
// 			SatelliteName: "PC08",
// 			Coordinates: &model.CoordsResults{
// 				X: "-1908.204600",
// 				Y: "21553.224987",
// 				Z: "36203.881809",
// 			},
// 		},
// 		&model.Gnss{
// 			ID:            "PC09",
// 			SatelliteID:   "PC09",
// 			SatelliteName: "PC09",
// 			Coordinates: &model.CoordsResults{
// 				X: "-11202.586298",
// 				Y: "28046.331947",
// 				Z: "-29182.143554",
// 			},
// 		},
// 		&model.Gnss{
// 			ID:            "PC10",
// 			SatelliteID:   "PC10",
// 			SatelliteName: "PC10",
// 			Coordinates: &model.CoordsResults{
// 				X: "-917.431406",
// 				Y: "41238.966109",
// 				Z: "-6711.991412",
// 			},
// 		},
// 		&model.Gnss{
// 			ID:            "PC11",
// 			SatelliteID:   "PC11",
// 			SatelliteName: "PC11",
// 			Coordinates: &model.CoordsResults{
// 				X: "-16138.177056",
// 				Y: "-3913.891460",
// 				Z: "-22348.411693",
// 			},
// 		},
// 		&model.Gnss{
// 			ID:            "PC12",
// 			SatelliteID:   "PC12",
// 			SatelliteName: "PC12",
// 			Coordinates: &model.CoordsResults{
// 				X: "-997.099233",
// 				Y: "-19759.345910",
// 				Z: "-19638.934483",
// 			},
// 		},
// 		&model.Gnss{
// 			ID:            "PC13",
// 			SatelliteID:   "PC13",
// 			SatelliteName: "PC13",
// 			Coordinates: &model.CoordsResults{
// 				X: "5858.392549",
// 				Y: "25505.986419",
// 				Z: "33308.911170",
// 			},
// 		},
// 		&model.Gnss{
// 			ID:            "PC14",
// 			SatelliteID:   "PC14",
// 			SatelliteName: "PC14",
// 			Coordinates: &model.CoordsResults{
// 				X: "-17706.605729",
// 				Y: "-14691.268566",
// 				Z: "15829.680477",
// 			},
// 		},
// 		&model.Gnss{
// 			ID:            "PC16",
// 			SatelliteID:   "PC16",
// 			SatelliteName: "PC16",
// 			Coordinates: &model.CoordsResults{
// 				X: "-22387.055407",
// 				Y: "28560.640995",
// 				Z: "-21454.026667",
// 			},
// 		},
// 	}

// 	return jsonData
// }
