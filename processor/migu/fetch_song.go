package migu

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const (
	ApiUrlSong = "https://c.musicapp.migu.cn/MIGUM2.0/v1.0/content/resourceinfo.do?copyrightId=%s&resourceType=2"
)

func (c *Core) fetchFromSong() (mediaMeta *meta.MediaMeta, err error) {
	songId := utils.RegexSingleMatchIgnoreError(c.Opts.Url, `song/(.+)`, "0")
	if songId == "0" {
		return
	}

	songJson, err := utils.HttpGet(fmt.Sprintf(ApiUrlSong, songId), map[string]string{
		"user-agent": consts.UAMac,
		"referer":    c.Opts.Url,
	})
	if err != nil {
		return nil, err
	}

	songInfo := &SongInfo{}
	if err = json.Unmarshal([]byte(songJson), songInfo); err != nil {
		return nil, err
	}
	if len(songInfo.Resource) == 0 {
		return nil, err
	}

	songResource := songInfo.Resource[0]

	mediaMeta = &meta.MediaMeta{
		Title:        songResource.SongName,
		Description:  songResource.SongName,
		Duration:     utils.DurationStr2Second(songResource.Length),
		Album:        songResource.Album,
		Artist:       songResource.Singer,
		CoverUrl:     songResource.AlbumImgs[1].Img,
		Audios:       []meta.Audio{{Url: getSongUrlFromResources(songResource.RateFormats)}},
		ResourceType: consts.ResourceTypeAudio,
		Headers: map[string]string{
			"user-agent": consts.UAMac,
			"referer":    c.Opts.Url,
		},
	}

	return
}

func getSongUrlFromResources(resources []RateFormat) string {
	for _, resource := range resources {
		if resource.FormatType == "HQ" {
			url := "https://freetyst.nf.migu.cn/" + utils.RegexSingleMatchIgnoreError(resource.Url, `ftp:\/\/[^/]+\/(.*)`, "")
			return strings.ReplaceAll(url, "+", "%2B")
		}
	}
	return ""
}
