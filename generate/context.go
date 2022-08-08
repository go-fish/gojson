package generate

import "context"

// ContextKey defines the type of context key
type ContextKey string

const (
	GenerateTypeContextKey ContextKey = "gojson.io/generateType"
	ArgvNameContextKey     ContextKey = "gojson.io/argvName"
	PointerContextKey      ContextKey = "gojson.io/pointer"
	AliasContextKey        ContextKey = "gojson.io/alias"
	OmitemptyContextKey    ContextKey = "gojson.io/omitempty"
	EscapeContextKey       ContextKey = "gojson.io/escape"
	TagNameContextKey      ContextKey = "gojson.io/tagName"
	LastFieldContextKey    ContextKey = "gojson.io/lastField"
)

// ContextWithGenerateType returns the context with GenerateType value
func ContextWithGenerateType(ctx context.Context, genType GenerateType) context.Context {
	return context.WithValue(ctx, GenerateTypeContextKey, genType)
}

// GenerateTypeInContext returns the GenerateType get from context
func GenerateTypeInContext(ctx context.Context) GenerateType {
	if value := ctx.Value(GenerateTypeContextKey); value != nil {
		return value.(GenerateType)
	}

	return ""
}

// ContextWithArgvName returns the context with ArgvName
func ContextWithArgvName(ctx context.Context, argv string) context.Context {
	return context.WithValue(ctx, ArgvNameContextKey, argv)
}

// ArgvNameInContext returns the ArgvName get from context
func ArgvNameInContext(ctx context.Context) string {
	if value := ctx.Value(ArgvNameContextKey); value != nil {
		return value.(string)
	}

	return ""
}

// ContextWithPointer returns the context with pointer
func ContextWithPointer(ctx context.Context, pointer bool) context.Context {
	return context.WithValue(ctx, PointerContextKey, pointer)
}

// IsPointer returns whether pointer is set in context
func IsPointer(ctx context.Context) bool {
	if value := ctx.Value(PointerContextKey); value != nil {
		return value.(bool)
	}

	return false
}

// ContextWithAlias returns the context with alias
func ContextWithAlias(ctx context.Context, alias bool) context.Context {
	return context.WithValue(ctx, AliasContextKey, alias)
}

// IsAlias returns whether alias is set in context
func IsAlias(ctx context.Context) bool {
	if value := ctx.Value(AliasContextKey); value != nil {
		return value.(bool)
	}

	return false
}

// ContextWithEscape returns the context with escape
func ContextWithEscape(ctx context.Context, escape bool) context.Context {
	return context.WithValue(ctx, EscapeContextKey, escape)
}

// IsEscape returns whether omitempty is set in context
func IsEscape(ctx context.Context) bool {
	if value := ctx.Value(EscapeContextKey); value != nil {
		return value.(bool)
	}

	return false
}
