package netease

import (
	"strconv"
	"strings"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

func (c *Core) fetchFromSong() (mediaMeta *meta.MediaMeta, audios []*meta.Audio, err error) {
	url := strings.Replace(c.Opts.Url, "#", "/", 1)
	songIdStr, err := utils.RegexSingleMatch(url, `id=([\d]+)`)
	if err != nil {
		return
	}
	songId, err := strconv.Atoi(songIdStr)
	if err != nil {
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
		Title:       songName,
		Description: utils.RegexSingleMatchIgnoreError(html, `"title": "(.+?)"`, ""),
	}

	audios = append(audios, &meta.Audio{
		Title:  songName,
		Artist: utils.RegexSingleMatchIgnoreError(html, `og:music:artist" content="(.+?)"[\s]*/>`, ""),
		Album:  utils.RegexSingleMatchIgnoreError(html, `og:music:album" content="(.+?)"[\s]*/>`, ""),
		Resource: &meta.Resource{
			Url: getSongUrl(songId),
			Headers: map[string]string{
				"user-agent": consts.UAMac,
				"referer":    c.Opts.Url,
			},
		},
	})
	return
}
