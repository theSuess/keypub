package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/theSuess/keypub/pkg/graph/generated"
	logf "github.com/theSuess/keypub/pkg/log"
	"github.com/theSuess/keypub/pkg/model"
)

var log = logf.Log.WithName("server")

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

func (r *userResolver) Keys(ctx context.Context, obj *model.User) ([]*model.PublicKey, error) {
	keys := []*model.PublicKey{}
	err := r.DB.Where(&model.PublicKey{UserID: obj.ID}).Find(&keys).Error
	return keys, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// PublicKey returns generated.PublicKeyResolver implementation.
func (r *Resolver) PublicKey() generated.PublicKeyResolver { return &publicKeyResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type publicKeyResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
