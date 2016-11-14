package models

import "github.com/smartystreets/detour"

const (
	TransactionIDHeader = "X-Transaction-Id"
	IfNoneMatchHeader   = "If-None-Match"
)

var (
	missingRequiredTransactionID = detour.SimpleInputError("Missing transaction.", TransactionIDHeader)
	missingRequiredKey           = detour.SimpleInputError("Missing key.", TransactionIDHeader)
)
