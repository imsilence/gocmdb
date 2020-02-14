package cycle

import (
	"time"

	"github.com/imsilence/gocmdb/agent/ens"
	"github.com/imsilence/gocmdb/agent/gconf"
	"github.com/imsilence/gocmdb/agent/entity"
)

type ResourcePlugin struct {
	config   *gconf.Config
	ens      *ens.ENS
	nextTime time.Time
	interval time.Duration
}

func (p *ResourcePlugin) Name() string {
	return "resource"
}

func (p *ResourcePlugin) Init(c *gconf.Config, e *ens.ENS) bool {
	p.config = c
	p.ens = e
	p.nextTime = time.Now()
	p.interval = time.Second * 5
	return true
}

func (p *ResourcePlugin) NextTime() time.Time {
	return p.nextTime
}

func (p *ResourcePlugin) Call() (interface{}, error) {
	p.nextTime = p.nextTime.Add(p.interval)
	return entity.NewLog(p.config.UUID, entity.RESOURCE, entity.NewResource(p.config.UUID)), nil
}

func (p *ResourcePlugin) Pipeline() chan interface{} {
	return p.ens.Log
}
