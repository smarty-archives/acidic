package handlers

type BlockingCallingContext struct {
	Result interface{}
	Waiter WaitGroup
}

func NewBlockingCallingContext(waiter WaitGroup) *BlockingCallingContext {
	waiter.Add(1)
	return &BlockingCallingContext{Waiter: waiter}
}

func (this *BlockingCallingContext) Complete(result interface{}) {
	this.Result = result
	this.Waiter.Done()
}
