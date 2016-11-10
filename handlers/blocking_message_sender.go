package handlers

import (
	"sync"

	"github.com/smartystreets/acidic/contracts"
)

type BlockingMessageSender struct {
	output chan<- contracts.ContextEnvelope
}

func NewBlockingMessageSender(output chan<- contracts.ContextEnvelope) *BlockingMessageSender {
	return &BlockingMessageSender{output: output}
}

func (this *BlockingMessageSender) Send(message interface{}) interface{} {
	waiter := &sync.WaitGroup{}
	context := NewBlockingCallingContext(waiter)

	this.output <- contracts.ContextEnvelope{
		Message: message,
		Context: context,
	}

	waiter.Wait()

	return context.Written()
}
