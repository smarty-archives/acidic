package acidic

import (
	"github.com/smartystreets/acidic/contracts"
	"github.com/smartystreets/acidic/contracts/messages"
	"github.com/smartystreets/clock"
	"time"
)

type Transaction struct {
	publisher Publisher
	status    uint64
	operation uint64
	started   time.Time
	updated   time.Time
	ttl       time.Duration
}

func NewTransaction(publisher Publisher, started time.Time, ttl time.Duration) *Transaction {
	return &Transaction{
		publisher: publisher,
		operation: 0,
		started:   started,
		updated:   started,
		ttl:       ttl,
	}
}

func (this *Transaction) Handle(message interface{}) error {
	// TODO: if timed out, return TimeoutError and raise TransactionFailedEvent

	switch message := message.(type) {

	case messages.StoreItemCommand:
		return this.handleStoreItem(message)
	case messages.ItemStoredEvent:
		return this.handleItemStored(message)
	case messages.ItemStoreFailedEvent:
		return this.handleItemStoreFailed(message)

	case messages.DeleteItemCommand:
		return this.handleDeleteItem(message)
	case messages.ItemDeletedEvent:
		return this.handleItemDeleted(message)
	case messages.ItemDeleteFailedEvent:
		return this.handleItemDeleteFailed(message)

	case messages.CommitTransactionCommand:
		return this.handleCommitTransaction(message)
	case messages.TransactionCommittedEvent:
		return this.handleTransactionCommitted(message)
	case messages.TransactionCommitFailedEvent:
		return this.handleTransactionCommitFailed(message)

	case messages.AbortTransactionCommand:
		return this.handleAbortTransaction(message)

	default:
		return nil
	}
}

func (this *Transaction) handleStoreItem(message messages.StoreItemCommand) error {
	return nil
}
func (this *Transaction) handleItemStored(message messages.ItemStoredEvent) error {
	return nil
}
func (this *Transaction) handleItemStoreFailed(message messages.ItemStoreFailedEvent) error {
	return nil
}

func (this *Transaction) handleDeleteItem(message messages.DeleteItemCommand) error {
	return nil
}
func (this *Transaction) handleItemDeleted(message messages.ItemDeletedEvent) error {
	return nil
}
func (this *Transaction) handleItemDeleteFailed(message messages.ItemDeleteFailedEvent) error {
	return nil
}

func (this *Transaction) handleCommitTransaction(message messages.CommitTransactionCommand) error {
	return nil
}
func (this *Transaction) handleTransactionCommitted(message messages.TransactionCommittedEvent) error {
	return nil
}
func (this *Transaction) handleTransactionCommitFailed(message messages.TransactionCommitFailedEvent) error {
	return nil
}

func (this *Transaction) handleAbortTransaction(message messages.AbortTransactionCommand) error {
	switch this.status {
	case TransactionStateReady, TransactionStateWriting:
		this.publisher.Raise(messages.TransactionAbortedEvent{
			Timestamp:     clock.UTCNow(),
			TransactionID: message.TransactionID,
		})

		return nil
	default:
		return contracts.InvalidTransitionError
	}
}

func (this *Transaction) Apply(message interface{}) {
	switch message := message.(type) {

	case messages.StoringItemEvent:
		this.applyStoringItem(message)
	case messages.ItemStoredEvent:
		this.applyItemStored(message)
	case messages.ItemStoreFailedEvent:
		this.applyItemStoreFailed(message)

	case messages.DeletingItemEvent:
		this.applyDeletingItem(message)
	case messages.ItemDeletedEvent:
		this.applyItemDeleted(message)
	case messages.ItemDeleteFailedEvent:
		this.applyItemDeleteFailed(message)

	case messages.TransactionCommittingEvent:
		this.applyTransactionCommitting(message)
	}
}

func (this *Transaction) applyStoringItem(message messages.StoringItemEvent) {
}
func (this *Transaction) applyItemStored(message messages.ItemStoredEvent) {
}
func (this *Transaction) applyItemStoreFailed(message messages.ItemStoreFailedEvent) {
}

func (this *Transaction) applyDeletingItem(message messages.DeletingItemEvent) {
}
func (this *Transaction) applyItemDeleted(message messages.ItemDeletedEvent) {
}
func (this *Transaction) applyItemDeleteFailed(message messages.ItemDeleteFailedEvent) {
}

func (this *Transaction) applyTransactionCommitting(message messages.TransactionCommittingEvent) {
	if this.status == TransactionStateReady {
		this.status = TransactionStateCommitting
	} else {
		this.status = TransactionStateWritingCommitting
	}

	// park any contexts needed to respond to
}

const (
	TransactionStateReady = iota
	TransactionStateWriting
	TransactionStateWritingCommitting
	TransactionStateCommitting
	// TransactionStateCommitted
	// TransactionStateAborted
	// TransactionStateFailed
)
