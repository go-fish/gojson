package generate

import (
	"context"
	"fmt"
)

type stringType struct{}

var _ Underlying = new(stringType)

// NewStringType returns the instance of stringType
func NewStringType() Underlying {
	return new(stringType)
}

func (st *stringType) Omitempty(value string) string {
	return fmt.Sprintf("if %s != \"\"", value)
}

func (st *stringType) Type() string {
	return "string"
}

func (st *stringType) GenerateDecode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	if IsEscape(ctx) {
		w.line("%s, err = dec.ReadStringWithEscape()", argv)
	} else {
		w.line("%s, err = dec.ReadString()", argv)
	}

	w.line("if err != nil {")
	w.line("return err")
	w.line("}")
	w.line("")
	return nil
}

func (st *stringType) GenerateEncode(ctx context.Context, w writer) error {
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

	w.line("enc.EncodeString(%s)", argv)
	return nil
}
