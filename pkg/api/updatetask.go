package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"go_final_project/pkg/db"
)

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Некорректный JSON"}, http.StatusBadRequest)
		return
	}

	if task.ID == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор"}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeJSON(w, map[string]string{"error": "Не указан заголовок задачи"}, http.StatusBadRequest)
		return
	}

	task.Date, err = checkDate(task.Date, task.Repeat)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			writeJSON(w, map[string]string{"error": "Задача не найдена"}, http.StatusNotFound)
			return
		}

		log.Printf("Ошибка обновления задачи в БД: %v", err)
		writeJSON(w, map[string]string{"error": "Внутренняя ошибка сервера"}, http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{}, http.StatusOK)
}
