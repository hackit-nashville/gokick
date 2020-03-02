package cmd

import (
	lib "github.com/hackit-nashville/gokick/cli-template/lib"

	"github.com/spf13/cobra"
)

// Run Command ./pentaho-cli run
var HelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Run: func(cmd *cobra.Command, args []string) {
		lib.HelloWorld()
	},
}
