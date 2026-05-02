package model

import (
	"time"
)

type Log struct {
	Person     Person    `json:"Person"`
	Timestamps time.Time `json:"timestamps"`
	Status     string    `json:"status"`
	Accuracy   float64   `json:"accuracy"`
}
