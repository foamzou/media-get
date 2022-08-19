package douyin

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const SearchPage = "https://www.douyin.com/search/%s"

//nolint
const APISearchAPI = "https://www.douyin.com/aweme/v1/web/general/search/single/?"
const APISearchParams = "device_platform=webapp&aid=6383&channel=channel_pc_web&search_channel=aweme_general&sort_type=0&publish_time=0&keyword=%s&search_source=normal_search&query_correct_type=1&is_filter_search=0&from_group_id=&offset=0&count=10&cookie_enabled=true&screen_width=1920&screen_height=1080&browser_language=zh-CN&browser_platform=MacIntel&browser_name=Chrome&browser_version=99.0.4844.83&browser_online=true"

func (c *Core) SearchSong() ([]*meta.SearchSongItem, error) {

	var searchSongItems []*meta.SearchSongItem

	kw := url.QueryEscape(c.Opts.Search.Keyword)
	searchUrl := fmt.Sprintf(SearchPage, kw)
	apiQuery := fmt.Sprintf(APISearchParams, kw)
	xb, _ := genXB(apiQuery, consts.UAMac)
	apiQuery += "&X-Bogus=" + xb
	api := APISearchAPI + apiQuery

	cookie, err := utils.GetCookie(searchUrl, map[string]string{
		"User-Agent": consts.UAMac,
	}, false)
	if err != nil {
		return nil, err
	}

	cookie, err = utils.GetCookie(searchUrl, map[string]string{
		"User-Agent": consts.UAMac,
		"Referer":    searchUrl,
		"Cookie":     cookie,
	}, true)
	if err != nil {
		return nil, err
	}

	jsonStr, err := utils.HttpGet(api, map[string]string{
		"User-Agent": consts.UAMac,
		"Referer":    searchUrl,
		"Cookie":     cookie,
	})

	if err != nil {
		return nil, err
	}

	var searchSongResponse SearchSongResponse
	err = json.Unmarshal([]byte(jsonStr), &searchSongResponse)
	if err != nil {
		return nil, err
	}

	for _, item := range searchSongResponse.Data {
		searchSongItems = append(searchSongItems, &meta.SearchSongItem{
			Name:     utils.RemoveTagFromString(item.AwemeInfo.Desc),
			Artist:   item.AwemeInfo.Author.Nickname,
			Duration: item.AwemeInfo.Video.Duration / 1000,
			Url:      fmt.Sprintf("https://www.douyin.com/video/%s", item.AwemeInfo.AwemeId),
			Source:   consts.SourceNameDouyin,
		})
	}

	return searchSongItems, nil
}
