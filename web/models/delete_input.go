package models

import (
	"net/http"
	"strings"

	"github.com/smartystreets/acidic/contracts/messages"
	"github.com/smartystreets/detour"
)

type DeleteInput struct {
	TransactionID   string // optional (start new transaction when not provided)
	ConditionalETag string
	Key             string
}

func (this *DeleteInput) Bind(request *http.Request) {
	this.TransactionID = request.Header.Get(TransactionIDHeader)
	this.ConditionalETag = request.Header.Get(IfNoneMatchHeader)
	this.Key = request.URL.Path
}

func (this *DeleteInput) Sanitize() {
	this.TransactionID = strings.TrimSpace(this.TransactionID)
	this.ConditionalETag = strings.TrimSpace(this.ConditionalETag)
	this.Key = strings.TrimSpace(this.Key)
}

func (this *DeleteInput) Validate() error {
	var errors detour.Errors

	errors.AppendIf(missingRequiredKey, len(this.Key) <= len("/"))

	return errors
}

func (this *DeleteInput) ToMessage(stuff interface{}) messages.DeleteItemCommand {
	// TODO: anything else?

	return messages.DeleteItemCommand{
		TransactionID: this.TransactionID,
		Key:           this.Key,
		ETag:          this.ConditionalETag,
	}
}
