package generate

import (
	"context"
	"fmt"
)

type uintType struct {
	biz int
}

var _ Underlying = new(intType)

// NewUintType returns the instance of intType
func NewUintType(biz int) Underlying {
	return &uintType{biz}
}

func (ut *uintType) Omitempty(value string) string {
	return fmt.Sprintf("if %s != 0", value)
}

func (ut *uintType) Type() string {
	if ut.biz == 0 {
		return "uint"
	}

	return fmt.Sprintf("uint%d", ut.biz)
}

func (ut *uintType) GenerateDecode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	switch ut.biz {
	case 8:
		w.line("%s, err := d.ReadUint8()", argv)

	case 16:
		w.line("%s, err := d.ReadUint16()", argv)

	case 32:
		w.line("%s, err := d.ReadUint32()", argv)

	case 64:
		w.line("%s, err := d.ReadUint64()", argv)

	default:
		w.line("%s, err := d.ReadUint()", argv)
	}

	w.line("if err != nil {")
	w.line("return err")
	w.line("}")
	w.line("")
	return nil
}

func (ut *uintType) GenerateEncode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	if IsPointer(ctx) {
		argv = "*" + argv
	}

	if IsAlias(ctx) {
		argv = fmt.Sprintf("%s(%s)", ut.Type(), argv)
	}

	switch ut.biz {
	case 8:
		w.line("enc.EncodeUint8(%s)", argv)

	case 16:
		w.line("enc.EncodeUint16(%s)", argv)

	case 32:
		w.line("enc.EncodeUint32(%s)", argv)

	case 64:
		w.line("enc.EncodeUint64(%s)", argv)

	default:
		w.line("enc.EncodeUint(%s)", argv)
	}

	return nil
}
