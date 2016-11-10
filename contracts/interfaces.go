package contracts

type MessageSender interface {
	Send(interface{}) interface{}
}

type ContextEnvelope struct {
	Message interface{}
	Context CallingContext
}

type CallingContext interface {
	Write(interface{})
	Close()
}

type WaitGroup interface {
	Add(delta int)
	Done()
}

type ApplicationHandler interface {
	Handle(interface{}) interface{}
}

type CorrelatedMessage interface {
	CorrelationID() string
}
