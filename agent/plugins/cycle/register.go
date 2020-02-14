package cycle

import (
	"time"

	"github.com/imsilence/gocmdb/agent/ens"
	"github.com/imsilence/gocmdb/agent/entity"
	"github.com/imsilence/gocmdb/agent/gconf"
)

type RegisterPlugin struct {
	config   *gconf.Config
	ens      *ens.ENS
	nextTime time.Time
	interval time.Duration
}

func (p *RegisterPlugin) Name() string {
	return "register"
}

func (p *RegisterPlugin) Init(c *gconf.Config, e *ens.ENS) bool {
	p.config = c
	p.ens = e
	p.nextTime = time.Now()
	p.interval = time.Minute * 30
	return true
}

func (p *RegisterPlugin) NextTime() time.Time {
	return p.nextTime
}

func (p *RegisterPlugin) Call() (interface{}, error) {
	p.nextTime = p.nextTime.Add(p.interval)
	return entity.NewRegister(p.config.UUID), nil
}

func (p *RegisterPlugin) Pipeline() chan interface{} {
	return p.ens.Register
}
