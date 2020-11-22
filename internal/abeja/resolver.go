package abeja

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/graph-gophers/graphql-go"
)

type resolver struct {
	db Database
}

func (r *resolver) Todos(ctx context.Context) ([]*todoResolver, error) {
	todos, err := r.db.Objects(ctx, "todo")
	if err != nil {
		return nil, fmt.Errorf("getting todos from db: %w", err)
	}

	resolvers := make([]*todoResolver, 0, len(todos))
	for id, encoded := range todos {
		var todo Todo
		if err := json.Unmarshal(encoded, &todo); err != nil {
			return nil, fmt.Errorf("decoding todo %d", id)
		}

		resolvers = append(resolvers, &todoResolver{&todo})
	}

	return resolvers, nil
}

func (r *resolver) CreateTodo(ctx context.Context, args struct{ Input newTodo }) (*todoResolver, error) {
	id, err := r.db.NewID(ctx, "todo")
	if err != nil {
		return nil, fmt.Errorf("creating new todo id: %w", err)
	}

	userID, err := strconv.Atoi(string(args.Input.UserID))
	if err != nil {
		return nil, fmt.Errorf("invalid user id `%s`", args.Input.UserID)
	}

	todo := &Todo{
		ID:     id,
		Text:   args.Input.Text,
		Done:   false,
		UserID: userID,
	}

	raw, err := json.Marshal(todo)
	if err != nil {
		return nil, fmt.Errorf("encoding todo: %w", err)
	}

	if err := r.db.Update(ctx, "todo", id, raw); err != nil {
		return nil, fmt.Errorf("saving todo: %w", err)
	}

	return &todoResolver{todo}, nil
}

// Todo is something, that has to be done.
type Todo struct {
	ID     int
	Text   string
	Done   bool
	UserID int
}

type todoResolver struct {
	t *Todo
}

func (r *todoResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(r.t.ID))
}

func (r *todoResolver) Text() string {
	return r.t.Text
}

func (r *todoResolver) Done() bool {
	return r.t.Done
}

func (r *todoResolver) User() (*userResolver, error) {
	if r.t.UserID != 1 {
		return nil, fmt.Errorf("user with id %d does not exist", r.t.UserID)
	}
	return &userResolver{&User{ID: 1, Name: "Anja"}}, nil
}

// User is someone who can do something.
type User struct {
	ID   int
	Name string
}

type userResolver struct {
	u *User
}

func (r *userResolver) ID() graphql.ID {
	return graphql.ID(strconv.Itoa(r.u.ID))
}

func (r *userResolver) Name() string {
	return r.u.Name
}

type newTodo struct {
	Text   string
	UserID graphql.ID
}
