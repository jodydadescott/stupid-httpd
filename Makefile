default:
	$(MAKE) all

windows:
	mkdir -p dist/windows
	env GOOS=windows GOARCH=amd64 go build -o dist/windows/stupid-httpd.exe main.go

linux:
	mkdir -p dist/linux
	env GOOS=linux GOARCH=amd64 go build -o dist/linux/stupid-httpd main.go

darwin:
	mkdir -p dist/darwin
	env GOOS=darwin GOARCH=amd64 go build -o dist/darwin/stupid-httpd main.go

all:
	$(MAKE) windows
	$(MAKE) darwin
	$(MAKE) linux