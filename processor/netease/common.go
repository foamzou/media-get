package netease

import (
	"fmt"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/utils"
)

func getSongUrl(songId int) string {
	return fmt.Sprintf("https://music.163.com/song/media/outer/url?id=%d.mp3", songId)
}

func fetchHtml(url string) (string, error) {
	return utils.HttpGet(url, map[string]string{
		"user-agent": consts.UAMac,
	})
}
