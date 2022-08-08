package backend

import (
	"encoding/json"
	"testing"

	"github.com/go-fish/gojson/backend/errors"
	"github.com/stretchr/testify/assert"
)

func TestDecodeStringBasic(t *testing.T) {
	dobj := NewDecoder([]byte(`"本书内容共 17 章"`), WithUnsafe())
	defer ReleaseDecoder(dobj)
	v, err := dobj.ReadStringWithEscape()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "本书内容共 17 章", v, "v must be equal to the value expected")
}

func TestDecoderStringComplex(t *testing.T) {
	dobj := NewDecoder([]byte(`  "string with spaces and \"escape\"d \"quotes\" and escaped line returns \\n and escaped \\\\ escaped char"`), WithUnsafe())
	defer ReleaseDecoder(dobj)
	v, err := dobj.ReadString()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, `string with spaces and "escape"d "quotes" and escaped line returns \n and escaped \\ escaped char`, v, "v is not equal to the value expected")
}

func TestDecoderStringNull(t *testing.T) {
	dobj := NewDecoder([]byte(`  null`), WithUnsafe())
	defer ReleaseDecoder(dobj)
	v, err := dobj.ReadString()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "", v, "v must be equal to ''")
}

func TestDecoderStringQuotaNull(t *testing.T) {
	dobj := NewDecoder([]byte("  \n\"null\""), WithUnsafe())
	defer ReleaseDecoder(dobj)

	v, err := dobj.ReadString()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, "null", v, "v must be equal to 'null'")
}

func TestDecoderStringInvalidJSON(t *testing.T) {
	dobj := NewDecoder([]byte(`  "invalid JSONs`), WithUnsafe())
	defer ReleaseDecoder(dobj)

	d := dobj.(*decoder)

	_, err := dobj.ReadString()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewParseError(d.length-1, "\"", d.data[d.length-1]), err, "err message must equal to the value expected")
}

func TestDecoderStringInvalidType(t *testing.T) {
	dobj := NewDecoder([]byte(`  123333`), WithUnsafe())
	defer ReleaseDecoder(dobj)

	_, err := dobj.ReadString()
	assert.NotNil(t, err, "Err must not be nil")
	assert.IsType(t, errors.NewParseError(3, `"`, '1'), err, "err message must euqal to the value expected")
}

func TestReadString(t *testing.T) {
	dobj := NewDecoder([]byte(`    "string with spaces and \"escape\"d \"quotes\" and escaped line returns \\n and escaped \\\\ escaped char"`), WithUnsafe())
	defer ReleaseDecoder(dobj)

	data, err := dobj.ReadString()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, `string with spaces and "escape"d "quotes" and escaped line returns \n and escaped \\ escaped char`, string(data), "data must be equal to the value expected")
}

func TestDecoderInt8Basic(t *testing.T) {
	dobj := NewDecoder([]byte("127"))
	defer ReleaseDecoder(dobj)

	v, err := dobj.ReadInt8()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int8(127), v, "v must be equal to the value expected")
}

func TestDecoderNegativeInt8Basic(t *testing.T) {
	dobj := NewDecoder([]byte("-127"))
	defer ReleaseDecoder(dobj)

	v, err := dobj.ReadInt8()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int8(-127), v, "v must be equal to the value expected")
}

func TestDecoderInt8Overflow(t *testing.T) {
	dobj := NewDecoder([]byte("128"))
	defer ReleaseDecoder(dobj)

	_, err := dobj.ReadInt8()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewOverflowError(0, []byte("128"), "int8"), err, "err must be equal to the value expected")
}

func TestDecoderNegativeInt8Overflow(t *testing.T) {
	dobj := NewDecoder([]byte("-128"))
	defer ReleaseDecoder(dobj)

	_, err := dobj.ReadInt8()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewOverflowError(0, []byte("-128"), "int8"), err, "err must be equal to the value expected")
}

func TestDecoderInt8Null(t *testing.T) {
	dobj := NewDecoder([]byte("null"))

	v, err := dobj.ReadInt8()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int8(0), v, "v must be equal to the value expected")
}

func TestDecoderInt8(t *testing.T) {
	dobj := NewDecoder([]byte(" \n\n 127"))

	v, err := dobj.ReadInt8()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int8(127), v, "v must be equal to the value expected")
}
func TestDecoderInt16Basic(t *testing.T) {
	dobj := NewDecoder([]byte("32767"))

	v, err := dobj.ReadInt16()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int16(32767), v, "v must be equal to the value expected")
}

