package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Accounts(ctx context.Context, pagination *PaginationInput, id *string) ([]*Account, error) {
	if id != nil && *id != "" {
		acc, err := r.server.accountClient.GetAccount(ctx, *id)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				return []*Account{}, nil
			}
			return nil, err
		}
		return []*Account{{ID: acc.ID, Name: acc.Name}}, nil
	}

	var skip uint64
	var take uint64 = 100
	if pagination != nil {
		if pagination.Skip != nil && *pagination.Skip >= 0 {
			skip = uint64(*pagination.Skip)
		}
		if pagination.Take != nil && *pagination.Take > 0 {
			t := uint64(*pagination.Take)
			if t < 100 {
				take = t
			}
		}
	}

	list, err := r.server.accountClient.ListAccounts(ctx, skip, take)
	if err != nil {
		return nil, err
	}

	accounts := make([]*Account, len(list))
	for i, acc := range list {
		accounts[i] = &Account{ID: acc.ID, Name: acc.Name}
	}
	return accounts, nil
}

func (r *queryResolver) Products(ctx context.Context, pagination *PaginationInput, id *string) ([]*Product, error) {
	return nil, nil
}

func (r *queryResolver) Orders(ctx context.Context, pagination *PaginationInput, id *string) ([]*Order, error) {
	return nil, nil
}
