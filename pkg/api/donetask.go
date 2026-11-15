package api

import (
	"errors"
	"log"
	"net/http"
	"time"

	"go_final_project/pkg/db"
	"go_final_project/pkg/nextdate"
)

func DoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, map[string]string{
			"error": "Метод не поддерживается",
		}, http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		writeJSON(w, map[string]string{
			"error": "Не указан идентификатор",
		}, http.StatusBadRequest)
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			writeJSON(w, map[string]string{
				"error": "Задача не найдена",
			}, http.StatusNotFound)
			return
		}

		log.Printf("Ошибка получения задачи из БД: %v", err)
		writeJSON(w, map[string]string{
			"error": "Внутренняя ошибка сервера",
		}, http.StatusInternalServerError)
		return
	}

	if task.Repeat == "" {
		err = db.DeleteTask(id)
	} else {
		var next string
		next, err = nextdate.NextDate(time.Now(), task.Date, task.Repeat)
		if err == nil {
			err = db.UpdateDate(next, id)
		}
	}

	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			writeJSON(w, map[string]string{
				"error": "Задача не найдена",
			}, http.StatusNotFound)
			return
		}

		log.Printf("Ошибка выполнения задачи: %v", err)
		writeJSON(w, map[string]string{
			"error": "Внутренняя ошибка сервера",
		}, http.StatusInternalServerError)
		return
	}

	writeJSON(w, map[string]string{}, http.StatusOK)
}
