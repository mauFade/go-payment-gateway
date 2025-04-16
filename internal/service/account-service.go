package service

import (
	"github.com/mauFade/go-payment-gateway/internal/domain"
	"github.com/mauFade/go-payment-gateway/internal/dto"
)

type AccountService struct {
	repository domain.AccountRepository
}

func NewAccountService(r domain.AccountRepository) *AccountService {
	return &AccountService{
		repository: r,
	}
}

func (s *AccountService) CreateAccount(input dto.CreateAccountRequest) (*dto.AccountOutput, error) {
	account := dto.ToAccount(input)

	existingAccount, err := s.repository.FindByAPIKey(account.APIKey)

	if err != nil && err != domain.ErrAccountNotFound {
		return nil, err
	}

	if existingAccount != nil {
		return nil, domain.ErrDuplicateAPIKey
	}

	err = s.repository.Save(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountOutput, error) {
	acc, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	acc.AddBalance(amount)
	err = s.repository.UpdateBalance(acc)

	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(acc)
	return &output, nil
}

func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountOutput, error) {
	acc, err := s.repository.FindByAPIKey(apiKey)

	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(acc)
	return &output, nil
}

func (s *AccountService) FindByID(id string) (*dto.AccountOutput, error) {
	acc, err := s.repository.FindByID(id)

	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(acc)
	return &output, nil
}
