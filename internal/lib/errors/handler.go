package errors

import (
	"log"
	"net/http"
)

func WriteHttp(w http.ResponseWriter, e error) {
	switch e.(type) {
	case Error:
		log.Printf("%+v\n", e)
		err := e.(Error)
		http.Error(w, err.Error(), err.Code())
	default:
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}
}
