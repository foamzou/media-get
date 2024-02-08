package kuwo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const APISearch = "https://kuwo.cn/search/searchMusicBykeyWord?vipver=1&ft=music&client=kt&cluster=0&pn=0&rn=50&rformat=json&encoding=utf8&strategy=2012&mobi=1&issubtitle=1&show_copyright_off=1&all=%s"

func (c *Core) SearchSong() ([]*meta.SearchSongItem, error) {
	var searchSongItems []*meta.SearchSongItem
	api := fmt.Sprintf(APISearch, url.QueryEscape(c.Opts.Search.Keyword))

	jsonStr, err := utils.HttpGet(consts.SourceNameKuwo, api, map[string]string{
		"User-Agent": consts.UAMac,
		"Referer":    "https://m.kuwo.cn/",
	})
	if err != nil {
		return nil, err
	}

	var searchSongResponse SearchSongResponse
	err = json.Unmarshal([]byte(jsonStr), &searchSongResponse)
	if err != nil {
		return nil, err
	}

	for _, item := range searchSongResponse.Abslist {
		duration, err := strconv.Atoi(item.DURATION)
		if err != nil {
			duration = 0
		}
		searchSongItems = append(searchSongItems, &meta.SearchSongItem{
			Name:     item.NAME,
			Artist:   item.ARTIST,
			Album:    item.ALBUM,
			Duration: duration,
			Url:      fmt.Sprintf("https://www.kuwo.cn/play_detail/%s", strings.ReplaceAll(item.MUSICRID, "MUSIC_", "")),
			Source:   consts.SourceNameKuwo,
		})
	}

	return searchSongItems, nil
}
