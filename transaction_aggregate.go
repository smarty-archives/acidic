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
		message.TransactionID = this.tryStartTransaction(message.TransactionID)
		return this.tryHandle(message.TransactionID, message)
	case ItemStoredEvent:
		return this.tryHandle(message.TransactionID, message)
	case ItemStoreFailedEvent:
		return this.tryHandle(message.TransactionID, message)

	case DeleteItemCommand:
		message.TransactionID = this.tryStartTransaction(message.TransactionID)
		return this.tryHandle(message.TransactionID, message)
	case ItemDeletedEvent:
		return this.tryHandle(message.TransactionID, message)
	case ItemDeleteFailedEvent:
		return this.tryHandle(message.TransactionID, message)

	case CommitTransactionCommand:
		return this.tryHandle(message.TransactionID, message)
	case TransactionCommittedEvent:
		return this.tryHandle(message.TransactionID, message)
	case TransactionCommitFailedEvent:
		return this.tryHandle(message.TransactionID, message)

	case AbortTransactionCommand:
		return this.tryHandle(message.TransactionID, message)

	default:
		return nil
	}
}
func (this *TransactionAggregate) tryHandle(transactionID string, message interface{}) error {
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
		this.tryApply(message.TransactionID, message)
	case ItemStoredEvent:
		this.tryApply(message.TransactionID, message)
	case ItemStoreFailedEvent:
		this.tryApply(message.TransactionID, message)

	case DeletingItemEvent:
		this.tryApply(message.TransactionID, message)
	case ItemDeletedEvent:
		this.tryApply(message.TransactionID, message)
	case ItemDeleteFailedEvent:
		this.tryApply(message.TransactionID, message)

	case TransactionCommittingEvent:
		this.tryApply(message.TransactionID, message)
	case TransactionCommittedEvent:
		this.tryApply(message.TransactionID, message)

	case TransactionFailedEvent:
		this.removeTransaction(message.TransactionID)

	case TransactionAbortedEvent:
		this.removeTransaction(message.TransactionID)
	}
}

func (this *TransactionAggregate) applyTransactionStarted(message TransactionStartedEvent) {
	this.open[message.TransactionID] = NewTransaction(this, message.Timestamp, message.TTL)
}
func (this *TransactionAggregate) tryApply(transactionID string, message interface{}) {
	if transaction, contains := this.open[transactionID]; contains {
		transaction.Apply(message)
	}
}

func (this *TransactionAggregate) tryStartTransaction(transactionID string) string {
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
