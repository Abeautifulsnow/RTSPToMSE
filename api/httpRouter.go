package api

import (
	"net/http"
	"os"
	"time"

	"github.com/Abeautifulsnow/RTSPToMSE/config"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Message struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload"`
}

// HTTPAPIServer start http server routes
func HTTPAPIServer() {
	//Set HTTP API mode
	config.Log.WithFields(logrus.Fields{
		"module": "http_server",
		"func":   "RTSPServer",
		"call":   "Start",
	}).Infoln("Server HTTP start")
	var public *gin.Engine
	if !config.Storage.ServerHTTPDebug() {
		gin.SetMode(gin.ReleaseMode)
		public = gin.New()
	} else {
		gin.SetMode(gin.DebugMode)
		public = gin.Default()
	}

	public.Use(CrossOrigin())
	//Add private login password protect methods
	privat := public.Group("/")
	if config.Storage.ServerHTTPLogin() != "" && config.Storage.ServerHTTPPassword() != "" {
		privat.Use(gin.BasicAuth(gin.Accounts{config.Storage.ServerHTTPLogin(): config.Storage.ServerHTTPPassword()}))
	}

	/*
		Static HTML Files Demo Mode
	*/

	if config.Storage.ServerHTTPDemo() {
		public.GET("/pages/stream/list", HTTPAPIStreamList)
		public.GET("/pages/player/hls/:uuid/:channel", HTTPAPIPlayHls)
		public.GET("/pages/player/mse/:uuid/:channel", HTTPAPIPlayMse)
	}

	/*
		Stream Control elements
	*/

	privat.GET("/streams", HTTPAPIServerStreams)
	privat.POST("/stream/:uuid/add", HTTPAPIServerStreamAdd)
	privat.POST("/stream/:uuid/edit", HTTPAPIServerStreamEdit)
	privat.GET("/stream/:uuid/delete", HTTPAPIServerStreamDelete)
	privat.GET("/stream/:uuid/reload", HTTPAPIServerStreamReload)
	privat.GET("/stream/:uuid/info", HTTPAPIServerStreamInfo)

	/*
		Stream Channel elements
	*/

	privat.GET("/stream/:uuid/channel/:channel/codec", HTTPAPIServerStreamChannelCodec)
	privat.GET("/stream/:uuid/channel/:channel/reload", HTTPAPIServerStreamChannelReload)
	privat.GET("/stream/:uuid/channel/:channel/info", HTTPAPIServerStreamChannelInfo)

	//MSE
	public.GET("/stream/:uuid/channel/:channel/mse", HTTPAPIServerStreamMSE)
	/*
		HTTPS Mode Cert
		# Key considerations for algorithm "RSA" ≥ 2048-bit
		openssl genrsa -out server.key 2048

		# Key considerations for algorithm "ECDSA" ≥ secp384r1
		# List ECDSA the supported curves (openssl ecparam -list_curves)
		#openssl ecparam -genkey -name secp384r1 -out server.key
		#Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)

		openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
	*/
	if config.Storage.ServerHTTPS() {
		if config.Storage.ServerHTTPSAutoTLSEnable() {
			go func() {
				err := autotls.Run(public, config.Storage.ServerHTTPSAutoTLSName()+config.Storage.ServerHTTPSPort())
				if err != nil {
					config.Log.Println("Start HTTPS Server Error", err)
				}
			}()
		} else {
			go func() {
				err := public.RunTLS(config.Storage.ServerHTTPSPort(), config.Storage.ServerHTTPSCert(), config.Storage.ServerHTTPSKey())
				if err != nil {
					config.Log.WithFields(logrus.Fields{
						"module": "http_router",
						"func":   "HTTPSAPIServer",
						"call":   "ServerHTTPSPort",
					}).Fatalln(err.Error())
					os.Exit(1)
				}
			}()
		}
	}
	err := public.Run(config.Storage.ServerHTTPPort())
	if err != nil {
		config.Log.WithFields(logrus.Fields{
			"module": "http_router",
			"func":   "HTTPAPIServer",
			"call":   "ServerHTTPPort",
		}).Fatalln(err.Error())
		os.Exit(1)
	}

}

func HTTPAPIStreamList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"port":    config.Storage.ServerHTTPPort(),
		"streams": config.Storage.Streams,
		"version": time.Now().String(),
		"page":    "stream_list",
	})
}

func HTTPAPIPlayHls(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"port":    config.Storage.ServerHTTPPort(),
		"streams": config.Storage.Streams,
		"version": time.Now().String(),
		"page":    "play_hls",
		"uuid":    c.Param("uuid"),
		"channel": c.Param("channel"),
	})
}
func HTTPAPIPlayMse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"port":    config.Storage.ServerHTTPPort(),
		"streams": config.Storage.Streams,
		"version": time.Now().String(),
		"page":    "play_mse",
		"uuid":    c.Param("uuid"),
		"channel": c.Param("channel"),
	})
}

type MultiViewOptions struct {
	Grid   int                             `json:"grid"`
	Player map[string]MultiViewOptionsGrid `json:"player"`
}
type MultiViewOptionsGrid struct {
	UUID       string `json:"uuid"`
	Channel    int    `json:"channel"`
	PlayerType string `json:"playerType"`
}

// CrossOrigin Access-Control-Allow-Origin any methods
func CrossOrigin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
