package ffmpeg

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ConvertToMp3(inputFile, outputFile string, tag *MetaTag) error {
	var tagParams []string

	if tag != nil {
		tagParams = []string{
			"-metadata", "title=" + tag.Title,
			"-metadata", "artist=" + tag.Artist,
			"-metadata", "album=" + tag.Album,
		}
	}

	baseParams := []string{
		"-i", inputFile,
		"-c:v", "copy",
		"-strict", "experimental",
	}

	params := append(baseParams, tagParams...)
	params = append(params, outputFile)

	cmd := exec.Command("ffmpeg", params...)
	fmt.Println("------------------")
	fmt.Println(cmd.Args)

	var out bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		fmt.Println(stdErr.String())
		return err
	}
	return nil
}
