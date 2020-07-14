package handlers

import (
	"encoding/json"
	"net/http"
)

const (
	ERR_MSG = "ERROR_MESSAGE"
	MSG     = "MESSAGE"
)

func jsonifyMessage(msg string, msgType string, httpCode int) ([]byte, int) {
	var data []byte
	var Obj struct {
		Status   string `json:"status"`
		HTTPCode int    `json:"httpCode"`
		Message  string `json:"message"`
	}
	Obj.Message = msg
	Obj.HTTPCode = httpCode
	switch msgType {
	case ERR_MSG:
		Obj.Status = "FAILED"

	case MSG:
		Obj.Status = "SUCCESS"
	}
	data, _ = json.Marshal(Obj)
	return data, httpCode
}

func writeJSONMessage(msg string, msgType string, httpCode int, w http.ResponseWriter) {
	d, code := jsonifyMessage(msg, msgType, httpCode)
	writeJSONResponse(d, code, w)
}

func writeJSONStruct(v interface{}, code int, w http.ResponseWriter) {
	d, err := json.Marshal(v)
	if err != nil {
		writeJSONMessage("Unable to marshal data. Err: "+err.Error(), ERR_MSG, http.StatusInternalServerError, w)
		return
	}
	writeJSONResponse(d, code, w)
}

func writeJSONResponse(d []byte, code int, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(d)
}
