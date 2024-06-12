build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/linux_amd64 main.go
build-linux-arm:
	GOOS=linux GOARCH=arm64 go build -o bin/linux/linux_arm64 main.go
build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/darwin_amd64 main.go
build-darwin-arm:
	GOOS=darwin GOARCH=arm64 go build -o bin/darwin/darwin_arm64 main.go
build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/windows/windows_64.exe main.go
build: build-linux build-linux-arm build-darwin build-darwin-arm build-windows

