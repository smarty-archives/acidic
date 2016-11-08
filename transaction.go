package acidic

import "time"

type Transaction struct {
	raised   MessageContainer
	sequence uint64
	started  time.Time
	updated  time.Time
	ttl      time.Duration
}

func NewTransaction(raised MessageContainer, started time.Time, ttl time.Duration) *Transaction {
	return &Transaction{
		raised:   raised,
		sequence: 0,
		started:  started,
		updated:  started,
		ttl:      ttl,
	}
}
