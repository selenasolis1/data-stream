package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	generated1 "github.com/selenasolis1/data-stream/graph/generated"
	"github.com/selenasolis1/data-stream/graph/model"
)

func (r *queryResolver) Name(ctx context.Context) (*model.Test, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) HelloWorld(ctx context.Context) (string, error) {
	return "hello world", nil
}

func (r *subscriptionResolver) Counting(ctx context.Context, machine []string, sensor []string) (<-chan int, error) {
	rid := rand.Int63()
	c := make(chan int, 1)
	hitmap := make(map[string]struct{})
	r.Forwarder.Subscribers[rid] = &DAQSubscriber{
		SubChan: c,
		Filter: func(name string, _ int, _ int) bool {
			if _, ok := hitmap[sensor]; ok {
				return true
			}
			return false
		},
	}

	go func() {
		<-ctx.Done()
		r.Forwarder.DeleteSubscriber(rid)
	}()
	return c, nil
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

// Subscription returns generated1.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated1.SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
