package generate

// ObjectType defines the type of object
type ObjectType string

const (
	FieldObjectType  ObjectType = "Field"
	NormalObjectType ObjectType = "Normal"
)

// GenerateType defines the tpye of generate
type GenerateType string

const (
	DecodeGenerateType GenerateType = "Decode"
	EncodeGenerateType GenerateType = "Encode"
	BothGenerateType   GenerateType = "Both"
	NoneGenerateType   GenerateType = "None"
)

// MetaObject describes the Metadata of object
type MetaObject interface {
	SetPointer()
	GetPointer() bool
	SetAlias(alias ...string)
	GetAlias() []string
	SetParent(obj Object)
	GetParent() Object

	// Field Info
	SetName(name string)
	GetName() string
	SetType(typ ObjectType)
	GetType() ObjectType

	// FieldTag
	SetOmitempty()
	GetOmitempty() bool
	SetInline()
	GetInline() bool
	SetIngore()
	GetIngore() bool
	SetEscape()
	GetEscape() bool
	SetTagName(name string)
	GetTagName() string

	// generate info
	SetGenerateType(gType GenerateType)
	GetGenerateType() GenerateType
}

type metaObject struct {
	pointer bool
	alias   []string
	parent  Object

	// fieldInfo
	name       string
	objectType ObjectType

	// field tag info
	tagName   string
	ingore    bool
	inline    bool
	omitempty bool
	escape    bool

	// generate info
	genType GenerateType
}

var _ MetaObject = new(metaObject)

// NewMetaObject returns the instance of MetaObject
func NewMetaObject() MetaObject {
	return new(metaObject)
}

func (mo *metaObject) SetPointer() {
	mo.pointer = true
}

func (mo *metaObject) GetPointer() bool {
	return mo.pointer
}

func (mo *metaObject) SetAlias(alias ...string) {
	mo.alias = append(mo.alias, alias...)
}

func (mo *metaObject) GetAlias() []string {
	return mo.alias
}

func (mo *metaObject) SetName(name string) {
	mo.name = name
}

func (mo *metaObject) GetName() string {
	return mo.name
}

func (mo *metaObject) SetParent(obj Object) {
	mo.parent = obj
}

func (mo *metaObject) GetParent() Object {
	return mo.parent
}

func (mo *metaObject) SetType(typ ObjectType) {
	mo.objectType = typ
}

func (mo *metaObject) GetType() ObjectType {
	return mo.objectType
}

func (mo *metaObject) SetOmitempty() {
	mo.omitempty = true
}

func (mo *metaObject) GetOmitempty() bool {
	return mo.omitempty
}

func (mo *metaObject) SetInline() {
	mo.inline = true
}

func (mo *metaObject) GetInline() bool {
	return mo.inline
}

func (mo *metaObject) SetIngore() {
	mo.ingore = true
}

func (mo *metaObject) GetIngore() bool {
	return mo.ingore
}

func (mo *metaObject) SetEscape() {
	mo.escape = true
}

func (mo *metaObject) GetEscape() bool {
	return mo.escape
}

func (mo *metaObject) SetTagName(name string) {
	mo.tagName = name
}

func (mo *metaObject) GetTagName() string {
	if mo.tagName != "" {
		return mo.tagName
	}

	return mo.name
}

func (mo *metaObject) SetGenerateType(gType GenerateType) {
	mo.genType = gType
}

func (mo *metaObject) GetGenerateType() GenerateType {
	return mo.genType
}
