package lib

import (
	"fmt"
	"os"
	"strings"

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

		// var buf bytes.Buffer
		// fmt.Fprintf(&buf, "Size: %d MB.", info.Size())
		// s := buf.String())
		// if _, err := io.Copy(buf, f.Read()); err != nil {
		// 	return err
		// }
		// fmt.Println(buf.toString())

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

}
