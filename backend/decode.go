package backend

import (
	"encoding/base64"
	"io"
	"math"
	"sync"

	"github.com/go-fish/gojson/backend/errors"
)

// Decoder defines the methods to decode message to struct
type Decoder interface {
	// string
	ReadString() (string, error)
	ReadStringWithEscape() (string, error)
	SkipString() error
	ReadBytes() ([]byte, error)
	SkipBytes() error

	// number
	ReadInt() (int, error)
	ReadInt8() (int8, error)
	ReadInt16() (int16, error)
	ReadInt32() (int32, error)
	ReadInt64() (int64, error)
	ReadUint() (uint, error)
	ReadUint8() (uint8, error)
	ReadUint16() (uint16, error)
	ReadUint32() (uint32, error)
	ReadUint64() (uint64, error)
	SkipNumber() error

	// double
	ReadFloat32() (float32, error)
	ReadFloat64() (float64, error)
	SkipFloat() error

	// bool
	ReadBool() (bool, error)
	SkipBool() error

	// interface
	ReadValue() (interface{}, error)
	ReadValueBytes() ([]byte, error)
	SkipValue() error

	// null
	AssertNull() bool

	// array
	AssertArray() (bool, error)
	ReadArray() ([]interface{}, error)
	SkipArray() error

	// object
	AssertObject() (bool, error)
	ReadObject() (map[string]interface{}, error)
	SkipObject() error

	ExpectByte(b byte) bool
	MustByte(b byte) error
}

type decoder struct {
	decoderOption

	data   []byte
	cursor int
	length int
}

type decoderOption struct {
	unsafe bool
}

// DecoderOptionFunc defines the function to set configuration of Decoder
type DecoderOptionFunc func(d *decoder)

var decoderpool *sync.Pool

func init() {
	decoderpool = &sync.Pool{
		New: func() interface{} {
			return new(decoder)
		},
	}
}

// WithUnsafe set unsafe to true
// which means decoder will use user's data directly without copy
func WithUnsafe() DecoderOptionFunc {
	return func(d *decoder) {
		d.unsafe = true
	}
}

// NewDecoder returns the instance of Decoder
func NewDecoder(data []byte, opts ...DecoderOptionFunc) Decoder {
	d := decoderpool.Get().(*decoder)

	for _, opt := range opts {
		opt(d)
	}

	if d.unsafe {
		d.data = data
		d.cursor = 0
		d.length = len(data)
		return d
	}

	d.data = make([]byte, len(data))
	d.length = copy(d.data, data)
	d.cursor = 0
	return d
}

// ReleaseDecoder gives decoder back to object pool
func ReleaseDecoder(obj Decoder) {
	d, ok := obj.(*decoder)
	if ok {
		d.data = nil
		d.cursor = 0
		d.length = 0
		d.unsafe = false

		decoderpool.Put(d)
	}
}

func (d *decoder) nextChar() (byte, error) {
	for d.cursor < d.length {
		switch d.data[d.cursor] {
		case ' ', '\n', '\t', '\r':
			d.cursor++

		default:
			return d.data[d.cursor], nil
		}
	}

	return 0, io.EOF
}

func (d *decoder) assertChar(b byte) bool {
	c, err := d.nextChar()
	if err == io.EOF {
		return false
	}

	return c == b
}

func (d *decoder) MustByte(b byte) error {
	c, err := d.nextChar()
	if err != nil {
		return err
	}

	if c == b {
		d.cursor++
		return nil
	}

	return errors.NewParseError(d.cursor, string(b), c)
}

func (d *decoder) ExpectByte(b byte) bool {
	c, err := d.nextChar()
	if err != nil {
		return false
	}

	if c == b {
		d.cursor++
		return true
	}

	return false
}

func (d *decoder) AssertArray() (bool, error) {
	c, err := d.nextChar()
	if err != nil {
		return false, err
	}

	switch c {
	case 'n', 'N':
		if d.AssertStringIngoreCase("null") {
			return true, nil
		}

		return false, errors.NewParseError(d.cursor, "null", c)
	case '[':
		d.cursor++

		if d.ExpectByte(']') {
			return true, nil
		}

		return false, nil

	default:
		return false, errors.NewParseError(d.cursor, "n or [", c)
	}
}

