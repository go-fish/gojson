package generate

import (
	"context"
	"fmt"
)

type floatType struct {
	biz int
}

var _ Underlying = new(floatType)

// NewFloatType returns the instance of floatType
func NewFloatType(biz int) Underlying {
	return &floatType{biz}
}

func (ft *floatType) Omitempty(value string) string {
	return fmt.Sprintf("if %s != 0", value)
}

func (ft *floatType) Type() string {
	if ft.biz == 0 {
		return "float64"
	}

	return fmt.Sprintf("float%d", ft.biz)
}

func (ft *floatType) GenerateDecode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	switch ft.biz {
	case 32:
		w.line("%s, err = dec.ReadFloat32()", argv)

	default:
		w.line("%s, err = dec.ReadFloat64()", argv)
	}

	w.line("if err != nil {")
	w.line("return err")
	w.line("}")
	w.line("")

	return nil
}

func (ft *floatType) GenerateEncode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	if IsPointer(ctx) {
		argv = "*" + argv
	}

	if IsAlias(ctx) {
		argv = fmt.Sprintf("%s(%s)", ft.Type(), argv)
	}

	switch ft.biz {
	case 32:
		w.line("enc.EncodeFloat32(%s)", argv)

	default:
		w.line("enc.EncodeFloat64(%s)", argv)
	}

	return nil
}
