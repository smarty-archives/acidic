package acidic

type MessageContainer interface {
	Add(interface{})
}

type Publisher interface {
	Raise(interface{})
}
