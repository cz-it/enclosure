package cmd

import (
	"fmt"
	"os"

	"github.com/cz-it/enclosure/wallmap"
	"github.com/spf13/cobra"
)

var source string
var destination string
var wall string

func init() {
	ParseCMD.Flags().StringVarP(&source, "src", "s", "", "the source png file.")
	ParseCMD.Flags().StringVarP(&destination, "dst", "d", "", "the output file.default is a.json")
	ParseCMD.Flags().StringVarP(&wall, "wall", "w", "", "the color of map  use format like :#282828")

	destination = "a.out"
}

// ParseCMD is cmd for parse
var ParseCMD = &cobra.Command{
	Use:   "parse",
	Short: "parse a map .",
	Run:   parseMain,
}

func parseMain(cmd *cobra.Command, args []string) {
	if source == "" {
		fmt.Printf("source for parse should be set.  see help .\n")
		os.Exit(1)
	}
	wallmap.Parse(source, destination, wall)
}
