package acidic

type MessageContainer interface {
	Add(interface{})
}

type Dispatcher interface {
	Raise(interface{})
}
