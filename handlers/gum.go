package handlers

import (
	"net/http"
	"strings"
)

func GumHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	gumAddr := parts[2]
	requestURI := "/" + strings.Join(parts[3:], "/")

	handler(w, r, gumAddr, requestURI)
}
