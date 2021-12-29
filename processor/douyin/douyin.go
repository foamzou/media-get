package douyin

import (
	"encoding/json"
	"fmt"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

// https://v.douyin.com/8BxvWBm/

// curl -H"user-agent:Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Mobile Safari/537.36"
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
		Album:       "抖音Video",
		Artist:      metaItem.Author.Nickname,
		Audio:       meta.Audio{Url: metaItem.Music.PlayUrl.Uri},
		Headers: map[string]string{
			"user-agent": consts.UAAndroid,
			"referer":    redirectUrl,
		},
	}

	return
}

func getVideoId(inputUrl string) (string, error) {
	return utils.RegexSingleMatch(inputUrl, `video/(.+?)/`)
}
