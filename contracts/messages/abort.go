package messages

import "time"

type AbortTransactionCommand struct {
	TransactionID string
}

type TransactionAbortedEvent struct {
	Timestamp     time.Time
	TransactionID string
}

func (this AbortTransactionCommand) CorrelationID() string {
	return "" // TODO
}

func (this TransactionAbortedEvent) CorrelationID() string {
	return "" // TODO
}
