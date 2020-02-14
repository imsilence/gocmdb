package entity

import (
	"time"
	"encoding/json"
)

const (
	RESOURCE = 0X0001
)

type Log struct {
	UUID    string    `json:"agent"`
	Type    int       `json:"type"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func NewLog(uuid string, typ int, msg interface{}) Log {
	bytes, _ := json.Marshal(msg)
	return Log{
		UUID:    uuid,
		Type:    typ,
		Message: string(bytes),
		Time:    time.Now(),
	}
}
