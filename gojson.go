package gojson

import (
	"encoding/json"

	"github.com/go-fish/gojson/backend"
)

// Unmarshal decode the json info to object
func Unmarshal(data []byte, v interface{}) error {
	switch t := v.(type) {
	case *string:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadString()
		if err != nil {
			return err
		}

		*t = v

	case *int:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadInt()
		if err != nil {
			return err
		}

		*t = v

	case *int8:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadInt8()
		if err != nil {
			return err
		}

		*t = v

	case *int16:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadInt16()
		if err != nil {
			return err
		}

		*t = v

	case *int32:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadInt32()
		if err != nil {
			return err
		}

		*t = v

	case *int64:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadInt64()
		if err != nil {
			return err
		}

		*t = v

	case *uint:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadUint()
		if err != nil {
			return err
		}

		*t = v

	case *uint8:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadUint8()
		if err != nil {
			return err
		}

		*t = v

	case *uint16:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadUint16()
		if err != nil {
			return err
		}

		*t = v

	case *uint32:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadUint32()
		if err != nil {
			return err
		}

		*t = v

	case *uint64:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadUint64()
		if err != nil {
			return err
		}

		*t = v

	case *[]byte:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadBytes()
		if err != nil {
			return err
		}

		*t = v

	case []byte:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadBytes()
		if err != nil {
			return err
		}

		t = v

	case *map[string]interface{}:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadObject()
		if err != nil {
			return err
		}

		*t = v

	case map[string]interface{}:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadObject()
		if err != nil {
			return err
		}

		t = v

	case *[]interface{}:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadArray()
		if err != nil {
			return err
		}

		*t = v

	case []interface{}:
		dec := backend.NewDecoder(data)
		defer backend.ReleaseDecoder(dec)

		v, err := dec.ReadArray()
		if err != nil {
			return err
		}

		t = v

	case json.Unmarshaler:
		return t.UnmarshalJSON(data)

	default:
		return json.Unmarshal(data, v)
	}

	return nil
}

// Marshal encode the object to json info
func Marshal(v interface{}) ([]byte, error) {
	switch t := v.(type) {
	case string:
		enc := backend.NewEncoder()
		defer backend.ReleaseEncoder(enc)

		enc.EncodeString(t)
		return enc.Bytes(), nil

	case int:
		enc := backend.NewEncoder()
		defer backend.ReleaseEncoder(enc)

		enc.EncodeInt(t)
		return enc.Bytes(), nil

	case int8:
		enc := backend.NewEncoder()
		defer backend.ReleaseEncoder(enc)

		enc.EncodeInt8(t)
		return enc.Bytes(), nil

	case int16:
		enc := backend.NewEncoder()
		defer backend.ReleaseEncoder(enc)

		enc.EncodeInt16(t)
		return enc.Bytes(), nil

	case int32:
		enc := backend.NewEncoder()
		defer backend.ReleaseEncoder(enc)

		enc.EncodeInt32(t)
		return enc.Bytes(), nil

	case int64:
		enc := backend.NewEncoder()
		defer backend.ReleaseEncoder(enc)

		enc.EncodeInt64(t)
		return enc.Bytes(), nil

	case uint:
		enc := backend.NewEncoder()
		defer backend.ReleaseEncoder(enc)

		enc.EncodeUint(t)
		return enc.Bytes(), nil

	case uint8:
		enc := backend.NewEncoder()
		defer backend.ReleaseEncoder(enc)

		enc.EncodeUint8(t)
		return enc.Bytes(), nil

	case uint16:
		enc := backend.NewEncoder()
		defer backend.ReleaseEncoder(enc)

		enc.EncodeUint16(t)
		return enc.Bytes(), nil

	case uint32:
		enc := backend.NewEncoder()
		defer backend.ReleaseEncoder(enc)

		enc.EncodeUint32(t)
		return enc.Bytes(), nil

	case uint64:
		enc := backend.NewEncoder()
		defer backend.ReleaseEncoder(enc)

		enc.EncodeUint64(t)
		return enc.Bytes(), nil

	case []byte:
		enc := backend.NewEncoder()
		defer backend.ReleaseEncoder(enc)

		enc.EncodeBytes(t)
		return enc.Bytes(), nil

	case map[string]interface{}:
		enc := backend.NewEncoder()
		defer backend.ReleaseEncoder(enc)

		enc.EncodeInterfaceMap(t)
		return enc.Bytes(), nil

	case []interface{}:
		enc := backend.NewEncoder()
		defer backend.ReleaseEncoder(enc)

		enc.EncodeInterfaceArray(t)
		return enc.Bytes(), nil

	case json.Marshaler:
		return t.MarshalJSON()

	default:
		return json.Marshal(v)
	}
}
