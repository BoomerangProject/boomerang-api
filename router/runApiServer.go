package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func something() {

	router := httprouter.New()
	router.GET("/", wrapHandler(http.HandlerFunc(Index)))
	http.ListenAndServe(":8080", router)
}

func Index(responseWriter http.ResponseWriter, httpRequest *http.Request) {
	fmt.Println("you did it")
}
