package graphql

import (
	"context"
	"errors"
	"fmt"
	"github.com/Gokert/gnss-radar/internal/delivery/graphql/generated"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/pkg/utils"
	authorization "github.com/Gokert/gnss-radar/internal/service"
	"github.com/Gokert/gnss-radar/internal/store"
)

type authorizationMutationsResolver struct { *Resolver }

// AuthorizationMutations returns generated.AuthorizationMutationsResolver implementation.
func (r *Resolver) AuthorizationMutations() generated.AuthorizationMutationsResolver { return &authorizationMutationsResolver{r} }

// Authorization is the resolver for the authorization field.
func (r *mutationResolver) Authorization(ctx context.Context) (*model.AuthorizationMutations,  error){
	return &model.AuthorizationMutations{}, nil
}

// Signup is the resolver for the signup field.
func (r *authorizationMutationsResolver) Signup(ctx context.Context, obj *model.AuthorizationMutations, input model.SignupInput) ( *model.SignupOutput,  error){
	if err := utils.ValidateSignup(&input); err != nil {
		return nil, fmt.Errorf("validation.ValidateStruct %w", err)
	}

	result, err := r.authService.Signup(ctx, authorization.SignupRequest{Login: input.Login, Password: input.Password})
	if err != nil {
		switch {
		case errors.Is(err, store.ErrEntityAlreadyExist):
			return nil, model.ErrorAlreadyExists
		default:
			return nil, fmt.Errorf("authService.Signup %w", err)
		}

	return &model.SignupOutput{UserID: result.ID}, nil
}

// Signin is the resolver for the signin field.
func (r *authorizationMutationsResolver) Signin(ctx context.Context, obj *model.AuthorizationMutations, input model.SigninInput) ( *model.SigninOutput,  error){
	if err := utils.ValidateSignin(&input); err != nil {
		return nil, fmt.Errorf("validation.ValidateStruct %w", err)
	}

	result, err := r.authService.Signin(ctx, authorization.SigninRequest{Login: input.Login, Password: input.Password})
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			return nil, model.ErrorNotFound
		default:
			return nil, fmt.Errorf("authService.Signin %w", err)
		}
	}

	utils.SetCookie(ctx, result.SID)

	return &model.SigninOutput{}, nil
}

// Logout is the resolver for the logout field.
func (r *authorizationMutationsResolver) Logout(ctx context.Context, obj *model.AuthorizationMutations, input model.LogoutInput) ( *model.LogoutOutput,  error){
	cookie, err := utils.GetRequest(ctx).Cookie("session_id")
	if err != nil {
		return nil, model.ErrorBadRequest
	}

	result, err := r.authService.Logout(ctx, cookie.Value)
	if err != nil {
		return nil, fmt.Errorf("authService.Logout: %w", err)
	}
	if !result {
		return nil, model.ErrorInternalError
	}

	return nil, nil
}

// Authcheck is the resolver for the authcheck field.
func (r *queryResolver) Authcheck(ctx context.Context, input *model.AuthcheckInput) ( *model.AuthcheckOutput,  error){
	cookie, err := utils.GetRequest(ctx).Cookie("session_id")
	if err != nil {
		return nil, model.ErrorBadRequest
	}

	result, err := r.authService.Authcheck(ctx, cookie.Value)
	if err != nil {
		return nil, fmt.Errorf("authService.Authcheck: %w", err)
	}

	if !result {
		return nil, model.ErrorNotAuthorized
	}

	return &model.AuthcheckOutput{}, nil
}