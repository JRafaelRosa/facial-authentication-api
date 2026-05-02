package storage

import (
	"database/sql"
	"fmt"
	"net/http"
)

func ListLogs(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	query := "SELECT name, email, position, accuracy, status FROM access_log"
	executeSimpleList(w, db, query)
}

func ListAccept(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	query := "SELECT name, email, position, accuracy, status FROM access_log WHERE status = 'Accepted'"
	executeSimpleList(w, db, query)
}

func ListRefuse(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	query := "SELECT name, email, position, accuracy, status FROM access_log WHERE status = 'Refused'"
	executeSimpleList(w, db, query)
}

func executeSimpleList(w http.ResponseWriter, db *sql.DB, query string) {
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Erro ao consultar banco", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "--- LISTA DE ACESSOS ---")

	for rows.Next() {
		var name, email, position, status string
		var accuracy float64
		if err := rows.Scan(&name, &email, &position, &accuracy, &status); err == nil {
			fmt.Fprintf(w, "[%s] Nome: %s | Cargo: %s | Precisão: %.2f\n", status, name, position, accuracy)
		}
	}
}
