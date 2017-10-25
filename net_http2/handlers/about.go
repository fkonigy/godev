package handlers

import (
	"fmt"
	"net/http"
)

type AboutHandler struct{}

func (h AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About this app.")
}
