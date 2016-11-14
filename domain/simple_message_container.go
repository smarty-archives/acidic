package acidic

type SimpleMessageContainer struct {
	messages []interface{}
}

func (this *SimpleMessageContainer) Add(message interface{}) {
	this.messages = append(this.messages, message)
}
