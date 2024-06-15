build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/linux_amd64 ./cmd/linux

build-linux-arm:
	GOOS=linux GOARCH=arm64 go build -o bin/linux/linux_arm64 ./cmd/linux

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/darwin_amd64 ./cmd/darwin

build-darwin-arm:
	GOOS=darwin GOARCH=arm64 go build -o bin/darwin/darwin_arm64 ./cmd/darwin

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/windows/windows_64.exe ./cmd/windows
