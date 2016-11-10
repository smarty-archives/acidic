package web

import (
	"github.com/smartystreets/acidic/contracts"
	"github.com/smartystreets/acidic/web/models"
	"github.com/smartystreets/detour"
)

type Controller struct {
	sender contracts.MessageSender
}

func NewController(sender contracts.MessageSender) *Controller {
	return &Controller{sender: sender}
}

func (this *Controller) Load(input *models.LoadInput) detour.Renderer {
	return this.handle(input.ToMessage())
}

func (this *Controller) Store(input *models.StoreInput) detour.Renderer {
	defer input.Close()
	return this.handle(input.ToMessage())
}
func (this *Controller) Delete(input *models.DeleteInput) detour.Renderer {
	return this.handle(input.ToMessage())
}

func (this *Controller) Commit(input *models.TransactionInput) detour.Renderer {
	return this.handle(input.ToCommitMessage())
}
func (this *Controller) Abort(input *models.TransactionInput) detour.Renderer {
	return this.handle(input.ToAbortMessage())
}

func (this *Controller) handle(message interface{}) detour.Renderer {
	result := this.sender.Send(message)
	return NewApplicationResultRenderer(result)
}