func TestDecoderNegativeInt16Basic(t *testing.T) {
	dobj := NewDecoder([]byte("-32767"))

	v, err := dobj.ReadInt16()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int16(-32767), v, "v must be equal to the value expected")
}

func TestDecoderInt16Overflow(t *testing.T) {
	dobj := NewDecoder([]byte("32768"))

	_, err := dobj.ReadInt16()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewOverflowError(0, []byte("32768"), "int16"), err, "err must be equal to the value expected")
}

func TestDecoderNegativeInt16Overflow(t *testing.T) {
	dobj := NewDecoder([]byte("-32768"))

	_, err := dobj.ReadInt16()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewOverflowError(0, []byte("-32768"), "int16"), err, "err must be equal to the value expected")
}

func TestDecoderInt16Null(t *testing.T) {
	dobj := NewDecoder([]byte("null"))

	v, err := dobj.ReadInt16()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int16(0), v, "v must be equal to the value expected")
}

func TestDecoderInt16(t *testing.T) {
	dobj := NewDecoder([]byte(" \n\n -32767"))

	v, err := dobj.ReadInt16()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int16(-32767), v, "v must be equal to the value expected")
}

func TestDecoderInt32Basic(t *testing.T) {
	dobj := NewDecoder([]byte("2147483647"))

	v, err := dobj.ReadInt32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int32(2147483647), v, "v must be equal to the value expected")
}

func TestDecoderNegativeInt32Basic(t *testing.T) {
	dobj := NewDecoder([]byte("-2147483647"))

	v, err := dobj.ReadInt32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int32(-2147483647), v, "v must be equal to the value expected")
}

func TestDecoderInt32Overflow(t *testing.T) {
	dobj := NewDecoder([]byte("2147483648"))

	_, err := dobj.ReadInt32()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewOverflowError(0, []byte("2147483648"), "int32"), err, "err must be equal to the value expected")
}

func TestDecoderNegativeInt32Overflow(t *testing.T) {
	dobj := NewDecoder([]byte("-2147483648"))

	_, err := dobj.ReadInt32()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewOverflowError(0, []byte("-2147483648"), "int32"), err, "err must be equal to the value expected")
}

func TestDecoderInt32Null(t *testing.T) {
	dobj := NewDecoder([]byte("null"))

	v, err := dobj.ReadInt32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int32(0), v, "v must be equal to the value expected")
}

func TestDecoderInt32(t *testing.T) {
	dobj := NewDecoder([]byte(" \n\n -2147483647"))

	v, err := dobj.ReadInt32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int32(-2147483647), v, "v must be equal to the value expected")
}

func TestDecoderInt64Basic(t *testing.T) {
	dobj := NewDecoder([]byte("9223372036854775807"))

	v, err := dobj.ReadInt64()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int64(9223372036854775807), v, "v must be equal to the value expected")
}

func TestDecoderNegativeInt64Basic(t *testing.T) {
	dobj := NewDecoder([]byte("-9223372036854775807"))

	v, err := dobj.ReadInt64()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int64(-9223372036854775807), v, "v must be equal to the value expected")
}

func TestDecoderInt64Overflow(t *testing.T) {
	dobj := NewDecoder([]byte("9223372036854775808"))

	_, err := dobj.ReadInt64()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewOverflowError(0, []byte("9223372036854775808"), "int64"), err, "err must be equal to the value expected")
}

func TestDecoderNegativeInt64Overflow(t *testing.T) {
	dobj := NewDecoder([]byte("-9223372036854775808"))

	_, err := dobj.ReadInt64()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewOverflowError(0, []byte("-9223372036854775808"), "int64"), err, "err must be equal to the value expected")
}

func TestDecoderInt64Null(t *testing.T) {
	dobj := NewDecoder([]byte("null"))

	v, err := dobj.ReadInt64()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int64(0), v, "v must be equal to the value expected")
}

func TestDecoderInt64(t *testing.T) {
	dobj := NewDecoder([]byte(" \n\n -9223372036854775807"))

	v, err := dobj.ReadInt64()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, int64(-9223372036854775807), v, "v must be equal to the value expected")
}

func TestDecoderUint8Basic(t *testing.T) {
	dobj := NewDecoder([]byte("255"))

	v, err := dobj.ReadUint8()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint8(255), v, "v must be equal to the value expected")
}

func TestDecoderUint8Overflow(t *testing.T) {
	dobj := NewDecoder([]byte("256"))

	_, err := dobj.ReadUint8()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewOverflowError(0, []byte("256"), "uint8"), err, "err must be equal to the value expected")
}

