package contracts

import "errors"

var (
	TransactionNotFoundError = errors.New("")
	TransactionTimeoutError  = errors.New("")
	ConcurrenyError          = errors.New("")
	KeyNotFoundError         = errors.New("")
	AccessDeniedError        = errors.New("")
	StorageUnavailable       = errors.New("")
	WriteFailedError         = errors.New("")
	InvalidTransitionError   = errors.New("Transaction cannot transition to that state.")
)
