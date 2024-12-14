package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"

	"github.com/Gokert/gnss-radar/internal/pkg/model"
)

// GenerateRecieverCode is the resolver for the generateRecieverCode field.
func (r *queryResolver) GenerateRecieverCode(ctx context.Context, filter model.CodeRecieverInput) (*model.CodeReciever, error) {
	reciever, err := r.gnssSevice.CodeGen(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("r.gnssSevice.CodeGen: %w", err)
	}
	return &reciever, nil
}