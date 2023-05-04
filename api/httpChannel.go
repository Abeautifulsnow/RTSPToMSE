package api

import (
	"github.com/Abeautifulsnow/RTSPToMSE/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// HTTPAPIServerStreamChannelCodec function return codec info struct
func HTTPAPIServerStreamChannelCodec(c *gin.Context) {
	requestLogger := config.Log.WithFields(logrus.Fields{
		"module":  "http_stream",
		"stream":  c.Param("uuid"),
		"channel": c.Param("channel"),
		"func":    "HTTPAPIServerStreamChannelCodec",
	})

	if !config.Storage.StreamChannelExist(c.Param("uuid"), c.Param("channel")) {
		c.IndentedJSON(500, Message{Status: 0, Payload: config.ErrorStreamNotFound.Error()})
		requestLogger.WithFields(logrus.Fields{
			"call": "StreamChannelExist",
		}).Errorln(config.ErrorStreamNotFound.Error())
		return
	}
	codecs, err := config.Storage.StreamChannelCodecs(c.Param("uuid"), c.Param("channel"))
	if err != nil {
		c.IndentedJSON(500, Message{Status: 0, Payload: err.Error()})
		requestLogger.WithFields(logrus.Fields{
			"call": "StreamChannelCodec",
		}).Errorln(err.Error())
		return
	}
	c.IndentedJSON(200, Message{Status: 1, Payload: codecs})
}

// HTTPAPIServerStreamChannelInfo function return stream info struct
func HTTPAPIServerStreamChannelInfo(c *gin.Context) {
	info, err := config.Storage.StreamChannelInfo(c.Param("uuid"), c.Param("channel"))
	if err != nil {
		c.IndentedJSON(500, Message{Status: 0, Payload: err.Error()})
		config.Log.WithFields(logrus.Fields{
			"module":  "http_stream",
			"stream":  c.Param("uuid"),
			"channel": c.Param("channel"),
			"func":    "HTTPAPIServerStreamChannelInfo",
			"call":    "StreamChannelInfo",
		}).Errorln(err.Error())
		return
	}
	c.IndentedJSON(200, Message{Status: 1, Payload: info})
}

// HTTPAPIServerStreamChannelReload function reload stream
func HTTPAPIServerStreamChannelReload(c *gin.Context) {
	err := config.Storage.StreamChannelReload(c.Param("uuid"), c.Param("channel"))
	if err != nil {
		c.IndentedJSON(500, Message{Status: 0, Payload: err.Error()})
		config.Log.WithFields(logrus.Fields{
			"module":  "http_stream",
			"stream":  c.Param("uuid"),
			"channel": c.Param("channel"),
			"func":    "HTTPAPIServerStreamChannelReload",
			"call":    "StreamChannelReload",
		}).Errorln(err.Error())
		return
	}
	c.IndentedJSON(200, Message{Status: 1, Payload: config.Success})
}
