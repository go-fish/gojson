package generate

import (
	"context"
	"fmt"
)

type intType struct {
	biz int
}

var _ Underlying = new(intType)

// NewIntType returns the instance of intType
func NewIntType(biz int) Underlying {
	return &intType{biz}
}

func (it *intType) Omitempty(value string) string {
	return fmt.Sprintf("if %s != 0", value)
}

func (it *intType) Type() string {
	if it.biz == 0 {
		return "int"
	}

	return fmt.Sprintf("int%d", it.biz)
}

func (it *intType) GenerateDecode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	switch it.biz {
	case 8:
		w.line("%s, err = dec.ReadInt8()", argv)

	case 16:
		w.line("%s, err = dec.ReadInt16()", argv)

	case 32:
		w.line("%s, err = dec.ReadInt32()", argv)

	case 64:
		w.line("%s, err = dec.ReadInt64()", argv)

	default:
		w.line("%s, err = dec.ReadInt()", argv)
	}

	w.line("if err != nil {")
	w.line("return err")
	w.line("}")
	w.line("")

	return nil
}

func (it *intType) GenerateEncode(ctx context.Context, w writer) error {
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

	switch it.biz {
	case 8:
		w.line("enc.EncodeInt8(%s)", argv)

	case 16:
		w.line("enc.EncodeInt16(%s)", argv)

	case 32:
		w.line("enc.EncodeInt32(%s)", argv)

	case 64:
		w.line("enc.EncodeInt64(%s)", argv)

	default:
		w.line("enc.EncodeInt(%s)", argv)
	}

	return nil
}
