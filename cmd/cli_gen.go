package cmd

import (
	"fmt"
	"os"

	"path/filepath"

	"github.com/hackit-nashville/gokick/lib"
	"github.com/spf13/cobra"
)

var (
	name, directory string
	commands        []string
	templateConfig  lib.TemplateConfig
)

// CliGeneratorCMD generates
var CliGeneratorCMD = &cobra.Command{
	Use:   "cli",
	Short: "generate a CLI template",
	Run: func(cmd *cobra.Command, args []string) {

		for _, dir := range commands {
			fmt.Println(dir)
		}

		directory, err := filepath.Abs(directory)
		if err != nil {
			fmt.Println("Error creating absolute path.", err)
			os.Exit(1)
		}

		directory = filepath.Join(filepath.Clean(directory))

		lib.Generate(directory, templateConfig)

	},
}

func init() {
	CliGeneratorCMD.Flags().SetInterspersed(false)
	CliGeneratorCMD.PersistentFlags().StringVarP(&templateConfig.Name, "name", "n", "example-cli", "name of the project")
	CliGeneratorCMD.PersistentFlags().StringVarP(&directory, "directory", "d", "", "directory to init the project")
	CliGeneratorCMD.PersistentFlags().StringSliceVar(&commands, "commands", []string{}, "comma delimitted list of commands to seed")
}
