package service

import "github.com/mauFade/go-payment-gateway/internal/domain"

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(r domain.AccountRepository) *AccountService {
	return &AccountService{
		repository: r,
	}
}

func (s *AccountService) CreateAccount()
