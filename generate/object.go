package generate

import (
	"context"
	"fmt"
	"strings"
)

// Object defines the basic type to generate code
type Object interface {
	Underlying

	Meta() MetaObject
	SetUnderlying(underlying Underlying)
	GetUnderlying() Underlying
}

// Underlying defines the underlying of object
type Underlying interface {
	Omitempty(value string) string
	Type() string
	GenerateDecode(ctx context.Context, w writer) error
	GenerateEncode(ctx context.Context, w writer) error
}

type object struct {
	metaObject MetaObject
	underlying Underlying
}

var _ Object = new(object)

// NewObject returns the object instance
func NewObject() Object {
	return &object{
		metaObject: NewMetaObject(),
	}
}

func (o *object) Meta() MetaObject {
	return o.metaObject
}

func (o *object) Omitempty(value string) string {
	if o.Meta().GetPointer() {
		return fmt.Sprintf("if %s != nil", value)
	}

	if !o.Meta().GetOmitempty() {
		return ""
	}

	return o.underlying.Omitempty(value)
}

func (o *object) Type() string {
	typ := o.GetUnderlying().Type()

	if alias := o.Meta().GetAlias(); len(alias) > 0 {
		typ = alias[0]
	}

	if o.Meta().GetPointer() {
		typ = "*" + typ
	}

	return typ
}

func (o *object) GenerateDecode(ctx context.Context, w writer) (err error) {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	fn := func() error {
		switch o.GetUnderlying().(type) {
		case *structType:
			argv = strings.TrimPrefix(argv, "*")
			ctx = ContextWithArgvName(ctx, argv)

			if err := o.GetUnderlying().GenerateDecode(ctx, w); err != nil {
				return err
			}

		default:
			if alias := o.Meta().GetAlias(); len(alias) > 0 {
				sargv := genRandomSuffix("argv")
				ctx = ContextWithArgvName(ctx, sargv)

				// initialize sargv
				w.line("var %s %s", sargv, o.GetUnderlying().Type())
				w.line("")

				if err := o.GetUnderlying().GenerateDecode(ctx, w); err != nil {
					return err
				}

				w.line("%s = %s", argv, convertAlias(sargv, o.Meta().GetAlias()...))
				return nil
			}

			return o.GetUnderlying().GenerateDecode(ctx, w)
		}

		return nil
	}

	if o.Meta().GetPointer() {
		w.line("var %s %s", argv, o.Type())
		w.line("")
		w.line("if dec.AssertNull() {")
		w.line("%s = nil", argv)
		w.line("} else {")
		w.line("if %s == nil {", argv)
		w.line("%s = new(%s)", argv, strings.TrimPrefix(o.Type(), "*"))

		w.line("}")
		w.line("")

		argv = "*" + argv
		ctx = ContextWithArgvName(ctx, argv)

		if err := fn(); err != nil {
			return err
		}

		w.line("}")
		w.line("")
		return nil
	}

	w.line("var %s %s", argv, o.Type())
	w.line("")
	return fn()
}

func (o *object) GenerateEncode(ctx context.Context, w writer) error {
	argv := ArgvNameInContext(ctx)
	if argv == "" {
		return fmt.Errorf("Missing ArgvName in Context")
	}

	ctx = ContextWithAlias(ctx, len(o.Meta().GetAlias()) > 0)
	ctx = ContextWithPointer(ctx, o.Meta().GetPointer())
	return o.GetUnderlying().GenerateEncode(ctx, w)
}

func (o *object) SetUnderlying(obj Underlying) {
	o.underlying = obj
}

func (o *object) GetUnderlying() Underlying {
	return o.underlying
}
