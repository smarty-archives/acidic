package messages

import "time"

type TransactionStartedEvent struct {
	Timestamp     time.Time
	TransactionID string
	TTL           time.Duration
}

type TransactionFailedEvent struct {
	Timestamp     time.Time
	TransactionID string
	Reason        error
}

func (this TransactionFailedEvent) CorrelationID() string {
	return "" // TODO
}
