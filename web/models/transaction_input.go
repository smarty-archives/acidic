package models

import (
	"net/http"
	"strings"

	"github.com/smartystreets/acidic/contracts/messages"
	"github.com/smartystreets/detour"
)

type TransactionInput struct {
	TransactionID string
}

func (this *TransactionInput) Bind(request *http.Request) {
	this.TransactionID = request.Header.Get(TransactionIDHeader)
}

func (this *TransactionInput) Sanitize() {
	this.TransactionID = strings.TrimSpace(this.TransactionID)
}

func (this *TransactionInput) Validate() error {
	var errors detour.Errors

	errors.AppendIf(missingRequiredTransactionID, len(this.TransactionID) == 0)

	return errors
}

func (this *TransactionInput) ToCommitMessage() messages.CommitTransactionCommand {
	return messages.CommitTransactionCommand{TransactionID: this.TransactionID}
}

func (this *TransactionInput) ToAbortMessage() messages.AbortTransactionCommand {
	return messages.AbortTransactionCommand{TransactionID: this.TransactionID}
}
