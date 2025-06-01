MAIN_FILE_PATH = ./cmd/app/main.go

native: #build on current machine
	go build -trimpath -o crypt $(MAIN_FILE_PATH)

windows: #build for windows
	GOOS=windows GOARCH=amd64 go build -trimpath -o crypt.exe $(MAIN_FILE_PATH)

mac: #build for mac
	GOOS=darwin GOARCH=arm64 go build -trimpath -o crypt $(MAIN_FILE_PATH)
