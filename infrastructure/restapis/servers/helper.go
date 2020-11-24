package servers

import (
	"log"
	"net/http"
)

func renderSuccess(w http.ResponseWriter, data []byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func renderError(w http.ResponseWriter, err error, output []byte) {
	log.Printf("Error: %s\n", err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(output)
}
