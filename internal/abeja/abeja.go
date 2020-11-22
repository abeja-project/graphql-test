package abeja

import (
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// Abeja holds the state of the service.
type Abeja struct {
	schema  *graphql.Schema
	handler http.Handler
}

// New creats a new Abeja.
func New() *Abeja {
	schema := graphql.MustParseSchema(graphQLSchema, new(resolver))

	return &Abeja{
		schema:  schema,
		handler: &relay.Handler{Schema: schema},
	}
}

func (a *Abeja) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler.ServeHTTP(w, r)
}