func (d *decoder) AssertObject() (bool, error) {
	c, err := d.nextChar()
	if err != nil {
		return false, err
	}

	switch c {
	case 'n', 'N':
		if d.AssertStringIngoreCase("null") {
			return true, nil
		}

		return false, errors.NewParseError(d.cursor, "null", c)
	case '{':
		d.cursor++

		if d.ExpectByte('}') {
			return true, nil
		}

		return false, nil

	default:
		return false, errors.NewParseError(d.cursor, "n or {", c)
	}
}

func (d *decoder) AssertNull() bool {
	return d.AssertStringIngoreCase("null")
}

func (d *decoder) AssertNaN() bool {
	return d.AssertStringIngoreCase("nan")
}

func (d *decoder) AssertInf() bool {
	return d.AssertStringIngoreCase("inf")
}

func (d *decoder) AssertTrue() bool {
	return d.AssertString("true")
}

func (d *decoder) AssertFalse() bool {
	return d.AssertString("false")
}

func (d *decoder) AssertBytes(data []byte) bool {
	if d.cursor+len(data) > d.length {
		return false
	}

	for i := d.cursor; i < d.cursor+len(data); i++ {
		if d.data[i] != data[i-d.cursor] {
			return false
		}
	}

	d.cursor += len(data)
	return true
}

func (d *decoder) AssertString(data string) bool {
	if d.cursor+len(data) > d.length {
		return false
	}

	for i := d.cursor; i < d.cursor+len(data); i++ {
		if d.data[i] != data[i-d.cursor] {
			return false
		}
	}

	d.cursor += len(data)
	return true
}

func (d *decoder) AssertBytesIngoreCase(data []byte) bool {
	count := len(data)

	if d.cursor+count > d.length {
		return false
	}

	for i := d.cursor; i < d.cursor+count; i++ {
		c := d.data[i]

		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}

		if c != data[i-d.cursor] {
			return false
		}
	}

	d.cursor += count
	return true
}

func (d *decoder) AssertStringIngoreCase(data string) bool {
	count := len(data)

	if d.cursor+count > d.length {
		return false
	}

	for i := d.cursor; i < d.cursor+count; i++ {
		c := d.data[i]

		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}

		if c != data[i-d.cursor] {
			return false
		}
	}

	d.cursor += count
	return true
}

func (d *decoder) ReadString() (string, error) {
	c, err := d.nextChar()
	if err != nil {
		return "", errors.NewEOFError(d.cursor, `"`)
	}

	switch c {
	case 'n':
		if d.AssertNull() {
			return "", nil
		}

		return "", errors.NewParseError(d.cursor, "null", c)

	case '"':
		d.cursor++
		data := d.data[d.cursor:]
		index := 0

		for index < len(data) {
			c := data[index]

			if c == '"' {
				d.cursor = d.cursor + index + 1
				return bytesToString(data[:index]), nil
			} else if c == '\\' {
				if next := index + 1; next < len(data) && (data[next] == '\\' || data[next] == '"' || data[next] == 'u') {
					//copy(data[index:], data[next:])
					data = append(data[:index], data[next:]...)
				}
			}

			index++
		}

		return "", errors.NewParseError(d.length-1, `"`, d.data[d.length-1])

	default:
		return "", errors.NewParseError(d.cursor, `"`, c)
	}
}

