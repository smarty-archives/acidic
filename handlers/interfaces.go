package handlers

type WaitGroup interface {
	Add(delta int)
	Done()
}

type CallingContext interface {
	Complete(interface{})
}

type ContextEnvelope struct {
	Message interface{}
	Context CallingContext
}

type ApplicationHandler interface {
	Handle(interface{}) interface{}
}

type CorrelatedMessage interface {
	CorrelationID() string
}
