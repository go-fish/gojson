package parser

import (
	"context"
	"fmt"
	gparser "go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/go-fish/gojson/log"
)

// Parser describe the gojson parser
type Parser interface {
	Parse(inputs ...string) error
	Packages() []*Package
}

type parser struct {
	packages []*Package
}

var _ Parser = new(parser)

// NewParser returns the instance of Parser
func NewParser() Parser {
	return new(parser)
}

func (p *parser) Parse(inputs ...string) error {
	for _, input := range inputs {
		fi, err := os.Stat(input)
		if err != nil {
			log.Errorf("Invalid input %s, error: %s", input, err)
			continue
		}

		if !fi.IsDir() {
			input = filepath.Dir(input)
		}

		if err := p.parseDirectory(input); err != nil {
			return err
		}

		log.Infof("Parsed input %s", input)
	}

	return nil
}

func (p *parser) parseDirectory(dir string) error {
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			return nil
		}

		log.Infof("Parsing directory %s", path)
		defer log.ErrorOrInfof(
			err,
			func() { log.Infof("Successed to parse directory %s", path) },
			func() { log.Errorf("Failed to parse directory %s, error: %s", path, err) },
		)

		pkgs, err := gparser.ParseDir(token.NewFileSet(), path, nil, gparser.ParseComments)
		if err != nil {
			return err
		}

		if len(pkgs) > 1 {
			return fmt.Errorf("More than 1 package found in directory %s", dir)
		}

		// parse package
		for _, pkg := range pkgs {
			pi := NewPackage(path, pkg)

			if err = pi.Parse(context.TODO()); err != nil {
				return err
			}

			p.packages = append(p.packages, pi)
		}

		return nil
	})
}

func (p *parser) Packages() []*Package {
	return p.packages
}
