default:
	$(MAKE) all

windows:
	mkdir -p dist/windows
	env GOOS=windows GOARCH=amd64 go build -o dist/windows/stupid-http.exe main.go

linux:
	mkdir -p dist/linux
	env GOOS=linux GOARCH=amd64 go build -o dist/linux/stupid-http main.go

darwin:
	mkdir -p dist/darwin
	env GOOS=darwin GOARCH=amd64 go build -o dist/darwin/stupid-http main.go

docker:
	$(MAKE) linux
	rm -rf dist/docker
	cp -r stuff/docker dist
	cp dist/linux/stupid-http dist/docker
	cd dist/docker && $(MAKE)

push:
	cd dist/docker && $(MAKE) push

all:
	$(MAKE) windows
	$(MAKE) darwin
	$(MAKE) linux