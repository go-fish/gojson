package test

import "time"

type Test struct {
	TestInline `json:",inline,omitempty"`
	TestTime   *time.Time `json:"testTime,omitempty"`
}

type TestInline struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
