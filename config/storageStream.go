package config

import (
	"github.com/Abeautifulsnow/RTSPToMSE/utils"
	"github.com/liip/sheriff"
)

// MarshalledStreamsList lists all streams and includes only fields which are safe to serialize.
func (obj *StorageST) MarshalledStreamsList() (interface{}, error) {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	val, err := sheriff.Marshal(&sheriff.Options{
		Groups: []string{"api"},
	}, obj.Streams)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// StreamReload reload stream
func (obj *StorageST) StopAll() {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	for _, st := range obj.Streams {
		for _, i2 := range st.Channels {
			if i2.runLock {
				i2.signals <- utils.SignalStreamStop
			}
		}
	}
}

// StreamReload reload stream
func (obj *StorageST) StreamReload(uuid string) error {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	if tmp, ok := obj.Streams[uuid]; ok {
		for _, i2 := range tmp.Channels {
			if i2.runLock {
				i2.signals <- utils.SignalStreamRestart
			}
		}
		return nil
	}
	return ErrorStreamNotFound
}

// StreamInfo return stream info
func (obj *StorageST) StreamInfo(uuid string) (*StreamST, error) {
	obj.mutex.RLock()
	defer obj.mutex.RUnlock()
	if tmp, ok := obj.Streams[uuid]; ok {
		return &tmp, nil
	}
	return nil, ErrorStreamNotFound
}
