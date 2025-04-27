package service

import (
	"github.com/mauFade/go-payment-gateway/internal/domain"
	"github.com/mauFade/go-payment-gateway/internal/dto"
)

type InvoiceService struct {
	invoiceRepository domain.InvoiceRepository
	accountService    AccountService
}

func NewInvoiceService(ir domain.InvoiceRepository, as AccountService) *InvoiceService {
	return &InvoiceService{
		invoiceRepository: ir,
		accountService:    as,
	}
}

func (s *InvoiceService) Create(input *dto.CreateInvoiceRequest) (*dto.InvoiceOutput, error) {
	acc, err := s.accountService.FindByAPIKey(input.APIKey)
	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoice(input, acc.ID)
	if err != nil {
		return nil, err
	}

	if err := invoice.Process(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusApproved {
		_, err := s.accountService.UpdateBalance(input.APIKey, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}
	if err := s.invoiceRepository.Save(invoice); err != nil {
		return nil, err
	}

	return dto.FromInvoice(invoice), nil
}

func (s *InvoiceService) FindByID(id, apiKey string) (*dto.InvoiceOutput, error) {
	i, err := s.invoiceRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	acc, err := s.accountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}
	if i.AccountID != acc.ID {
		return nil, domain.ErrUnaithorizedAccess
	}

	return dto.FromInvoice(i), nil
}

func (s *InvoiceService) ListByAccountID(accId string) ([]*dto.InvoiceOutput, error) {
	invoices, err := s.invoiceRepository.FindByAccountID(accId)
	if err != nil {
		return nil, err
	}

	output := make([]*dto.InvoiceOutput, len(invoices))
	for i, invoice := range invoices {
		output[i] = dto.FromInvoice(invoice)
	}
	return output, nil
}

func (s *InvoiceService) ListByAPIKey(apiKey string) ([]*dto.InvoiceOutput, error) {
	acc, err := s.accountService.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	return s.ListByAccountID(acc.ID)
}
