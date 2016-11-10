package web

import (
	"net/http"

	"github.com/smartystreets/detour"
)

type ErrorRenderer struct {
	err error
}

func NewErrorRenderer(err error) detour.Renderer {
	return ErrorRenderer{err: err}
}

func (this ErrorRenderer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO: based upon the type of err, render different errors to the response
}
