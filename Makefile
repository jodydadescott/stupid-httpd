build:
	dep ensure -v
	mkdir -p dist/darwin
	mkdir -p dist/linux
	go build -o dist/darwin/stupid-http main/main.go
	env GOOS=linux GOARCH=amd64 go build -o dist/linux/stupid-http main/main.go
