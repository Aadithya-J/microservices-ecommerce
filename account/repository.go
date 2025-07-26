package account

import (
	"context"
	"database/sql"
)

type Repository interface {
	Close()
	PutAccount(ctx context.Context, account *Account) error
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	if r.db != nil {
		r.db.Close()
	}
}

func (r *postgresRepository) ping() error {
	if r.db == nil {
		return sql.ErrConnDone
	}
	return r.db.Ping()
}

func (r *postgresRepository) PutAccount(ctx context.Context, account *Account) error {
	if err := r.ping(); err != nil {
		return err
	}

	query := `INSERT INTO accounts (id, name) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, account.ID, account.Name)
	return err
}

func (r *postgresRepository) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	if err := r.ping(); err != nil {
		return nil, err
	}

	query := `SELECT id, name FROM accounts WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)

	account := &Account{}
	if err := row.Scan(&account.ID, &account.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return account, nil
}

func (r *postgresRepository) ListAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	if err := r.ping(); err != nil {
		return nil, err
	}

	query := `SELECT id, name FROM accounts ORDER BY id LIMIT $1 OFFSET $2`
	rows, err := r.db.QueryContext(ctx, query, take, skip)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []Account
	for rows.Next() {
		account := Account{}
		if err := rows.Scan(&account.ID, &account.Name); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
