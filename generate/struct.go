package generate

import (
	"context"
	"fmt"
	"strings"
)

type structType struct {
	name   string
	fields []Object
}

var _ Underlying = new(structType)

// NewStructType returns an struct object instance
func NewStructType(name string, fields ...Object) Underlying {
	return &structType{
		name:   name,
		fields: fields,
	}
}

func (st *structType) Omitempty(value string) string {
	return ""
}

func (st *structType) Type() string {
	return st.name
}

func (st *structType) GenerateDecode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	w.line("if isNil, err := dec.AssertObject(); err != nil {")
	w.line("return err")
	w.line("} else if !isNil {")
	w.line("for {")

	// read next key
	key := genRandomSuffix("key")
	w.line("%s, err := dec.ReadString()", key)
	w.line("if err != nil {")
	w.line("return err")
	w.line("}")
	w.line("")
	w.line("if err := dec.MustByte(':'); err != nil {")
	w.line("return err")
	w.line("}")
	w.line("")

	w.line("switch %s {", key)

	for _, field := range st.fields {
		w.line("case %q:", field.Meta().GetTagName())

		sargv := genRandomSuffix("value")
		ctx = ContextWithArgvName(context.TODO(), sargv)
		ctx = ContextWithEscape(ctx, field.Meta().GetEscape())

		// initialize parent if its pointer
		for parent := field.Meta().GetParent(); parent != nil; {
			if parent.Meta().GetPointer() {
				w.line("if %s.%s == nil {", argv, parent.Meta().GetName())
				w.line("%s.%s = new(%s)", argv, parent.Meta().GetName(), strings.TrimPrefix(parent.Type(), "*"))
				w.line("}")
				w.line("")
			}

			parent = parent.Meta().GetParent()
		}

		if err := field.GenerateDecode(ctx, w); err != nil {
			return err
		}

		w.line("%s.%s = %s", argv, field.Meta().GetName(), sargv)
		w.line("")
	}

	// set default case
	w.line("default:")
	w.line("if err := dec.SkipValue(); err != nil {")
	w.line("return err")
	w.line("}")
	w.line("}")

	// check whether object is closed
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

func (st *structType) GenerateEncode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	w.line("enc.WriteByte('{')")

	for id, field := range st.fields {
		sargv := argv + "." + field.Meta().GetName()
		ctx = ContextWithArgvName(ctx, sargv)

		omitemptyStr := field.Omitempty(sargv)

		if omitemptyStr != "" {
			w.line("%s {", omitemptyStr)
		}

		index := 0
		for parent := field.Meta().GetParent(); parent != nil; {
			if parent.Meta().GetPointer() {
				w.line("if %s.%s != nil {", argv, parent.Meta().GetName())
				index++
			}

			parent = parent.Meta().GetParent()
		}

		w.line("enc.EncodeKey(%q)", field.Meta().GetTagName())

		if err := field.GenerateEncode(ctx, w); err != nil {
			return err
		}

		for i := 0; i < index; i++ {
			w.line("}")
		}

		if id < len(st.fields)-1 {
			w.line("enc.WriteByte(',')")
		}

		if omitemptyStr != "" {
			w.line("}")
		}
	}

	w.line("enc.WriteByte('}')")

	return nil
}
