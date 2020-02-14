package entity

import "time"

type Task struct {
}

type TaskResult struct {
	UUID   int         `json:"uuid"`
	Type   int         `json:"type"`
	Result interface{} `json:"result"`
	Time   time.Time   `json:"time"`
}
