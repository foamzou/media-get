package netease

import (
	urlLib "net/url"
	"strings"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

func (c *Core) fetchFromMV() (mediaMeta *meta.MediaMeta, err error) {
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
		Duration:    utils.RegexSingleMatchIntIgnoreError(html, `property="video:duration" content="(.+?)"/>`, 0),
		Album:       "网易Video",
		Artist: utils.RegexSingleMatchIgnoreError(html, `artistName=(.+?)&`,
			utils.RegexSingleMatchIgnoreError(html, `data-author="(.+?)"`, ""),
		),
		Audios:       []meta.Audio{{Url: videoUrl}},
		Videos:       []meta.Video{{Url: videoUrl}},
		ResourceType: consts.ResourceTypeVideo,
		Headers: map[string]string{
			"user-agent": consts.UAMac,
			"referer":    c.Opts.Url,
		},
	}

	return
}
