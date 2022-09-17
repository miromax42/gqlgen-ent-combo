package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gqlgen-ent-combo/ent"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input ent.CreateTaskInput) (*ent.Task, error) {
	return r.client.Task.Create().SetInput(input).Save(ctx)
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, id int, input ent.UpdateTaskInput) (*ent.Task, error) {
	return r.client.Task.UpdateOneID(id).SetInput(input).Save(ctx)
}
