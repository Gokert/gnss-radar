package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"time"

	"github.com/Gokert/gnss-radar/internal/delivery/graphql/generated"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/store"
)

// SatelliteName is the resolver for the satelliteName field.
func (r *taskResolver) SatelliteName(ctx context.Context, obj *model.Task) (string, error) {
	if obj == nil {
		return "", nil
	}

	satellites, err := r.gnssSevice.ListSatellites(ctx, store.ListSatellitesFilter{
		Ids: []string{obj.SatelliteID},
	})
	if err != nil {
		return "", fmt.Errorf("gnssSevice.ListSatellites: %w", err)
	}

	if len(satellites) == 0 {
		return "", nil
	}

	return satellites[0].SatelliteName, nil
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *taskResolver) UpdatedAt(ctx context.Context, obj *model.Task) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updatedAt"))
}

// Task returns generated.TaskResolver implementation.
func (r *Resolver) Task() generated.TaskResolver { return &taskResolver{r} }

type taskResolver struct{ *Resolver }
