package migu

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const (
	ApiUrlSong = "https://c.musicapp.migu.cn/MIGUM2.0/v1.0/content/resourceinfo.do?copyrightId=%s&resourceType=2"
)

func (c *Core) fetchFromSong(url string) (mediaMeta *meta.MediaMeta, err error) {
	songId := utils.RegexSingleMatchIgnoreError(url, `song/(.+)`, "0")
	if songId == "0" {
		return
	}

	songJson, err := utils.HttpGet(consts.SourceNameMigu, fmt.Sprintf(ApiUrlSong, songId), map[string]string{
		"user-agent": consts.UAMac,
		"referer":    url,
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

	songUrl, audioUrlErr := getAudioUrl(songResource.CopyrightId)

	mediaMeta = &meta.MediaMeta{
		Title:        songResource.SongName,
		Description:  songResource.SongName,
		Duration:     utils.DurationStr2Second(songResource.Length),
		Album:        songResource.Album,
		Artist:       songResource.Singer,
		CoverUrl:     songResource.AlbumImgs[1].Img,
		Audios:       []meta.Audio{{Url: songUrl, NotAvailable: audioUrlErr != nil}},
		ResourceType: consts.ResourceTypeAudio,
		Headers: map[string]string{
			"user-agent": consts.UAMac,
			"referer":    url,
		},
	}

	return
}

func getAudioUrl(copyrightID string) (string, error) {
	songResourceInfoUrl := fmt.Sprintf("https://c.musicapp.migu.cn/MIGUM2.0/v1.0/content/resourceinfo.do?copyrightId=%s&resourceType=0", copyrightID)

	header := map[string]string{
		"user-agent": consts.UAAndroid,
	}

	resp, err := utils.HttpGet(consts.SourceNameMigu, songResourceInfoUrl, header)
	if err != nil {
		return "", err
	}
	var songResources SongResources
	if err = json.Unmarshal([]byte(resp), &songResources); err != nil {
		return "", err
	}
	if len(songResources.Resource) == 0 {
		return "", err
	}

	audioUrl := songResources.Resource[0].AudioUrl

	pattern := `\/彩铃\/[^\/]+\/`
	regex := regexp.MustCompile(pattern)

	audioUrl = regex.ReplaceAllString(audioUrl, "/标清高清/MP3_320_16_Stero/")

	audioUrl = "https://freetyst.nf.migu.cn/" + utils.RegexSingleMatchIgnoreError(audioUrl, `ftp:\/\/[^/]+\/(.*)`, "")
	audioUrl = strings.ReplaceAll(audioUrl, "+", "%2B")
	// 咪咕改版需要登录后，部分音源已经不可用，需要检查一下
	if err = utils.HttpHead(consts.SourceNameMigu, audioUrl, header); err != nil {
		return audioUrl, err
	}

	return audioUrl, nil
}
