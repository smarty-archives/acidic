package acidic

import (
	"time"

	"github.com/smartystreets/clock"
	"github.com/smartystreets/random"
)

type TransactionAggregate struct {
	raised MessageContainer
	open   map[string]*Transaction
	ttl    time.Duration
}

func NewTransactionAggregate(raised MessageContainer) *TransactionAggregate {
	return &TransactionAggregate{
		raised: raised,
		open:   make(map[string]*Transaction),
		ttl:    time.Minute * 5,
	}
}

func (this *TransactionAggregate) Handle(message interface{}) error {
	switch message := message.(type) {

	case StoreItemCommand:
		message.TransactionID = this.startTransaction(message.TransactionID)
		return this.handle(message.TransactionID, message)
	case ItemStoredEvent:
		return this.handle(message.TransactionID, message)
	case ItemStoreFailedEvent:
		return this.handle(message.TransactionID, message)

	case DeleteItemCommand:
		message.TransactionID = this.startTransaction(message.TransactionID)
		return this.handle(message.TransactionID, message)
	case ItemDeletedEvent:
		return this.handle(message.TransactionID, message)
	case ItemDeleteFailedEvent:
		return this.handle(message.TransactionID, message)

	case CommitTransactionCommand:
		return this.handle(message.TransactionID, message)
	case TransactionCommittedEvent:
		return this.handle(message.TransactionID, message)
	case TransactionCommitFailedEvent:
		return this.handle(message.TransactionID, message)

	case AbortTransactionCommand:
		return this.handle(message.TransactionID, message)

	default:
		return nil
	}
}
func (this *TransactionAggregate) handle(transactionID string, message interface{}) error {
	if transaction, contains := this.open[transactionID]; contains {
		return transaction.Handle(message)
	} else {
		return TransactionNotFoundError
	}
}

func (this *TransactionAggregate) Raise(message interface{}) {
	this.raised.Add(message)
	this.apply(message)
}
func (this *TransactionAggregate) apply(message interface{}) {
	switch message := message.(type) {

	case TransactionStartedEvent:
		this.applyTransactionStarted(message)

	case StoringItemEvent:
		this.applyMessage(message.TransactionID, message)
	case ItemStoredEvent:
		this.applyMessage(message.TransactionID, message)
	case ItemStoreFailedEvent:
		this.applyMessage(message.TransactionID, message)

	case DeletingItemEvent:
		this.applyMessage(message.TransactionID, message)
	case ItemDeletedEvent:
		this.applyMessage(message.TransactionID, message)
	case ItemDeleteFailedEvent:
		this.applyMessage(message.TransactionID, message)

	case TransactionCommittingEvent:
		this.applyMessage(message.TransactionID, message)
	case TransactionCommittedEvent:
		this.applyMessage(message.TransactionID, message)

	case TransactionFailedEvent:
		this.applyMessage(message.TransactionID, message)

	case TransactionAbortedEvent:
		this.applyMessage(message.TransactionID, message)
	case TransactionAbortFailedEvent:
		this.applyMessage(message.TransactionID, message)
	}
}

func (this *TransactionAggregate) applyTransactionStarted(message TransactionStartedEvent) {
	this.open[message.TransactionID] = NewTransaction(this, message.Timestamp, message.TTL)
}
func (this *TransactionAggregate) applyTransactionAborted(message TransactionAbortedEvent) {
	this.removeTransaction(message.TransactionID)
}
func (this *TransactionAggregate) applyTransactionFailed(message TransactionFailedEvent) {
	this.removeTransaction(message.TransactionID)
}
func (this *TransactionAggregate) applyMessage(transactionID string, message interface{}) {
	if transaction, contains := this.open[transactionID]; contains {
		transaction.Apply(message)
	}
}

func (this *TransactionAggregate) startTransaction(transactionID string) string {
	if len(transactionID) > 0 {
		return transactionID
	}

	this.Raise(TransactionStartedEvent{
		Timestamp:     clock.UTCNow(),
		TransactionID: random.GUIDString(),
		TTL:           this.ttl,
	})

	return transactionID
}
func (this *TransactionAggregate) removeTransaction(transactionID string) {
	delete(this.open, transactionID)
}
