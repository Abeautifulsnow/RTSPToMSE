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

// HTTPAPIServerStreamAdd function add new stream
func HTTPAPIServerStreamAdd(c *gin.Context) {
	var payload config.StreamST
	err := c.BindJSON(&payload)
	if err != nil {
		c.IndentedJSON(400, Message{Status: 0, Payload: err.Error()})
		config.Log.WithFields(logrus.Fields{
			"module": "http_stream",
			"stream": c.Param("uuid"),
			"func":   "HTTPAPIServerStreamAdd",
			"call":   "BindJSON",
		}).Errorln(err.Error())
		return
	}
	err = config.Storage.StreamAdd(c.Param("uuid"), payload)
	if err != nil {
		c.IndentedJSON(500, Message{Status: 0, Payload: err.Error()})
		config.Log.WithFields(logrus.Fields{
			"module": "http_stream",
			"stream": c.Param("uuid"),
			"func":   "HTTPAPIServerStreamAdd",
			"call":   "StreamAdd",
		}).Errorln(err.Error())
		return
	}
	c.IndentedJSON(200, Message{Status: 1, Payload: config.Success})
}

// HTTPAPIServerStreamEdit function edit stream
func HTTPAPIServerStreamEdit(c *gin.Context) {
	var payload config.StreamST
	err := c.BindJSON(&payload)
	if err != nil {
		c.IndentedJSON(400, Message{Status: 0, Payload: err.Error()})
		config.Log.WithFields(logrus.Fields{
			"module": "http_stream",
			"stream": c.Param("uuid"),
			"func":   "HTTPAPIServerStreamEdit",
			"call":   "BindJSON",
		}).Errorln(err.Error())
		return
	}
	err = config.Storage.StreamEdit(c.Param("uuid"), payload)
	if err != nil {
		c.IndentedJSON(500, Message{Status: 0, Payload: err.Error()})
		config.Log.WithFields(logrus.Fields{
			"module": "http_stream",
			"stream": c.Param("uuid"),
			"func":   "HTTPAPIServerStreamEdit",
			"call":   "StreamEdit",
		}).Errorln(err.Error())
		return
	}
	c.IndentedJSON(200, Message{Status: 1, Payload: config.Success})
}

// HTTPAPIServerStreamDelete function delete stream
func HTTPAPIServerStreamDelete(c *gin.Context) {
	err := config.Storage.StreamDelete(c.Param("uuid"))
	if err != nil {
		c.IndentedJSON(500, Message{Status: 0, Payload: err.Error()})
		config.Log.WithFields(logrus.Fields{
			"module": "http_stream",
			"stream": c.Param("uuid"),
			"func":   "HTTPAPIServerStreamDelete",
			"call":   "StreamDelete",
		}).Errorln(err.Error())
		return
	}
	c.IndentedJSON(200, Message{Status: 1, Payload: config.Success})
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
