package generate

import (
	"context"
	"fmt"
	"strings"
)

type selectorType struct {
	typ string
}

var _ Underlying = new(selectorType)

// NewSelectorType returns the instance of selectorType
func NewSelectorType(selector string) Underlying {
	return &selectorType{selector}
}

func (st *selectorType) Omitempty(value string) string {
	return ""
}

func (st *selectorType) Type() string {
	return st.typ
}

func (st *selectorType) GenerateDecode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	if strings.HasPrefix(argv, "*") {
		argv = strings.TrimPrefix(argv, "*")
	} else {
		argv = "&" + argv
	}

	w.line("data, err := dec.ReadValueBytes()")
	w.line("if err != nil {")
	w.line("return err")
	w.line("}")
	w.line("")

	w.line("if decoder, ok := backend.IsDecoder(%s); ok {", argv)
	w.line("if err := decoder.UnmarshalJSON(data); err != nil {")
	w.line("return err")
	w.line("}")
	w.line("} else {")
	w.line("if err := json.Unmarshal(data, %s); err != nil {", argv)
	w.line("return err")
	w.line("}")
	w.line("}")
	return nil
}

func (st *selectorType) GenerateEncode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	if IsPointer(ctx) {
		argv = "*" + argv
	}

	if IsAlias(ctx) {
		argv = fmt.Sprintf("%s(%s)", st.Type(), argv)
	}

	value := genRandomSuffix("value")
	w.line("%s := %s", value, argv)
	w.line("if encoder, ok := backend.IsEncoder(&%s); ok {", value)
	w.line("data, err := encoder.MarshalJSON()")
	w.line("if err != nil {")
	w.line("return nil, err")
	w.line("}")
	w.line("")
	w.line("enc.Write(data)")
	w.line("} else {")
	w.line("data, err := json.Marshal(%s)", value)
	w.line("if err != nil {")
	w.line("return nil, err")
	w.line("}")
	w.line("")
	w.line("enc.Write(data)")
	w.line("}")
	w.line("")
	return nil
}
