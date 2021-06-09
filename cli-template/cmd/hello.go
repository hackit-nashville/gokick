package cmd

import (
	lib "{{.Name}}/lib"

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
