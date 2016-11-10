package messages

import (
	"io"
	"time"
)

type StoreItemCommand struct {
	TransactionID string // optional (blank = new tx)
	Key           string
	ETag          string // optional
	Metadata      map[string]string
	Payload       io.Reader
}
type StoringItemEvent struct {
	Sequence      uint64 // incremented for each mutating operation; this helps us to know which Store was the most recent one
	TransactionID string
	Key           string
	ETag          string // optional
	Payload       io.Reader
	Metadata      map[string]string
}
type ItemStoredEvent struct {
	Timestamp     time.Time
	Sequence      uint64
	TransactionID string
	CanonicalKey  string
	Key           string
	Revision      string
	ETag          string
}
type ItemStoreFailedEvent struct {
	Timestamp     time.Time
	Sequence      uint64
	TransactionID string
	Key           string
}
