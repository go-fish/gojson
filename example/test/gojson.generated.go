// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY go-fish/gojson.
// ************************************************************

package test

import (
	json "encoding/json"
	time "time"

	backend "github.com/go-fish/gojson/backend"
)

func (t *Test) UnmarshalJSON(data []byte) error {
	dec := backend.NewDecoder(data, backend.WithUnsafe())
	defer backend.ReleaseDecoder(dec)

	if isNil, err := dec.AssertObject(); err != nil {
		return err
	} else if !isNil {
		for {
			key1bjmjtfw, err := dec.ReadString()
			if err != nil {
				return err
			}

			if err := dec.MustByte(':'); err != nil {
				return err
			}

			switch key1bjmjtfw {
			case "createdAt":
				var valuewa3vtjyt time.Time

				data, err := dec.ReadValueBytes()
				if err != nil {
					return err
				}

				if decoder, ok := backend.IsDecoder(&valuewa3vtjyt); ok {
					if err := decoder.UnmarshalJSON(data); err != nil {
						return err
					}
				} else {
					if err := json.Unmarshal(data, &valuewa3vtjyt); err != nil {
						return err
					}
				}
				t.CreatedAt = valuewa3vtjyt

			case "testTime":
				var valueem1hv1mo *time.Time

				if dec.AssertNull() {
					valueem1hv1mo = nil
				} else {
					if valueem1hv1mo == nil {
						valueem1hv1mo = new(time.Time)
					}

					data, err := dec.ReadValueBytes()
					if err != nil {
						return err
					}

					if decoder, ok := backend.IsDecoder(valueem1hv1mo); ok {
						if err := decoder.UnmarshalJSON(data); err != nil {
							return err
						}
					} else {
						if err := json.Unmarshal(data, valueem1hv1mo); err != nil {
							return err
						}
					}
				}

				t.TestTime = valueem1hv1mo

			default:
				if err := dec.SkipValue(); err != nil {
					return err
				}
			}
			if dec.ExpectByte('}') {
				break
			}

			if err := dec.MustByte(','); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Test) MarshalJSON() ([]byte, error) {
	enc := backend.NewEncoder()
	defer backend.ReleaseEncoder(enc)

	enc.WriteByte('{')
	enc.EncodeKey("createdAt")
	value31xmswml := t.CreatedAt
	if encoder, ok := backend.IsEncoder(&value31xmswml); ok {
		data, err := encoder.MarshalJSON()
		if err != nil {
			return nil, err
		}

		enc.Write(data)
	} else {
		data, err := json.Marshal(value31xmswml)
		if err != nil {
			return nil, err
		}

		enc.Write(data)
	}

	enc.WriteByte(',')
	if t.TestTime != nil {
		enc.EncodeKey("testTime")
		value6g2ast6w := *t.TestTime
		if encoder, ok := backend.IsEncoder(&value6g2ast6w); ok {
			data, err := encoder.MarshalJSON()
			if err != nil {
				return nil, err
			}

			enc.Write(data)
		} else {
			data, err := json.Marshal(value6g2ast6w)
			if err != nil {
				return nil, err
			}

			enc.Write(data)
		}

	}
	enc.WriteByte('}')
	return enc.Bytes(), nil
}

func (t *TestInline) UnmarshalJSON(data []byte) error {
	dec := backend.NewDecoder(data, backend.WithUnsafe())
	defer backend.ReleaseDecoder(dec)

	if isNil, err := dec.AssertObject(); err != nil {
		return err
	} else if !isNil {
		for {
			keycjp67hq7, err := dec.ReadString()
			if err != nil {
				return err
			}

			if err := dec.MustByte(':'); err != nil {
				return err
			}

			switch keycjp67hq7 {
			case "createdAt":
				var valuev3ruhfur time.Time

				data, err := dec.ReadValueBytes()
				if err != nil {
					return err
				}

				if decoder, ok := backend.IsDecoder(&valuev3ruhfur); ok {
					if err := decoder.UnmarshalJSON(data); err != nil {
						return err
					}
				} else {
					if err := json.Unmarshal(data, &valuev3ruhfur); err != nil {
						return err
					}
				}
				t.CreatedAt = valuev3ruhfur

			default:
				if err := dec.SkipValue(); err != nil {
					return err
				}
			}
			if dec.ExpectByte('}') {
				break
			}

			if err := dec.MustByte(','); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *TestInline) MarshalJSON() ([]byte, error) {
	enc := backend.NewEncoder()
	defer backend.ReleaseEncoder(enc)

	enc.WriteByte('{')
	enc.EncodeKey("createdAt")
	valuek4r3blhp := t.CreatedAt
	if encoder, ok := backend.IsEncoder(&valuek4r3blhp); ok {
		data, err := encoder.MarshalJSON()
		if err != nil {
			return nil, err
		}

		enc.Write(data)
	} else {
		data, err := json.Marshal(valuek4r3blhp)
		if err != nil {
			return nil, err
		}

		enc.Write(data)
	}

	enc.WriteByte('}')
	return enc.Bytes(), nil
}