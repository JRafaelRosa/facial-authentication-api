package storage

import (
	"Servidor-Go/internal/model"
	"database/sql"
	"fmt"
)

func StoreLog(db *sql.DB, l model.Log) error {
	query := "INSERT INTO access_log (name, email, position, accuracy, status) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, l.Person.Name, l.Person.Email, l.Person.Position, l.Accuracy, l.Status)

	if err != nil {
		fmt.Println("Erro ao cadastrar log no Banco")
		return err
	}

	fmt.Println("Log salvo no banco com sucesso")
	return nil
}
