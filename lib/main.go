package lib

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path"

	"github.com/markbates/pkger"
)

func init() {
	pkger.Include("/cli-template")
}

type Cfg struct {
	x string
}

// Generate says hello
func Generate(workingDir string) {

	err := pkger.Walk("/cli-template", func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// Render Template
		fmt.Println(filePath)
		fmt.Println(info.Sys)
		t, err := template.ParseFiles(filePath.Path())
		if err != nil {
			log.Print(err)
			return err
		}

		f, err := os.Create(path.Join(workingDir, filePath.Path()))
		if err != nil {
			log.Println("create file: ", err)
			return err
		}
		var config Cfg
		err = t.Execute(f, config)
		if err != nil {
			log.Print("execute: ", err)
			return err
		}

		f.Close()

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

}
