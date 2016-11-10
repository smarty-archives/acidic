package messages

import "time"

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
