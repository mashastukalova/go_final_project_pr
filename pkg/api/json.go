package api

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, data any, statusCode int) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "ошибка формирования JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	_, _ = w.Write(append(response, '\n'))
}
