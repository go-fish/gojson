package generate

import (
	"context"
	"fmt"
)

type mapType struct {
	key   Object
	value Object
}

var _ Underlying = new(mapType)

// NewMapType returns the instance of mapType
func NewMapType(key, value Object) Underlying {
	return &mapType{
		key:   key,
		value: value,
	}
}

func (mt *mapType) Omitempty(value string) string {
	return fmt.Sprintf("if len(%s) > 0", value)
}

func (mt *mapType) Type() string {
	return fmt.Sprintf("map[%s]%s", mt.key.Type(), mt.value.Type())
}

func (mt *mapType) GenerateDecode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	w.line("if isNil, err := dec.AssertObject(); err != nil {")
	w.line("return err")
	w.line("} else if isNil {")
	w.line("%s = nil", argv)
	w.line("} else {")
	w.line("%s = make(%s)", argv, mt.Type())
	w.line("")

	key := genRandomSuffix("key")
	sargv := genRandomSuffix("value")
	w.line("")
	w.line("for {")

	// read key
	if err := mt.key.GenerateDecode(ContextWithArgvName(ctx, key), w); err != nil {
		return err
	}

	w.line("")
	w.line("if err := dec.MustByte(':'); err != nil {")
	w.line("return err")
	w.line("}")
	w.line("")

	// read value
	if err := mt.value.GenerateDecode(ContextWithArgvName(ctx, sargv), w); err != nil {
		return err
	}

	w.line("%s[%s] = %s", argv, key, sargv)
	w.line("")
	w.line("if dec.ExpectByte('}') {")
	w.line("break")
	w.line("}")
	w.line("")
	w.line("if err := dec.MustByte(','); err != nil {")
	w.line("return err")
	w.line("}")
	w.line("}")
	w.line("}")
	w.line("")

	return nil
}

func (mt *mapType) GenerateEncode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	if IsPointer(ctx) {
		argv = "*" + argv
	}

	if IsAlias(ctx) {
		argv = fmt.Sprintf("%s(%s)", mt.Type(), argv)
	}

	w.line("enc.WriteByte('{')")
	index := genRandomSuffix("index")
	w.line("%s := 0", index)
	w.line("for key, value := range %s {", argv)
	w.line("if %s > 0 {", index)
	w.line("enc.WriteByte(',')")
	w.line("}")
	w.line("")

	ctx = ContextWithArgvName(ctx, "key")

	if err := mt.key.GenerateEncode(ctx, w); err != nil {
		return err
	}

	w.line("enc.WriteByte(':')")

	ctx = ContextWithArgvName(ctx, "value")

	// check elem omitempty
	omitemptyStr := mt.value.Omitempty("value")
	if omitemptyStr != "" {
		w.line("%s {", omitemptyStr)
	}
	if err := mt.value.GenerateEncode(ctx, w); err != nil {
		return err
	}

	if omitemptyStr != "" {
		w.line("} else {")
		w.line("enc.WriteNull()")
		w.line("}")
	}

	w.line("%s++", index)
	w.line("}")
	w.line("enc.WriteByte('}')")
	w.line("")
	return nil
}
