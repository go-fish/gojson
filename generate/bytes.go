package generate

import (
	"context"
	"fmt"
)

type bytesType struct{}

var _ Underlying = new(bytesType)

// NewBytesType returns the instance of bytesType
func NewBytesType() Underlying {
	return new(bytesType)
}

func (bt *bytesType) Omitempty(value string) string {
	return fmt.Sprintf("if len(%s) > 0", value)
}

func (bt *bytesType) Type() string {
	return "[]byte"
}

func (bt *bytesType) GenerateDecode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	w.line("%s, err = dec.ReadBytes()", argv)
	w.line("if err != nil {")
	w.line("return err")
	w.line("}")
	w.line("")

	return nil
}

func (bt *bytesType) GenerateEncode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	if IsPointer(ctx) {
		argv = "*" + argv
	}

	if IsAlias(ctx) {
		argv = fmt.Sprintf("%s(%s)", bt.Type(), argv)
	}

	w.line("enc.EncodeBytes(%s)", argv)
	return nil
}
