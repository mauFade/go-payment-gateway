package domain

type AccountRepository interface {
	Save(acc *Account) error
	FindByAPIKey(key string) (*Account, error)
	FindByID(id string) (*Account, error)
	UpdateBalance(acc *Account) error
}

type InvoiceRepository interface {
	Save(i *Invoice) error
	FindByID(id string) (*Invoice, error)
	FindByAccountID(accID string) ([]*Invoice, error)
	UpdateStatus(i *Invoice) error
}
