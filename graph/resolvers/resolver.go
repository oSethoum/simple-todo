package resolvers

import (
	"app/graph/generated"
	"app/graph/models"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Todos        []*models.Todo
	TodoChannels map[*chan *models.Todo]chan *models.Todo
}

var schema *graphql.ExecutableSchema

// Singleton pattern applied on schema
func ExecutableSchema() graphql.ExecutableSchema {
	if schema == nil {
		schema = new(graphql.ExecutableSchema)
		*schema = generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{
			Todos:        []*models.Todo{},
			TodoChannels: make(map[*chan *models.Todo]chan *models.Todo),
		}})
	}

	return *schema
}
