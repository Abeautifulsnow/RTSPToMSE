package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Abeautifulsnow/RTSPToMSE/api"
	"github.com/Abeautifulsnow/RTSPToMSE/config"
	"github.com/sirupsen/logrus"
)

func main() {
	config.Log.WithFields(logrus.Fields{
		"module": "main",
		"func":   "main",
	}).Info("Server CORE start")
	go api.HTTPAPIServer()
	go config.RTSPServer()
	go config.Storage.StreamChannelRunAll()
	signalChanel := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(signalChanel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signalChanel
		config.Log.WithFields(logrus.Fields{
			"module": "main",
			"func":   "main",
		}).Info("Server receive signal", sig)
		done <- true
	}()
	config.Log.WithFields(logrus.Fields{
		"module": "main",
		"func":   "main",
	}).Info("Server start success a wait signals")
	<-done
	config.Storage.StopAll()
	time.Sleep(2 * time.Second)
	config.Log.WithFields(logrus.Fields{
		"module": "main",
		"func":   "main",
	}).Info("Server stop working by signal")
}
