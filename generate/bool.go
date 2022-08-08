package generate

import (
	"context"
	"fmt"
)

type boolType struct{}

var _ Underlying = new(boolType)

// NewBoolType returns the instance of boolType
func NewBoolType() Underlying {
	return new(boolType)
}

func (bt *boolType) Omitempty(value string) string {
	return fmt.Sprintf("if %s", value)
}

func (bt *boolType) Type() string {
	return "bool"
}

func (bt *boolType) GenerateDecode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	w.line("%s, err = dec.ReadBool()", argv)
	w.line("if err != nil {")
	w.line("return err")
	w.line("}")
	w.line("")

	return nil
}

func (bt *boolType) GenerateEncode(ctx context.Context, w writer) error {
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

	w.line("enc.EncodeBool(%s)", argv)
	return nil
}
