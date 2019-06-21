package main

import (
	"fmt"
	"os"

	"github.com/cz-it/enclosure/cmd"
	"github.com/spf13/cobra"
)

func init() {
	enclosureCMD.AddCommand(cmd.VersionCMD)
	enclosureCMD.AddCommand(cmd.ParseCMD)
}

var enclosureCMD = &cobra.Command{
	Use:   "enclosure parse",
	Short: "enclosure is map parser",
	Long: `
	enclosure:
		A map tool which parse a map(png) and output a JSON with pixel coordinate of legal position
	`,
	Run: enclosureMain,
}

func enclosureMain(cmd *cobra.Command, args []string) {
	cmd.Help()
	return
}

func cmdExecute() {
	if err := enclosureCMD.Execute(); err != nil {
		fmt.Printf("CMD execute error:%s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	cmdExecute()
}
