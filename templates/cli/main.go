package main

import (
	"fmt"
	"os"

	cmd "{{.Name}}/cmd"

	"github.com/spf13/cobra"
)

// Version of {{.Name}}. Overwritten during build
var version = "development"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "{{.Name}}",
	Short: "{{.Name}} is used as a starter cli project",
}

func main() {
	Execute(version)
}

// Import other commands
func init() {
	cmd.Import(rootCmd)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	rootCmd.Version = version
	rootCmd.SetVersionTemplate(`{{"{{"}}printf "%s" .Version{{"}}"}}
`)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
