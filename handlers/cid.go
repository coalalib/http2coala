package handlers

import (
	"log"
	"net/http"
	"strings"

	"http-coala/utils"
)

func CidHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	cid := parts[2]
	gumAddr, err := utils.GetGum(cid)
	if err != nil {
		log.Println("Get gum error: ", err)
		return
	}

	requestURI := "/" + strings.Join(parts[3:], "/")
	handler(w, r, gumAddr, requestURI)
}
