package main

import (
	"errors"
	"os"

	"github.com/go-fish/gojson/builder"
	"github.com/go-fish/gojson/log"
	"github.com/spf13/cobra"
)

var rootcmd *cobra.Command

func init() {
	rootcmd = &cobra.Command{
		Use:  "gojson [flags] [file|directory]...",
		Long: "gojson used to generate the Marshal/Unmarshal json code for custom struct",
	}
}

func main() {
	var unsafe, escapeHTML, escapeUnicode bool
	var output string

	rootcmd.Flags().BoolVar(&unsafe, "unsafe", true, "use input data directly without copy when decode json bytes")
	rootcmd.Flags().BoolVar(&escapeHTML, "escapehtml", false, "escape html char when encode object to json")
	rootcmd.Flags().BoolVar(&escapeUnicode, "escapeunicode", false, "escape unicode rune when decode object to json")
	rootcmd.Flags().StringVarP(&output, "output", "o", "gojson.generated.go", "the filename of output file")

	rootcmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Missing inputs when generate gojson code")
		}

		mgr := builder.NewBuilder(
			unsafe,
			escapeHTML,
			escapeUnicode,
			output,
		)

		return mgr.Execute(args...)
	}

	if err := rootcmd.Execute(); err != nil {
		log.Errorf(err.Error())
		os.Exit(1)
	}
}
