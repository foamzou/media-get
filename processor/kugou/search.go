package kugou

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const APISearch = "https://mobiles.kugou.com/api/v3/search/song?format=json&keyword=%s&page=1&pagesize=30&showtype=1"

func (c *Core) SearchSong() ([]*meta.SearchSongItem, error) {
	var searchSongItems []*meta.SearchSongItem
	api := fmt.Sprintf(APISearch, url.QueryEscape(c.Opts.Search.Keyword))

	jsonStr, err := utils.HttpGet(api, map[string]string{
		"User-Agent": consts.UAAndroid,
		"Referer":    "https://m3ws.kugou.com/",
	})
	if err != nil {
		return nil, err
	}

	var searchSongResponse SearchSongResponse
	err = json.Unmarshal([]byte(jsonStr), &searchSongResponse)
	if err != nil {
		return nil, err
	}

	for _, item := range searchSongResponse.Data.Info {
		searchSongItems = append(searchSongItems, &meta.SearchSongItem{
			Name:              item.Songname,
			Artist:            item.Singername,
			Album:             item.AlbumName,
			Duration:          item.Duration,
			ResourceForbidden: item.Privilege >= 10,
			Url:               fmt.Sprintf("https://www.kugou.com/song/#hash=%s&album_id=%s", item.Hash, item.AlbumId),
			Source:            consts.SourceNameKugou,
		})
	}

	return searchSongItems, nil
}
