package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graph/generated"
	"app/graph/models"
	"context"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input models.CreateTodoInput) (*models.Todo, error) {
	todo := models.Todo{
		ID:   len(r.Todos) + 1,
		Text: input.Text,
		Done: input.Done,
	}

	r.Todos = append(r.Todos, &todo)

	// notifying the listenners when todo is created

	for key := range r.TodoChannels {
		r.TodoChannels[key] <- &todo
	}

	return &todo, nil
}

// GetTodos is the resolver for the getTodos field.
func (r *queryResolver) GetTodos(ctx context.Context) ([]*models.Todo, error) {
	return r.Todos, nil
}

// TodoCreated is the resolver for the todoCreated field.
func (r *subscriptionResolver) TodoCreated(ctx context.Context) (<-chan *models.Todo, error) {
	// regisster new listnner
	// channel has to be buffered so it doesn't block the execution
	createdChannel := make(chan *models.Todo, 1)

	// append the channel to the list of lstenners
	r.TodoChannels[&createdChannel] = createdChannel
	println("client: ", &createdChannel, " connected")

	// this is like thread, it will be waiting until the ctx is done which means the client has disconnected
	go func() {
		<-ctx.Done()
		// remove the channel when the connection is closed from the client
		println("client: ", &createdChannel, " disconnected")
		delete(r.TodoChannels, &createdChannel)
	}()

	return createdChannel, nil
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
