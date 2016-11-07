package kvacid

import "time"

type Index struct {
	// add writers here, e.g. CommitWriter, ApplyWriter, IndexWriter, etc.
	head       uint64
	tail       uint64
	opened     map[string]*OpenTransaction
	closed     map[string]*CommittedItem // map[canonical key]item (across all committed transactions)
	incomplete map[uint64]uint64         // number of items remaining to be completed for each commit sequence
	completed  map[string]*CompletedItem
	history    map[string]struct{} // list of N recent transactions so that Commit() can return success on each call.
}

type OpenTransaction struct {
	items         map[string]*OpenItem
	id            string    // unique, random identifier of the transaction
	pendingWrites uint64    // number of writes in flight across all items within this transaction
	expiration    time.Time // point at which transaction will timeout if nothing happens
	ttl           time.Duration
	status        uint64        // uncommitted, committing, rolled back, concurrency issue
	contexts      []interface{} // parked contexts of callers who requested a commit; can have multiple
}

type OpenItem struct {
	key           string
	revision      string
	status        uint64 // deleted, writing, written
	pendingWrites uint64 // allows for concurrent, outstanding writes to the same key.
}

type CommittedItem struct {
	key            string
	revision       string
	commitSequence uint64
}

type CompletedItem struct {
	key        string
	revision   string
	deleted    bool
	expiration time.Time
}

func (this *Index) Apply(commit interface{}) {
	// absorb the state of the commit and overwrite any committed or pending state
	// and/or invalidate any conflicting uncommitted state
}

type CommitWriter interface {
}

type ApplyWriter interface {
}

type IndexWriter interface {
}