// ReadStringWithEscape defines the function to readstring more effective when read the string which is larger than thousand or more than thousand escaped char
func (d *decoder) ReadStringWithEscape() (string, error) {
	c, err := d.nextChar()
	if err != nil {
		return "", errors.NewEOFError(d.cursor, `"`)
	}

	switch c {
	case 'n':
		if d.AssertNull() {
			return "", nil
		}

		return "", errors.NewParseError(d.cursor, "null", c)

	case '"':
		d.cursor++
		data := d.data[d.cursor:]
		prev := -1
		escapedCount := 0
		index := 0
		cursor := 0

		for index < len(data) {
			c := data[index]

			if c == '"' {
				if prev != -1 {
					cursor += copy(data[cursor:], data[prev+1:index])
					escapedCount++
					d.cursor += cursor + 1

					if d.cursor+escapedCount < d.length {
						d.data = append(d.data[:d.cursor], d.data[d.cursor+escapedCount:]...)
						d.length -= escapedCount
					}
				} else {
					d.cursor += index + 1
				}

				return bytesToString(data[:index-escapedCount]), nil
			} else if c == '\\' {
				if next := index + 1; next < len(data) && (data[next] == '\\' || data[next] == '"' || data[next] == 'u') {
					if prev != -1 {
						cursor += copy(data[cursor:], data[prev+1:index])
						escapedCount++
					} else {
						cursor = index
					}

					prev = index
					index++
				}
			}

			index++
		}

		return "", errors.NewParseError(d.length-1, `"`, d.data[d.length-1])

	default:
		return "", errors.NewParseError(d.cursor, `"`, c)
	}
}

func (d *decoder) SkipString() error {
	c, err := d.nextChar()
	if err != nil {
		return errors.NewEOFError(d.cursor, `"`)
	}

	switch c {
	case 'n':
		if d.AssertNull() {
			return nil
		}

		return errors.NewParseError(d.cursor, "null", c)

	case '"':
		for i := d.cursor + 1; i < d.length; i++ {
			if d.data[i] == '"' {
				if prev := i - 1; d.data[prev] == '\\' {
					continue
				}

				d.cursor = i + 1
				return nil
			}
		}

		return errors.NewParseError(d.cursor, `"`, d.data[d.length-1])
	}

	return errors.NewParseError(d.cursor, `"`, c)
}

func (d *decoder) ReadBytes() ([]byte, error) {
	str, err := d.ReadString()
	if err != nil {
		return nil, err
	}

	if str == "" {
		return nil, nil
	}

	return base64.StdEncoding.DecodeString(str)
}

func (d *decoder) SkipBytes() error {
	return d.SkipString()
}

func (d *decoder) ReadInt() (int, error) {
	c, err := d.nextChar()
	if err != nil {
		return 0, errors.NewEOFError(d.cursor, "0-9")
	}

	start := d.cursor

	switch c {
	case 'n':
		if d.AssertNull() {
			return 0, nil
		}

		return 0, errors.NewParseError(d.cursor, "null", c)

	case '-':
		d.cursor++
		fallthrough

	default:
		for d.cursor < d.length {
			b := d.data[d.cursor]

			if isNumber(b) || isSkip(b) {
				d.cursor++
				continue
			}

			// whether number is closed
			if b == '.' || b == ']' ||
				b == '}' || b == ',' {
				break
			}

			return 0, errors.NewParseError(d.cursor, "0-9", b)
		}

		data := d.data[start:d.cursor]

		num, err := bytesToInt64(data)
		if err != nil && err == overflowErr {
			return 0, errors.NewOverflowError(start, data, "int")
		}

		return int(num), err
	}
}

func (d *decoder) ReadInt8() (int8, error) {
	c, err := d.nextChar()
	if err != nil {
		return 0, errors.NewEOFError(d.cursor, "0-9")
	}

	start := d.cursor

	switch c {
	case 'n':
		if d.AssertNull() {
			return 0, nil
		}

		return 0, errors.NewParseError(d.cursor, "null", c)

	case '-':
		d.cursor++
		fallthrough

	default:
		for d.cursor < d.length {
			b := d.data[d.cursor]

			if isNumber(b) || isSkip(b) {
				d.cursor++
				continue
			}

			// whether number is closed
			if b == '.' || b == ']' ||
				b == '}' || b == ',' {
				break
			}

			return 0, errors.NewParseError(d.cursor, "0-9", b)
		}

		data := d.data[start:d.cursor]

		num, err := bytesToInt8(data)
		if err != nil && err == overflowErr {
			return 0, errors.NewOverflowError(start, data, "int8")
		}

		return num, err
	}
}

