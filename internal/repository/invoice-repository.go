package repository

import (
	"database/sql"

	"github.com/mauFade/go-payment-gateway/internal/domain"
)

type InvoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) *InvoiceRepository {
	return &InvoiceRepository{
		db: db,
	}
}

func (r *InvoiceRepository) Save(i *domain.Invoice) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO invoices (id, account_id, amount, status, description, last_card_digits, payment_type, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		i.ID,
		i.AccountId,
		i.Amount,
		i.Status,
		i.Description,
		i.LastCardDigits,
		i.PaymentType,
		i.CreatedAt,
		i.UpdatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}
