package main

import (
	"context"
)

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Accounts(ctx context.Context, pagination *PaginationInput, id *string) ([]*Account, error) {
	list, err := r.server.accountClient.ListAccounts(ctx, 0, 100) // TODO: Use pagination input
	if err != nil {
		return nil, err
	}

	var accounts []*Account
	for _, acc := range list {
		accounts = append(accounts, &Account{ID: acc.ID, Name: acc.Name})
	}

	return accounts, nil
}

func (r *queryResolver) Products(ctx context.Context, pagination *PaginationInput, id *string) ([]*Product, error) {
	return nil, nil // Not implemented
}

func (r *queryResolver) Orders(ctx context.Context, pagination *PaginationInput, id *string) ([]*Order, error) {
	return nil, nil // Not implemented
}
