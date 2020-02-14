package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/imsilence/gocmdb/agent/gconf"
	"github.com/imsilence/gocmdb/agent/plugins"
	"github.com/imsilence/gocmdb/agent/ens"
	_ "github.com/imsilence/gocmdb/agent/plugins/init"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	conf := gconf.NewConfig()

	ens := ens.NewENS(conf)
	ens.Start()

	logrus.Info("agent is staring...")
	plugins.DefaultManager.Init(conf, ens)
	plugins.DefaultManager.Start()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<- ch
}