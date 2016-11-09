package acidic

import (
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

func (this *Transaction) handleStoreItem(message StoreItemCommand) error {
	return nil
}
func (this *Transaction) handleItemStored(message ItemStoredEvent) error {
	return nil
}
func (this *Transaction) handleItemStoreFailed(message ItemStoreFailedEvent) error {
	return nil
}

func (this *Transaction) handleDeleteItem(message DeleteItemCommand) error {
	return nil
}
func (this *Transaction) handleItemDeleted(message ItemDeletedEvent) error {
	return nil
}
func (this *Transaction) handleItemDeleteFailed(message ItemDeleteFailedEvent) error {
	return nil
}

func (this *Transaction) handleCommitTransaction(message CommitTransactionCommand) error {
	return nil
}
func (this *Transaction) handleTransactionCommitted(message TransactionCommittedEvent) error {
	return nil
}
func (this *Transaction) handleTransactionCommitFailed(message TransactionCommitFailedEvent) error {
	return nil
}

func (this *Transaction) handleAbortTransaction(message AbortTransactionCommand) error {
	switch this.status {
	case TransactionStateReady, TransactionStateWriting:
		this.publisher.Raise(TransactionAbortedEvent{
			Timestamp:     clock.UTCNow(),
			TransactionID: message.TransactionID,
		})

		return nil
	default:
		return InvalidTransitionError
	}
}

func (this *Transaction) Apply(message interface{}) {
	switch message := message.(type) {

	case StoringItemEvent:
		this.applyStoringItem(message)
	case ItemStoredEvent:
		this.applyItemStored(message)
	case ItemStoreFailedEvent:
		this.applyItemStoreFailed(message)

	case DeletingItemEvent:
		this.applyDeletingItem(message)
	case ItemDeletedEvent:
		this.applyItemDeleted(message)
	case ItemDeleteFailedEvent:
		this.applyItemDeleteFailed(message)

	case TransactionCommittingEvent:
		this.applyTransactionCommitting(message)
	}
}

func (this *Transaction) applyStoringItem(message StoringItemEvent) {
}
func (this *Transaction) applyItemStored(message ItemStoredEvent) {
}
func (this *Transaction) applyItemStoreFailed(message ItemStoreFailedEvent) {
}

func (this *Transaction) applyDeletingItem(message DeletingItemEvent) {
}
func (this *Transaction) applyItemDeleted(message ItemDeletedEvent) {
}
func (this *Transaction) applyItemDeleteFailed(message ItemDeleteFailedEvent) {
}

func (this *Transaction) applyTransactionCommitting(message TransactionCommittingEvent) {
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
