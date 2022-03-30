BINARY_NAME=media-get
VERSION_NAME := $(shell grep "BuildName" version/version.go | head -n1 | awk -F'"' '{print $$2}')

MAC_BIN=.release/${BINARY_NAME}-${VERSION_NAME}-darwin
LINUX_BIN=.release/${BINARY_NAME}-${VERSION_NAME}-linux
WIN_BIN=.release/${BINARY_NAME}-${VERSION_NAME}-win.exe

release:
	@printf "${VERSION_NAME}" > ./LATEST_VERSION
	mkdir -p .release
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -trimpath -ldflags "-s -w" -o ${MAC_BIN} main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -trimpath -ldflags "-s -w" -o ${LINUX_BIN} main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -trimpath -ldflags "-s -w" -o ${WIN_BIN} main.go
	upx ${MAC_BIN}
	upx ${LINUX_BIN}
	upx ${WIN_BIN}
	chmod +x ${MAC_BIN} ${LINUX_BIN}

build:
	go build main.go

clean:
	go clean
	rm .release/*

test:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

lint:
	golangci-lint run