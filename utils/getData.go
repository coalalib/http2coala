package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetData(r *http.Request) (data []byte, err error) {
	contentType := r.Header.Get("Content-Type")

	if strings.HasPrefix(contentType, "multipart/form-data") {
		err = r.ParseMultipartForm(10 << 20)
		if err != nil {
			return
		}

		formData := make(map[string]interface{})
		for key, values := range r.MultipartForm.Value {
			if len(values) == 1 {
				formData[key] = values[0]
			} else {
				formData[key] = values
			}
		}

		return json.Marshal(formData)
	} else {
		data, err = ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		defer r.Body.Close()

		return
	}
}