func (d *decoder) ReadInt16() (int16, error) {
	c, err := d.nextChar()
	if err != nil {
		return 0, errors.NewEOFError(d.cursor, "0-9")
	}

	start := d.cursor

	switch c {
	case 'n':
		if d.AssertNull() {
			return 0, nil
		}

		return 0, errors.NewParseError(d.cursor, "null", c)

	case '-':
		d.cursor++
		fallthrough

	default:
		for d.cursor < d.length {
			b := d.data[d.cursor]

			if isNumber(b) || isSkip(b) {
				d.cursor++
				continue
			}

			// whether number is closed
			if b == '.' || b == ']' ||
				b == '}' || b == ',' {
				break
			}

			return 0, errors.NewParseError(d.cursor, "0-9", b)
		}

		data := d.data[start:d.cursor]

		num, err := bytesToInt16(data)
		if err != nil && err == overflowErr {
			return 0, errors.NewOverflowError(start, data, "int16")
		}

		return num, err
	}
}

func (d *decoder) ReadInt32() (int32, error) {
	c, err := d.nextChar()
	if err != nil {
		return 0, errors.NewEOFError(d.cursor, "0-9")
	}

	start := d.cursor

	switch c {
	case 'n':
		if d.AssertNull() {
			return 0, nil
		}

		return 0, errors.NewParseError(d.cursor, "null", c)

	case '-':
		d.cursor++
		fallthrough

	default:
		for d.cursor < d.length {
			b := d.data[d.cursor]

			if isNumber(b) || isSkip(b) {
				d.cursor++
				continue
			}

			// whether number is closed
			if b == '.' || b == ']' ||
				b == '}' || b == ',' {
				break
			}

			return 0, errors.NewParseError(d.cursor, "0-9", b)
		}

		data := d.data[start:d.cursor]

		num, err := bytesToInt32(data)
		if err != nil && err == overflowErr {
			return 0, errors.NewOverflowError(start, data, "int32")
		}

		return num, err
	}
}

func (d *decoder) ReadInt64() (int64, error) {
	c, err := d.nextChar()
	if err != nil {
		return 0, errors.NewEOFError(d.cursor, "0-9")
	}

	start := d.cursor

	switch c {
	case 'n':
		if d.AssertNull() {
			return 0, nil
		}

		return 0, errors.NewParseError(d.cursor, "null", c)

	case '-':
		d.cursor++
		fallthrough

	default:
		for d.cursor < d.length {
			b := d.data[d.cursor]

			if isNumber(b) || isSkip(b) {
				d.cursor++
				continue
			}

			// whether number is closed
			if b == '.' || b == ']' ||
				b == '}' || b == ',' {
				break
			}

			return 0, errors.NewParseError(d.cursor, "0-9", b)
		}

		data := d.data[start:d.cursor]

		num, err := bytesToInt64(data)
		if err != nil && err == overflowErr {
			return 0, errors.NewOverflowError(start, data, "int64")
		}

		return num, err
	}
}

func (d *decoder) ReadUint() (uint, error) {
	c, err := d.nextChar()
	if err != nil {
		return 0, errors.NewEOFError(d.cursor, "0-9")
	}

	start := d.cursor

	switch c {
	case 'n':
		if d.AssertNull() {
			return 0, nil
		}

		return 0, errors.NewParseError(d.cursor, "null", c)

	default:
		for d.cursor < d.length {
			b := d.data[d.cursor]

			if isNumber(b) || isSkip(b) {
				d.cursor++
				continue
			}

			// whether number is closed
			if b == '.' || b == ']' ||
				b == '}' || b == ',' {
				break
			}

			return 0, errors.NewParseError(d.cursor, "0-9", b)
		}

		data := d.data[start:d.cursor]

		num, err := bytesToUint64(data)
		if err != nil && err == overflowErr {
			return 0, errors.NewOverflowError(start, data, "uint")
		}

		return uint(num), err
	}
}

