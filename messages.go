package acidic

type LoadItemRequest struct{}
type LoadItemResponse struct{}

type StoreItemCommand struct{}
type TransactionStartedEvent struct{}
type StoringItemEvent struct{}
type ItemStoredEvent struct{}
type ItemStoreFailedEvent struct{}

type DeleteItemCommand struct{}
type DeletingItemEvent struct{}
type ItemDeletedEvent struct{}
type ItemDeleteFailedEvent struct{}

type CommitTransactionCommand struct{}
type TransactionCommittingEvent struct{}
type TransactionCommittedEvent struct{}

type TransactionFailedEvent struct{}

type AbortTransactionCommand struct{}
type TransactionAbortedEvent struct{}
type TransactionAbortFailedEvent struct{}

type MessageContainer interface {
	Add(interface{})
}
