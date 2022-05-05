package ffmpeg

import (
	"bytes"
	"os/exec"

	"github.com/foamzou/audio-get/logger"
	"github.com/foamzou/audio-get/utils"
)

func ConvertSingleInput(inputFile, outputFile string, tag *MetaTag, isAudio bool) error {
	return ConvertMultiInput([]string{inputFile}, outputFile, tag, isAudio)
}

func ConvertMultiInput(inputFiles []string, outputFile string, tag *MetaTag, isAudio bool) error {
	var tagParams []string
	shouldAddCover := isAudio && len(inputFiles) == 1 && tag != nil && tag.Cover != ""

	if tag != nil {
		tagParams = []string{
			"-metadata", "title=" + utils.FilterUnexpectedChar(tag.Title),
			"-metadata", "artist=" + tag.Artist,
			"-metadata", "album=" + tag.Album,
		}
		if shouldAddCover {
			// audio at 0, cover/video at 1
			tagParams = append(tagParams, "-map", "0:a", "-map", "1:v")
		}
	}

	var inputParams []string
	for _, file := range inputFiles {
		inputParams = append(inputParams, "-i", file)
	}

	if tag != nil && shouldAddCover {
		inputParams = append(inputParams, "-i", tag.Cover)
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
	logger.Debug("------------------")
	logger.Debug(cmd.Args)

	var out bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		logger.Error(stdErr.String())
		return err
	}
	return nil
}
