package ens

import (
	"fmt"
	"time"

	"github.com/imroc/req"
	"github.com/sirupsen/logrus"

	"github.com/imsilence/gocmdb/agent/entity"
	"github.com/imsilence/gocmdb/agent/gconf"
)

type ENS struct {
	config     *gconf.Config
	Heartbeat  chan interface{}
	Register   chan interface{}
	Task       chan interface{}
	TaskResult chan interface{}
	Log        chan interface{}
}

func NewENS(c *gconf.Config) *ENS {
	return &ENS{
		config:     c,
		Heartbeat:  make(chan interface{}, 16),
		Register:   make(chan interface{}, 16),
		Task:       make(chan interface{}, 64),
		TaskResult: make(chan interface{}, 128),
		Log:        make(chan interface{}, 10240),
	}
}

func (e *ENS) Start() {
	headers := req.Header{"Token": e.config.Token}

	go func() {
		endpoint := fmt.Sprintf("%s/heartbeat/%s/", e.config.Endpoint, e.config.UUID)
		for evt := range e.Heartbeat {
			if body, ok := evt.(entity.Heartbeat); ok {
				response, err := req.New().Post(endpoint, req.BodyJSON(body), headers)
				if err != nil {
					logrus.Error(response, err)
				} else {
					logrus.Debug(response, err)
				}
			}
		}
	}()

	go func() {
		endpoint := fmt.Sprintf("%s/register/%s/", e.config.Endpoint, e.config.UUID)
		for evt := range e.Register {
			if body, ok := evt.(entity.Register); ok {
				response, err := req.New().Post(endpoint, req.BodyJSON(body), headers)
				if err != nil {
					logrus.Error(response, err)
				} else {
					logrus.Debug(response, err)
				}
			}
		}
	}()

	go func() {
		endpoint := fmt.Sprintf("%s/log/%s/", e.config.Endpoint, e.config.UUID)
		for evt := range e.Log {
			if body, ok := evt.(entity.Log); ok {
				response, err := req.New().Post(endpoint, req.BodyJSON(body), headers)
				if err != nil {
					logrus.Error(response, err)
				} else {
					logrus.Debug(response, err)
				}
			}
		}
	}()

	go func() {
		endpoint := fmt.Sprintf("%s/result/%s/", e.config.Endpoint, e.config.UUID)
		for evt := range e.TaskResult {
			if body, ok := evt.(entity.TaskResult); ok {
				response, err := req.New().Post(endpoint, req.BodyJSON(body), headers)
				if err != nil {
					logrus.Error(response, err)
				} else {
					logrus.Debug(response, err)
				}
			}
		}
	}()

	go func() {
		endpoint := fmt.Sprintf("%s/task/%s/", e.config.Endpoint, e.config.UUID)
		for now := range time.Tick(10 * time.Second) {
			response, err := req.New().Get(endpoint, req.QueryParam{"time": now.Unix()}, headers)
			if err != nil {
				logrus.Error(response, err)
			} else {
				logrus.Debug(response, err)
			}
		}
	}()
}
