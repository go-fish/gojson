package builder

import (
	"testing"

	"github.com/go-fish/gojson/generate"
	"github.com/go-fish/gojson/parser"
)

func TestBuilder_Execute(t *testing.T) {
	type fields struct {
		parser    parser.Parser
		generator generate.Generator
	}
	type args struct {
		inputs []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				parser:    parser.NewParser(),
				generator: generate.NewGenerator(true, true, true, "gojson.genreated.go"),
			},
			args: args{
				inputs: []string{"../example"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{
				parser:    tt.fields.parser,
				generator: tt.fields.generator,
			}
			if err := b.Execute(tt.args.inputs...); (err != nil) != tt.wantErr {
				t.Errorf("Builder.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
