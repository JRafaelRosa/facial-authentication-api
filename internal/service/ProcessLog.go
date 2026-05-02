package service

import (
	"Servidor-Go/internal/model"
	"errors"
	"fmt"
)

func ProcessLog(l model.Log) (string, error) {
	if l.Accuracy < 0.8 {
		fmt.Print("Acesso negado pessoa desconhecida")
		return "Refused", errors.New("Pessoa desconhecida")
	}

	fmt.Println("Bem vindo: " + l.Person.Name)
	return "Accepted", nil

}
