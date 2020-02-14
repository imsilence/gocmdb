package gconf

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Config struct {
	Endpoint string
	Token    string

	PID     int
	PidFile string

	UUID     string
	UuidFile string
}

func NewConfig() *Config {
	UuidFile := "agentd.uuid"
	PidFile := "agentd.pid"
	UUID := ""
	if cxt, err := ioutil.ReadFile(UuidFile); err == nil {
		UUID = string(cxt)
	} else if os.IsNotExist(err) {
		UUID = strings.ReplaceAll(uuid.New().String(), "-", "")
		ioutil.WriteFile(UuidFile, []byte(UUID), os.ModePerm)
	}

	pid := os.Getpid()
	ioutil.WriteFile(PidFile, []byte(strconv.Itoa(pid)), os.ModePerm)

	return &Config{
		Endpoint: "http://localhost:8080/v2/api",
		Token:    "0e3b482aa4998ee86945f0cc868882df",
		PID:      pid,
		PidFile:  PidFile,
		UUID:     UUID,
		UuidFile: UuidFile,
	}
}
