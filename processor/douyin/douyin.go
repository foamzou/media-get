package douyin

import (
	"encoding/json"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const ApiGetItemInfo = "https://www.douyin.com/aweme/v1/web/aweme/detail/?device_platform=webapp"

type Core struct {
	Opts *args.Options
}

func (c *Core) IsMusicPlatform() bool {
	return false
}

func (c *Core) Domains() []string {
	return []string{"douyin.com"}
}

func (c *Core) GetSourceName() string {
	return consts.SourceNameDouyin
}

func (c *Core) FetchMetaAndResourceInfo() (mediaMeta *meta.MediaMeta, err error) {
	var videoId string
	var redirectUrl string

	videoId, _ = getVideoId(c.Opts.Url)
	redirectUrl = c.Opts.Url
	var cookie string

	// from share url
	if videoId == "" {
		videoId, cookie, redirectUrl, err = getRedirectInfoFromShareUrl(c.Opts.Url)
		if err != nil {
			return nil, err
		}
	} else {
		// from pc url, use the url to get the cookie contains the __ac_signature
		// please let me know if you have a better way to get the __ac_signature
		hardCodeShareUrl := "https://v.douyin.com/8BxvWBm/"
		_, cookie, _, err = getRedirectInfoFromShareUrl(hardCodeShareUrl)
		if err != nil {
			return nil, err
		}
	}

	// get detail
	query := "channel=channel_pc_web&pc_client_type=1&aweme_id=" + videoId + "&msToken="
	xb, _ := genXB(query, consts.UAAndroid)
	query += "&X-Bogus=" + xb

	metaInfo, err := utils.HttpGet(consts.SourceNameDouyin, ApiGetItemInfo+query, map[string]string{
		"User-Agent":      consts.UAAndroid,
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"Accept-Encoding": "deflate, gzip",
		"Accept":          "*/*",
		"Cookie":          cookie,
		"Referer":         redirectUrl,
	})
	if err != nil {
		return
	}
	metaResponse := &MetaResponse{}
	if err = json.Unmarshal([]byte(metaInfo), metaResponse); err != nil {
		return
	}

	metaItem := metaResponse.Detail

	mediaMeta = &meta.MediaMeta{
		Title:       metaItem.Desc,
		Description: metaItem.Desc,
		Duration:    metaItem.Duration / 1000,
		Album:       "抖音Video",
		Artist:      metaItem.Author.Nickname,
		Audios:      []meta.Audio{{Url: metaItem.Music.PlayUrl.Uri}},
		CoverUrl:    metaItem.Video.OriginCover.UrlList[0],
		Videos: []meta.Video{{
			Url:    metaItem.Video.PlayAddr.UrlList[0],
			Width:  metaItem.Video.Width,
			Height: metaItem.Video.Height,
			Ratio:  metaItem.Video.Ratio,
		}},
		ResourceType: consts.ResourceTypeVideo,
		Headers: map[string]string{
			"user-agent": consts.UAAndroid,
			"referer":    redirectUrl,
		},
	}

	return
}

func getRedirectInfoFromShareUrl(url string) (string, string, string, error) {
	redirectUrl, err := utils.GetLocation(consts.SourceNameDouyin, url, map[string]string{
		"user-agent": consts.UAAndroid,
	})

	if err != nil {
		return "", "", "", err
	}

	videoId, err := getVideoId(redirectUrl)
	if err != nil {
		return "", "", "", err
	}

	cookie, err := utils.GetCookie(consts.SourceNameDouyin, redirectUrl, map[string]string{
		"User-Agent": consts.UAAndroid,
	}, false)
	if err != nil {
		return "", "", "", err
	}

	return videoId, cookie, redirectUrl, nil
}

func getVideoId(inputUrl string) (string, error) {
	return utils.RegexSingleMatch(inputUrl, `video/([^/?#]+)`)
}
