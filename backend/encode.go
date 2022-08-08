package backend

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
	"sync"
	"unicode/utf16"
	"unicode/utf8"
)

const quota = '"'
const comma = ','
const colon = ':'
const objectOpen = '{'
const objectClose = '}'
const arrayOpen = '['
const arrayClose = ']'

// Encoder defines the method to encode each type to []byte
type Encoder interface {
	// key
	EncodeKey(key string)

	// String
	EncodeString(value string)

	// Number
	EncodeInt(value int)
	EncodeInt8(value int8)
	EncodeInt16(value int16)
	EncodeInt32(value int32)
	EncodeInt64(value int64)
	EncodeUint(value uint)
	EncodeUint8(value uint8)
	EncodeUint16(value uint16)
	EncodeUint32(value uint32)
	EncodeUint64(value uint64)

	// Double
	EncodeFloat32(value float32)
	EncodeFloat64(value float64)

	// bool
	EncodeBool(value bool)

	// Bytes
	EncodeBytes(value []byte)

	// Interface
	EncodeInterface(value interface{})

	// Array
	EncodeInterfaceArray(values []interface{})

	// Map
	EncodeInterfaceMap(values map[string]interface{})

	Write(data []byte)
	WriteByte(char byte)
	WriteNull()

	Bytes() []byte
	String() string
}

type encoder struct {
	encoderOption

	data []byte
}

type encoderOption struct {
	escapeHTML    bool
	escapeUnicode bool
}

// EncoderOptionFunc defines the function to set option of encoder
type EncoderOptionFunc func(e *encoder)

var encoderpool *sync.Pool

func init() {
	encoderpool = &sync.Pool{
		New: func() interface{} {
			return new(encoder)
		},
	}
}

func acquireEncoder() *encoder {
	return encoderpool.Get().(*encoder)
}

// WithEscapeHTML set escapeHtml to true
// which will escape html char, such as < > &
func WithEscapeHTML() EncoderOptionFunc {
	return func(e *encoder) {
		e.escapeHTML = true
	}
}

// WithEscapeUnicode set escapeUnicode to true
// which will convert rune to unicode char
func WithEscapeUnicode() EncoderOptionFunc {
	return func(e *encoder) {
		e.escapeUnicode = true
	}
}

// NewEncoder returns the new instance of encoder from object pool
func NewEncoder(opts ...EncoderOptionFunc) Encoder {
	e := acquireEncoder()

	for _, opt := range opts {
		opt(e)
	}

	e.data = make([]byte, 0, 1024)

	return e
}

// ReleaseEncoder gives Encoder back to object pool
func ReleaseEncoder(obj Encoder) {
	e, ok := obj.(*encoder)
	if ok {
		encoderpool.Put(e)
	}
}

func (e *encoder) Write(data []byte) {
	e.write(data)
}

func (e *encoder) WriteByte(char byte) {
	e.writeByte(char)
}

func (e *encoder) EncodeKey(key string) {
	e.writeByte(quota)
	e.writeString(key)
	e.writeByte(quota)
	e.writeByte(colon)
}

const lowhex = "0123456789abcdef"

// https://datatracker.ietf.org/doc/html/rfc7159#section-7
func (e *encoder) EncodeString(value string) {
	e.writeByte(quota)

	// write escape string
	for i := 0; i < len(value); {
		c := value[i]

		// just append valid ASCII byte to encoder
		if c < utf8.RuneSelf {
			switch c {
			case '"', '\\', '\r', '\n', '\f', '\b', '\t':
				e.writeByte('\\')

			default:
				// This encodes bytes < 0x20 except for \t, \n and \r.
				// If escapeHTML is set, it also escapes <, >, and &
				if c < 0x20 || (e.escapeHTML && isEscapeHtmlByte(c)) {
					e.writeString(`\u00`)
					e.writeByte(lowhex[c>>4])
					e.writeByte(lowhex[c&0xF])
					i++
					continue
				}
			}
		} else if e.escapeUnicode {
			// https://datatracker.ietf.org/doc/html/rfc7159#section-8.1
			r, size := utf8.DecodeRuneInString(value[i:])
			if r == utf8.RuneError && size == 1 {
				e.writeString(`\ufffd`)
				i++
				continue
			}

			if r < 0x10000 {
				e.writeRune(r)
				i += size
				continue
			}

			hi, lo := utf16.EncodeRune(r)
			e.writeRune(hi)
			e.writeRune(lo)
			i += size
			continue
		}

		e.writeByte(c)
		i++
	}

	e.writeByte(quota)
}

