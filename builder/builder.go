package builder

import (
	"github.com/go-fish/gojson/generate"
	"github.com/go-fish/gojson/parser"
)

// Builder is a builder to parse and generate code
type Builder struct {
	parser    parser.Parser
	generator generate.Generator
}

// NewBuilder returns the instance of builder
func NewBuilder(
	unsafe bool,
	escapeHTML bool,
	escapeUnicode bool,
	output string,
) *Builder {
	return &Builder{
		parser:    parser.NewParser(),
		generator: generate.NewGenerator(unsafe, escapeHTML, escapeUnicode, output),
	}
}

// Execute runs the builder
func (b *Builder) Execute(inputs ...string) error {
	if err := b.parser.Parse(inputs...); err != nil {
		return err
	}

	for _, pkg := range b.parser.Packages() {
		if err := b.generator.Generate(
			pkg.Name,
			pkg.Path,
			pkg.Imports,
			pkg.Objects...,
		); err != nil {
			return err
		}
	}

	return nil
}
