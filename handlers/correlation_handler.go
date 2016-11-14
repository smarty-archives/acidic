package handlers

type CorrelationHandler struct {
	application ApplicationHandler
	parked      map[string][]CallingContext
}

// NOTE TO SELF FOR TDD:

// The intention of this handler is to facilitate "parking" a request while awaiting a future event.
// A simple example of this might be a CommitTransactionCommand being sent. It takes time for the commit to be durably
// written to storage. Once we get a TransactionCommittedEvent or TransactionCommitFailed message, we can release
// this context furthermore, it might be that there are multiple Commit() instructions for a given transaction
// and a single TransactionCommittedEvent (etc.) can release all parked contexts awaiting that result.
// To further the example, once a given transaction is in a failed state, additional instructions such as Commit()
// will return a failure error back to the caller.

func NewCorrelationHandler(application ApplicationHandler) *CorrelationHandler {
	return &CorrelationHandler{
		application: application,
		parked:      make(map[string][]CallingContext),
	}
}

func (this *CorrelationHandler) Handle(message interface{}) {
	switch message := message.(type) {
	case ContextEnvelope:
		this.handleContext(message)
	default:
		this.handle(message)
	}
}

func (this *CorrelationHandler) handleContext(envelope ContextEnvelope) {
	result := this.application.Handle(envelope.Message) // typically command messages which may result in an error
	correlationID := extractCorrelationID(envelope.Message)

	if len(correlationID) > 0 && result == nil {
		this.park(correlationID, envelope.Context)
	} else {
		writeResult(envelope.Context, result)
	}
}
func (this *CorrelationHandler) handle(message interface{}) {
	this.application.Handle(message) // typically event messages coming in from backend writers (app shouldn't return a result)

	if correlationID := extractCorrelationID(message); len(correlationID) > 0 {
		this.release(correlationID, message)
	}
}

func (this *CorrelationHandler) park(id string, context CallingContext) {
	items := this.parked[id]
	items = append(items, context)
	this.parked[id] = items
}
func (this *CorrelationHandler) release(id string, message interface{}) {
	for _, item := range this.parked[id] {
		writeResult(item, message)
	}
}

func extractCorrelationID(message interface{}) string {
	if correlated, ok := message.(CorrelatedMessage); ok {
		return correlated.CorrelationID()
	}

	return ""
}
func writeResult(context CallingContext, result interface{}) {
	context.Write(result)
	context.Close()
}
