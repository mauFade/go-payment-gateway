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
		i.AccountID,
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

func (r *InvoiceRepository) FindByID(id string) (*domain.Invoice, error) {
	var invoice domain.Invoice

	err := r.db.QueryRow(`
    SELECT id, account_id, amount, status, description, payment_type, card_last_digits, created_at, updated_at
    FROM invoices
    WHERE id = $1
`, id).Scan(
		&invoice.ID,
		&invoice.AccountID,
		&invoice.Amount,
		&invoice.Status,
		&invoice.Description,
		&invoice.PaymentType,
		&invoice.LastCardDigits,
		&invoice.CreatedAt,
		&invoice.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrInvoiceNotFound
	}

	if err != nil {
		return nil, err
	}

	return &invoice, nil
}
