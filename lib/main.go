package lib

import (
	"fmt"
	"log"
	"os"
	"path"

	"text/template"

	"github.com/hackit-nashville/gokick/lib/cli_template"
)

const templatePath = "/cli-template"

type TemplateConfig struct {
	Name string
}

// Generate says hello
func Generate(workingDir string, templateConfig TemplateConfig) {

	// List through each asset in the CLI template
	for _, assetName := range cli_template.AssetNames() {

		// Get the contents of the asset
		assetContent, err := cli_template.Asset(assetName)
		if err != nil {
			log.Fatal(fmt.Sprintf("Unable to read %s asset: %v", assetName, err))
		}

		// Consider each asset to be a possible template and parse it as such.
		var assetTemplate = template.Must(template.New(assetName).Parse(string(assetContent)))

		// Reconcile the generated out file and directory
		generatedAssetPath := path.Join(workingDir, assetName)
		generatedAssetDirectory := path.Dir(generatedAssetPath)

		fmt.Printf("Generating %s\n", generatedAssetPath)

		// Ensure dir is created
		err = os.MkdirAll(generatedAssetDirectory, 0700)
		if err != nil {
			log.Print(err)
		}

		// Create the output file
		generatedAsset, err := os.Create(generatedAssetPath)
		if err != nil {
			log.Println("create file: ", err)
		}

		// Execute the template and save the file
		err = assetTemplate.Execute(generatedAsset, templateConfig)
		if err != nil {
			log.Print("execute: ", err)
		}

	}
}
