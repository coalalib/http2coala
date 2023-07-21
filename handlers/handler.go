package handlers

import (
	"log"
	"net/http"
	"strings"

	"http-coala/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 1 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	gumAddr := parts[1]
	requestURI := "/" + strings.Join(parts[2:], "/")

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		rsp, err := utils.GetCoala(gumAddr, requestURI+"?"+r.URL.RawQuery)

		if err != nil {
			log.Println("Error send coala: ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if int(rsp.Code) > 128 {
			http.Error(w, string(rsp.Body), utils.CodeHandler(rsp.Code))
			// http.Error(w, rsp.Code.String(), utils.CodeHandler(rsp.Code))
		} else {
			w.Write(rsp.Body)
		}

	case "POST":
		data, err := utils.GetData(r)

		if err != nil {
			log.Println("Error get data:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		rsp, err := utils.PostCoala(gumAddr, requestURI+"?"+r.URL.RawQuery, data)

		if err != nil {
			log.Println("Error send coala: ", err)
			return
		}

		if int(rsp.Code) > 128 {
			http.Error(w, string(rsp.Body), utils.CodeHandler(rsp.Code))
		} else {
			w.Write(rsp.Body)
		}

	case "DELETE":
		data, err := utils.GetData(r)

		if err != nil {
			log.Println("Error get data:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		rsp, err := utils.DeleteCoala(gumAddr, requestURI+"?"+r.URL.RawQuery, data)

		if err != nil {
			log.Println("Error send coala: ", err)
			return
		}

		if int(rsp.Code) > 128 {
			http.Error(w, string(rsp.Body), utils.CodeHandler(rsp.Code))
		} else {
			w.Write(rsp.Body)
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
