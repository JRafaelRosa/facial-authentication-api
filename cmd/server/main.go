package main

import (
	"Servidor-Go/internal/handler"
	"Servidor-Go/internal/storage"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const port = ":8080"

func main() {

	db, err := sql.Open("mysql", "root:23014826@tcp(127.0.0.1:3306)/servidor_logs")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	defer db.Close()

	//cria o servidor
	router := http.NewServeMux()

	router.HandleFunc("/registrar", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateLogHandler(w, r, db)
	})

	router.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		storage.ListLogs(w, r, db)
	})

	router.HandleFunc("/logs/aceitos", func(w http.ResponseWriter, r *http.Request) {
		storage.ListAccept(w, r, db)
	})

	router.HandleFunc("/logs/recusados", func(w http.ResponseWriter, r *http.Request) {
		storage.ListRefuse(w, r, db)
	})

	fmt.Println("Listening on port " + port)
	err = http.ListenAndServe(port, router)

	if err != nil {
		fmt.Println("erro ao iniciar servidor", err)
	}

}
