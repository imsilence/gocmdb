package cycle

import (
	"github.com/imsilence/gocmdb/agent/plugins"
	"github.com/imsilence/gocmdb/agent/plugins/cycle"
)

func init() {
	plugins.DefaultManager.RegisterCycle(&cycle.HeartbeatPlugin{})
	plugins.DefaultManager.RegisterCycle(&cycle.RegisterPlugin{})
	plugins.DefaultManager.RegisterCycle(&cycle.ResourcePlugin{})
}
