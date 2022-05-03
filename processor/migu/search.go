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
	keyword := c.Opts.Search.Keyword
	// 如果最后一个字符是 . 则移除。
	// case：I Remember F.I.R 与 I Remember F.I.R.  的搜索结果不一样
	if keyword[len(keyword)-1:] == "." {
		keyword = keyword[:len(keyword)-1]
	}
	api := fmt.Sprintf(APISearch, url.QueryEscape(keyword))

	ua := consts.UAAndroid
	jsonStr, err := utils.HttpGet(api, map[string]string{
		"User-Agent": ua,
		"By":         utils.Md5(ua),
		"Referer":    "https://m.music.migu.cn/v4/search",
	})
	fmt.Println(jsonStr)
	if err != nil {
		return nil, err
	}

	var searchSongResponse SearchSongResponse
	err = json.Unmarshal([]byte(jsonStr), &searchSongResponse)
	if err != nil {
		return nil, err
	}

	for _, item := range searchSongResponse.Data.SongsData.Items {
		artist := ""
		if len(item.Singers) > 0 {
			artist = item.Singers[0].Name
		}
		searchSongItems = append(searchSongItems, &meta.SearchSongItem{
			Name:     item.Name,
			Artist:   artist,
			Album:    item.Album.Name,
			Duration: 0, // unknown in the API
			Url:      fmt.Sprintf("https://music.migu.cn/v3/music/song/%s", item.CopyrightId),
			Source:   consts.SourceNameMigu,
		})
	}

	return searchSongItems, nil
}
