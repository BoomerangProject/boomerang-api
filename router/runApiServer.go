package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/BoomerangProject/boomerang-api/handlers"
	"github.com/gorilla/context"
)
func wrapHandler(httpHandler http.Handler) httprouter.Handle {

	return func(responseWriter http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {

		context.Set(httpRequest, "params", params)
		httpHandler.ServeHTTP(responseWriter, httpRequest)
	}
}

func Run() {

	router := httprouter.New()
	router.POST("/registerAsBusiness", wrapHandler(http.HandlerFunc(handlers.RegisterAsBusiness)))
	http.ListenAndServe(":8080", router)
}
