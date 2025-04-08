package domain

type AccountRepository interface {
	Save(acc *Account) error
	FindByAPIKey(key string) (*Account, error)
	FindByID(id string) (*Account, error)
	Update(acc *Account) error
}
