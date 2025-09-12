package utils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse[K comparable, V any](w http.ResponseWriter, status int, response map[K]V) {
	resp, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unexpected internal server error. try again later"))
		return
	}

	w.WriteHeader(status)
	w.Write(resp)
}
