package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/selenasolis1/data-stream/users/graph/generated"
	"github.com/selenasolis1/data-stream/users/graph/model"
	"labix.org/v2/mgo/bson"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user model.User
	count, err := r.users.Find(bson.M{"email": input.Email}).Count()
	if err != nil {
		return &model.User{}, err
	} else if count > 0 {
		return &model.User{}, errors.New("user with that email already exists")
	}
	err = r.users.Insert(bson.M{"email": input.Email})
	if err != nil {
		return &model.User{}, err
	}
	err = r.users.Find(bson.M{"email": input.Email}).One(&user)
	if err != nil {
		return &model.User{}, err
	}
	return &user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	var fields = bson.M{}
	var user model.User
	update := false
	if input.First != nil && *input.First != "" {
		fields["first"] = *input.First
		update = true
	}
	if input.Last != nil && *input.Last != "" {
		fields["last"] = *input.Last
		update = true
	}
	if input.Email != nil && *input.Email != "" {
		fields["email"] = *input.Email
		update = true
	}
	if !update {
		return &model.User{}, errors.New("no fields present for updating data")
	}
	err := r.users.UpdateId(bson.ObjectIdHex(input.ID), fields)
	if err != nil {
		return &model.User{}, err
	}
	err = r.users.Find(bson.M{"_id": bson.ObjectIdHex(input.ID)}).One(&user)
	if err != nil {
		return &model.User{}, err
	}
	user.ID = bson.ObjectIdHex(user.ID).Hex()
	return &user, nil
}

func (r *mutationResolver) UpdateNotification(ctx context.Context, input *model.UpdateNotification) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) NotificationAdded(ctx context.Context, id string) (<-chan *model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
	r.users = db.GetCollection("users")
	return &mutationResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	r.users = db.GetCollection("users")
	return &queryResolver{r}
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver {
	r.users = db.GetCollection("users")
	return &subscriptionResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *Resolver) Mutation() generated.MutationResolver         { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver               { return &queryResolver{r} }
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }
