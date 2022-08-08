package parser

import (
	"context"
	"fmt"
	"go/ast"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-fish/gojson/generate"
	"github.com/go-fish/gojson/log"
)

// Package describe the package info
type Package struct {
	pkg *ast.Package

	Name    string
	Path    string
	Imports map[string]string
	Objects []generate.Object
}

// NewPackage returns the package info according to ast Package
func NewPackage(path string, pkg *ast.Package) *Package {
	return &Package{
		pkg:     pkg,
		Name:    pkg.Name,
		Path:    path,
		Imports: make(map[string]string),
		Objects: make([]generate.Object, 0, 8),
	}
}

// Parse convert the ast.Package to Package
func (p *Package) Parse(ctx context.Context) error {
	var err error

	log.Infof("Parsing package %s", p.Name)
	defer log.ErrorOrInfof(
		err,
		func() { log.Infof("Successed to parse package %s", p.Name) },
		func() { log.Errorf("Failed to parse package %s, error: %s", p.Name, err) },
	)

	for _, file := range p.pkg.Files {
		if err = p.parseFile(ctx, file); err != nil {
			return err
		}
	}

	return nil
}

func (p *Package) parseFile(ctx context.Context, file *ast.File) error {
	var err error

	log.Infof("Parsing file %s", file.Name)
	defer log.ErrorOrInfof(
		err,
		func() { log.Infof("Successed to parse file %s", file.Name) },
		func() { log.Errorf("Failed to parse file %s, error: %s", file.Name, err) },
	)

	// add import to package
	for _, imp := range file.Imports {
		var name, path string

		if imp.Name != nil {
			name = imp.Name.Name
		}

		if imp.Path != nil {
			path = imp.Path.Value
		}

		if name == "" {
			name = filepath.Base(strings.Trim(path, `"`))
		}

		if value, ok := p.Imports[path]; ok && value != name {
			return fmt.Errorf("Found import package %s with different name (%s, %s)", path, value, name)
		}

		p.Imports[path] = name
	}

	// parse struct in file
	for _, decl := range file.Decls {
		if t, ok := decl.(*ast.GenDecl); ok {
			doc := t.Doc

			for _, spec := range t.Specs {
				if t, ok := spec.(*ast.TypeSpec); ok {
					if st, ok := t.Type.(*ast.StructType); ok {
						var obj generate.Object

						if obj, err = p.parseStruct(ctx, doc, t.Name.Name, st); err != nil {
							return err
						}

						if obj != nil {
							p.Objects = append(p.Objects, obj)
						}
					}
				}
			}
		}
	}

	return nil
}

func (p *Package) parseStruct(ctx context.Context, doc *ast.CommentGroup, name string, st *ast.StructType) (generate.Object, error) {
	obj := generate.NewObject()

	// parse docment
	generateType := parseComment(doc)

	if generateType == generate.NoneGenerateType {
		log.Infof("Ingore struct %s, nogenerate comment found", name)
		return nil, nil
	}

	if st.Fields == nil || len(st.Fields.List) == 0 {
		log.Infof("Ingore struct %s, No field found", name)
		return nil, nil
	}

	// set generate type
	obj.Meta().SetGenerateType(generateType)

	// add struct name
	obj.Meta().SetName(name)

	fields := make(map[string]generate.Object)

	// parse fields
	for _, field := range st.Fields.List {
		if err := parseField(ctx, field, fields); err != nil {
			return nil, err
		}
	}

	obj.SetUnderlying(generate.NewStructType(name, mapToArray(fields)...))
	return obj, nil
}

func parseComment(doc *ast.CommentGroup) generate.GenerateType {
	if doc == nil {
		return generate.BothGenerateType
	}

	var noencode, nodecode bool

	for _, line := range strings.Split(doc.Text(), "\n") {
		if strings.HasPrefix(line, "+gojson:") {
			for _, field := range strings.Fields(line) {
				switch strings.ToLower(field) {
				case "noencode":
					noencode = true

				case "nodecode":
					nodecode = true

				case "none":
					noencode = true
					nodecode = true
				}
			}
		}
	}

	switch {
	case noencode && nodecode:
		return generate.NoneGenerateType

	case noencode:
		return generate.DecodeGenerateType

	case nodecode:
		return generate.EncodeGenerateType

	default:
		return generate.BothGenerateType
	}
}

