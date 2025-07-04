package account

import (
	"context"

	"github.com/segmentio/ksuid"
)

type Service interface {
	PostAccount(ctx context.Context, name string) (Account, error)
	GetAccount(ctx context.Context, id string) (Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type accountService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &accountService{repository: repository}
}

func (s *accountService) PostAccount(ctx context.Context, name string) (Account, error) {
	acc := &Account{
		Name: name,
		ID:   ksuid.New().String(),
	}
	if err := s.repository.PutAccount(ctx, acc); err != nil {
		return Account{}, err
	}
	return *acc, nil
}

func (s *accountService) GetAccount(ctx context.Context, id string) (Account, error) {
	acc, err := s.repository.GetAccountByID(ctx, id)
	if err != nil {
		return Account{}, err
	}
	return *acc, nil
}

func (s *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	if take > 100 || (skip == 0 && take == 0) {
		return nil, nil
	}
	accounts, err := s.repository.ListAccounts(ctx, skip, take)
	if err != nil {
		return nil, err
	}
	result := make([]Account, len(accounts))
	for i, acc := range accounts {
		result[i] = Account{
			ID:   acc.ID,
			Name: acc.Name,
		}
	}
	return result, nil
}
