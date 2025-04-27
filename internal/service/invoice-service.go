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
	o := dto.FromInvoice(invoice)

	return &o, nil
}
