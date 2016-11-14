package messages

import "time"

type AbortTransactionCommand struct {
	TransactionID string
}

type TransactionAbortedEvent struct {
	Timestamp     time.Time
	TransactionID string
}

func (this TransactionAbortedEvent) CorrelationID() string {
	return this.TransactionID
}