func parseField(ctx context.Context, field *ast.Field, fields map[string]generate.Object) (err error) {
	obj := generate.NewObject()

	// parse field tag
	if err := parseFieldTag(field.Tag, obj); err != nil {
		return err
	}

	// parse field
	if len(field.Names) > 0 {
		obj.Meta().SetName(field.Names[0].Name)
	}

	if err := parseType(ctx, field.Type, obj); err != nil {
		return err
	}

	// parse inline field
	if obj.Meta().GetInline() {
		if generate.IsStructUnderlying(obj.GetUnderlying()) {
			for _, field := range generate.FieldsInUnderlying(obj.GetUnderlying()) {
				if field.Meta().GetParent() != nil {
					tmp := field.Meta().GetParent()
					obj.Meta().SetParent(tmp)
				}

				field.Meta().SetParent(obj)
				if err := appendField(fields, field.Meta().GetTagName(), field); err != nil {
					return err
				}
			}

			return nil
		}

		return fmt.Errorf("inline tag only support StructType")
	}

	return appendField(fields, obj.Meta().GetTagName(), obj)
}

func parseFieldTag(tag *ast.BasicLit, obj generate.Object) error {
	if tag == nil {
		return nil
	}

	// get json tag
	tagStr := reflect.StructTag(strings.Trim(tag.Value, "`")).Get("json")
	if tagStr == "-" {
		obj.Meta().SetIngore()
		return nil
	}

	for index, field := range strings.Split(tagStr, ",") {
		if index == 0 && field != "" {
			obj.Meta().SetTagName(field)
			continue
		}

		switch field {
		case "inline":
			obj.Meta().SetInline()

		case "omitempty":
			obj.Meta().SetOmitempty()

		case "escape":
			obj.Meta().SetEscape()
		}
	}

	return nil
}

func parseInlineField(ctx context.Context, field *ast.Field, obj generate.Object, fields map[string]generate.Object) error {
	var st *ast.StructType
	ftyp := field.Type

	for {
		switch t := ftyp.(type) {
		case *ast.Ident:
			if t.Obj != nil {
				ts, ok := t.Obj.Decl.(*ast.TypeSpec)
				if ok {
					st, ok = ts.Type.(*ast.StructType)
					if ok {
						goto ParseStruct
					}

					ftyp = ts.Type
				}
			}

		case *ast.StarExpr:
			ftyp = t.X
			fobj := generate.NewObject()

			if err := parseType(ctx, t.X, fobj); err != nil {
				return err
			}

		default:
			return fmt.Errorf("Inline tag only support StructType")
		}
	}

ParseStruct:
	for _, field := range st.Fields.List {
		if err := parseField(ctx, field, fields); err != nil {
			return err
		}
	}

	return nil
}

func appendField(fields map[string]generate.Object, fkey string, fobj generate.Object) error {
	if _, ok := fields[fkey]; ok {
		return fmt.Errorf("Conflict field %s", fkey)
	}

	fields[fkey] = fobj
	return nil
}

func mapToArray(fields map[string]generate.Object) []generate.Object {
	ret := make([]generate.Object, 0, len(fields))

	for _, field := range fields {
		ret = append(ret, field)
	}

	return ret
}

func parseType(ctx context.Context, typ ast.Expr, obj generate.Object) error {
	switch t := typ.(type) {
	case *ast.Ident:
		if obj.Meta().GetName() == "" {
			obj.Meta().SetName(t.String())
		}

		return parseIdent(ctx, t, obj)

	case *ast.StarExpr:
		obj.Meta().SetPointer()
		return parseType(ctx, t.X, obj)

	case *ast.ArrayType:
		return parseArrayType(ctx, t, obj)

	case *ast.MapType:
		return parseMapType(ctx, t, obj)

	case *ast.InterfaceType:
		obj.SetUnderlying(generate.NewInterfaceType())
		return nil

	case *ast.SelectorExpr:
		obj.SetUnderlying(generate.NewSelectorType(fmt.Sprintf("%s.%s", t.X, t.Sel.String())))
		return nil

	default:
		return fmt.Errorf("Unsupported Field Type %T", typ)
	}
}

