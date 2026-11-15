package api

import (
	"errors"
	"log"
	"net/http"

	"go_final_project/pkg/db"
)

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		writeJSON(w, map[string]string{
			"error": "Не указан идентификатор",
		}, http.StatusBadRequest)
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			writeJSON(w, map[string]string{
				"error": "Задача не найдена",
			}, http.StatusNotFound)
			return
		}

		log.Printf("Ошибка удаления задачи из БД: %v", err)
		writeJSON(w, map[string]string{
			"error": "Внутренняя ошибка сервера",
		}, http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{}, http.StatusOK)
}
