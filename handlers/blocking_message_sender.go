package handlers

import "sync"

type BlockingMessageSender struct {
	output chan<- ContextEnvelope
}

func NewBlockingMessageSender(output chan<- ContextEnvelope) *BlockingMessageSender {
	return &BlockingMessageSender{output: output}
}

func (this *BlockingMessageSender) Send(message interface{}) interface{} {
	waiter := &sync.WaitGroup{}
	context := NewBlockingCallingContext(waiter)

	this.output <- ContextEnvelope{
		Message: message,
		Context: context,
	}

	waiter.Wait()

	return context.Result
}
