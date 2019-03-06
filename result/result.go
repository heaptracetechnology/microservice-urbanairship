package result

import (
	"encoding/json"
	"net/http"
)

func WriteErrorResponse(responseWriter http.ResponseWriter, err error) {
	messageBytes, _ := json.Marshal(err)
	WriteJsonResponse(responseWriter, messageBytes, http.StatusBadRequest)
	return
}

func WriteJsonResponse(responseWriter http.ResponseWriter, bytes []byte, statusCode int) {
	responseWriter.WriteHeader(statusCode)
	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	responseWriter.Write(bytes)
}
