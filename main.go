package main

import (
	"go-cli-starter-template/cmd"
)

// Version of go-cli-starter-template. Overwritten during build
var Version = "development"

func main() {
	cmd.Execute(Version)
}
