package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"http-coala/utils"
)

// Функция для кодирования значений формы в URL-строку
func encodeFormValues(data url.Values) string {
	var encodedValues []string
	for key, values := range data {
		for _, value := range values {
			encodedValues = append(encodedValues, fmt.Sprintf("%s=%s", key, url.QueryEscape(value)))
		}
	}
	return strings.Join(encodedValues, "&")
}

func handler(w http.ResponseWriter, r *http.Request, gumAddr, requestURI string) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		rsp, err := utils.GetCoala(gumAddr, requestURI+"?"+r.URL.RawQuery)

		if err != nil {
			log.Println("Error send coala: ", err)
			return
		}

		w.Write(rsp.Body)
	case "POST":
		// Парсинг параметров формы POST
		err := r.ParseForm()
		if err != nil {
			log.Println("Error parsing form:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Получение списка параметров POST
		postParams := r.PostForm
		_ = postParams

		rsp, err := utils.PostCoala(gumAddr, requestURI+"?"+r.URL.RawQuery, nil)

		if err != nil {
			log.Println("Error send coala: ", err)
			return
		}

		w.Write(rsp.Body)

	case "DELETE":
		rsp, err := utils.DeleteCoala(gumAddr, requestURI+"?"+r.URL.RawQuery, nil)

		if err != nil {
			log.Println("Error send coala: ", err)
			return
		}

		w.Write(rsp.Body)

	// case "POST":
	// 	// Парсинг параметров формы POST
	// 	err := r.ParseForm()
	// 	if err != nil {
	// 		log.Println("Error parsing form:", err)
	// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 		return
	// 	}

	// 	// Получение списка параметров POST
	// 	postParams := r.PostForm

	// 	// Преобразование в JSON
	// 	jsonData := make(map[string]interface{})
	// 	for key, values := range postParams {
	// 		if len(values) == 1 {
	// 			jsonData[key] = values[0]
	// 		} else {
	// 			jsonData[key] = values
	// 		}
	// 	}

	// 	jsonBytes, err := json.Marshal(jsonData)
	// 	if err != nil {
	// 		log.Println("Error encoding to JSON:", err)
	// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 		return
	// 	}

	// 	rsp, err := utils.PostCoala(gumAddr, requestURI+"?"+r.URL.RawQuery, jsonBytes)

	// 	if err != nil {
	// 		log.Println("Error send coala: ", err)
	// 		return
	// 	}

	// 	w.Write(rsp.Body)
	// 	log.Println("Success")

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
