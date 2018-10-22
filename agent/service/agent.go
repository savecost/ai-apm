package service

import (
	"github.com/mafanr/g"
	"github.com/mafanr/vgo/util"

	"go.uber.org/zap"
)

// Agent ...
type Agent struct {
	quitC     chan bool
	uploadC   chan *util.VgoPacket
	downloadC chan *util.VgoPacket
	client    *TcpClient
}

var gAgent *Agent

// New ...
func New() *Agent {
	gAgent = &Agent{
		quitC:     make(chan bool, 1),
		uploadC:   make(chan *util.VgoPacket, 1000),
		downloadC: make(chan *util.VgoPacket, 100),
		client:    NewTcpClient(),
	}
	return gAgent
}

// Start ...
func (a *Agent) Start() error {
	// 启动upload
	go a.upload()

	// 初始化处理下行命令等
	go a.download()

	// 初始化tcp client
	go a.client.Init()

	// 启动本地接收采集信息端口
	//a.pinpoint.Start()

	return nil
}

// Close ...
func (a *Agent) Close() error {

	return nil
}

func (a *Agent) upload() {
	defer func() {
		if err := recover(); err != nil {
			g.L.Warn("report:.", zap.Stack("server"), zap.Any("err", err))
		}
	}()

	for {
		select {
		case p, ok := <-a.uploadC:
			if ok {
				if err := a.client.WritePacket(p); err != nil {
					g.L.Warn("report:client.WritePacket", zap.String("error", err.Error()))
				}
			}
			break
		}
	}
}

func (a *Agent) download() {
	for {
		select {
		case p, ok := <-a.downloadC:
			if ok {
				g.L.Info("cmd", zap.Any("msg", p))
			}
		case <-a.quitC:
			return
		}
	}
}