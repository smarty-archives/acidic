package acidic

import "sync"

type MessageContext struct {
	waiter  *sync.WaitGroup
	message interface{}
	result  interface{}
	error   error
}

func NewMessageContext(message interface{}) *MessageContext {
	waiter := &sync.WaitGroup{}
	waiter.Add(1)

	return &MessageContext{
		waiter:  waiter,
		message: message,
	}
}

func (this *MessageContext) Wait() {
	this.waiter.Wait()
}

func (this *MessageContext) Complete(result interface{}, err error) {
	this.result = result
	this.error = err
	this.waiter.Done()
}
