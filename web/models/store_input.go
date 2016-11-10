package models

import (
	"io"
	"net/http"
	"strings"

	"github.com/smartystreets/acidic/contracts/messages"
	"github.com/smartystreets/detour"
)

type StoreInput struct {
	TransactionID   string // optional (start new transaction when not provided)
	ConditionalETag string
	Key             string
	Metadata        map[string]string
	Payload         io.ReadCloser
}

func (this *StoreInput) Bind(request *http.Request) {
	this.TransactionID = request.Header.Get(TransactionIDHeader)
	this.ConditionalETag = request.Header.Get(IfNoneMatchHeader)
	this.Key = request.URL.Path
	this.Metadata = make(map[string]string, len(request.Header)) // TODO: copy all values from keys
	this.Payload = request.Body                                  // TODO: ensure body is read/closed fully
}

func (this *StoreInput) Sanitize() {
	this.TransactionID = strings.TrimSpace(this.TransactionID)
	this.ConditionalETag = strings.TrimSpace(this.ConditionalETag)
	this.Key = strings.TrimSpace(this.Key)
}

func (this *StoreInput) Validate() error {
	var errors detour.Errors

	errors.AppendIf(missingRequiredKey, len(this.Key) <= len("/"))

	return errors
}

func (this *StoreInput) Close() error {
	if this.Payload != nil {
		return this.Payload.Close()
	}

	return nil
}

func (this *StoreInput) ToMessage() messages.StoreItemCommand {
	return messages.StoreItemCommand{
		TransactionID: this.TransactionID,
		Key:           this.Key,
		ETag:          this.ConditionalETag,
		Metadata:      this.Metadata,
		Payload:       this.Payload,
	}
}
