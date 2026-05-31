package utils

import (
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	http.Error(w, message, code)
	log.Printf("HTTP: %d: %s", code, message)
}
