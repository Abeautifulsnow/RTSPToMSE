package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/imdario/mergo"
	"github.com/sirupsen/logrus"
)

// Command line flag global variables
var Debug bool
var ConfigFile string

var Log = logrus.New()
var Storage = NewStreamCore()

func init() {
	//TODO: next add write to file
	if !Debug {
		Log.SetOutput(io.Discard)
	}

	// Enable output log content containing line number.
	Log.SetReportCaller(true)
	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	Log.SetLevel(Storage.ServerLogLevel())
}

// NewStreamCore do load config file
func NewStreamCore() *StorageST {
	flag.BoolVar(&Debug, "Debug", true, "set Debug mode")
	flag.StringVar(&ConfigFile, "config", "config.json", "config patch (/etc/server/config.json or config.json)")
	flag.Parse()

	var tmp StorageST
	data, err := os.ReadFile(ConfigFile)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"module": "config",
			"func":   "NewStreamCore",
			"call":   "ReadFile",
		}).Errorln(err.Error())
		os.Exit(1)
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"module": "config",
			"func":   "NewStreamCore",
			"call":   "Unmarshal",
		}).Errorln(err.Error())
		os.Exit(1)
	}
	Debug = tmp.Server.Debug
	for i, i2 := range tmp.Streams {
		for i3, i4 := range i2.Channels {
			channel := tmp.ChannelDefaults
			err = mergo.Merge(&channel, i4)
			if err != nil {
				Log.WithFields(logrus.Fields{
					"module": "config",
					"func":   "NewStreamCore",
					"call":   "Merge",
				}).Errorln(err.Error())
				os.Exit(1)
			}
			channel.clients = make(map[string]ClientST)
			channel.ack = time.Now().Add(-255 * time.Hour)
			channel.hlsSegmentBuffer = make(map[int]SegmentOld)
			channel.signals = make(chan int, 100)
			i2.Channels[i3] = channel
		}
		tmp.Streams[i] = i2
	}
	return &tmp
}
