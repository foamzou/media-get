package bilibili

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const APISearch = "https://api.bilibili.com/x/web-interface/search/all/v2?page=1&page_size=10&platform=pc&single_column=0&keyword=%s&preload=true"
const APIFinger = "https://api.bilibili.com/x/frontend/finger/spi"
const SearchPageUrl = "https://search.bilibili.com/all"

func (c *Core) SearchSong() ([]*meta.SearchSongItem, error) {
	cookie, _ := utils.GetCookie(consts.SourceNameBilibili, SearchPageUrl, map[string]string{
		"User-Agent": consts.UAMac,
		"Referer":    "https://search.bilibili.com/",
	}, false)
	// In some case, the cookie is not required
	//if err != nil {
	//	return nil, err
	//}

	fingerJsonStr, err := utils.HttpGet(consts.SourceNameBilibili, APIFinger, map[string]string{
		"User-Agent": consts.UAMac,
		"Referer":    "https://search.bilibili.com/all",
		"Cookie":     cookie,
	})
	if err != nil {
		return nil, err
	}
	var fingerInfo FingerInfo
	err = json.Unmarshal([]byte(fingerJsonStr), &fingerInfo)
	if err != nil {
		return nil, err
	}
	cookie = fmt.Sprintf("buvid3=%s; buvid4=%s; %s", fingerInfo.Data.B3, fingerInfo.Data.B4, cookie)

	var searchSongItems []*meta.SearchSongItem
	api := fmt.Sprintf(APISearch, url.QueryEscape(c.Opts.Search.Keyword))
	jsonStr, err := utils.HttpGet(consts.SourceNameBilibili, api, map[string]string{
		"User-Agent": consts.UAMac,
		"Referer":    "https://search.bilibili.com/all",
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

	for _, item := range searchSongResponse.Data.Result {
		if item.ResultType != "video" {
			continue
		}
		for _, videoItem := range item.Data {
			searchSongItems = append(searchSongItems, &meta.SearchSongItem{
				Name:     utils.RemoveTagFromString(videoItem.Title),
				Artist:   videoItem.Author,
				Duration: utils.DurationStr2Second(videoItem.Duration),
				Url:      fmt.Sprintf("https://www.bilibili.com/video/%s", videoItem.Bvid),
				Source:   consts.SourceNameBilibili,
			})
		}
	}

	return searchSongItems, nil
}
