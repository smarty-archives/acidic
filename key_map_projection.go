package acidic

import "time"

type KeyMapProjection struct {
	Head      uint64           `json:"head"`
	Tail      uint64           `json:"tail"`
	Committed map[string]*Item `json:"committed"`
	open      map[string]map[string]*Item
}

func NewKeyMapProjection() *KeyMapProjection {
	return &KeyMapProjection{}
}

func (this *KeyMapProjection) Translate(item LoadItemRequest) (LoadItemRequest, error) {
	// TODO
	if len(item.TransactionID) > 0 {
		if items, contains := this.open[item.TransactionID]; !contains {
			return LoadItemRequest{}, TransactionNotFoundError
		} else if item, contains := items[item.Key]; contains {
			return LoadItemRequest{
				Key:      item.Key,
				Revision: item.Revision,
			}, nil
		}
	}

	return LoadItemRequest{}, nil
}

func (this *KeyMapProjection) Apply(message interface{}) {
	switch message := message.(type) {

	case TransactionStartedEvent:
		this.applyTransactionStarted(message)

	case StoringItemEvent:
		this.applyStoringItem(message)
	case ItemStoredEvent:
		this.applyItemStored(message)

	case DeletingItemEvent:
		this.applyDeletingItem(message)
	case ItemDeletedEvent:
		this.applyItemDeleted(message)

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

func (this *KeyMapProjection) applyDeletingItem(message DeletingItemEvent) {
}
func (this *KeyMapProjection) applyItemDeleted(message ItemDeletedEvent) {
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

type Item struct {
	Key        string
	Revision   string
	Expiration time.Time
}
