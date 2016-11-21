package acidic

const (
	txReady = iota
	txWriting
	txWritingCommitting
	txCommitting
	txCommitted
	txAborted
	txFailed
)
