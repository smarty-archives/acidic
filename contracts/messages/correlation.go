package messages

import (
	"io"
	"strconv"
	"sync/atomic"
)

const base10 = 10

var correlationID uint64

func NextID() string {
	value := atomic.AddUint64(&correlationID, 1)
	return strconv.FormatUint(value, base10)
}

func NewDeleteItemCommand(transaction, key, etag string) DeleteItemCommand {
	return DeleteItemCommand{
		correlationID: NextID(),
		TransactionID: transaction,
		Key:           key,
		ETag:          etag,
	}
}

func NewStoreItemCommand(transaction, key, etag string, metadata map[string]string, payload io.Reader) StoreItemCommand {
	return StoreItemCommand{
		correlationID: NextID(),
		TransactionID: transaction,
		Key:           key,
		ETag:          etag,
		Metadata:      metadata,
		Payload:       payload,
	}
}

func (this TransactionAbortedEvent) CorrelationID() string   { return this.TransactionID }
func (this CommitTransactionCommand) CorrelationID() string  { return this.TransactionID }
func (this DeleteItemCommand) CorrelationID() string         { return this.correlationID }
func (this ItemDeletedEvent) CorrelationID() string          { return this.correlationID }
func (this ItemDeleteFailedEvent) CorrelationID() string     { return this.correlationID }
func (this StoreItemCommand) CorrelationID() string          { return this.correlationID }
func (this ItemStoredEvent) CorrelationID() string           { return this.correlationID }
func (this ItemStoreFailedEvent) CorrelationID() string      { return this.correlationID }
func (this TransactionFailedEvent) CorrelationID() string    { return this.TransactionID }
func (this TransactionCommittedEvent) CorrelationID() string { return this.TransactionID }
