package kuwo

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const (
	InfoAPI = "http://m.kuwo.cn/newh5/singles/songinfoandlrc?musicId=%d&httpsStatus=1&reqId=%s"
	PlayAPI = "http://www.kuwo.cn/api/v1/www/music/playUrl?mid=%d&type=music&httpsStatus=1&reqId=%s&plat=web_www"
)

type Core struct {
	Opts *args.Options
}

func (c *Core) IsMusicPlatform() bool {
	return true
}

func (c *Core) Domains() []string {
	return []string{"kuwo.cn"}
}

func (c *Core) GetSourceName() string {
	return consts.SourceNameKuwo
}

func (c *Core) FetchMetaAndResourceInfo() (mediaMeta *meta.MediaMeta, err error) {
	songID, err := getSongID(c.Opts.Url)
	if err != nil {
		return nil, err
	}
	songMeta, err := getSongMeta(songID)
	if err != nil {
		return nil, err
	}
	songUrl, err := c.getSongUrl(songID)
	if err != nil {
		return nil, err
	}

	songInfo := songMeta.Data.Songinfo
	mediaMeta = &meta.MediaMeta{
		Title:       songInfo.SongName,
		Description: songInfo.Album,
		Artist:      songInfo.Artist,
		Album:       songInfo.Album,
		Duration:    utils.DurationStr2Second(songInfo.SongTimeMinutes),
		CoverUrl:    songInfo.Pic,
		//IsTrial:      songInfo.IsFreePart == 1,
		ResourceType: consts.ResourceTypeAudio,
		Audios:       []meta.Audio{{Url: songUrl}},
	}
	return mediaMeta, err
}

func getSongMeta(songID int) (songMeta *SongMeta, err error) {
	reqID := utils.GenReqID()
	kuwoMetaJson, err := utils.HttpGet(consts.SourceNameKuwo, fmt.Sprintf(InfoAPI, songID, reqID), map[string]string{
		"user-agent": consts.UAMac,
	})
	if err != nil {
		return
	}
	songMeta = &SongMeta{}
	err = json.Unmarshal([]byte(kuwoMetaJson), &songMeta)
	if err != nil {
		return
	}
	return songMeta, nil
}

func (c *Core) getSongUrl(songID int) (songUrl string, err error) {
	reqID := utils.GenReqID()

	cookie, err := utils.GetCookie(consts.SourceNameKuwo, c.Opts.Url, map[string]string{
		"user-agent": consts.UAMac,
	}, false)
	if err != nil {
		return "", err
	}
	cookieKey := utils.RegexSingleMatchIgnoreError(cookie, "(Hm_.+?)=", "Hm_Iuvt_cdb524f42f0cer9b268e4v7y735ewrq2324")

	cp := &crypto{
		cookie: cookie,
		key:    cookieKey,
	}
	secret := cp.GenSecretHeader()

	kuwoMetaJson, err := utils.HttpGet(consts.SourceNameKuwo, fmt.Sprintf(PlayAPI, songID, reqID), map[string]string{
		"user-agent": consts.UAMac,
		"cookie":     cookie,
		"secret":     secret,
	})
	if err != nil {
		return
	}
	var songUrlMeta = struct {
		Data struct {
			Url string `json:"url"`
		} `json:"data"`
	}{}
	err = json.Unmarshal([]byte(kuwoMetaJson), &songUrlMeta)
	if err != nil {
		return
	}
	if songUrlMeta.Data.Url == "" {
		err = errors.New("the song play url is empty")
		return
	}
	return songUrlMeta.Data.Url, nil
}

func getSongID(url string) (songID int, err error) {
	songID = utils.RegexSingleMatchIntIgnoreError(url, "play_detail/([\\d]+)", 0)
	if songID == 0 {
		err = errors.New("can not match songID")
		return
	}
	return songID, nil
}
