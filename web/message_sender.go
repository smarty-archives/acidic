package web

type MessageSender interface {
	Send(interface{}) (interface{}, error)
}

type BlockingMessageSender struct {
}

func (this *BlockingMessageSender) Send(message interface{}) (interface{}, error) {
	return nil, nil
}
