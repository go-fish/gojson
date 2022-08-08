package backend

import (
	"encoding/json"
	"errors"
	"math"
	"unsafe"
)

var overflowErr = errors.New("number overflow")

func bytesToInt8(data []byte) (int8, error) {
	var num int8
	var negative bool

	if data[0] == '-' {
		data = data[1:]
		negative = true
	}

	for i := 0; i < len(data); i++ {
		num = (num << 3) + (num << 1)
		add := int8(data[i]) - '0'

		if num > math.MaxInt8-add {
			return 0, overflowErr
		}

		num += add
	}

	if negative {
		return -int8(num), nil
	}

	return int8(num), nil
}

func int8ToBytes(num int8) []byte {
	var data [32]byte
	var negative bool

	if num < 0 {
		negative = true
		num = 0 - num
	}

	cursor := 31

	for num >= 10 {
		v := num % 10
		num = num / 10
		data[cursor] = byte(v) + '0'
		cursor--
	}

	data[cursor] = byte(num) + '0'

	if negative {
		cursor--
		data[cursor] = '-'
	}

	return data[cursor:]
}

func bytesToInt16(data []byte) (int16, error) {
	var num int16
	var negative bool

	if data[0] == '-' {
		data = data[1:]
		negative = true
	}

	for i := 0; i < len(data); i++ {
		num = (num << 3) + (num << 1)
		add := int16(data[i]) - '0'

		if num > math.MaxInt16-add {
			return 0, overflowErr
		}

		num += add
	}

	if negative {
		return -num, nil
	}

	return num, nil
}

func int16ToBytes(num int16) []byte {
	var data [32]byte
	var negative bool

	if num < 0 {
		negative = true
		num = 0 - num
	}

	cursor := 31

	for num >= 10 {
		v := num % 10
		num = num / 10
		data[cursor] = byte(v) + '0'
		cursor--
	}

	data[cursor] = byte(num) + '0'

	if negative {
		cursor--
		data[cursor] = '-'
	}

	return data[cursor:]
}

func bytesToInt32(data []byte) (int32, error) {
	var num int32
	var negative bool

	if data[0] == '-' {
		data = data[1:]
		negative = true
	}

	for i := 0; i < len(data); i++ {
		num = (num << 3) + (num << 1)
		add := int32(data[i]) - '0'

		if num > math.MaxInt32-add {
			return 0, overflowErr
		}

		num += add
	}

	if negative {
		return -num, nil
	}

	return num, nil
}

func int32ToBytes(num int32) []byte {
	var data [32]byte
	var negative bool

	if num < 0 {
		negative = true
		num = 0 - num
	}

	cursor := 31

	for num >= 10 {
		v := num % 10
		num = num / 10
		data[cursor] = byte(v) + '0'
		cursor--
	}

	data[cursor] = byte(num) + '0'

	if negative {
		cursor--
		data[cursor] = '-'
	}

	return data[cursor:]
}

func bytesToInt64(data []byte) (int64, error) {
	var num int64
	var negative bool

	if data[0] == '-' {
		data = data[1:]
		negative = true
	}

	for i := 0; i < len(data); i++ {
		num = (num << 3) + (num << 1)
		add := int64(data[i]) - '0'

		if num > math.MaxInt64-add {
			return 0, overflowErr
		}

		num += add
	}

	if negative {
		return -num, nil
	}

	return num, nil
}

func int64ToBytes(num int64) []byte {
	var data [32]byte
	var negative bool

	if num < 0 {
		negative = true
		num = 0 - num
	}

	cursor := 31

	for num >= 10 {
		v := num % 10
		num = num / 10
		data[cursor] = byte(v) + '0'
		cursor--
	}

	data[cursor] = byte(num) + '0'

	if negative {
		cursor--
		data[cursor] = '-'
	}

	return data[cursor:]
}

func bytesToUint8(data []byte) (uint8, error) {
	var num uint8

	for i := 0; i < len(data); i++ {
		num = (num << 3) + (num << 1)
		add := uint8(data[i]) - '0'

		if num > math.MaxUint8-add {
			return 0, overflowErr
		}

		num += add
	}

	return num, nil
}

func uint8ToBytes(num uint8) []byte {
	var data [32]byte

	cursor := 31

	for num >= 10 {
		v := num % 10
		num = num / 10
		data[cursor] = byte(v) + '0'
		cursor--
	}

	data[cursor] = byte(num) + '0'

	return data[cursor:]
}

func bytesToUint16(data []byte) (uint16, error) {
	var num uint16

	for i := 0; i < len(data); i++ {
		num = (num << 3) + (num << 1)
		add := uint16(data[i]) - '0'

		if num > math.MaxUint16-add {
			return 0, overflowErr
		}

		num += add
	}

	return num, nil
}

func uint16ToBytes(num uint16) []byte {
	var data [32]byte

	cursor := 31

	for num >= 10 {
		v := num % 10
		num = num / 10
		data[cursor] = byte(v) + '0'
		cursor--
	}

	data[cursor] = byte(num) + '0'

	return data[cursor:]
}

func bytesToUint32(data []byte) (uint32, error) {
	var num uint32

	for i := 0; i < len(data); i++ {
		num = (num << 3) + (num << 1)
		add := uint32(data[i]) - '0'

		if num > math.MaxUint32-add {
			return 0, overflowErr
		}

		num += add
	}

	return num, nil
}

func uint32ToBytes(num uint32) []byte {
	var data [32]byte

	cursor := 31

	for num >= 10 {
		v := num % 10
		num = num / 10
		data[cursor] = byte(v) + '0'
		cursor--
	}

	data[cursor] = byte(num) + '0'

	return data[cursor:]
}

func bytesToUint64(data []byte) (uint64, error) {
	var num uint64

	for i := 0; i < len(data); i++ {
		num = (num << 3) + (num << 1)
		add := uint64(data[i]) - '0'

		if num > math.MaxUint64-add {
			return 0, overflowErr
		}

		num += add
	}

	return num, nil
}

func uint64ToBytes(num uint64) []byte {
	var data [32]byte

	cursor := 31

	for num >= 10 {
		v := num % 10
		num = num / 10
		data[cursor] = byte(v) + '0'
		cursor--
	}

	data[cursor] = byte(num) + '0'

	return data[cursor:]
}

func isEscapeHtmlByte(b byte) bool {
	switch b {
	case '<', '>', '&':
		return true

	default:
		return false
	}
}

func bytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

func isNumber(b byte) bool {
	return b > 0x2F && b < 0x3A
}

func isSkip(b byte) bool {
	return b == ' ' || b == '\n' || b == '\t' || b == '\r'
}

// IsDecoder checks whether obj implements json.Unmarshaler
func IsDecoder(obj interface{}) (json.Unmarshaler, bool) {
	decoder, ok := obj.(json.Unmarshaler)
	return decoder, ok
}

// IsEncoder checks whether obj implements json.Marshaler
func IsEncoder(obj interface{}) (json.Marshaler, bool) {
	encoder, ok := obj.(json.Marshaler)
	return encoder, ok
}
