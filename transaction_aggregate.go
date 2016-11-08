package acidic

type TransactionAggregate struct {
	messages MessageContainer
}

func NewTransactionAggregate(messages MessageContainer) *TransactionAggregate {
	return &TransactionAggregate{
		messages: messages,
	}
}

func (this *TransactionAggregate) Handle(message interface{}) {
	switch message := message.(type) {

	case StoreItemCommand:
		this.handleStoreItem(message)
	case ItemStoredEvent:
		this.handleItemStored(message)
	case ItemStoreFailedEvent:
		this.handleItemStoreFailed(message)

	case DeleteItemCommand:
		this.handleDeleteItem(message)
	case ItemDeletedEvent:
		this.handleItemDeleted(message)
	case ItemDeleteFailedEvent:
		this.handleItemDeleteFailed(message)

	case CommitTransactionCommand:
		this.handleCommitTransaction(message)
	case TransactionCommittedEvent:
		this.handleTransactionCommitted(message)

	case TransactionFailedEvent:
		this.handleTransactionFailed(message)

	case AbortTransactionCommand:
		this.handleAbortTransaction(message)
	}
}

func (this *TransactionAggregate) handleStoreItem(message StoreItemCommand) {
}
func (this *TransactionAggregate) handleItemStored(message ItemStoredEvent) {
}
func (this *TransactionAggregate) handleItemStoreFailed(message ItemStoreFailedEvent) {
}

func (this *TransactionAggregate) handleDeleteItem(message DeleteItemCommand) {
}
func (this *TransactionAggregate) handleItemDeleted(message ItemDeletedEvent) {
}
func (this *TransactionAggregate) handleItemDeleteFailed(message ItemDeleteFailedEvent) {
}

func (this *TransactionAggregate) handleCommitTransaction(message CommitTransactionCommand) {
}
func (this *TransactionAggregate) handleTransactionCommitted(message TransactionCommittedEvent) {
}

func (this *TransactionAggregate) handleTransactionFailed(message TransactionFailedEvent) {
}

func (this *TransactionAggregate) handleAbortTransaction(message AbortTransactionCommand) {
}

func (this *TransactionAggregate) raise(message interface{}) {
	this.messages.Add(message)
	this.Apply(message)
}
func (this *TransactionAggregate) Apply(message interface{}) {
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
	case TransactionAbortFailedEvent:
		this.applyTransactionAbortFailed(message)
	}
}

func (this *TransactionAggregate) applyTransactionStarted(message TransactionStartedEvent) {
}

func (this *TransactionAggregate) applyStoringItem(message StoringItemEvent) {
}
func (this *TransactionAggregate) applyItemStored(message ItemStoredEvent) {
}
func (this *TransactionAggregate) applyItemStoreFailed(message ItemStoreFailedEvent) {
}

func (this *TransactionAggregate) applyDeletingItem(message DeletingItemEvent) {
}
func (this *TransactionAggregate) applyItemDeleted(message ItemDeletedEvent) {
}
func (this *TransactionAggregate) applyItemDeleteFailed(message ItemDeleteFailedEvent) {
}

func (this *TransactionAggregate) applyTransactionCommitting(message TransactionCommittingEvent) {
}
func (this *TransactionAggregate) applyTransactionCommitted(message TransactionCommittedEvent) {
}

func (this *TransactionAggregate) applyTransactionFailed(message TransactionFailedEvent) {
}

func (this *TransactionAggregate) applyTransactionAborted(message TransactionAbortedEvent) {
}
func (this *TransactionAggregate) applyTransactionAbortFailed(message TransactionAbortFailedEvent) {
}