func (e *encoder) EncodeInt(value int) {
	e.write(int64ToBytes(int64(value)))
}

func (e *encoder) EncodeInt8(value int8) {
	e.write(int8ToBytes(value))
}

func (e *encoder) EncodeInt16(value int16) {
	e.write(int16ToBytes(value))
}

func (e *encoder) EncodeInt32(value int32) {
	e.write(int32ToBytes(value))
}

func (e *encoder) EncodeInt64(value int64) {
	e.write(int64ToBytes(value))
}

func (e *encoder) EncodeUint(value uint) {
	e.write(uint64ToBytes(uint64(value)))
}

func (e *encoder) EncodeUint8(value uint8) {
	e.write(uint8ToBytes(value))
}

func (e *encoder) EncodeUint16(value uint16) {
	e.write(uint16ToBytes(value))
}

func (e *encoder) EncodeUint32(value uint32) {
	e.write(uint32ToBytes(value))
}

func (e *encoder) EncodeUint64(value uint64) {
	e.write(uint64ToBytes(value))
}

func (e *encoder) EncodeFloat32(value float32) {
	e.writeString(strconv.FormatFloat(float64(value), 'g', -1, 32))
}

func (e *encoder) EncodeFloat64(value float64) {
	e.writeString(strconv.FormatFloat(value, 'g', -1, 64))
}

const (
	trueStr  = "true"
	falseStr = "false"
)

func (e *encoder) EncodeBool(value bool) {
	if value {
		e.writeString(trueStr)
		return
	}

	e.writeString(falseStr)
}

func (e *encoder) EncodeBytes(value []byte) {
	if value == nil {
		e.WriteNull()
		return
	}

	dst := acquireBytes()
	defer releaseBytes(dst)
	count := base64.StdEncoding.EncodedLen(len(value))
	base64.StdEncoding.Encode(dst, value)
	e.writeByte(quota)
	e.write(dst[:count])
	e.writeByte(quota)
}

func (e *encoder) EncodeInterface(value interface{}) {
	switch v := value.(type) {
	case string:
		e.EncodeString(v)

	case int:
		e.EncodeInt(v)

	case int8:
		e.EncodeInt8(v)

	case int16:
		e.EncodeInt16(v)

	case int32:
		e.EncodeInt32(v)

	case int64:
		e.EncodeInt64(v)

	case uint:
		e.EncodeUint(v)

	case uint8:
		e.EncodeUint8(v)

	case uint16:
		e.EncodeUint16(v)

	case uint32:
		e.EncodeUint32(v)

	case uint64:
		e.EncodeUint64(v)

	case float32:
		e.EncodeFloat32(v)

	case float64:
		e.EncodeFloat64(v)

	case bool:
		e.EncodeBool(v)

	case []byte:
		e.EncodeBytes(v)

	default:
		data, _ := json.Marshal(value)
		e.write(data)
	}
}

func (e *encoder) EncodeInterfaceArray(values []interface{}) {
	if len(values) == 0 {
		e.writeByte(arrayOpen)
		e.writeByte(arrayClose)
		return
	}

	e.writeByte(arrayOpen)

	for i := 0; i < len(values); i++ {
		if i > 0 {
			e.writeByte(comma)
		}

		e.EncodeInterface(values[i])
	}

	e.writeByte(arrayClose)
}

func (e *encoder) EncodeInterfaceMap(values map[string]interface{}) {
	if len(values) == 0 {
		e.writeByte(objectOpen)
		e.writeByte(objectClose)
		return
	}

	e.writeByte(objectOpen)
	index := 0

	for key, value := range values {
		if index > 0 {
			e.writeByte(comma)
		}

		e.EncodeKey(key)
		e.EncodeInterface(value)
	}

	e.writeByte(objectClose)
}

func (e *encoder) write(data []byte) {
	e.data = append(e.data, data...)
}

func (e *encoder) writeByte(b byte) {
	e.data = append(e.data, b)
}

func (e *encoder) writeRune(r rune) {
	e.writeString(`\u`)

	for i := 12; i >= 0; i -= 4 {
		e.writeByte(lowhex[r>>uint(i)&0xF])
	}
}

func (e *encoder) writeString(str string) {
	e.data = append(e.data, str...)
}

var jsonnull = []byte("null")

func (e *encoder) WriteNull() {
	e.write(jsonnull)
}

func (e *encoder) Bytes() []byte {
	return e.data
}

func (e *encoder) String() string {
	return bytesToString(e.data)
}
