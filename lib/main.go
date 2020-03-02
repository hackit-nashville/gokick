package lib

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/markbates/pkger"
	"text/template"
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
		filePath = strings.Split(filePath, ":")[1]
		fmt.Println(filePath)

		f, err := pkger.Open(filePath)
		if err != nil {
			return err
		}
		defer f.Close()

		info, err = f.Stat()
		if err != nil {
			return err
		}

		type Inventory struct {
			Material string
			Count    uint
		}

		type Inventory struct {
			Material string
			Count    uint
		}
		sweaters := Inventory{"wool", 17}
		tmpl, err := template.New("test").Parse(f.Read())
		if err != nil { panic(err) }
		err = tmpl.Execute(os.Stdout, sweaters)
		if err != nil { panic(err) }

		// f.Read()
		// if _, err := io.Copy(os.Stdout, f.Read()); err != nil {
		// 	return err
		// }

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

}
