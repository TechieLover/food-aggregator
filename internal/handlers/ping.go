package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func setPingRoutes(router *httprouter.Router) {
	router.GET("/ping", Ping)
}

func Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	writeJSONMessage("pong", MSG, http.StatusOK, w)
}
