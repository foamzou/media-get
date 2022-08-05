package version

import (
	"fmt"
)

const (
	BuildCode = 11
	BuildName = "0.2.6"
	Repo      = "https://github.com/foamzou/media-get"
)

func DisplayVersionInfo() {
	fmt.Printf("Version: %s\nBuild code: %d\nGithub: %s\n", BuildName, BuildCode, Repo)
}
