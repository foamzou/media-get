package kugou

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const (
	InfoAPI  = "https://www.kugou.com/yy/index.php?r=play/getdata&hash=%s&album_id=%d&_=%d&mid=%s"
	InfoAPI2 = "https://www.kugou.com/yy/index.php?r=play/getdata&hash=%s&album_id=%d&album_audio_id=%d&_=%d&mid=%s"
)

type Core struct {
	Opts *args.Options
}

func (c *Core) IsMusicPlatform() bool {
	return true
}

func (c *Core) Domains() []string {
	return []string{"kugou.com"}
}

func (c *Core) GetSourceName() string {
	return consts.SourceNameKugou
}

func (c *Core) FetchMetaAndResourceInfo() (mediaMeta *meta.MediaMeta, err error) {
	hash, albumID, mixSongID, err := getSongIdInfo(c.Opts.Url)
	if err != nil {
		return nil, err
	}
	songMeta, err := tryGetSongMeta(hash, albumID, mixSongID)
	if err != nil {
		return nil, err
	}

	songData := songMeta.Data
	mediaMeta = &meta.MediaMeta{
		Title:        songData.SongName,
		Description:  songData.AudioName,
		Artist:       songData.AuthorName,
		Album:        songData.AlbumName,
		Duration:     songData.Timelength / 1000,
		CoverUrl:     songData.Img,
		IsTrial:      songData.IsFreePart == 1,
		ResourceType: consts.ResourceTypeAudio,
		Audios:       []meta.Audio{{Url: songData.PlayUrl}},
	}
	return mediaMeta, err
}

func tryGetSongMeta(hash string, albumID, mixSongID int) (songMeta *SongMeta, err error) {
	ms := time.Now().UnixNano() / 1e6
	mid := utils.Md5(strconv.FormatInt(ms, 10))
	songMeta, err = getSongMeta(fmt.Sprintf(InfoAPI, hash, albumID, ms, mid))
	if err == nil {
		return songMeta, nil
	}

	songMeta, err = getSongMeta(fmt.Sprintf(InfoAPI2, hash, mixSongID, mixSongID, ms, mid))
	if err != nil {
		return nil, err
	}
	return songMeta, nil
}

func getSongMeta(url string) (songMeta *SongMeta, err error) {
	kugouMetaJson, err := utils.HttpGet(consts.SourceNameKugou, url, map[string]string{
		"user-agent": consts.UAMac,
	})
	if err != nil {
		return
	}
	songMeta = &SongMeta{}
	err = json.Unmarshal([]byte(kugouMetaJson), &songMeta)
	if err != nil {
		return
	}
	return songMeta, nil
}

func getSongIdInfo(url string) (hash string, albumID int, mixSongID int, err error) {
	if utils.RegexSingleMatchIgnoreError(url, "song/#(hash=.+?&album_id=\\d+)", "") != "" {
		hash = utils.RegexSingleMatchIgnoreError(url, "hash=(.+?)&album_id=\\d+", "")
		albumID = utils.RegexSingleMatchIntIgnoreError(url, "hash=.+?&album_id=(\\d+)", 0)
		if hash == "" || albumID == 0 {
			return "", 0, 0, errors.New("not parse param required")
		}
		mixSongID = albumID
		return
	}
	html, err := utils.HttpGet(consts.SourceNameKugou, url, map[string]string{
		"user-agent": consts.UAMac,
	})
	if err != nil {
		return
	}

	match, err := utils.RegexSingleMatch(html, "\\[({\"hash\".+?})\\]")
	if err != nil {
		return
	}

	var songIDInfo struct {
		Hash      string `json:"hash"`
		AlbumID   int    `json:"album_id"`
		MixSongID int    `json:"mixsongid"`
	}

	err = json.Unmarshal([]byte(match), &songIDInfo)
	if err != nil {
		return
	}
	return songIDInfo.Hash, songIDInfo.AlbumID, songIDInfo.MixSongID, nil
}
