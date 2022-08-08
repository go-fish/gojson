package generate

import (
	"math/rand"
	"strings"
)

const randomChars = "0123456789abcdefghijklmnopqrstuvwxyz"

func genRandomString(cap int) string {
	data := make([]byte, cap)
	for i := range data {
		data[i] = randomChars[rand.Int63()%int64(len(randomChars))]
	}
	return string(data)
}

func genRandomSuffix(data string) string {
	return data + genRandomString(8)
}

func shortName(obj Object) string {
	return strings.ToLower(obj.Meta().GetName())[:1]
}

func convertAlias(argv string, alias ...string) string {
	if len(alias) == 0 {
		return argv
	}

	var builder strings.Builder

	for _, v := range alias {
		builder.WriteString(v)
		builder.WriteByte('(')
	}

	builder.WriteString(argv)
	builder.WriteString(strings.Repeat(")", len(alias)))
	return builder.String()
}

// IsStructUnderlying returns whether the object underlying is struct type
func IsStructUnderlying(underlying Underlying) bool {
	_, ok := underlying.(*structType)
	return ok
}

// IsStringUnderlying returns whether the object underlying is string type
func IsStringUnderlying(underlying Underlying) bool {
	_, ok := underlying.(*stringType)
	return ok
}

// IsSelectorUnderlying returns whether the object underlying is selector type
func IsSelectorUnderlying(underlying Underlying) bool {
	_, ok := underlying.(*selectorType)
	return ok
}

// FieldsInUnderlying returns the field list if underlying is struct type
func FieldsInUnderlying(underlying Underlying) []Object {
	st, ok := underlying.(*structType)
	if !ok {
		return nil
	}

	return st.fields
}