func (d *decoder) ReadUint8() (uint8, error) {
	c, err := d.nextChar()
	if err != nil {
		return 0, errors.NewEOFError(d.cursor, "0-9")
	}

	start := d.cursor

	switch c {
	case 'n':
		if d.AssertNull() {
			return 0, nil
		}

		return 0, errors.NewParseError(d.cursor, "null", c)

	default:
		for d.cursor < d.length {
			b := d.data[d.cursor]

			if isNumber(b) || isSkip(b) {
				d.cursor++
				continue
			}

			// whether number is closed
			if b == '.' || b == ']' ||
				b == '}' || b == ',' {
				break
			}

			return 0, errors.NewParseError(d.cursor, "0-9", b)
		}

		data := d.data[start:d.cursor]

		num, err := bytesToUint8(data)
		if err != nil && err == overflowErr {
			return 0, errors.NewOverflowError(start, data, "uint8")
		}

		return num, err
	}
}

func (d *decoder) ReadUint16() (uint16, error) {
	c, err := d.nextChar()
	if err != nil {
		return 0, errors.NewEOFError(d.cursor, "0-9")
	}

	start := d.cursor

	switch c {
	case 'n':
		if d.AssertNull() {
			return 0, nil
		}

		return 0, errors.NewParseError(d.cursor, "null", c)

	default:
		for d.cursor < d.length {
			b := d.data[d.cursor]

			if isNumber(b) || isSkip(b) {
				d.cursor++
				continue
			}

			// whether number is closed
			if b == '.' || b == ']' ||
				b == '}' || b == ',' {
				break
			}

			return 0, errors.NewParseError(d.cursor, "0-9", b)
		}

		data := d.data[start:d.cursor]

		num, err := bytesToUint16(data)
		if err != nil && err == overflowErr {
			return 0, errors.NewOverflowError(start, data, "uint16")
		}

		return num, err
	}
}

func (d *decoder) ReadUint32() (uint32, error) {
	c, err := d.nextChar()
	if err != nil {
		return 0, errors.NewEOFError(d.cursor, "0-9")
	}

	start := d.cursor

	switch c {
	case 'n':
		if d.AssertNull() {
			return 0, nil
		}

		return 0, errors.NewParseError(d.cursor, "null", c)

	default:
		for d.cursor < d.length {
			b := d.data[d.cursor]

			if isNumber(b) || isSkip(b) {
				d.cursor++
				continue
			}

			// whether number is closed
			if b == '.' || b == ']' ||
				b == '}' || b == ',' {
				break
			}

			return 0, errors.NewParseError(d.cursor, "0-9", b)
		}

		data := d.data[start:d.cursor]

		num, err := bytesToUint32(data)
		if err != nil && err == overflowErr {
			return 0, errors.NewOverflowError(start, data, "uint32")
		}

		return num, err
	}
}

func (d *decoder) ReadUint64() (uint64, error) {
	c, err := d.nextChar()
	if err != nil {
		return 0, errors.NewEOFError(d.cursor, "0-9")
	}

	start := d.cursor

	switch c {
	case 'n':
		if d.AssertNull() {
			return 0, nil
		}

		return 0, errors.NewParseError(d.cursor, "null", c)

	default:
		for d.cursor < d.length {
			b := d.data[d.cursor]

			if isNumber(b) || isSkip(b) {
				d.cursor++
				continue
			}

			// whether number is closed
			if b == '.' || b == ']' ||
				b == '}' || b == ',' {
				break
			}

			return 0, errors.NewParseError(d.cursor, "0-9", b)
		}

		data := d.data[start:d.cursor]

		num, err := bytesToUint64(data)
		if err != nil && err == overflowErr {
			return 0, errors.NewOverflowError(start, data, "uint64")
		}

		return num, err
	}
}

