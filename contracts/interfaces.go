package contracts

type MessageSender interface {
	Send(interface{}) interface{}
}
