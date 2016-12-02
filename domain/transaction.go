package acidic

import (
	"time"

	"github.com/smartystreets/acidic/contracts/messages"
	"log"
	"reflect"
)

type Transaction struct {
	raised    Dispatcher
	id        string
	status    uint64
	operation uint64
	started   time.Time
	updated   time.Time
	ttl       time.Duration
}

func NewTransaction(raised Dispatcher, id string, started time.Time, ttl time.Duration) *Transaction {
	return &Transaction{
		raised:    raised,
		id:        id,
		operation: 0,
		started:   started,
		updated:   started,
		ttl:       ttl,
	}
}

func (this *Transaction) Handle(message interface{}) error {
	switch message := message.(type) {

	case messages.StoreItemCommand:
		return this.handleStoreItem(message)
	case messages.DeleteItemCommand:
		return this.handleDeleteItem(message)
	case messages.ItemStoredEvent, messages.ItemDeletedEvent:
		return this.handleWriteCompleted(message)
	case messages.ItemStoreFailedEvent, messages.ItemDeleteFailedEvent:
		return this.handleWriteFailed(message)

	case messages.CommitTransactionCommand:
		return this.handleCommitTransaction(message)
	case messages.TransactionCommittedEvent:
		return this.handleTransactionCommitted(message)
	case messages.TransactionCommitFailedEvent:
		return this.handleTransactionCommitFailed(message)

	case messages.AbortTransactionCommand:
		return this.handleAbortTransaction(message)

	default:
		log.Panicf("Unknown message type: %s", reflect.TypeOf(message).Name())
		return nil
	}
}

func (this *Transaction) handleStoreItem(message messages.StoreItemCommand) error {
	if this.status == txAborted || this.status == txFailed || this.status == txCommitted {
		log.Panicf("Invalid state")
	}
	//if this.status != txReady && this.status == this

	// if !(ready||writing), return error
	// if timed out, raise TransactionFailed and return timeout error
	// concurrency: check outstanding map for possible failures, if so, return concurrency error
	// concurrency: if etags match, return an error that releases the context but is a success, eg. AlreadyWrittenError
	// raise StoringItemEvent
	return nil
}
func (this *Transaction) handleDeleteItem(message messages.DeleteItemCommand) error {
	// if (!ready|writing), return error
	// if timed out, raise TransactionFailed and return timeout error
	// concurrency: check outstanding map for possible concurrency issues; if so, return concurrency error
	// concurrency: if etags match, return an error that releases the context but is a success, eg. AlreadyWrittenError
	// raise DeletingItemEvent
	return nil
}
func (this *Transaction) handleWriteCompleted(message interface{}) error {
	// if (aborted|failed), return
	// if tx timed out, raise TransactionFailed+reason=TransactionTimeoutError
	// apply message
	// if state == txWritingCommitting and > 1 outstanding writes, done
	// raise TransactionCommittingEvent
	return nil
}
func (this *Transaction) handleWriteFailed(message interface{}) error {
	// if (aborted|failed), return
	// raise TransactionFailed+reason=WriteError
	return nil
}

func (this *Transaction) handleCommitTransaction(message messages.CommitTransactionCommand) error {
	// if (txAborted|txFailed), return error
	// if (txWritingCommitting, txCommitting), return nil
	// if txCommitted, return AlreadyCommittedError (which is a HTTP 200/no-op) to the caller
	// if timed out, raise TransactionFailed and return timeout error

	// if state=txWriting, raise TransactionPendingCommitEvent (which transitions to txWritingCommitting)
	// if state=txReady, raise TransactionCommittingEvent
	return nil
}
func (this *Transaction) handleTransactionCommitted(message messages.TransactionCommittedEvent) error {
	this.Apply(message)
	return nil
}
func (this *Transaction) handleTransactionCommitFailed(message messages.TransactionCommitFailedEvent) error {
	// if (txAborted|txFailed), return
	// raise TransactionFailed
	return nil
}

func (this *Transaction) handleAbortTransaction(message messages.AbortTransactionCommand) error {
	// if (txAborted), return nil
	// if (txFailed), return TransactionFailedError
	// if (txCommitted), return AlreadyCommittedError
	// if (txCommitting|txWritingCommitting), return InvalidTransitionError
	// raise TransactionAbortedEvent
	return nil
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

	case messages.TransactionCommitAwaitingWritesEvent:
		this.applyTransactionAwaitingWrites(message)
	case messages.TransactionCommittingEvent:
		this.applyTransactionCommitting(message)

	default:
		log.Panicf("Unknown message type: %s", reflect.TypeOf(message).Name())
	}
}

func (this *Transaction) applyStoringItem(message messages.StoringItemEvent) {
	this.updated = message.Timestamp
	if this.status != txCommittingAwaitingWrites {
		this.status = txWriting
	}
}
func (this *Transaction) applyItemStored(message messages.ItemStoredEvent) {
	this.updated = message.Timestamp
}
func (this *Transaction) applyItemStoreFailed(message messages.ItemStoreFailedEvent) {
	this.updated = message.Timestamp
	this.status = txFailed
}

func (this *Transaction) applyDeletingItem(message messages.DeletingItemEvent) {
	this.updated = message.Timestamp
	if this.status != txCommittingAwaitingWrites {
		this.status = txWriting
	}
}
func (this *Transaction) applyItemDeleted(message messages.ItemDeletedEvent) {
	this.updated = message.Timestamp
}
func (this *Transaction) applyItemDeleteFailed(message messages.ItemDeleteFailedEvent) {
	this.updated = message.Timestamp
	this.status = txFailed
}

func (this *Transaction) applyTransactionAwaitingWrites(message messages.TransactionCommitAwaitingWritesEvent) {
	this.updated = message.Timestamp
	this.status = txCommittingAwaitingWrites
}
func (this *Transaction) applyTransactionCommitting(message messages.TransactionCommittingEvent) {
	this.updated = message.Timestamp
	this.status = txCommitting
}
