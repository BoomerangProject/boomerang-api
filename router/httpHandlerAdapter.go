package router

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

func wrapHandler(httpHandler http.Handler) httprouter.Handle {

	return func(responseWriter http.ResponseWriter, httpRequest *http.Request, params httprouter.Params) {

		context.Set(httpRequest, "params", params)
		httpHandler.ServeHTTP(responseWriter, httpRequest)
	}
}
