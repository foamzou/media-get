package qmkg

import (
	"fmt"
	"strings"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

type Core struct {
	Opts *args.Options
}

func (c *Core) SearchSong() (searchItems []*meta.SearchSongItem, err error) {
	return
}

func (c *Core) IsMusicPlatform() bool {
	return false
}

func (c *Core) Domains() []string {
	return []string{
		"kg.qq.com",
		"kg2.qq.com",
		"kg3.qq.com",
		"kg4.qq.com",
		"kg5.qq.com",
		"kg6.qq.com",
		"kg7.qq.com",
		"kg8.qq.com",
		"kg9.qq.com",
	}
}

func (c *Core) GetSourceName() string {
	return consts.SourceNameQMKG
}

func (c *Core) FetchMetaAndResourceInfo() (mediaMeta *meta.MediaMeta, err error) {

	html, err := utils.HttpGet(c.Opts.Url, map[string]string{
		"user-agent": consts.UAAndroid,
		"referer":    c.Opts.Url,
	})
	if err != nil {
		return
	}

	title, err := utils.RegexSingleMatch(html, "<title>([\\s\\S]+?)</title>")
	if err != nil {
		return
	}
	title = strings.TrimSpace(title)
	artistName := strings.Split(title, "-")[0]
	songName := strings.Split(strings.ReplaceAll(title, artistName+"-", ""), "-全民")[0]
	duration := utils.RegexSingleMatchIntIgnoreError(html, "\"segment_end\":([\\d]+),", 0)
	coverUrl := utils.RegexSingleMatchIgnoreError(html, "\"cover\":\"(.+?)\",", "")
	playUrl, err := utils.RegexSingleMatch(html, "\"playurl\":\"(.+?)\",")
	if err != nil {
		return
	}
	if playUrl == "" {
		err = fmt.Errorf("playUrl is empty")
		return
	}

	mediaMeta = &meta.MediaMeta{
		Title:        songName,
		Duration:     duration / 1000,
		Album:        "全民K歌",
		Artist:       artistName,
		Audios:       []meta.Audio{{Url: playUrl}},
		CoverUrl:     coverUrl,
		ResourceType: consts.ResourceTypeAudio,
		Headers: map[string]string{
			"user-agent": consts.UAAndroid,
			"referer":    c.Opts.Url,
		},
	}

	return
}
