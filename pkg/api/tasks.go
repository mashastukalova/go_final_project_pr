package api

import (
	"log"
	"net/http"

	"go_final_project/pkg/db"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	tasks, err := db.Tasks(50)
	if err != nil {
		log.Printf("Ошибка получения списка задач из БД: %v", err)
		writeJSON(w, map[string]string{
			"error": "Внутренняя ошибка сервера",
		}, http.StatusInternalServerError)
		return
	}

	writeJSON(w, TasksResp{
		Tasks: tasks,
	}, http.StatusOK)
}
