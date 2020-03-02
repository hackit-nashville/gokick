package lib

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"text/template"

	"github.com/markbates/pkger"
)

const templatePath = "/cli-template"

func init() {
	pkger.Include(templatePath)
}

type Cfg struct {
	x string
}

// Generate says hello
func Generate(workingDir string) {

	err := pkger.Walk(templatePath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// Render Template
		// fmt.Println(filePath)
		// fmt.Println(info.Name())
		// pkger.Open(info)
		filePath = strings.Split(filePath, ":")[1]
		f, err := pkger.Open(filePath)
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer(nil)
		io.Copy(buf, f)
		f.Close()
		s := string(buf.Bytes())
		// fmt.Println(s)

		var t = template.Must(template.New(filePath).Parse(s))
		if err != nil {
			log.Print(err)
			return err
		}

		outFile := path.Join(workingDir, strings.Replace(filePath, templatePath, "", 1))
		outDir := path.Dir(outFile)

		fmt.Printf("Generating %s\n", outFile)

		// Ensure dir is created
		err = os.MkdirAll(outDir, 0700)
		if err != nil {
			log.Print(err)
			return err
		}

		nf, err := os.Create(outFile)
		if err != nil {
			log.Println("create file: ", err)
			return err
		}
		var config Cfg
		err = t.Execute(nf, config)
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