func TestDecoderUint8Null(t *testing.T) {
	dobj := NewDecoder([]byte("null"))

	v, err := dobj.ReadUint8()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint8(0), v, "v must be equal to the value expected")
}

func TestDecoderUint8(t *testing.T) {
	dobj := NewDecoder([]byte(" \n\n 255"))

	v, err := dobj.ReadUint8()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint8(255), v, "v must be equal to the value expected")
}
func TestDecoderUint16Basic(t *testing.T) {
	dobj := NewDecoder([]byte("65535"))

	v, err := dobj.ReadUint16()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint16(65535), v, "v must be equal to the value expected")
}

func TestDecoderUint16Overflow(t *testing.T) {
	dobj := NewDecoder([]byte("65536"))

	_, err := dobj.ReadUint16()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewOverflowError(0, []byte("65536"), "uint16"), err, "err must be equal to the value expected")
}

func TestDecoderUint16Null(t *testing.T) {
	dobj := NewDecoder([]byte("null"))

	v, err := dobj.ReadUint16()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint16(0), v, "v must be equal to the value expected")
}

func TestDecoderUint16(t *testing.T) {
	dobj := NewDecoder([]byte(" \n\n 65535"))

	v, err := dobj.ReadUint16()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint16(65535), v, "v must be equal to the value expected")
}

func TestDecoderUint32Basic(t *testing.T) {
	dobj := NewDecoder([]byte("4294967295"))

	v, err := dobj.ReadUint32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint32(4294967295), v, "v must be equal to the value expected")
}

func TestDecoderUint32Overflow(t *testing.T) {
	dobj := NewDecoder([]byte("4294967296"))

	_, err := dobj.ReadUint32()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewOverflowError(0, []byte("4294967296"), "uint32"), err, "err must be equal to the value expected")
}

func TestDecoderUint32Null(t *testing.T) {
	dobj := NewDecoder([]byte("null"))

	v, err := dobj.ReadUint32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint32(0), v, "v must be equal to the value expected")
}

func TestDecoderUint32(t *testing.T) {
	dobj := NewDecoder([]byte(" \n\n 4294967295"))

	v, err := dobj.ReadUint32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint32(4294967295), v, "v must be equal to the value expected")
}

func TestDecoderUint64Basic(t *testing.T) {
	dobj := NewDecoder([]byte("18446744073709551615"))

	v, err := dobj.ReadUint64()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint64(18446744073709551615), v, "v must be equal to the value expected")
}

func TestDecoderUint64Overflow(t *testing.T) {
	dobj := NewDecoder([]byte("18446744073709551616"))

	_, err := dobj.ReadUint64()
	assert.NotNil(t, err, "Err must not be nil")
	assert.Equal(t, errors.NewOverflowError(0, []byte("18446744073709551616"), "uint64"), err, "err must be equal to the value expected")
}

func TestDecoderUint64Null(t *testing.T) {
	dobj := NewDecoder([]byte("null"))

	v, err := dobj.ReadUint64()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint64(0), v, "v must be equal to the value expected")
}

func TestDecoderUint64(t *testing.T) {
	dobj := NewDecoder([]byte(" \n\n 18446744073709551615"))

	v, err := dobj.ReadUint64()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, uint64(18446744073709551615), v, "v must be equal to the value expected")
}

func TestDecoderFloat32Basic(t *testing.T) {
	dobj := NewDecoder([]byte("127.11"))

	v, err := dobj.ReadFloat32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, float32(127.11), v, "v must be equal to the value expected")
}

func TestDecoderNegativeFloat32Basic(t *testing.T) {
	dobj := NewDecoder([]byte("-127.11"))

	v, err := dobj.ReadFloat32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, float32(-127.11), v, "v must be equal to the value expected")
}

func TestDecoderFloat32Null(t *testing.T) {
	dobj := NewDecoder([]byte("null"))

	v, err := dobj.ReadFloat32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, float32(0), v, "v must be equal to the value expected")
}

func TestDecoderFloat32(t *testing.T) {
	dobj := NewDecoder([]byte(" \n\n 127"))

	v, err := dobj.ReadFloat32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, float32(127), v, "v must be equal to the value expected")
}

func TestDecoderFloat32E(t *testing.T) {
	dobj := NewDecoder([]byte(" \n\n 1.27e+7"))

	v, err := dobj.ReadFloat32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, float32(1.27e+7), v, "v must be equal to the value expected")
}

