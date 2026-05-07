package service

import (
	"Servidor-Go/internal/model"
	"errors"
)

const accuracyThreshold = 0.8

func ProcessLog(l model.Log) (string, error) {
	if l.Accuracy < accuracyThreshold {
		return "Refused", errors.New("Erro de Validação")
	}

	return "Accepted", nil
}
