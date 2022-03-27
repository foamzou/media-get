package netease

import (
	"encoding/json"
	"fmt"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const APISearch = "https://music.163.com/weapi/cloudsearch/get/web"

func (c *Core) SearchSong() ([]*meta.SearchSongItem, error) {
	var searchSongItems []*meta.SearchSongItem

	params, encSecKey, err := encrypt([]byte(fmt.Sprintf(`{"s": "%s", "type": 1, "limit": "10", "offset": 0, "total": true}`, c.Opts.Search.Keyword)))
	if err != nil {
		return nil, err
	}
	jsonStr, err := utils.PostForm(APISearch,
		map[string]string{
			"params":    params,
			"encSecKey": encSecKey,
		},
		map[string]string{
			"user-agent": consts.UAMac,
			"referer":    "https://music.163.com/search/",
		})

	if err != nil {
		return nil, err
	}

	var searchSongResponse SearchSongResponse
	err = json.Unmarshal([]byte(jsonStr), &searchSongResponse)
	if err != nil {
		return nil, err
	}

	for _, item := range searchSongResponse.Result.Songs {
		searchSongItems = append(searchSongItems, &meta.SearchSongItem{
			Name:     item.Name,
			Artist:   item.Ar[0].Name,
			Album:    item.Al.Name,
			Duration: item.Dt / 1000,
			Url:      fmt.Sprintf("https://music.163.com/#/song?id=%d", item.Id),
			Source:   consts.SourceNameNetease,
		})
	}

	return searchSongItems, nil
}
