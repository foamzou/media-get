package version

import (
	"fmt"
)

const (
	BuildCode = 4
	BuildName = "0.1.3"
	Repo      = "https://github.com/foamzou/media-get"
)

func DisplayVersionInfo() {
	fmt.Printf("Version: %s\nBuild code: %d\nGithub: %s\n", BuildName, BuildCode, Repo)
}
