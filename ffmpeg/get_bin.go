package ffmpeg

import (
	"os"
	"os/exec"
	"runtime"
)

const (
	BuildInSubPath = "/media-get/bin/"
	FFMPEG         = "ffmpeg"
	FFPROBE        = "ffprobe"
)

func GetFfmpegBin() string {
	return getBin(FFMPEG)
}

func GetFfprobeBin() string {
	return getBin(FFPROBE)
}

func getBin(binName string) string {
	if commandExists(binName) {
		return binName
	}
	buildInPath := GetBuildInPath(binName)
	if fileExists(buildInPath) {
		return buildInPath
	}
	return ""
}

func GetBuildInPath(binName string) string {
	appTmpPath, err := os.UserConfigDir()
	if err != nil {
		return ""
	}
	binPath := appTmpPath + BuildInSubPath + binName
	if runtime.GOOS == "windows" {
		binPath = binPath + ".exe"
	}
	return binPath
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
func fileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
