package backend

import (
	"sync"
)

var bytespool *sync.Pool

func init() {
	bytespool = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024)
		},
	}
}

func acquireBytes() []byte {
	return bytespool.Get().([]byte)
}

func releaseBytes(data []byte) {
	bytespool.Put(data)
}
