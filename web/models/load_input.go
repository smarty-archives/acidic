package models

import (
	"net/http"
	"strings"

	"github.com/smartystreets/acidic/contracts/messages"
	"github.com/smartystreets/detour"
)

type LoadInput struct {
	TransactionID   string // optional (load latest committed version regardless of transaction)
	ConditionalETag string
	Key             string
}

func (this *LoadInput) Bind(request *http.Request) {
	this.TransactionID = request.Header.Get(TransactionIDHeader)
	this.ConditionalETag = request.Header.Get(IfNoneMatchHeader)
	this.Key = request.URL.Path
}

func (this *LoadInput) Sanitize() {
	this.TransactionID = strings.TrimSpace(this.TransactionID)
	this.ConditionalETag = strings.TrimSpace(this.ConditionalETag)
	this.Key = strings.TrimSpace(this.Key)
}

func (this *LoadInput) Validate() error {
	var errors detour.Errors

	errors.AppendIf(missingRequiredKey, len(this.Key) <= len("/"))

	return errors
}

func (this *LoadInput) ToMessage() messages.LoadItemRequest {
	return messages.LoadItemRequest{
		TransactionID: this.TransactionID,
		Key:           this.Key,
		ETag:          this.ConditionalETag,
	}
}
