package acidic

type TransactionAggregate struct {
	raised MessageContainer
}

func NewTransactionAggregate(raised MessageContainer) *TransactionAggregate {
	return &TransactionAggregate{
		raised: raised,
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

	case TransactionFailedEvent:
		return this.handleTransactionFailed(message)

	case AbortTransactionCommand:
		return this.handleAbortTransaction(message)

	default:
		return nil
	}
}

func (this *TransactionAggregate) handleStoreItem(message StoreItemCommand) error {
	return nil
}
func (this *TransactionAggregate) handleItemStored(message ItemStoredEvent) error {
	return nil
}
func (this *TransactionAggregate) handleItemStoreFailed(message ItemStoreFailedEvent) error {
	return nil
}

func (this *TransactionAggregate) handleDeleteItem(message DeleteItemCommand) error {
	return nil
}
func (this *TransactionAggregate) handleItemDeleted(message ItemDeletedEvent) error {
	return nil
}
func (this *TransactionAggregate) handleItemDeleteFailed(message ItemDeleteFailedEvent) error {
	return nil
}

func (this *TransactionAggregate) handleCommitTransaction(message CommitTransactionCommand) error {
	return nil
}
func (this *TransactionAggregate) handleTransactionCommitted(message TransactionCommittedEvent) error {
	return nil
}

func (this *TransactionAggregate) handleTransactionFailed(message TransactionFailedEvent) error {
	return nil
}

func (this *TransactionAggregate) handleAbortTransaction(message AbortTransactionCommand) error {
	return nil
}

func (this *TransactionAggregate) Replay(message interface{}) {
	// for this project, we won't be event sourcing, but this is here to indicate a minor breakthrough in how events
	// both internal and external--are processed.
	this.apply(message)
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
