package plugins

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/imsilence/gocmdb/agent/ens"
	"github.com/imsilence/gocmdb/agent/gconf"
)

type CyclePlugin interface {
	Name() string
	Init(*gconf.Config, *ens.ENS) bool
	NextTime() time.Time
	Call() (interface{}, error)
	Pipeline() chan interface{}
}

type PersistencePlugin interface {
	Name() string
	Init(*gconf.Config, *ens.ENS) bool
	Run() (<-chan interface{}, error)
	Stop() error
	Pipeline() chan interface{}
}

type TaskPlugin interface {
	Name() string
	Init(*gconf.Config, *ens.ENS) bool
	Call() (interface{}, error)
	Pipeline() chan interface{}
}

type Manager struct {
	config      *gconf.Config
	ens         *ens.ENS
	Cycles      map[string]CyclePlugin
	Persistencs map[string]PersistencePlugin
	Tasks       map[string]TaskPlugin
}

func NewManager() *Manager {
	return &Manager{
		Cycles:      map[string]CyclePlugin{},
		Persistencs: map[string]PersistencePlugin{},
		Tasks:       map[string]TaskPlugin{},
	}
}

func (m *Manager) RegisterCycle(p CyclePlugin) error {
	name := p.Name()
	if _, ok := m.Cycles[name]; ok {
		return fmt.Errorf("plugin %s is exists", name)
	}
	m.Cycles[name] = p
	return nil
}

func (m *Manager) RegisterPersistence(p PersistencePlugin) error {
	name := p.Name()
	if _, ok := m.Persistencs[name]; ok {
		return fmt.Errorf("plugin %s is exists", name)
	}
	m.Persistencs[name] = p
	return nil
}

func (m *Manager) RegisterTask(p TaskPlugin) error {
	name := p.Name()
	if _, ok := m.Tasks[name]; ok {
		return fmt.Errorf("plugin %s is exists", name)
	}
	m.Tasks[name] = p
	return nil
}

func (m *Manager) Init(c *gconf.Config, e *ens.ENS) {
	m.config = c
	m.ens = e
	for _, plugin := range m.Cycles {
		plugin.Init(c, e)
	}
	for _, plugin := range m.Persistencs {
		plugin.Init(c, e)
	}
	for _, plugin := range m.Tasks {
		plugin.Init(c, e)
	}
}

func (m *Manager) Start() {
	go m.startCycle()
}

func (m *Manager) startCycle() {
	logrus.Debug("start cycle")
	for now := range time.Tick(time.Second) {
		for _, plugin := range m.Cycles {
			if now.Before(plugin.NextTime()) {
				continue
			}
			logrus.WithFields(logrus.Fields{
				"plugin": plugin.Name(),
			}).Debug("call plugin")
			if evt, err := plugin.Call(); err == nil {
				plugin.Pipeline() <- evt
			}

		}
	}
}

var DefaultManager = NewManager()
