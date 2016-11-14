package acidic

type MessageContainer interface {
	Add(interface{})
}

type AggregateRoot interface {
	Raise(interface{})
}
