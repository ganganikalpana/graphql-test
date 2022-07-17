package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graphql/database"
	"graphql/graph/generated"
	"graphql/graph/model"
	"log"
	"github.com/thanhpk/randstr"
)

// CreateDog is the resolver for the createDog field.
func (r *mutationResolver) CreateDog(ctx context.Context, input model.NewDog) (*model.Dog, error) {
	return db.Save(&input), nil
}

// Dog is the resolver for the dog field.
func (r *queryResolver) Dog(ctx context.Context, id string) (*model.Dog, error) {
	return db.FindByID(id), nil
}

// Dogs is the resolver for the dogs field.
func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
	return db.All(), nil
}

// DogCreated is the resolver for the dogCreated field.
func (r *subscriptionResolver) DogCreated(ctx context.Context) (<-chan *model.Dog, error) {
	token := randstr.Hex(16)
	mc := make(chan *model.Dog, 1)
	r.mutex.Lock()
	r.messageChannels[token] = mc
	r.mutex.Unlock()

	go func() {
		<-ctx.Done()
		r.mutex.Lock()
		delete(r.messageChannels, token)
		r.mutex.Unlock()
		log.Println("Deleted")
	}()

	log.Println("Subscription: message created")

	return mc, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect()
