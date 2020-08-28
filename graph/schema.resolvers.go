package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/xaoirse/logbook/graph/generated"
	"github.com/xaoirse/logbook/graph/model"
)

func (r *mutationResolver) CreateAction(ctx context.Context, input model.NewAction) (*model.Action, error) {
	action := model.Action{
		Name: &input.Name,
	}
	r.DB.Create(&action)
	return &action, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Actions(ctx context.Context) ([]*model.Action, error) {
	var actions []*model.Action
	r.DB.Find(&actions)
	// var acts []*model.Action
	// for _, a := range actions {
	// 	acts = append(acts, &a)
	// }
	return actions, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