func TestDecoderFloat32EE(t *testing.T) {
	dobj := NewDecoder([]byte(" \n\n 1.27E-7"))

	v, err := dobj.ReadFloat32()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, float32(1.27e-7), v, "v must be equal to the value expected")
}

func TestDecodeBytes(t *testing.T) {
	var data = []byte("Test123")

	info, err := json.Marshal(data)
	assert.Nil(t, err, "Err must be nil")

	decoder := NewDecoder(info)

	v, err := decoder.ReadBytes()
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, data, v, "v must be equal to the value expected")
}

func TestDecoderStringSliceBasic(t *testing.T) {
	decoder := NewDecoder([]byte(`["TestString1","TestString2"]`))

	testArr, err := decoder.ReadArray()
	assert.Nil(t, err, "Err must be nil")
	assert.Len(t, testArr, 2, "testArr should be of len 2")
	assert.Equal(t, "TestString1", testArr[0], "testArr[0] must be equal to the value expected")
	assert.Equal(t, "TestString2", testArr[1], "testArr[1] must be equal to the value expected")
}

func TestDecoderTwoDimensionalStringSliceBasic(t *testing.T) {
	decoder := NewDecoder([]byte(`[["TestString1","TestString2"],["TestString3", "TestString4"]]`))

	testArr, err := decoder.ReadArray()
	assert.Nil(t, err, "Err must be nil")
	assert.Len(t, testArr, 2, "testArr should be of len 2")
	assert.Equal(t, []interface{}{"TestString1", "TestString2"}, testArr[0], "testArr[0] must be equal to the value expected")
	assert.Equal(t, []interface{}{"TestString3", "TestString4"}, testArr[1], "testArr[1] must be equal to the value expected")
}

func TestDecoderTwoDimensionalInterfaceSliceBasic(t *testing.T) {
	decoder := NewDecoder([]byte(`[["TestString1","TestString2"],[123, 456],[123.123,456.456]]`))

	testArr, err := decoder.ReadArray()
	assert.Nil(t, err, "Err must be nil")
	assert.Len(t, testArr, 3, "testArr should be of len 2")
	assert.Equal(t, []interface{}{"TestString1", "TestString2"}, testArr[0], "testArr[0] must be equal to the value expected")
	assert.Equal(t, []interface{}{float64(123), float64(456)}, testArr[1], "testArr[1] must be equal to the value expected")
	assert.Equal(t, []interface{}{123.123, 456.456}, testArr[2], "testArr[2] must be equal to the value expected")
}

func TestSkipArray(t *testing.T) {
	decoder := NewDecoder([]byte(`[["Test[String1]","TestString2"],[123, 456],[123.123,456.456]]]`))

	err := decoder.SkipArray()
	assert.Nil(t, err, "Err must be nil")
}

func TestDecodeObject(t *testing.T) {
	decoder := NewDecoder([]byte(`   {"ke\"y1\"": [["TestString1","TestString2"],[123, 456],[123.123,456.456]], "key2": {"key3": "1111", "key4":123.111}}`))

	testMap, err := decoder.ReadObject()
	assert.Nil(t, err, "Err must be nil")
	assert.Len(t, testMap, 2, "testMap should be of len 2")
	assert.Len(t, testMap["ke\"y1\""], 3, "testMap[key1] should be of len 3")

	testArr := testMap["ke\"y1\""].([]interface{})
	assert.Equal(t, []interface{}{"TestString1", "TestString2"}, testArr[0], "testArr[0] must be equal to the value expected")
	assert.Equal(t, []interface{}{float64(123), float64(456)}, testArr[1], "testArr[1] must be equal to the value expected")
	assert.Equal(t, []interface{}{123.123, 456.456}, testArr[2], "testArr[2] must be equal to the value expected")

	testMap2 := testMap["key2"].(map[string]interface{})
	assert.Equal(t, "1111", testMap2["key3"], "testMap2[key3] must be equal to the value expected")
	assert.Equal(t, 123.111, testMap2["key4"], "testMap2[key4] must be equal to the value expected")
}

func TestSkipObject(t *testing.T) {
	decoder := NewDecoder([]byte(`   {"ke\"y1\"": [["TestString1","TestString2"],[123, 456],[123.123,456.456]], "key2": {"key3": "1111", "key4":123.111}}`))

	err := decoder.SkipObject()
	assert.Nil(t, err, "Err must be nil")
}

