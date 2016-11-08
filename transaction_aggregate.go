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
		return this.handleStoreItem(message)
	case ItemStoredEvent:
		return this.handleItemStored(message)
	case ItemStoreFailedEvent:
		return this.handleItemStoreFailed(message)

	case DeleteItemCommand:
		return this.handleDeleteItem(message)
	case ItemDeletedEvent:
		return this.handleItemDeleted(message)
	case ItemDeleteFailedEvent:
		return this.handleItemDeleteFailed(message)

	case CommitTransactionCommand:
		return this.handleCommitTransaction(message)
	case TransactionCommittedEvent:
		return this.handleTransactionCommitted(message)
	case TransactionCommitFailedEvent:
		return this.handleTransactionCommitFailed(message)

	case AbortTransactionCommand:
		return this.handleAbortTransaction(message)

	default:
		return nil
	}
}

func (this *TransactionAggregate) handleStoreItem(message StoreItemCommand) error {
	message.TransactionID = this.startTransaction(message.TransactionID)

	// TODO
	return nil
}
func (this *TransactionAggregate) handleItemStored(message ItemStoredEvent) error {
	return nil
}
func (this *TransactionAggregate) handleItemStoreFailed(message ItemStoreFailedEvent) error {
	return this.raiseTransactionFailed(message.TransactionID)
}

func (this *TransactionAggregate) handleDeleteItem(message DeleteItemCommand) error {
	message.TransactionID = this.startTransaction(message.TransactionID)

	// TODO
	return nil
}
func (this *TransactionAggregate) handleItemDeleted(message ItemDeletedEvent) error {
	return nil
}
func (this *TransactionAggregate) handleItemDeleteFailed(message ItemDeleteFailedEvent) error {
	return this.raiseTransactionFailed(message.TransactionID)
}

func (this *TransactionAggregate) handleCommitTransaction(message CommitTransactionCommand) error {
	return nil
}
func (this *TransactionAggregate) handleTransactionCommitted(message TransactionCommittedEvent) error {
	return nil
}
func (this *TransactionAggregate) handleTransactionCommitFailed(message TransactionCommitFailedEvent) error {
	return this.raiseTransactionFailed(message.TransactionID)
}

func (this *TransactionAggregate) handleAbortTransaction(message AbortTransactionCommand) error {
	return nil
}

func (this *TransactionAggregate) raiseTransactionFailed(transactionID string) error {
	this.raise(TransactionFailedEvent{
		Timestamp:     clock.UTCNow(),
		TransactionID: transactionID,
		Reason:        WriteFailedError,
	})

	return WriteFailedError
}

func (this *TransactionAggregate) raise(message interface{}) {
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
	this.open[message.TransactionID] = NewTransaction(this.raised, message.Timestamp, message.TTL)
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

	this.raise(TransactionStartedEvent{
		Timestamp:     clock.UTCNow(),
		TransactionID: random.GUIDString(),
		TTL:           this.ttl,
	})

	return transactionID
}
func (this *TransactionAggregate) removeTransaction(transactionID string) {
	delete(this.open, transactionID)
}
