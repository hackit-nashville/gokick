default: build

build: bindata
	rm -rf build
	GOOS=linux CGO_ENABLED=0 go build -o build/gokick main.go 

bindata:
  # go get -u github.com/kevinburke/go-bindata/...
	go-bindata -pkg cli_template -o lib/cli_template/main.go -prefix templates/cli templates/cli/...

test:
	rm -rf ../go-test
	go run main.go cli --directory .. --name go-test
