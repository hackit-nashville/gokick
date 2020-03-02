package lib

import (
	"fmt"
	"os"

	"github.com/markbates/pkger"
)

func init() {
	pkger.Include("/cli-template")
}

// Generate says hello
func Generate() {

	err := pkger.Walk("/cli-template", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

}