func (d *decoder) SkipNumber() error {
	c, err := d.nextChar()
	if err != nil {
		return errors.NewEOFError(d.cursor, "0-9")
	}

	switch c {
	case 'n':
		if d.AssertNull() {
			return nil
		}

		return errors.NewParseError(d.cursor, "null", c)

	case '-':
		d.cursor++
		fallthrough

	default:
		for i := d.cursor; i < d.length; i++ {
			b := d.data[i]

			if isNumber(b) || isSkip(b) {
				continue
			}

			if b == '.' || b == ']' ||
				b == '}' || b == ',' {
				d.cursor = i - 1
				return nil
			}

			return errors.NewParseError(i, "0-9", b)
		}

		return errors.NewParseError(d.length-1, "0-9", d.data[d.length-1])
	}
}

func (d *decoder) ReadFloat32() (float32, error) {
	c, err := d.nextChar()
	if err != nil {
		return 0, errors.NewEOFError(d.cursor, "")
	}

	var negative bool

	if c == '-' {
		negative = true
		d.cursor++

		c, err = d.nextChar()
		if err != nil {
			return 0, errors.NewEOFError(d.cursor, "")
		}
	}

	switch c {
	case 'n', 'N':
		if d.AssertNull() {
			return 0, nil
		}

		if d.AssertNaN() {
			return float32(math.NaN()), nil
		}

		return 0, errors.NewParseError(d.cursor, "null or nan", c)

	case 'i', 'I':
		if d.AssertInf() {
			if negative {
				return float32(math.Inf(-1)), nil
			}

			return float32(math.Inf(1)), nil
		}

		return 0, errors.NewParseError(d.cursor, "", c)

	default:
		num, err := d.readFloat()
		if err != nil {
			return 0, err
		}

		if negative {
			return -float32(num), nil
		}

		return float32(num), nil
	}
}

func (d *decoder) ReadFloat64() (float64, error) {
	c, err := d.nextChar()
	if err != nil {
		return 0, errors.NewEOFError(d.cursor, "")
	}

	var negative bool

	if c == '-' {
		negative = true
		d.cursor++

		c, err = d.nextChar()
		if err != nil {
			return 0, errors.NewEOFError(d.cursor, "")
		}
	}

	switch c {
	case 'n', 'N':
		if d.AssertNull() {
			return 0, nil
		}

		if d.AssertNaN() {
			return math.NaN(), nil
		}

		return 0, errors.NewParseError(d.cursor, "null or nan", c)

	case 'i', 'I':
		if d.AssertInf() {
			if negative {
				return math.Inf(-1), nil
			}

			return math.Inf(1), nil
		}

		return 0, errors.NewParseError(d.cursor, "", c)

	default:
		num, err := d.readFloat()
		if err != nil {
			return 0, err
		}

		if negative {
			return -num, nil
		}

		return num, nil
	}
}

func (d *decoder) SkipFloat() error {
	c, err := d.nextChar()
	if err != nil {
		return errors.NewEOFError(d.cursor, "")
	}

	if c == '-' {
		d.cursor++

		c, err = d.nextChar()
		if err != nil {
			return errors.NewEOFError(d.cursor, "")
		}
	}

	switch c {
	case 'n', 'N':
		if d.AssertNull() {
			return nil
		}

		if d.AssertNaN() {
			return nil
		}

		return errors.NewParseError(d.cursor, "null or nan", c)

	case 'i', 'I':
		if d.AssertInf() {
			return nil
		}

		return errors.NewParseError(d.cursor, "", c)

	default:
		return d.skipFloat()
	}
}

