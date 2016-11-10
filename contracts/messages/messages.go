package messages

import (
	"io"
	"time"
)

type LoadItemRequest struct {
	TransactionID string // optional
	Key           string
	Revision      string // optional (provided by KeyMapProjection)
	ETag          string // optional
}
type LoadItemResponse struct {
	TransactionID string // optional
	ContentLength uint64
	ContentType   string
	Key           string
	ETag          string
	Metadata      map[string]string
	Payload       io.Reader
}

type StoreItemCommand struct {
	TransactionID string // optional (blank = new tx)
	Key           string
	ETag          string // optional
	Metadata      map[string]string
	Payload       io.Reader
}
type TransactionStartedEvent struct {
	Timestamp     time.Time
	TransactionID string
	TTL           time.Duration
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

type CommitTransactionCommand struct {
	TransactionID string
}
type TransactionCommittingEvent struct {
	TransactionID string
	Contents      map[string]*CommitItem
}
type TransactionCommittedEvent struct {
	Timestamp     time.Time
	Commit        uint64
	TransactionID string
}
type TransactionCommitFailedEvent struct {
	Timestamp     time.Time
	TransactionID string
}

type TransactionFailedEvent struct {
	Timestamp     time.Time
	TransactionID string
	Reason        error
}

type AbortTransactionCommand struct {
	TransactionID string
}
type TransactionAbortedEvent struct {
	Timestamp     time.Time
	TransactionID string
}

type CommitWrittenEvent struct {
	Timestamp      time.Time
	Commit         uint64
	TransactionIDs []string
}

type ItemMergedEvent struct {
	Timestamp     time.Time
	Commit        uint64
	Sequence      uint64
	TransactionID string
	Key           string
	Revision      string
	ETag          string
	Expiration    time.Time // the point at which the entry is considered fully merged by the underlying storage
}

type MessageContainer interface {
	Add(interface{})
}

type Publisher interface {
	Raise(interface{})
}

type CommitItem struct {
	Key      string
	Revision string
	ETag     string
}
