package cmd

import (
	"github.com/spf13/cobra"
)

// [[[command]]]Cmd run
var [[[command]]]Cmd = &cobra.Command{
	Use:   "[[[command]]]",
	Short: "Say [[[command]]]",
	Run: func(cmd *cobra.Command, args []string) {
		lib.[[[command]]]()
	},
}
