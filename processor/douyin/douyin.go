package douyin

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const ApiGetItemInfo = "https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids="

type Core struct {
	Opts *args.Options
}

func (c *Core) FetchMetaAndResourceInfo() (mediaMeta *meta.MediaMeta, err error) {
	redirectUrl, err := utils.GetLocation(c.Opts.Url, map[string]string{
		"user-agent": consts.UAAndroid,
	})

	if err != nil {
		return
	}

	videoId, err := getVideoId(redirectUrl)
	if err != nil {
		return
	}

	metaInfo, err := utils.HttpGet(ApiGetItemInfo+videoId, map[string]string{
		"user-agent": consts.UAAndroid,
		"referer":    redirectUrl,
	})
	if err != nil {
		return
	}

	metaResponse := &MetaResponse{}
	if err = json.Unmarshal([]byte(metaInfo), metaResponse); err != nil {
		return
	}
	if len(metaResponse.ItemList) == 0 {
		err = fmt.Errorf("not resource")
		return
	}

	metaItem := metaResponse.ItemList[0]

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
			Ratio:  formatRatio(metaItem.Video.Ratio),
		}},
		ResourceType: consts.ResourceTypeVideo,
		Headers: map[string]string{
			"user-agent": consts.UAAndroid,
			"referer":    redirectUrl,
		},
	}

	return
}

func formatRatio(dyRatio string) string {
	return strings.Replace(dyRatio, "p", "P", 1)
}

func getVideoId(inputUrl string) (string, error) {
	return utils.RegexSingleMatch(inputUrl, `video/(.+?)/`)
}
