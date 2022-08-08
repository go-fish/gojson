package generate

import (
	"context"
	"fmt"
)

type interfaceType struct{}

var _ Underlying = new(interfaceType)

// NewInterfaceType returns the instance of interfaceType
func NewInterfaceType() Underlying {
	return new(interfaceType)
}

func (it *interfaceType) Omitempty(value string) string {
	return fmt.Sprintf("if %s != nil", value)
}

func (it *interfaceType) Type() string {
	return "interface{}"
}

func (it *interfaceType) GenerateDecode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	if IsPointer(ctx) {
		argv = "*" + argv
	}

	if IsAlias(ctx) {
		argv = fmt.Sprintf("%s(%s)", it.Type(), argv)
	}

	w.line("%s, err = dec.ReadValue()", argv)
	w.line("if err != nil {")
	w.line("return err")
	w.line("}")
	w.line("")

	return nil
}

func (it *interfaceType) GenerateEncode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	w.line("enc.EncodeInterface(%s)", argv)
	return nil
}
