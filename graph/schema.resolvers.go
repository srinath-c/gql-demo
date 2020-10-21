package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/srinath-c/gql-demo/graph/generated"
	"github.com/srinath-c/gql-demo/graph/model"
)

var (
	todoStore = make(map[string]*model.Todo)
	userStore = make(map[string]*model.User)
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	//Assigned input variables to local input
	username := input.UserID
	todoText := input.Text

	//checking whether username is available in userStore map
	_, ok := userStore[username]
	if !ok {
		//unique uuid assigned to id
		id := uuid.New().String()
		//added username to userStore
		userStore[username] = &model.User{
			ID:   id,
			Name: username,
		}
	}

	//unique uuid assigned to id
	id := uuid.New().String()

	//Adding todo to todoStore
	todoStore[id] = &model.Todo{
		ID:   id,
		Text: todoText,
		User: userStore[username],
	}
	//returning the created todo
	return todoStore[id], nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	//iterating through todostore to get individual todos
	for _, todo := range todoStore {
		//appending it to output
		todos = append(todos, todo)
	}
	//returning todos
	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
