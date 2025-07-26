package main

import "context"

type accountResolver struct {
	server *Server
}

func (r *accountResolver) Orders(ctx context.Context, obj *Account) ([]*Order, error) {
	// We are not implementing this yet, as it requires the order service.
	return []*Order{}, nil
}

// func (r *accountResolver) Orders(ctx context.Context, obj *Account ) ([]*Order, error) {
// }