func parseIdent(ctx context.Context, t *ast.Ident, obj generate.Object) (err error) {
	switch t.Name {
	case "string":
		obj.SetUnderlying(generate.NewStringType())

	case "int":
		obj.SetUnderlying(generate.NewIntType(0))

	case "int8":
		obj.SetUnderlying(generate.NewIntType(8))

	case "int16":
		obj.SetUnderlying(generate.NewIntType(16))

	case "int32":
		obj.SetUnderlying(generate.NewIntType(32))

	case "int64":
		obj.SetUnderlying(generate.NewIntType(64))

	case "uint":
		obj.SetUnderlying(generate.NewUintType(0))

	case "uint8":
		obj.SetUnderlying(generate.NewUintType(8))

	case "uint16":
		obj.SetUnderlying(generate.NewUintType(16))

	case "uint32":
		obj.SetUnderlying(generate.NewUintType(32))

	case "uint64":
		obj.SetUnderlying(generate.NewUintType(64))

	case "float32":
		obj.SetUnderlying(generate.NewFloatType(32))

	case "float64":
		obj.SetUnderlying(generate.NewFloatType(64))

	case "bool":
		obj.SetUnderlying(generate.NewBoolType())

	default:
		if t.Obj != nil {
			ts, ok := t.Obj.Decl.(*ast.TypeSpec)
			if ok {
				return parseTypeSpec(ctx, ts, obj)
			}
		}

		return fmt.Errorf("Invalid ident type neither basic type nor typespec")
	}

	return nil
}

func parseTypeSpec(ctx context.Context, ts *ast.TypeSpec, obj generate.Object) (err error) {
	switch t := ts.Type.(type) {
	case *ast.StructType:
		fields := make(map[string]generate.Object)
		name := ts.Name.Name

		for _, field := range t.Fields.List {
			if err := parseField(ctx, field, fields); err != nil {
				return err
			}
		}

		obj.SetUnderlying(generate.NewStructType(name, mapToArray(fields)...))
		return nil

	default:
		// alias type
		obj.Meta().SetAlias(ts.Name.Name)
		return parseType(ctx, t, obj)
	}
}

func parseArrayType(ctx context.Context, t *ast.ArrayType, obj generate.Object) (err error) {
	var cap int
	var elem generate.Object

	if t.Len != nil {
		switch l := t.Len.(type) {
		case *ast.BasicLit:
			if cap, err = strconv.Atoi(l.Value); err != nil {
				return fmt.Errorf("unexpected Len %s in ArrayType, error: %s", l.Value, err)
			}

		default:
			return fmt.Errorf("unexpected type %T of Len in ArrayType", l)
		}
	}

	switch td := t.Elt.(type) {
	case *ast.Ident:
		if td.Name == "byte" {
			obj.SetUnderlying(generate.NewBytesType())
			return nil
		}
	}

	elem = generate.NewObject()

	if err := parseType(ctx, t.Elt, elem); err != nil {
		return err
	}

	obj.SetUnderlying(generate.NewArrayType(cap, elem))
	return nil
}

func parseMapType(ctx context.Context, t *ast.MapType, obj generate.Object) (err error) {
	key := generate.NewObject()
	value := generate.NewObject()

	// parse key
	if err = parseType(ctx, t.Key, key); err != nil {
		return err
	}

	// check key type
	if !generate.IsStringUnderlying(key.GetUnderlying()) {
		return fmt.Errorf("unexpected Key type %s in Map", key.GetUnderlying().Type())
	}

	if err = parseType(ctx, t.Value, value); err != nil {
		return err
	}

	obj.SetUnderlying(generate.NewMapType(key, value))
	return nil
}
