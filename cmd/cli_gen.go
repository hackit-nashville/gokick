package cmd

import (
	"fmt"

	"github.com/hackit-nashville/gokick/lib"
	"github.com/spf13/cobra"
)

var (
	name, directory string
	commands        []string
)

// CliGeneratorCMD generates
var CliGeneratorCMD = &cobra.Command{
	Use:   "cli",
	Short: "generate a CLI template",
	Run: func(cmd *cobra.Command, args []string) {

		for _, dir := range commands {
			fmt.Println(dir)
		}

		lib.Generate(directory)

	},
}

func init() {
	CliGeneratorCMD.Flags().SetInterspersed(false)
	CliGeneratorCMD.PersistentFlags().StringVarP(&name, "name", "n", "", "name of the project")
	CliGeneratorCMD.PersistentFlags().StringVarP(&directory, "directory", "d", "", "directory to init the project")
	CliGeneratorCMD.PersistentFlags().StringSliceVar(&commands, "commands", []string{}, "comma delimitted list of commands to seed")
}