func (d *decoder) readFloat() (float64, error) {
	var digit, decimal, exp uint64
	var fract, exponent, esign bool
	var dp int
	var err error

Loop:
	for d.cursor < d.length {
		switch c := d.data[d.cursor]; true {
		case c >= '0' && c <= '9':
			if fract && !exponent {
				digit = (digit << 3) + (digit << 1)
				decimal = (decimal << 3) + (decimal << 1) + uint64(c-'0')
				dp++
				d.cursor++
				continue
			}

			if exponent {
				exp = (exp << 3) + (exp << 1) + uint64(c-'0')
				d.cursor++
				continue
			}

			if !fract && !exponent {
				digit = (digit << 3) + (digit << 1) + uint64(c-'0')
				d.cursor++
			}

		case c == '.':
			if fract {
				return 0, errors.NewParseError(d.cursor, "", c)
			}

			fract = true
			d.cursor++

		case c == 'e' || c == 'E':
			if exponent {
				return 0, errors.NewParseError(d.cursor, "", c)
			}

			exponent = true
			d.cursor++

			if d.cursor >= d.length {
				return 0, errors.NewParseError(d.cursor, "", c)
			}

			switch b := d.data[d.cursor]; b {
			case '+', '-':
				esign = b == '+'
				d.cursor++

			default:
				return 0, errors.NewParseError(d.cursor, "+/-", b)
			}

		default:
			c, err = d.nextChar()
			if err != nil {
				return 0, err
			}

			if c == ']' ||
				c == '}' || c == ',' {
				break Loop
			}

			return 0, errors.NewParseError(d.cursor, "", c)
		}
	}

	num := float64(digit)

	if fract {
		num += float64(decimal)
		num = num / float64(math.Pow10(dp))
	}

	if exponent {
		if esign {
			num = num * float64(math.Pow10(int(exp)))
		} else {
			num = num / float64(math.Pow10(int(exp)))
		}
	}

	return num, nil
}

func (d *decoder) skipFloat() error {
	var fract, exponent bool
	var err error

Loop:
	for d.cursor < d.length {
		switch c := d.data[d.cursor]; true {
		case c >= '0' && c <= '9':
			d.cursor++

		case c == '.':
			if fract {
				return errors.NewParseError(d.cursor, "", c)
			}

			fract = true
			d.cursor++

		case c == 'e' || c == 'E':
			if exponent {
				return errors.NewParseError(d.cursor, "", c)
			}

			exponent = true
			d.cursor++

			if d.cursor >= d.length {
				return errors.NewParseError(d.cursor, "", c)
			}

			switch d.data[d.cursor] {
			case '+', '-':
				d.cursor++

			default:
				return errors.NewParseError(d.cursor, "+/-", c)
			}

		default:
			c, err = d.nextChar()
			if err != nil {
				return err
			}

			if c == ']' ||
				c == '}' || c == ',' {
				break Loop
			}

			return errors.NewParseError(d.cursor, "", c)
		}
	}

	return nil
}

func (d *decoder) ReadBool() (bool, error) {
	c, err := d.nextChar()
	if err != nil {
		return false, err
	}

	switch c {
	case 't':
		if d.AssertTrue() {
			return true, nil
		}

	case 'f':
		if d.AssertFalse() {
			return false, nil
		}

	}

	return false, errors.NewParseError(d.cursor, "", c)
}

func (d *decoder) SkipBool() error {
	c, err := d.nextChar()
	if err != nil {
		return err
	}

	switch c {
	case 't':
		if d.AssertTrue() {
			return nil
		}

	case 'f':
		if d.AssertFalse() {
			return nil
		}
	}

	return errors.NewParseError(d.cursor, "", c)
}

func (d *decoder) ReadArray() ([]interface{}, error) {
	v := make([]interface{}, 0, 8)

	c, err := d.nextChar()
	if err != nil {
		return nil, err
	}

	switch c {
	case 'n':
		if d.AssertNull() {
			return nil, nil
		}

		return nil, errors.NewParseError(d.cursor, "null", c)

	case '[':
		d.cursor++

		if d.ExpectByte(']') {
			return []interface{}{}, nil
		}

		for d.cursor < d.length {
			value, err := d.ReadValue()
			if err != nil {
				return nil, err
			}

			v = append(v, value)

			nextChar, err := d.nextChar()
			if err != nil {
				return nil, err
			}

			switch nextChar {
			case ',':
				d.cursor++

			case ']':
				d.cursor++
				return v, nil

			default:
				return nil, errors.NewParseError(d.cursor, ", or ]", c)
			}
		}
	}

	return nil, errors.NewParseError(d.cursor, "", d.data[d.cursor])
}

