package domain

import (
	"math/rand/v2"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

type Invoice struct {
	ID             string
	AccountID      string
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
		AccountID:      accId,
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

	randomSource := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))
	var NewStatus Status
	if randomSource.Float64() <= 0.7 {
		NewStatus = StatusApproved
	} else {
		NewStatus = StatusRejected
	}

	i.Status = NewStatus
	return nil
}

func (i *Invoice) UpdateStatus(newStatus Status) error {
	if i.Status != StatusPending {
		return ErrInvalidStatus
	}

	i.Status = newStatus
	i.UpdatedAt = time.Now()
	return nil
}
