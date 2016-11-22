package acidic

const (
	txReady = iota
	txWriting
	txCommittingAwaitingWrites
	txCommitting
	txCommitted
	txAborted
	txFailed
)
