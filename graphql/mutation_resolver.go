package main

import (
	"context"
)

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) CreateAccount(ctx context.Context, input AccountInput) (*Account, error) {
	acc, err := r.server.accountClient.CreateAccount(ctx, input.Name)
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:   acc.ID,
		Name: acc.Name,
	}, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input ProductInput) (*Product, error) {
	return nil, nil // Not implemented
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input OrderInput) (*Order, error) {
	return nil, nil // Not implemented
}
