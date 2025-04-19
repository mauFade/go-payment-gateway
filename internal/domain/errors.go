package domain

import "errors"

var (
	// When a account is not found in DB
	ErrAccountNotFound = errors.New("account not found")
	// When API key already exists in DB
	ErrDuplicateAPIKey = errors.New("this API key already exists")
	// When a invoice is not found in DB
	ErrInvoiceNotFound = errors.New("invoice not found")
	// When a user is not able to perform that action
	ErrUnaithorizedAccess = errors.New("unauthorized access")
	// When invalid amount is provided
	ErrInvalidAmount = errors.New("invalid amount")
	// When a invalid status is provided
	ErrInvalidStatus = errors.New("invalid status")
)
