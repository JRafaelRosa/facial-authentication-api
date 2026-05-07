package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Log struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Position string  `json:"position"`
	Accuracy float64 `json:"accuracy"`
	Status   string  `json:"status"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func ListLogs(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := "SELECT name, email, position, accuracy, status FROM access_log"

	statusFilter := r.URL.Query().Get("status")

	var rows *sql.Rows
	var err error

	if statusFilter != "" {
		query += " WHERE status = ?"
		rows, err = db.Query(query, statusFilter)
	} else {
		rows, err = db.Query(query)
	}

	if err != nil {
		http.Error(w, "Erro ao consultar banco", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var logs []Log

	for rows.Next() {
		var l Log

		err := rows.Scan(&l.Name, &l.Email, &l.Position, &l.Accuracy, &l.Status)
		if err != nil {
			http.Error(w, "Erro ao ler dados", http.StatusInternalServerError)
			return
		}

		logs = append(logs, l)
	}

	WriteJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    logs,
	})
}

func WriteJSON(w http.ResponseWriter, status int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
