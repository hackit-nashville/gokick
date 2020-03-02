package cmd

import (
	lib "github.com/hackit-nashville/gokick/lib"

	"github.com/spf13/cobra"
)

// HelloCmd run
var HelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Run: func(cmd *cobra.Command, args []string) {
		lib.HelloWorld()
	},
}
