package handlers

import "github.com/smartystreets/acidic/contracts"

type CorrelationHandler struct {
	application contracts.ApplicationHandler
	parked      map[string][]contracts.CallingContext
}

func NewCorrelationHandler(application contracts.ApplicationHandler) *CorrelationHandler {
	return &CorrelationHandler{
		application: application,
		parked:      make(map[string][]contracts.CallingContext),
	}
}

func (this *CorrelationHandler) Handle(message interface{}) {
	switch message := message.(type) {
	case contracts.ContextEnvelope:
		this.handleContext(message)
	default:
		this.handle(message)
	}
}

func (this *CorrelationHandler) handleContext(envelope contracts.ContextEnvelope) {
	result := this.application.Handle(envelope.Message)
	correlationID := extractCorrelationID(envelope.Message)

	if len(correlationID) > 0 && result == nil {
		this.park(correlationID, envelope.Context)
	} else {
		writeResult(envelope.Context, result)
	}
}
func (this *CorrelationHandler) handle(message interface{}) {
	this.application.Handle(message)

	if correlationID := extractCorrelationID(message); len(correlationID) > 0 {
		this.release(correlationID, message)
	}
}

func extractCorrelationID(message interface{}) string {
	if correlated, ok := message.(contracts.CorrelatedMessage); ok {
		return correlated.CorrelationID()
	}

	return ""
}
func writeResult(context contracts.CallingContext, result interface{}) {
	context.Write(result)
	context.Close()
}

func (this *CorrelationHandler) park(id string, context contracts.CallingContext) {
	items := this.parked[id]
	items = append(items, context)
	this.parked[id] = items
}
func (this *CorrelationHandler) release(id string, message interface{}) {
	for _, item := range this.parked[id] {
		writeResult(item, message)
	}
}
