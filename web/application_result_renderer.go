package web

import (
	"net/http"

	"github.com/smartystreets/detour"
)

type ApplicationResultRenderer struct {
	result interface{}
}

func NewApplicationResultRenderer(result interface{}) detour.Renderer {
	return ApplicationResultRenderer{result: result}
}

func (this ApplicationResultRenderer) Render(response http.ResponseWriter, request *http.Request) {
	// TODO: based upon the type of result, render different errors to the response
	// TODO: custom content result which lets us return a binary stream and custom headers
}
