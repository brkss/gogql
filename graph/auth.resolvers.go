package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/brkss/gogql/graph/model"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input *model.LoginUserInput) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input *model.RegisterUserInput) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented: Register - register"))
}

// Me is the resolver for the Me field.
func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Me - Me"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }