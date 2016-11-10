package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/smartystreets/detour"
)

func RegisterRoutes(router httprouter.Router, controller *Controller) {
	load := detour.New(controller.Load)
	store := detour.New(controller.Store)
	delete := detour.New(controller.Delete)
	commit := detour.New(controller.Commit)
	abort := detour.New(controller.Abort)

	router.GET("/", func(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		load.ServeHTTP(response, request)
	})
	router.PUT("/", func(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		store.ServeHTTP(response, request)
	})
	router.DELETE("/", func(response http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		delete.ServeHTTP(response, request)
	})
	router.POST("/tx/:tx", func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
		request.Header.Set("X-Transaction-Id", params.ByName("tx"))
		commit.ServeHTTP(response, request)
	})
	router.DELETE("/tx/:tx", func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
		request.Header.Set("X-Transaction-Id", params.ByName("tx"))
		abort.ServeHTTP(response, request)
	})
}
