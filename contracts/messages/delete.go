package messages

import "time"

type DeleteItemCommand struct {
	correlationID string // required
	TransactionID string // optional (blank = new tx)
	Key           string
	ETag          string // optional
}
type DeletingItemEvent struct {
	correlationID string // required
	Timestamp     time.Time
	Sequence      uint64
	TransactionID string
	Key           string
	ETag          string // optional
}
type ItemDeletedEvent struct {
	correlationID string // required
	Timestamp     time.Time
	Sequence      uint64
	TransactionID string
	Key           string
}
type ItemDeleteFailedEvent struct {
	correlationID string // required
	Timestamp     time.Time
	Sequence      uint64
	TransactionID string
	Key           string
}
