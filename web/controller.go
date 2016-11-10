package web

import (
	"github.com/smartystreets/acidic/contracts/messages"
	"github.com/smartystreets/acidic/web/models"
	"github.com/smartystreets/detour"
)

type Controller struct {
	sender MessageSender
}

// TODO: protect sensitive files, e.g. /tx/... from load/store/delete
func NewController(sender MessageSender) *Controller {
	return &Controller{sender: sender}
}

func (this *Controller) Load(input *models.LoadInput) detour.Renderer {
	if _, err := this.sender.Send(messages.LoadItemRequest{}); err != nil {
		return translateError(err)
	} else {
		return nil // custom content result which lets us return a binary stream and custom headers
	}
}

func (this *Controller) Store(input *models.StoreInput) detour.Renderer {
	input.Close()
	return nil
}

func (this *Controller) Delete(input *models.DeleteInput) detour.Renderer {
	return nil
}

// POST /tx/id
func (this *Controller) Commit(input *models.TransactionInput) detour.Renderer {
	message := messages.CommitTransactionCommand{TransactionID: input.TransactionID}
	if _, err := this.sender.Send(message); err != nil {
		return translateError(err)
	}

	return nil // TODO: what kind of success do we return?
}

// DELETE /tx/id
func (this *Controller) Abort(input *models.TransactionInput) detour.Renderer {
	message := messages.AbortTransactionCommand{TransactionID: input.TransactionID}
	if _, err := this.sender.Send(message); err != nil {
		return translateError(err)
	}

	return nil // TODO: what kind of success do we return?
}

func translateError(error) detour.Renderer {
	return nil // TODO
}
