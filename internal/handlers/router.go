package handlers

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/julienschmidt/httprouter"
)

func GetRouter() http.Handler {
	router := httprouter.New()
	router.PanicHandler = PanicHandler
	setPingRoutes(router)
	setFoodAggregatorRoutes(router)
	return router
}

func PanicHandler(w http.ResponseWriter, r *http.Request, c interface{}) {
	log.Printf("Recovering from panic, Reason: %+v", c.(error))
	debug.PrintStack()
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(c.(error).Error()))
}
