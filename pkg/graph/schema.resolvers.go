package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/theSuess/keypub/pkg/auth"
	"github.com/theSuess/keypub/pkg/graph/generated"
	logf "github.com/theSuess/keypub/pkg/log"
	"github.com/theSuess/keypub/pkg/model"
)

func (r *groupResolver) Users(ctx context.Context, obj *model.Group) ([]*model.User, error) {
	users := []*model.User{}
	err := r.DB.Model(obj).Related(&users, "Users").Error
	return users, err
}

func (r *groupResolver) Owners(ctx context.Context, obj *model.Group) ([]*model.User, error) {
	owners := []*model.User{}
	err := r.DB.Model(obj).Related(&owners, "Owners").Error
	return owners, err
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	id, _ := uuid.NewRandom()
	user := &model.User{
		ID:       id.String(),
		Name:     input.Name,
		Username: input.Username,
	}
	err := r.DB.Create(user).Error
	return user, err
}

func (r *mutationResolver) AddKey(ctx context.Context, input model.NewKey) (*model.PublicKey, error) {
	id, _ := uuid.NewRandom()
	key := &model.PublicKey{
		ID:      id.String(),
		Name:    input.Name,
		Content: input.Content,
		UserID:  input.UserID,
	}
	err := r.DB.Create(&key).Error
	return key, err
}

func (r *mutationResolver) CreateGroup(ctx context.Context, input model.NewGroup) (*model.Group, error) {
	claim := auth.ForContext(ctx)
	if claim == nil {
		return nil, errors.New("only authenticated users can perform this action")
	}
	id, _ := uuid.NewRandom()
	group := &model.Group{
		ID:   id.String(),
		Name: input.Name,
		Users: []*model.User{
			{ID: claim.UserID},
		},
		Owners: []*model.User{
			{ID: claim.UserID},
		},
	}
	if err := r.DB.Create(group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func (r *mutationResolver) JoinGroup(ctx context.Context, input model.JoinGroup) (*model.Group, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *publicKeyResolver) User(ctx context.Context, obj *model.PublicKey) (*model.User, error) {
	user := &model.User{}
	err := r.DB.Where(&model.User{ID: obj.UserID}).First(user).Error
	return user, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	log.Info("Fetching all users")
	users := []*model.User{}
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *queryResolver) User(ctx context.Context, username string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Groups(ctx context.Context) ([]*model.Group, error) {
	log.Info("Fetching all groups")
	groups := []*model.Group{}
	err := r.DB.Find(&groups).Error
	return groups, err
}

func (r *userResolver) Keys(ctx context.Context, obj *model.User) ([]*model.PublicKey, error) {
	keys := []*model.PublicKey{}
	err := r.DB.Where(&model.PublicKey{UserID: obj.ID}).Find(&keys).Error
	return keys, err
}

func (r *userResolver) Groups(ctx context.Context, obj *model.User, first *int) ([]*model.Group, error) {
	groups := []*model.Group{}
	var err error
	if first != nil {
		err = r.DB.Model(obj).Limit(*first).Related(&groups, "Groups").Error
	} else {
		err = r.DB.Model(obj).Related(&groups, "Groups").Error
	}
	return groups, err
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var log = logf.Log.WithName("server")
