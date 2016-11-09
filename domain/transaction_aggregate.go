package acidic

import (
	"time"

	"github.com/smartystreets/acidic/contracts"
	"github.com/smartystreets/acidic/contracts/messages"
	"github.com/smartystreets/clock"
	"github.com/smartystreets/random"
)

type TransactionAggregate struct {
	raised messages.MessageContainer
	open   map[string]*Transaction
	ttl    time.Duration
}

func NewTransactionAggregate(raised messages.MessageContainer) *TransactionAggregate {
	return &TransactionAggregate{
		raised: raised,
		open:   make(map[string]*Transaction),
		ttl:    time.Minute * 5,
	}
}

func (this *TransactionAggregate) Handle(message interface{}) error {
	switch message := message.(type) {

	case messages.StoreItemCommand:
		message.TransactionID = this.tryStartTransaction(message.TransactionID)
		return this.tryHandle(message.TransactionID, message)
	case messages.ItemStoredEvent:
		return this.tryHandle(message.TransactionID, message)
	case messages.ItemStoreFailedEvent:
		return this.tryHandle(message.TransactionID, message)

	case messages.DeleteItemCommand:
		message.TransactionID = this.tryStartTransaction(message.TransactionID)
		return this.tryHandle(message.TransactionID, message)
	case messages.ItemDeletedEvent:
		return this.tryHandle(message.TransactionID, message)
	case messages.ItemDeleteFailedEvent:
		return this.tryHandle(message.TransactionID, message)

	case messages.CommitTransactionCommand:
		return this.tryHandle(message.TransactionID, message)
	case messages.TransactionCommittedEvent:
		return this.tryHandle(message.TransactionID, message)
	case messages.TransactionCommitFailedEvent:
		return this.tryHandle(message.TransactionID, message)

	case messages.AbortTransactionCommand:
		return this.tryHandle(message.TransactionID, message)

	default:
		return nil
	}
}
func (this *TransactionAggregate) tryHandle(transactionID string, message interface{}) error {
	if transaction, contains := this.open[transactionID]; contains {
		return transaction.Handle(message)
	} else {
		return contracts.TransactionNotFoundError
	}
}

func (this *TransactionAggregate) Raise(message interface{}) {
	this.raised.Add(message)
	this.apply(message)
}
func (this *TransactionAggregate) apply(message interface{}) {
	switch message := message.(type) {

	case messages.TransactionStartedEvent:
		this.applyTransactionStarted(message)

	case messages.StoringItemEvent:
		this.tryApply(message.TransactionID, message)
	case messages.ItemStoredEvent:
		this.tryApply(message.TransactionID, message)
	case messages.ItemStoreFailedEvent:
		this.tryApply(message.TransactionID, message)

	case messages.DeletingItemEvent:
		this.tryApply(message.TransactionID, message)
	case messages.ItemDeletedEvent:
		this.tryApply(message.TransactionID, message)
	case messages.ItemDeleteFailedEvent:
		this.tryApply(message.TransactionID, message)

	case messages.TransactionCommittingEvent:
		// TODO: consider keys that are committing/committed in other transactions which might conflict here
		// TODO: maybe we take a dependency on the KeyMapProjection here???
		this.tryApply(message.TransactionID, message)
	case messages.TransactionCommittedEvent:
		this.removeTransaction(message.TransactionID)

	case messages.TransactionFailedEvent:
		this.removeTransaction(message.TransactionID)

	case messages.TransactionAbortedEvent:
		this.removeTransaction(message.TransactionID)
	}
}

func (this *TransactionAggregate) applyTransactionStarted(message messages.TransactionStartedEvent) {
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

	this.Raise(messages.TransactionStartedEvent{
		Timestamp:     clock.UTCNow(),
		TransactionID: random.GUIDString(),
		TTL:           this.ttl,
	})

	return transactionID
}
func (this *TransactionAggregate) removeTransaction(transactionID string) {
	delete(this.open, transactionID)
}
