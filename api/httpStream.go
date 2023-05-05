package api

import (
	"github.com/Abeautifulsnow/RTSPToMSE/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// HTTPAPIServerStreams function return stream list
func HTTPAPIServerStreams(c *gin.Context) {
	list, err := config.Storage.MarshalledStreamsList()
	if err != nil {
		c.IndentedJSON(500, Message{Status: 0, Payload: err.Error()})
		return
	}
	c.IndentedJSON(200, Message{Status: 1, Payload: list})
}

// HTTPAPIServerStreamDelete function reload stream
func HTTPAPIServerStreamReload(c *gin.Context) {
	err := config.Storage.StreamReload(c.Param("uuid"))
	if err != nil {
		c.IndentedJSON(500, Message{Status: 0, Payload: err.Error()})
		config.Log.WithFields(logrus.Fields{
			"module": "http_stream",
			"stream": c.Param("uuid"),
			"func":   "HTTPAPIServerStreamReload",
			"call":   "StreamReload",
		}).Errorln(err.Error())
		return
	}
	c.IndentedJSON(200, Message{Status: 1, Payload: config.Success})
}

// HTTPAPIServerStreamInfo function return stream info struct
func HTTPAPIServerStreamInfo(c *gin.Context) {
	info, err := config.Storage.StreamInfo(c.Param("uuid"))
	if err != nil {
		c.IndentedJSON(500, Message{Status: 0, Payload: err.Error()})
		config.Log.WithFields(logrus.Fields{
			"module": "http_stream",
			"stream": c.Param("uuid"),
			"func":   "HTTPAPIServerStreamInfo",
			"call":   "StreamInfo",
		}).Errorln(err.Error())
		return
	}
	c.IndentedJSON(200, Message{Status: 1, Payload: info})
}
