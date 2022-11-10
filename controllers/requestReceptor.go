package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RequestReceptor(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)
	if body != nil {
		json.NewEncoder(w).Encode(body)
	}
	// headers, _ := ioutil.ReadAll(r.Header)
	// if headers != nil {
	// 	utils.SetResponseError(w, r, "could not read body")
	// 	return
	// }
}
