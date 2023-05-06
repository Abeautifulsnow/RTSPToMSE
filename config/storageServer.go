package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-version"
	"github.com/liip/sheriff"
	"github.com/sirupsen/logrus"
)

var (
	//Default www static file dir
	DefaultHTTPDir = "web"
)

// ClientDelete Delete Client
func (obj *StorageST) SaveConfig() error {
	Log.WithFields(logrus.Fields{
		"module": "config",
		"func":   "NewStreamCore",
	}).Debugln("Saving configuration to", ConfigFile)
	v2, err := version.NewVersion("2.0.0")
	if err != nil {
		return err
	}
	data, err := sheriff.Marshal(&sheriff.Options{
		Groups:     []string{"config"},
		ApiVersion: v2,
	}, obj)
	if err != nil {
		return err
	}
	res, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(ConfigFile, res, 0644)
	if err != nil {
		Log.WithFields(logrus.Fields{
			"module": "config",
			"func":   "SaveConfig",
			"call":   "WriteFile",
		}).Errorln(err.Error())
		return err
	}
	return nil
}

// ServerHTTPDir
func (obj *StorageST) ServerHTTPDir() string {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	if filepath.Clean(obj.Server.HTTPDir) == "." {
		return DefaultHTTPDir
	}
	return filepath.Clean(obj.Server.HTTPDir)
}

// ServerHTTPDebug read debug options
func (obj *StorageST) ServerHTTPDebug() bool {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	return obj.Server.HTTPDebug
}

// ServerLogLevel read debug options
func (obj *StorageST) ServerLogLevel() logrus.Level {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	return obj.Server.LogLevel
}

// ServerHTTPPort read HTTP Port options
func (obj *StorageST) ServerHTTPPort() string {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	return obj.Server.HTTPPort
}

// ServerRTSPPort read HTTP Port options
func (obj *StorageST) ServerRTSPPort() string {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	return obj.Server.RTSPPort
}

// ServerHTTPS read HTTPS Port options
func (obj *StorageST) ServerHTTPS() bool {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	return obj.Server.HTTPS
}

// ServerHTTPSPort read HTTPS Port options
func (obj *StorageST) ServerHTTPSPort() string {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	return obj.Server.HTTPSPort
}

// ServerHTTPSAutoTLSEnable read HTTPS Port options
func (obj *StorageST) ServerHTTPSAutoTLSEnable() bool {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	return obj.Server.HTTPSAutoTLSEnable
}

// ServerHTTPSAutoTLSName read HTTPS Port options
func (obj *StorageST) ServerHTTPSAutoTLSName() string {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	return obj.Server.HTTPSAutoTLSName
}

// ServerHTTPSCert read HTTPS Cert options
func (obj *StorageST) ServerHTTPSCert() string {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	return obj.Server.HTTPSCert
}

// ServerHTTPSKey read HTTPS Key options
func (obj *StorageST) ServerHTTPSKey() string {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	return obj.Server.HTTPSKey
}

// ServerICEServers read ICE servers
func (obj *StorageST) ServerICEServers() []string {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()
	return obj.Server.ICEServers
}

// ServerICEServers read ICE username
func (obj *StorageST) ServerICEUsername() string {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()
	return obj.Server.ICEUsername
}

// ServerICEServers read ICE credential
func (obj *StorageST) ServerICECredential() string {
	obj.mutex.Lock()
	defer obj.mutex.Unlock()
	return obj.Server.ICECredential
}
