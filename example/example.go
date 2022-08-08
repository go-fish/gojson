package example

import "github.com/go-fish/gojson/example/test"

type StringType string
type StringType1 StringType

// Test111
type ExampleStruct struct {
	TestString            *string                                   `json:"testString,omitempty"`
	TestInt               int16                                     `json:"testInt,omitempty"`
	TestFloat32           float32                                   `json:"testFloat32,omitempty"`
	TestBool              bool                                      `json:"testBool,omitempty"`
	TestBytes             []byte                                    `json:"testBytes,omitempty"`
	TestAlias             *StringType1                              `json:"testAlias,omitempty"`
	TestMap               map[string]map[StringType]AliasTypeStruct `json:"testMap,omitempty"`
	TestAliasMap          map[string]AliasMap                       `json:"testAliasMap,omitempty"`
	TestObject            map[StringType]interface{}                `json:"testObject,omitempty"`
	TestSlice             []StringType                              `json:"testSlice,omitempty"`
	TestArray             [4]*test.Test                             `json:"testArray,omitempty"`
	CommonTestStruct      TestStruct                                `json:"testStruct,omitempty"`
	CommonAliasTestStruct *AliasTypeStruct                          `json:"testAliasStruct"`
	TestPtr               *TestPtrStruct                            `json:"testPtr,omitempty"`
	TestSelector          *test.Test                                `json:"testSelector,omitempty"`
	*AliasTypeStruct      `json:",inline,omitempty"`
}

type AliasTypeStruct TestStruct

type AliasMap map[StringType]AliasTypeStruct

// +gojson: none
type TestStruct struct {
	*TestPtrStruct `json:",inline,omitempty"`
}

// +gojson: none
type TestPtrStruct struct {
	TestBoolStruct `json:"testBoolStruct,omitempty"`
	// TestString     string `json:"testStringStr,omitempty"`
}

// +gojson: none
type TestBoolStruct struct {
	TestBool bool `json:"testBool,omitempty"`
}
