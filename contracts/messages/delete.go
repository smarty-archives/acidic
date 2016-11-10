package messages

import "time"

type DeleteItemCommand struct {
	TransactionID string // optional (blank = new tx)
	Key           string
	ETag          string // optional
}
type DeletingItemEvent struct {
	Sequence      uint64
	TransactionID string
	Key           string
	ETag          string // optional
}
type ItemDeletedEvent struct {
	Timestamp     time.Time
	Sequence      uint64
	TransactionID string
	Key           string
}
type ItemDeleteFailedEvent struct {
	Timestamp     time.Time
	Sequence      uint64
	TransactionID string
	Key           string
}
