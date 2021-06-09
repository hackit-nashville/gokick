default:
	pkger
	go build .

test:
	rm -rf ../go-test
	go run main.go cli --directory .. --name go-test