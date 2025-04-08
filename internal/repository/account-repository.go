package repository

import (
	"database/sql"
	"time"

	"github.com/mauFade/go-payment-gateway/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) Save(acc *domain.Account) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO accounts (id, name, email, balance, api_key, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		acc.ID,
		acc.Name,
		acc.Email,
		acc.Balance,
		acc.APIKey,
		acc.CreatedAt,
		acc.UpdatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}

func (r *AccountRepository) FindByAPIKey(key string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, balance, api_key, created_at, updated_at 
		FROM accounts 
		WHERE api_key = $1
		`, key).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.Balance,
		&account.APIKey,
		&createdAt,
		&updatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}

	if err != nil {
		return nil, err
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt

	return &account, nil
}

func (r *AccountRepository) FindByID(id string) (*domain.Account, error) {
	var account domain.Account
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, balance, api_key, created_at, updated_at 
		FROM accounts 
		WHERE id = $1
		`, id).Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.Balance,
		&account.APIKey,
		&createdAt,
		&updatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrAccountNotFound
	}

	if err != nil {
		return nil, err
	}

	account.CreatedAt = createdAt
	account.UpdatedAt = updatedAt

	return &account, nil
}
