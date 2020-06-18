package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/theSuess/keypub/pkg/auth"
	"github.com/theSuess/keypub/pkg/graph/generated"
	"github.com/theSuess/keypub/pkg/model"
)

func (r *groupResolver) Users(ctx context.Context, obj *model.Group) ([]*model.User, error) {
	return r.GroupService.UsersOf(obj)
}

func (r *groupResolver) Owners(ctx context.Context, obj *model.Group) ([]*model.User, error) {
	return r.GroupService.OwnersOf(obj)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		Name:     input.Name,
		Username: input.Username,
	}
	err := r.UserService.Register(user)
	return user, err
}

func (r *mutationResolver) AddKey(ctx context.Context, input model.NewKey) (*model.PublicKey, error) {
	key := &model.PublicKey{
		Name:    input.Name,
		Content: input.Content,
		UserID:  input.UserID,
	}
	err := r.KeyService.AddKey(key)
	return key, err
}

func (r *mutationResolver) CreateGroup(ctx context.Context, input model.NewGroup) (*model.Group, error) {
	claim := auth.ForContext(ctx)
	if claim == nil {
		return nil, errors.New("only authenticated users can perform this action")
	}
	group := &model.Group{
		Name: input.Name,
	}
	user, err := r.UserService.ByID(claim.UserID)
	if err != nil {
		return nil, err
	}
	err = r.GroupService.CreateGroup(group, user)
	return group, err
}

func (r *mutationResolver) JoinGroup(ctx context.Context, input model.JoinGroup) (*model.Group, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *publicKeyResolver) User(ctx context.Context, obj *model.PublicKey) (*model.User, error) {
	return r.KeyService.Owner(obj)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	log.Info("Fetching all users")
	return r.UserService.FindAll(-1, -1)
}

func (r *queryResolver) User(ctx context.Context, username string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Groups(ctx context.Context) ([]*model.Group, error) {
	log.Info("Fetching all groups")
	return r.GroupService.FindAll(-1, -1)
}

func (r *userResolver) Keys(ctx context.Context, obj *model.User) ([]*model.PublicKey, error) {
	limit, offset := resolvePagination(nil, nil)
	return r.UserService.KeysOf(obj, limit, offset)
}

func (r *userResolver) Groups(ctx context.Context, obj *model.User, first *int) ([]*model.Group, error) {
	limit, offset := resolvePagination(first, nil)
	return r.UserService.GroupsOf(obj, limit, offset)
}

// Group returns generated.GroupResolver implementation.
func (r *Resolver) Group() generated.GroupResolver { return &groupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// PublicKey returns generated.PublicKeyResolver implementation.
func (r *Resolver) PublicKey() generated.PublicKeyResolver { return &publicKeyResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type groupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type publicKeyResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
