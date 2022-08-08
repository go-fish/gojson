package generate

import (
	"bytes"
	"fmt"
	"sync"
)

type writer interface {
	line(format string, args ...interface{})
	bytes() []byte
}

type buffer struct {
	w *bytes.Buffer
}

var bufferpool *sync.Pool

func init() {
	bufferpool = &sync.Pool{
		New: func() interface{} {
			return &buffer{
				w: bytes.NewBuffer(make([]byte, 2048)),
			}
		},
	}
}

func acquireWriter() writer {
	b := bufferpool.Get().(*buffer)
	b.w.Reset()
	return b
}

func releaseWriter(w writer) {
	b, ok := w.(*buffer)
	if !ok {
		return
	}

	b.w.Reset()
	bufferpool.Put(b)
}

func (b *buffer) line(format string, args ...interface{}) {
	b.w.WriteString(fmt.Sprintf(format, args...))
	b.w.WriteByte('\n')
}

func (b *buffer) bytes() []byte {
	return b.w.Bytes()
}
