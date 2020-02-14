package cycle

import (
	"time"

	"github.com/imsilence/gocmdb/agent/ens"
	"github.com/imsilence/gocmdb/agent/entity"
	"github.com/imsilence/gocmdb/agent/gconf"
)

type HeartbeatPlugin struct {
	config   *gconf.Config
	ens      *ens.ENS
	nextTime time.Time
	interval time.Duration
}

func (p *HeartbeatPlugin) Name() string {
	return "heartbeat"
}

func (p *HeartbeatPlugin) Init(c *gconf.Config, e *ens.ENS) bool {
	p.config = c
	p.ens = e
	p.nextTime = time.Now()
	p.interval = time.Second * 10
	return true
}

func (p *HeartbeatPlugin) NextTime() time.Time {
	return p.nextTime
}

func (p *HeartbeatPlugin) Call() (interface{}, error) {
	p.nextTime = p.nextTime.Add(p.interval)
	return entity.NewHeartbeat(p.config.UUID), nil
}

func (p *HeartbeatPlugin) Pipeline() chan interface{} {
	return p.ens.Heartbeat
}
