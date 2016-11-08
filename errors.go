package acidic

import "errors"

var (
	TransactionNotFoundError = errors.New("")
	TransactionTimeoutError  = errors.New("")
	ConcurrenyError          = errors.New("")
	KeyNotFoundError         = errors.New("")
	AccessDeniedError        = errors.New("")
	StorageUnavailable       = errors.New("")
	WriteFailedError         = errors.New("")
)
