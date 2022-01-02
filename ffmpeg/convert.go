package ffmpeg

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ConvertSingleInput(inputFile, outputFile string, tag *MetaTag) error {
	return ConvertMultiInput([]string{inputFile}, outputFile, tag)
}

func ConvertMultiInput(inputFiles []string, outputFile string, tag *MetaTag) error {
	var tagParams []string

	if tag != nil {
		tagParams = []string{
			"-metadata", "title=" + tag.Title,
			"-metadata", "artist=" + tag.Artist,
			"-metadata", "album=" + tag.Album,
		}
	}

	var inputParams []string
	for _, file := range inputFiles {
		inputParams = append(inputParams, "-i")
		inputParams = append(inputParams, file)
	}

	baseParams := []string{
		"-y",
		"-c:v", "copy",
		"-strict", "experimental",
	}

	params := inputParams
	params = append(params, baseParams...)
	params = append(params, tagParams...)
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
