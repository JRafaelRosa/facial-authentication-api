package handler

import (
	"Servidor-Go/internal/model"
	"Servidor-Go/internal/service"
	"Servidor-Go/internal/storage"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateLogHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var log model.Log

	err := json.NewDecoder(r.Body).Decode(&log)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Erro ao decodar JSON: ", err)
		return
	}

	log.Status, err = service.ProcessLog(log)

	if err != nil {
		_ = storage.StoreLog(db, log)

		w.WriteHeader(http.StatusForbidden)
		http.Error(w, "Falha ao validar pessoa "+err.Error(), http.StatusForbidden)
		return
	}

	err = storage.StoreLog(db, log)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Falha ao Salvar log no banco", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Log registrado! Status: %s. Bem-vindo, %s", log.Status, log.Person.Name)
}
