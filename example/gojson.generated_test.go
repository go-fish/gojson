package example

import (
	"testing"
	"time"

	"github.com/go-fish/gojson/example/test"
	"github.com/stretchr/testify/assert"
)

func stringPtr(str string) *string {
	return &str
}

func timePtr() *time.Time {
	t := time.Now()
	return &t
}

func TestExampleStruct_MarshalAndUnmarshalJSON(t *testing.T) {
	now := timePtr()

	type fields struct {
		TestString            *string
		TestInt               int16
		TestFloat32           float32
		TestBool              bool
		TestBytes             []byte
		TestAlias             *StringType1
		TestMap               map[string]map[StringType]AliasTypeStruct
		TestAliasMap          map[string]AliasMap
		TestObject            map[StringType]interface{}
		TestSlice             []StringType
		TestArray             [4]*test.Test
		CommonTestStruct      TestStruct
		CommonAliasTestStruct *AliasTypeStruct
		TestPtr               *TestPtrStruct
		TestSelector          *test.Test
		AliasTypeStruct       *AliasTypeStruct
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				TestString:  stringPtr("testString"),
				TestInt:     123,
				TestFloat32: 123.123,
				TestBool:    true,
				TestBytes:   []byte("test"),
				TestAlias: func() *StringType1 {
					str := StringType1(StringType("test"))
					return &str
				}(),
				TestMap: map[string]map[StringType]AliasTypeStruct{
					"test": map[StringType]AliasTypeStruct{
						"test": AliasTypeStruct{
							&TestPtrStruct{
								TestBoolStruct: TestBoolStruct{true},
							},
						},
					},
				},
				TestAliasMap: map[string]AliasMap{
					"test": AliasMap{
						"test": AliasTypeStruct{
							&TestPtrStruct{
								TestBoolStruct: TestBoolStruct{true},
							},
						},
					},
				},
				TestObject: map[StringType]interface{}{
					"test": "test",
				},
				TestSlice: []StringType{"test1", "test2"},
				TestArray: [4]*test.Test{
					&test.Test{
						TestTime: now,
					},
				},
				// CommonTestStruct: TestStruct{},
				// CommonAliasTestStruct: &AliasTypeStruct{},
				TestPtr: &TestPtrStruct{
					TestBoolStruct: TestBoolStruct{true},
				},
				TestSelector: &test.Test{
					TestTime: now,
				},
				AliasTypeStruct: &AliasTypeStruct{
					&TestPtrStruct{
						TestBoolStruct: TestBoolStruct{true},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ExampleStruct{
				TestString:            tt.fields.TestString,
				TestInt:               tt.fields.TestInt,
				TestFloat32:           tt.fields.TestFloat32,
				TestBool:              tt.fields.TestBool,
				TestBytes:             tt.fields.TestBytes,
				TestAlias:             tt.fields.TestAlias,
				TestMap:               tt.fields.TestMap,
				TestAliasMap:          tt.fields.TestAliasMap,
				TestObject:            tt.fields.TestObject,
				TestSlice:             tt.fields.TestSlice,
				TestArray:             tt.fields.TestArray,
				CommonTestStruct:      tt.fields.CommonTestStruct,
				CommonAliasTestStruct: tt.fields.CommonAliasTestStruct,
				TestPtr:               tt.fields.TestPtr,
				TestSelector:          tt.fields.TestSelector,
				AliasTypeStruct:       tt.fields.AliasTypeStruct,
			}
			got, err := e.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ExampleStruct.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var e1 ExampleStruct

			if err := e1.UnmarshalJSON(got); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, &e1, e, "Marshal must equal to Unmarshal object")
		})
	}
}
