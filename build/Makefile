default:
	@echo "Make what? (all, clean, windows-amd64, windows-arm64, linux-amd64, linux-arm, linux-arm64, darwin-amd64, darwin-arm64)"
	exit 2

windows-amd64:
	env GOOS=windows GOARCH=amd64 go build -o windows-amd64 main.go

windows-arm64:
	env GOOS=windows GOARCH=arm64 go build -o windows-arm64 main.go

linux-amd64:
	env GOOS=linux GOARCH=amd64 go build -o linux-amd64 main.go

linux-arm:
	env GOOS=linux GOARCH=arm go build -o linux-arm main.go

linux-arm64:
	env GOOS=linux GOARCH=arm64 go build -o linux-arm64 main.go

darwin-amd64:
	env GOOS=darwin GOARCH=amd64 go build -o darwin-amd64 main.go

darwin-arm64:
	env GOOS=darwin GOARCH=arm64 go build -o darwin-arm64 main.go

all:
	$(MAKE) windows-amd64
	$(MAKE) windows-arm64
	$(MAKE) linux-amd64
	$(MAKE) linux-arm
	$(MAKE) linux-arm64
	$(MAKE) darwin-amd64
	$(MAKE) darwin-arm64

clean:
	$(RM) windows-amd64
	$(RM) windows-arm64
	$(RM) linux-amd64
	$(RM) linux-arm
	$(RM) linux-arm64
	$(RM) darwin-amd64
	$(RM) darwin-arm64