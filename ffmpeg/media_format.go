package ffmpeg

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"strings"
)

func GetMediaFormat(filePath string) (*MediaFormat, error) {
	cmdList := strings.Split("ffprobe -v quiet -show_format -print_format json", " ")

	cmdList = append(cmdList, filePath)

	cmd := exec.Command(cmdList[0], cmdList[1:]...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	mediaFormat := &MediaFormat{}
	err = json.Unmarshal(out.Bytes(), mediaFormat)
	if err != nil {
		return nil, err
	}
	return mediaFormat, nil
}
