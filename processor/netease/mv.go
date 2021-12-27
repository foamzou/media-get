package netease

import (
	urlLib "net/url"
	"strings"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

func (c *Core) fetchFromMV() (mediaMeta *meta.MediaMeta, audios []*meta.Audio, err error) {
	url := strings.Replace(c.Opts.Url, "#", "/", 1)
	html, err := fetchHtml(url)
	if err != nil {
		return
	}
	videoUrl, err := utils.RegexSingleMatch(html, `property="og:video" content="(.+?)"[\s]*/>`)
	if err != nil {
		return
	}
	videoUrl, err = urlLib.QueryUnescape(videoUrl)
	if err != nil {
		return
	}

	songName := utils.RegexSingleMatchIgnoreError(html, `"og:title" content="(.+?)"[\s]*/>`,
		utils.RegexSingleMatchIgnoreError(html, `"title": "(.+?)",`, ""),
	)

	mediaMeta = &meta.MediaMeta{
		Title:       songName,
		Description: utils.RegexSingleMatchIgnoreError(html, `"description": "(.+?)"`, ""),
	}

	audios = append(audios, &meta.Audio{
		Title: songName,
		Artist: utils.RegexSingleMatchIgnoreError(html, `artistName=(.+?)&`,
			utils.RegexSingleMatchIgnoreError(html, `data-author="(.+?)"`, ""),
		),
		Album: "网易Video",
		Resource: &meta.Resource{
			Url: videoUrl,
			Headers: map[string]string{
				"user-agent": consts.UAMac,
				"referer":    c.Opts.Url,
			},
		},
	})
	return
}
