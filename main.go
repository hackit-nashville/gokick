package main

import (
	"fmt"
	"os"

	cmd "github.com/hackit-nashville/gokick/cmd"

	"github.com/spf13/cobra"
)

// Version of gokick. Overwritten during build
var version = "development"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gokick",
	Short: "generate common Golang projects",
}

func main() {
	rootCmd.Version = version
	rootCmd.SetVersionTemplate(`{{printf "%s" .version}}
`)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Import other commands
func init() {
	rootCmd.AddCommand(cmd.HelloCmd)
}
