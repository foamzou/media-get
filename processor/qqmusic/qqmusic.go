package qqmusic

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const (
	SongInfoAPI = "https://c.y.qq.com/v8/fcg-bin/fcg_play_single_song.fcg?format=json&platform=yqq&songmid=%s"
	SongHtmlURL = "https://i.y.qq.com/v8/playsong.html?ADTAG=ryqq.songDetail&songmid=%s&songid=0&songtype=0"
)

type Core struct {
	Opts *args.Options
}

func (c *Core) IsMusicPlatform() bool {
	return true
}

func (c *Core) Domain() string {
	return "qq.com"
}

func (c *Core) GetSourceName() string {
	return consts.SourceNameQq
}

func (c *Core) FetchMetaAndResourceInfo() (mediaMeta *meta.MediaMeta, err error) {
	songMid, err := getSongMid(c.Opts.Url)
	if err != nil {
		return nil, err
	}
	songMeta, err := getSongMeta(songMid)
	if err != nil {
		return nil, err
	}

	songUrl, coverUrl, err := getSongUrlAndCoverUrl(songMid)
	if err != nil {
		return nil, err
	}

	songInfo := songMeta.Data[0]
	mediaMeta = &meta.MediaMeta{
		Title:       songInfo.Title,
		Description: songInfo.Subtitle,
		Artist:      songInfo.Singer[0].Name,
		Album:       songInfo.Album.Name,
		Duration:    songInfo.Interval,
		CoverUrl:    coverUrl,
		//IsTrial:      songInfo.IsFreePart == 1,
		ResourceType: consts.ResourceTypeAudio,
		Audios:       []meta.Audio{{Url: songUrl}},
	}
	return mediaMeta, err
}

func getSongUrlAndCoverUrl(songMid string) (songUrl, coverUrl string, err error) {
	html, err := utils.HttpGet(fmt.Sprintf(SongHtmlURL, songMid), map[string]string{
		"user-agent": consts.UAAndroid,
	})
	if err != nil {
		return
	}
	jsonStr := utils.RegexSingleMatchIgnoreError(html, "window.__ssrFirstPageData__ =([\\s\\S]+?)</script", "")

	var jsonStruct struct {
		SongList []struct {
			Url string `json:"url"`
		} `json:"songList"`
		MetaData struct {
			Image string `json:"image"`
		} `json:"metaData"`
	}
	err = json.Unmarshal([]byte(jsonStr), &jsonStruct)
	if err == nil {
		if len(jsonStruct.SongList) > 0 {
			songUrl = jsonStruct.SongList[0].Url
		}
		coverUrl = jsonStruct.MetaData.Image
	}
	if songUrl == "" {
		// fallback
		songUrl = utils.RegexSingleMatchIgnoreError(html, "\"url\":[\\s]*\"(http.+?music.+?)\"", "")
		songUrl = strings.ReplaceAll(songUrl, "\\u002F", "/")
	}
	if songUrl == "" {
		err = errors.New("url not found")
		return
	}
	return songUrl, coverUrl, nil
}

func getSongMid(url string) (string, error) {
	var songMid string
	songMid = utils.RegexSingleMatchIgnoreError(url, "songDetail\\/(\\w+)", "")
	if songMid != "" {
		return songMid, nil
	}
	songMid = utils.RegexSingleMatchIgnoreError(url, "songmid=(\\w+)", "")
	if songMid != "" {
		return songMid, nil
	}
	return "", errors.New("songMid not found")
}

func getSongMeta(songMid string) (songMeta *SongMeta, err error) {
	songAPI := fmt.Sprintf(SongInfoAPI, songMid)
	qqMetaJson, err := utils.HttpGet(songAPI, map[string]string{
		"user-agent": consts.UAMac,
	})
	if err != nil {
		return
	}
	songMeta = &SongMeta{}
	err = json.Unmarshal([]byte(qqMetaJson), &songMeta)
	if err != nil {
		return
	}
	return songMeta, nil
}
