package main

import (
	"fmt"
	"os"

	cmd "github.com/hackit-nashville/gokick/cli-template/cmd"

	"github.com/spf13/cobra"
)

// Version of go-cli-starter-template. Overwritten during build
var version = "development"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cli-starter-template",
	Short: "go-cli-starter-template is used as a starter cli project",
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
	rootCmd.SetVersionTemplate(`{{"{{"}}printf "%s" .version{{"}}"}}
`)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
