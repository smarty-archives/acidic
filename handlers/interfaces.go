package handlers

type WaitGroup interface {
	Add(delta int)
	Done()
}

type CallingContext interface {
	Write(interface{})
	Close()
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
