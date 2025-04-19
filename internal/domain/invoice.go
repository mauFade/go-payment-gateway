package domain

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type Invoice struct {
	ID             string
	AccountId      string
	Amount         float64
	Status         Status
	Description    string
	LastCardDigits string
	PaymentType    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type CreditCard struct {
	CardNumber     string
	CVV            string
	ExpiryMonth    int
	ExpirytYear    int
	CardholderName string
}

func NewInvoice(accId, description, paymentType string, amount float64, card CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	lastDigits := card.CardNumber[len(card.CardNumber)-4:]

	return &Invoice{
		ID:             uuid.NewString(),
		AccountId:      accId,
		Status:         StatusPending,
		Amount:         amount,
		Description:    description,
		PaymentType:    paymentType,
		LastCardDigits: lastDigits,
	}, nil
}

func (i *Invoice) Process() error {
	if i.Amount > 10000 {
		return nil
	}

	randomSource := rand.New(rand.NewSource(time.Now().Unix()))
	var NewStatus Status
	if randomSource.Float64() <= 0.7 {
		NewStatus = StatusApproved
	} else {
		NewStatus = StatusRejected
	}

	i.Status = NewStatus
	return nil
}
