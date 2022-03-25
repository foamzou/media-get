package migu

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const APISearch = "https://m.music.migu.cn/migumusic/h5/search/all?text=%s&pageNo=1&pageSize=30"

func (c *Core) SearchSong() ([]*meta.SearchSongItem, error) {
	var searchSongItems []*meta.SearchSongItem
	api := fmt.Sprintf(APISearch, url.QueryEscape(c.Opts.Search.Keyword))

	ua := consts.UAAndroid
	jsonStr, err := utils.HttpGet(api, map[string]string{
		"User-Agent": ua,
		"By":         utils.Md5(ua),
		"Referer":    "https://m.music.migu.cn/v4/search",
	})
	if err != nil {
		return nil, err
	}

	var searchSongResponse SearchSongResponse
	err = json.Unmarshal([]byte(jsonStr), &searchSongResponse)
	if err != nil {
		return nil, err
	}

	for _, item := range searchSongResponse.Data.SongsData.Items {
		searchSongItems = append(searchSongItems, &meta.SearchSongItem{
			Name:     item.Name,
			Artist:   item.Singers[0].Name,
			Album:    item.Album.Name,
			Duration: 0, // unknown in the API
			Url:      fmt.Sprintf("https://music.migu.cn/v3/music/song/%s", item.CopyrightId),
			Source:   "migu",
		})
	}

	return searchSongItems, nil
}
