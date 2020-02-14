package persistence

import (
	"github.com/imsilence/gocmdb/agent/gconf"
)

type LogPlugin struct {
}

func (p *LogPlugin) Name() string {
	return "log"
}

func (p *LogPlugin) Init(c gconf.Config) bool {
	return true
}

func (p *LogPlugin) Run() (<-chan interface{}, error) {
	return nil, nil
}

func (p *LogPlugin) Stop() error {
	return nil
}
