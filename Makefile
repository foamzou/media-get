BINARY_NAME=media-get
MAC_BIN=.release/${BINARY_NAME}-darwin
LINUX_BIN=.release/${BINARY_NAME}-linux
WIN_BIN=.release/${BINARY_NAME}-win

release:
	mkdir -p .release
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o ${MAC_BIN} main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o ${LINUX_BIN} main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -ldflags "-s -w" -o ${WIN_BIN} main.go
	upx ${MAC_BIN}
	upx ${LINUX_BIN}
	upx ${WIN_BIN}

build:
	go build main.go

clean:
	go clean
	rm .release/*

test:
	go test ./...

lint:
	golangci-lint run