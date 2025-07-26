package account

import (
	"context"

	pb "github.com/Aadithya-J/microservices-ecommerce/account/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.AccountServiceClient
}

func NewClient(conn *grpc.ClientConn) *Client {
	return &Client{
		conn:    conn,
		service: pb.NewAccountServiceClient(conn),
	}
}

func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *Client) CreateAccount(ctx context.Context, name string) (*Account, error) {
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "name cannot be empty")
	}

	resp, err := c.service.PostAccount(ctx, &pb.PostAccountRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:   resp.Account.Id,
		Name: resp.Account.Name,
	}, nil
}

func (c *Client) GetAccount(ctx context.Context, id string) (*Account, error) {
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "id cannot be empty")
	}

	resp, err := c.service.GetAccount(ctx, &pb.GetAccountRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:   resp.Account.Id,
		Name: resp.Account.Name,
	}, nil
}

func (c *Client) ListAccounts(ctx context.Context, skip, take uint64) ([]*Account, error) {
	if take > 100 {
		take = 100
	}

	resp, err := c.service.GetAccounts(ctx, &pb.GetAccountsRequest{
		Skip: skip,
		Take: take,
	})
	if err != nil {
		return nil, err
	}

	accounts := make([]*Account, len(resp.Accounts))
	for i, acc := range resp.Accounts {
		accounts[i] = &Account{
			ID:   acc.Id,
			Name: acc.Name,
		}
	}

	return accounts, nil
}
