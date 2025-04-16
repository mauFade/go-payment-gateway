package dto

import (
	"time"

	"github.com/mauFade/go-payment-gateway/internal/domain"
)

type CreateAccountRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AccountOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	APIKey    string    `json:"api_key"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToAccount(input CreateAccountRequest) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}

func FromAccount(acc *domain.Account) AccountOutput {
	return AccountOutput{
		ID:        acc.ID,
		Name:      acc.Name,
		Email:     acc.Email,
		APIKey:    acc.APIKey,
		Balance:   acc.Balance,
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
	}
}
