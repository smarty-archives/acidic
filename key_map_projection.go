package acidic

type KeyMapProjection struct {
}

func NewKeyMapProjection() *KeyMapProjection {
	return &KeyMapProjection{}
}

func (this *KeyMapProjection) Apply(message interface{}) {
	switch message := message.(type) {

	case TransactionStartedEvent:
		this.applyTransactionStarted(message)

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

	case ItemMergedEvent:
		this.applyItemMerged(message)
	}
}

func (this *KeyMapProjection) applyTransactionStarted(message TransactionStartedEvent) {
}

func (this *KeyMapProjection) applyStoringItem(message StoringItemEvent) {
}
func (this *KeyMapProjection) applyItemStored(message ItemStoredEvent) {
}
func (this *KeyMapProjection) applyItemStoreFailed(message ItemStoreFailedEvent) {
}

func (this *KeyMapProjection) applyDeletingItem(message DeletingItemEvent) {
}
func (this *KeyMapProjection) applyItemDeleted(message ItemDeletedEvent) {
}
func (this *KeyMapProjection) applyItemDeleteFailed(message ItemDeleteFailedEvent) {
}

func (this *KeyMapProjection) applyTransactionCommitting(message TransactionCommittingEvent) {
}
func (this *KeyMapProjection) applyTransactionCommitted(message TransactionCommittedEvent) {
}

func (this *KeyMapProjection) applyTransactionFailed(message TransactionFailedEvent) {
}

func (this *KeyMapProjection) applyTransactionAborted(message TransactionAbortedEvent) {
}

func (this *KeyMapProjection) applyItemMerged(message ItemMergedEvent) {
}