func Test_decoder_AssertObject(t *testing.T) {
	type fields struct {
		decoderOption decoderOption
		data          []byte
		cursor        int
		length        int
	}
	tests := []struct {
		name      string
		fields    fields
		wantIsNil bool
		err       error
	}{
		{
			name: "3121",
			fields: fields{
				data:   []byte("{ }"),
				cursor: 0,
				length: 3,
			},
			wantIsNil: true,
			err:       nil,
		},
		{
			name: "2222",
			fields: fields{
				data:   []byte("null"),
				cursor: 0,
				length: 4,
			},
			wantIsNil: true,
			err:       nil,
		},
		{
			name: "333",
			fields: fields{
				data:   []byte("a"),
				cursor: 0,
				length: 4,
			},
			wantIsNil: false,
			err:       errors.NewParseError(0, "n or {", 'a'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &decoder{
				decoderOption: tt.fields.decoderOption,
				data:          tt.fields.data,
				cursor:        tt.fields.cursor,
				length:        tt.fields.length,
			}
			if gotIsNil, err := d.AssertObject(); gotIsNil != tt.wantIsNil || err != tt.err {
				t.Errorf("decoder.AssertNilObject() = %v, want %v", err, tt.err)
			}
		})
	}
}

func TestReadStringAndEscape(t *testing.T) {
	readString := func() (string, error) {
		dobj := NewDecoder([]byte(`"b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcd""`), WithUnsafe())
		defer ReleaseDecoder(dobj)
		return dobj.ReadString()
	}

	readStringEscape := func() (string, error) {
		dobj := NewDecoder([]byte(`"b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcd""`), WithUnsafe())
		defer ReleaseDecoder(dobj)
		return dobj.ReadStringWithEscape()
	}

	str1, err := readString()
	assert.Nil(t, err, "Err must be nil")

	str2, err := readStringEscape()
	assert.Nil(t, err, "Err must be nil")

	assert.Equal(t, str1, str2, "str1 must equal to str2")
}

func BenchmarkReadString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//dobj := NewDecoder([]byte(`"b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcd""`), WithUnsafe())
		dobj := NewDecoder([]byte(`"本书内容共 17 章，由浅入深地讲解 Rust 相关的知识，涉及基础语法、软件包管理器、测试工具、类型系统、内存管理、异常处理、高级类型、并发模型、宏、外部函数接口、网络编程、 HTTP、数据库、 WebAssembly、 GTK+框架和 GDB 调试等重要知识点。本书适合想学习 Rust 编程的读者阅读，希望读者能够对 C、 C++或者 Python 有一些了解。书中丰富的代码示例和详细的讲解能够帮助读者快速上手，高效率掌握 Rust 编程。我们知道通过Linux的network namespace技术可以自定义一个独立的网络栈，简单到只有loopback设备，复杂到具备系统完整的网络能力，这就使得network namespace成为Linux网络虚拟化技术的基石——不论是虚拟机还是容器时代。network namespace的另一个隔离功能在于，系统管理员一旦禁用namespace中的网络设备，即使里面的进程拿到了一些系统特权，也无法和外界通信。最后，网络对安全较为敏感，即使network namespace能够提供网络资源隔离的机制，用户还是会结合其他类型的namespace一起使用，以提供更好的安全隔离能力。"`), WithUnsafe())
		dobj.ReadString()
		ReleaseDecoder(dobj)
	}
}

func BenchmarkReadStringWithEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//dobj := NewDecoder([]byte(`"b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcd"`), WithUnsafe())
		dobj := NewDecoder([]byte(`"本书内容共 17 章，由浅入深地讲解 Rust 相关的知识，涉及基础语法、软件包管理器、测试工具、类型系统、内存管理、异常处理、高级类型、并发模型、宏、外部函数接口、网络编程、 HTTP、数据库、 WebAssembly、 GTK+框架和 GDB 调试等重要知识点。本书适合想学习 Rust 编程的读者阅读，希望读者能够对 C、 C++或者 Python 有一些了解。书中丰富的代码示例和详细的讲解能够帮助读者快速上手，高效率掌握 Rust 编程。我们知道通过Linux的network namespace技术可以自定义一个独立的网络栈，简单到只有loopback设备，复杂到具备系统完整的网络能力，这就使得network namespace成为Linux网络虚拟化技术的基石——不论是虚拟机还是容器时代。network namespace的另一个隔离功能在于，系统管理员一旦禁用namespace中的网络设备，即使里面的进程拿到了一些系统特权，也无法和外界通信。最后，网络对安全较为敏感，即使network namespace能够提供网络资源隔离的机制，用户还是会结合其他类型的namespace一起使用，以提供更好的安全隔离能力。"`), WithUnsafe())
		dobj.ReadStringWithEscape()
		ReleaseDecoder(dobj)
	}
}
