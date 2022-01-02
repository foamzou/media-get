package netease

import (
	"strings"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

func (c *Core) fetchFromSong() (mediaMeta *meta.MediaMeta, err error) {
	url := strings.Replace(c.Opts.Url, "#", "/", 1)
	songId := utils.RegexSingleMatchIntIgnoreError(url, `id=([\d]+)`, 0)
	if songId == 0 {
		return
	}

	html, err := fetchHtml(url)

	if err != nil {
		return
	}

	songName := utils.RegexSingleMatchIgnoreError(html, `"og:title" content="(.+?)"[\s]*/>`,
		utils.RegexSingleMatchIgnoreError(html, `data-res-name="(.+?)"`, ""),
	)

	mediaMeta = &meta.MediaMeta{
		Title:        songName,
		Description:  utils.RegexSingleMatchIgnoreError(html, `"title": "(.+?)"`, ""),
		Duration:     utils.RegexSingleMatchIntIgnoreError(html, `property="music:duration" content="(.+?)"`, 0),
		Album:        utils.RegexSingleMatchIgnoreError(html, `og:music:album" content="(.+?)"[\s]*/>`, ""),
		Artist:       utils.RegexSingleMatchIgnoreError(html, `og:music:artist" content="(.+?)"[\s]*/>`, ""),
		Audios:       []meta.Audio{{Url: getSongUrl(songId)}},
		ResourceType: consts.ResourceTypeAudio,
		Headers: map[string]string{
			"user-agent": consts.UAMac,
			"referer":    c.Opts.Url,
		},
	}

	return
}
