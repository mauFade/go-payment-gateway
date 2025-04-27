package dto

import (
	"time"

	"github.com/mauFade/go-payment-gateway/internal/domain"
)

const (
	statusApproved = string(domain.StatusApproved)
	statusPending  = string(domain.StatusPending)
	statusRejected = string(domain.StatusRejected)
)

type CreateInvoiceRequest struct {
	APIKey         string
	Amount         float64 `json:"amount"`
	Description    string  `json:"description"`
	PaymentType    string  `json:"payment_type"`
	CardNumber     string  `json:"card_number"`
	CVV            string  `json:"cvv"`
	ExpiryMonth    int     `json:"expiry_month"`
	ExpiryYear     int     `json:"expiry_year"`
	CardHolderName string  `json:"cardholder_name"`
}

type InvoiceOutput struct {
	ID             string    `json:"id"`
	AccountID      string    `json:"account_id"`
	Amount         float64   `json:"amount"`
	Status         string    `json:"status"`
	Description    string    `json:"description"`
	PaymentType    string    `json:"payment_type"`
	CardLastDigits string    `json:"card_last_digits"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ToInvoice(i *CreateInvoiceRequest, accId string) (*domain.Invoice, error) {
	card := domain.CreditCard{
		CardNumber:     i.CardNumber,
		CVV:            i.CVV,
		ExpiryMonth:    i.ExpiryMonth,
		ExpiryYear:     i.ExpiryYear,
		CardholderName: i.CardHolderName,
	}

	return domain.NewInvoice(accId, i.Description, i.PaymentType, i.Amount, card)
}

func FromInvoice(i *domain.Invoice) InvoiceOutput {
	return InvoiceOutput{
		ID:             i.ID,
		AccountID:      i.AccountID,
		Amount:         i.Amount,
		Status:         string(i.Status),
		Description:    i.Description,
		PaymentType:    i.PaymentType,
		CardLastDigits: i.LastCardDigits,
		CreatedAt:      i.CreatedAt,
		UpdatedAt:      i.UpdatedAt,
	}
}
