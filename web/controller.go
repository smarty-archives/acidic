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
		return NewErrorRenderer(err)
	} else {
		return nil // TODO: custom content result which lets us return a binary stream and custom headers
	}
}

func (this *Controller) Store(input *models.StoreInput) detour.Renderer {
	defer input.Close()

	if _, err := this.sender.Send(input.ToMessage()); err != nil {
		return NewErrorRenderer(err)
	}

	return nil // TODO
}

func (this *Controller) Delete(input *models.DeleteInput) detour.Renderer {
	if _, err := this.sender.Send(input.ToMessage(nil)); err != nil {
		return NewErrorRenderer(err)
	}

	return nil // TODO
}

// POST /tx/id
func (this *Controller) Commit(input *models.TransactionInput) detour.Renderer {
	if _, err := this.sender.Send(input.ToCommitMessage()); err != nil {
		return NewErrorRenderer(err)
	}

	return nil // TODO
}

// DELETE /tx/id
func (this *Controller) Abort(input *models.TransactionInput) detour.Renderer {
	if _, err := this.sender.Send(input.ToAbortMessage()); err != nil {
		return NewErrorRenderer(err)
	}

	return nil // TODO
}
