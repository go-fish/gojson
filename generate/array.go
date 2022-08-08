package generate

import (
	"context"
	"fmt"
)

type arrayType struct {
	cap  int
	elem Object
}

var _ Underlying = new(arrayType)

// NewArrayType returns the arrayType instance
func NewArrayType(cap int, elem Object) Underlying {
	return &arrayType{
		cap:  cap,
		elem: elem,
	}
}

func (at *arrayType) Omitempty(value string) string {
	return fmt.Sprintf("if len(%s) > 0", value)
}

func (at *arrayType) Type() string {
	if at.cap > 0 {
		return fmt.Sprintf("[%d]%s", at.cap, at.elem.Type())
	}

	return fmt.Sprintf("[]%s", at.elem.Type())
}

func (at *arrayType) GenerateDecode(ctx context.Context, w writer) error {
	if at.cap > 0 {
		return at.gArrayDecode(ctx, w)
	}

	return at.gSliceDecode(ctx, w)
}

func (at *arrayType) gArrayDecode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	sargv := genRandomSuffix("value")
	w.line("if isNil, err := dec.AssertArray(); err != nil {")
	w.line("return err")
	w.line("} else if !isNil {")
	w.line("for i := 0; i < %d; i++ {", at.cap)
	if err := at.elem.GenerateDecode(ContextWithArgvName(ctx, sargv), w); err != nil {
		return err
	}

	// append subvalue to value
	w.line("%s[i] = %s", argv, sargv)
	w.line("")
	w.line("if i < %d-1 {", at.cap)
	w.line("if err := dec.MustByte(','); err != nil {")
	w.line("return err")
	w.line("}")
	w.line("}")
	w.line("}")

	// check whether array is end
	w.line("if !dec.ExpectByte(']') {")
	w.line("return fmt.Errorf(\"index out of range when decode Array %s\")", at.Type())
	w.line("}")
	w.line("}")

	return nil
}

func (at *arrayType) gSliceDecode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	sargv := genRandomSuffix("value")
	w.line("if isNil, err := dec.AssertArray(); err != nil {")
	w.line("return err")
	w.line("} else if isNil {")
	w.line("%s = nil", argv)
	w.line("} else {")
	w.line("for {")

	if err := at.elem.GenerateDecode(ContextWithArgvName(ctx, sargv), w); err != nil {
		return err
	}

	// // append subvalue to value
	w.line("%s = append(%s, %s)", argv, argv, sargv)
	w.line("")

	// check whether array is end
	w.line("if dec.ExpectByte(']') {")
	w.line("break")
	w.line("}")
	w.line("")
	w.line("if err := dec.MustByte(','); err != nil {")
	w.line("return err")
	w.line("}")
	w.line("}")
	w.line("}")

	return nil
}

func (at *arrayType) GenerateEncode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	if IsPointer(ctx) {
		argv = "*" + argv
	}

	if IsAlias(ctx) {
		argv = fmt.Sprintf("%s(%s)", at.Type(), argv)
	}

	w.line("enc.WriteByte('[')")
	w.line("for index, value := range %s {", argv)
	w.line("if index > 0 {")
	w.line("enc.WriteByte(',')")
	w.line("}")
	w.line("")

	ctx = ContextWithArgvName(ctx, "value")

	// check elem omitempty
	omitemptyStr := at.elem.Omitempty("value")
	if omitemptyStr != "" {
		w.line("%s {", omitemptyStr)
	}

	if err := at.elem.GenerateEncode(ctx, w); err != nil {
		return err
	}

	if omitemptyStr != "" {
		w.line("} else {")
		w.line("enc.WriteNull()")
		w.line("}")
	}

	w.line("}")
	w.line("enc.WriteByte(']')")
	w.line("")
	return nil
}
