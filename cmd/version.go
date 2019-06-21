package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

const (
	majorVer = 0
	minorVer = 0
	fixVer = 1
)

// VersionCMD is: dmtrade version
var VersionCMD = &cobra.Command {
	Use: "version",
	Short: "Version number.",
	Long: "Version number.",
	Run: versionMain,
}

func versionMain(cmd *cobra.Command, args []string) {
	fmt.Printf("Version:%d.%d.%d\n", majorVer, minorVer, fixVer )
}