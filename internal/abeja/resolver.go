package abeja

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
)

type resolver struct {
	todos []*Todo
}

func (r *resolver) Todos(ctx context.Context) ([]*todoResolver, error) {
	var resolvers []*todoResolver
	for _, todo := range r.todos {
		resolvers = append(resolvers, &todoResolver{todo})
	}
	return resolvers, nil
}

func (r *resolver) CreateTodo(ctx context.Context, args struct{ Input newTodo }) (*todoResolver, error) {
	todo := &Todo{
		ID:     "1",
		Text:   args.Input.Text,
		Done:   false,
		UserID: "1",
	}
	r.todos = append(r.todos, todo)
	return &todoResolver{todo}, nil
}

// Todo is something, that has to be done.
type Todo struct {
	ID     graphql.ID
	Text   string
	Done   bool
	UserID graphql.ID
}

type todoResolver struct {
	t *Todo
}

func (r *todoResolver) ID() graphql.ID {
	return r.t.ID
}

func (r *todoResolver) Text() string {
	return r.t.Text
}

func (r *todoResolver) Done() bool {
	return r.t.Done
}

func (r *todoResolver) User() (*userResolver, error) {
	if r.t.UserID != "1" {
		return nil, fmt.Errorf("user with id %s does not exist", r.t.UserID)
	}
	return &userResolver{&User{ID: "1", Name: "Anja"}}, nil
}

// User is someone who can do something.
type User struct {
	ID   graphql.ID
	Name string
}

type userResolver struct {
	u *User
}

func (r *userResolver) ID() graphql.ID {
	return r.u.ID
}

func (r *userResolver) Name() string {
	return r.u.Name
}

type newTodo struct {
	Text   string
	UserID graphql.ID
}
