package acidic

import "time"

type Transaction struct {
	raised   Raiser
	sequence uint64
	started  time.Time
	updated  time.Time
	ttl      time.Duration
}

func NewTransaction(raised Raiser, started time.Time, ttl time.Duration) *Transaction {
	return &Transaction{
		raised:   raised,
		sequence: 0,
		started:  started,
		updated:  started,
		ttl:      ttl,
	}
}

func (this *Transaction) Handle(message interface{}) error {
	return nil
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
	case TransactionCommittedEvent:
		this.applyTransactionCommitted(message)

	case TransactionFailedEvent:
		this.applyTransactionFailed(message)

	case TransactionAbortedEvent:
		this.applyTransactionAborted(message)
	case TransactionAbortFailedEvent:
		this.applyTransactionAbortFailed(message)
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
}
func (this *Transaction) applyTransactionCommitted(message TransactionCommittedEvent) {
}

func (this *Transaction) applyTransactionFailed(message TransactionFailedEvent) {
}

func (this *Transaction) applyTransactionAborted(message TransactionAbortedEvent) {

}
func (this *Transaction) applyTransactionAbortFailed(message TransactionAbortFailedEvent) {
}
