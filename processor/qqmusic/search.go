package qqmusic

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

//nolint
const APISearch = "https://shc.y.qq.com/soso/fcgi-bin/search_for_qq_cp?format=json&inCharset=utf-8&outCharset=utf-8&notice=0&platform=h5&needNewCode=1&w=%s&perpage=20&n=20&p=1&remoteplace=txt.mqq.all"

func (c *Core) SearchSong() ([]*meta.SearchSongItem, error) {
	var searchSongItems []*meta.SearchSongItem
	api := fmt.Sprintf(APISearch, url.QueryEscape(c.Opts.Search.Keyword))

	jsonStr, err := utils.HttpGet(api, map[string]string{
		"User-Agent": consts.UAAndroid,
		"Referer":    "https://i.y.qq.com/",
	})
	if err != nil {
		return nil, err
	}

	var searchSongResponse SearchSongResponse
	err = json.Unmarshal([]byte(jsonStr), &searchSongResponse)
	if err != nil {
		return nil, err
	}

	for _, item := range searchSongResponse.Data.Song.List {
		searchSongItems = append(searchSongItems, &meta.SearchSongItem{
			Name:           item.Songname,
			Artist:         item.Singer[0].Name,
			Album:          item.Albumname,
			Duration:       item.Interval,
			Url:            fmt.Sprintf("https://y.qq.com/n/ryqq/songDetail/%s", item.Songmid),
			CannotDownload: item.Pay.Payplay == 1,
			Source:         consts.SourceNameQq,
		})
	}

	return searchSongItems, nil
}
