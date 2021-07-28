# go-cli-starter-template

## Usage

```
➜ gokick cli --name helloworld -d /tmp/helloworld                      
Generating /tmp/helloworld/Dockerfile
Generating /tmp/helloworld/LICENSE
Generating /tmp/helloworld/cmd/hello.go
Generating /tmp/helloworld/cmd/import.go
Generating /tmp/helloworld/go.mod
Generating /tmp/helloworld/go.sum
Generating /tmp/helloworld/CHANGELOG.md
Generating /tmp/helloworld/Makefile
Generating /tmp/helloworld/README.md
Generating /tmp/helloworld/lib/main.go
Generating /tmp/helloworld/main.go
➜ cd /tmp/helloworld 
➜ helloworld make run
go run main.go hello
Hello world


➜ gokick cli --help
generate a CLI template

Usage:
  gokick cli [flags]

Flags:
      --commands strings   comma delimitted list of commands to seed
  -d, --directory string   directory to init the project
  -h, --help               help for cli
  -n, --name string        name of the project (default "example-cli")
```
