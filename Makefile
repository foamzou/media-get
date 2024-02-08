BINARY_NAME=media-get
VERSION_NAME := $(shell grep "BuildName" version/version.go | head -n1 | awk -F'"' '{print $$2}')

MAC_AMD64_BIN=.release/${BINARY_NAME}-${VERSION_NAME}-darwin
MAC_ARM64_BIN=.release/${BINARY_NAME}-${VERSION_NAME}-darwin-arm64
LINUX_AMD64_BIN=.release/${BINARY_NAME}-${VERSION_NAME}-linux
LINUX_ARM64_BIN=.release/${BINARY_NAME}-${VERSION_NAME}-linux-arm64
WIN_BIN=.release/${BINARY_NAME}-${VERSION_NAME}-win.exe

release:
	@printf "${VERSION_NAME}" > ./LATEST_VERSION
	mkdir -p .release
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -trimpath -ldflags "-s -w" -o ${MAC_AMD64_BIN} main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -trimpath -ldflags "-s -w" -o ${MAC_ARM64_BIN} main.go

	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -trimpath -ldflags "-s -w" -o ${LINUX_AMD64_BIN} main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -trimpath -ldflags "-s -w" -o ${LINUX_ARM64_BIN} main.go

	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -trimpath -ldflags "-s -w" -o ${WIN_BIN} main.go
#   upx 尚不支持 macos 最新版本，先不做压缩
#	upx ${MAC_AMD64_BIN}
#	upx ${MAC_ARM64_BIN}
	upx ${LINUX_AMD64_BIN}
	upx ${LINUX_ARM64_BIN}
	upx ${WIN_BIN}

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