func (d *decoder) SkipArray() error {
	c, err := d.nextChar()
	if err != nil {
		return err
	}

	switch c {
	case 'n':
		if d.AssertNull() {
			return nil
		}

		return errors.NewParseError(d.cursor, "null", c)

	case '[':
		d.cursor++

		if d.ExpectByte(']') {
			return nil
		}

		for d.cursor < d.length {
			if err := d.SkipValue(); err != nil {
				return err
			}

			nextChar, err := d.nextChar()
			if err != nil {
				return err
			}

			switch nextChar {
			case ',':
				d.cursor++

			case ']':
				d.cursor++
				return nil

			default:
				return errors.NewParseError(d.cursor, ", or ]", c)
			}
		}
	}

	return errors.NewParseError(d.cursor, "", d.data[d.cursor])
}

func (d *decoder) ReadObject() (map[string]interface{}, error) {
	v := make(map[string]interface{})

	c, err := d.nextChar()
	if err != nil {
		return nil, err
	}

	switch c {
	case 'n':
		if d.AssertNull() {
			return nil, nil
		}

		return nil, errors.NewParseError(d.cursor, "null", c)

	case '{':
		d.cursor++

		if d.ExpectByte('}') {
			return map[string]interface{}{}, nil
		}

		for d.cursor < d.length {
			key, err := d.ReadString()
			if err != nil {
				return nil, err
			}

			if !d.ExpectByte(':') {
				return nil, errors.NewParseError(d.cursor, ":", c)
			}

			value, err := d.ReadValue()
			if err != nil {
				return nil, err
			}

			v[key] = value

			nextChar, err := d.nextChar()
			if err != nil {
				return nil, err
			}

			switch nextChar {
			case ',':
				d.cursor++

			case '}':
				d.cursor++
				return v, nil

			default:
				return nil, errors.NewParseError(d.cursor, ", or }", c)
			}
		}
	}

	return nil, errors.NewParseError(d.cursor, "", d.data[d.cursor])
}

func (d *decoder) SkipObject() error {
	c, err := d.nextChar()
	if err != nil {
		return err
	}

	switch c {
	case 'n':
		if d.AssertNull() {
			return nil
		}

		return errors.NewParseError(d.cursor, "null", c)

	case '{':
		d.cursor++

		if d.ExpectByte('}') {
			return nil
		}

		for d.cursor < d.length {
			if err := d.SkipString(); err != nil {
				return err
			}

			if !d.ExpectByte(':') {
				return errors.NewParseError(d.cursor, ":", c)
			}

			if err := d.SkipValue(); err != nil {
				return err
			}

			nextChar, err := d.nextChar()
			if err != nil {
				return err
			}

			switch nextChar {
			case ',':
				d.cursor++

			case '}':
				d.cursor++
				return nil

			default:
				return errors.NewParseError(d.cursor, ", or }", c)
			}
		}
	}

	return errors.NewParseError(d.cursor, "", d.data[d.cursor])
}

func (d *decoder) ReadValue() (interface{}, error) {
	c, err := d.nextChar()
	if err != nil {
		return nil, err
	}

	switch c {
	case 'n':
		if d.AssertNull() {
			return nil, nil
		}

		return nil, errors.NewParseError(d.cursor, "null", c)

	case '"':
		return d.ReadString()

	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return d.ReadFloat64()

	case '[':
		return d.ReadArray()

	case '{':
		return d.ReadObject()

	case 't', 'f':
		return d.ReadBool()

	default:
		return nil, errors.NewParseError(d.cursor, "", c)
	}
}

func (d *decoder) SkipValue() error {
	c, err := d.nextChar()
	if err != nil {
		return err
	}

	switch c {
	case 'n':
		if !d.AssertNull() {
			return errors.NewParseError(d.cursor, "null", c)
		}

		return nil

	case '"':
		return d.SkipString()

	case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return d.SkipFloat()

	case '[':
		return d.SkipArray()

	case '{':
		return d.SkipObject()

	case 't', 'f':
		return d.SkipBool()

	default:
		return errors.NewParseError(d.cursor, "", c)
	}
}

func (d *decoder) ReadValueBytes() ([]byte, error) {
	start := d.cursor

	if err := d.SkipValue(); err != nil {
		return nil, err
	}

	return d.data[start:d.cursor], nil
}
