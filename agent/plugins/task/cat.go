package task

import (
	"github.com/imsilence/gocmdb/agent/gconf"
)

type CatPlugin struct {
}

func (p *CatPlugin) Name() string {
	return "cat"
}

func (p *CatPlugin) Init(c gconf.Config) bool {
	return true
}

func (p *CatPlugin) Call() (interface{}, error) {
	return "cat", nil
}
