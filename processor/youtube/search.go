package youtube

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const (
	SearchPage = "https://www.youtube.com/results?search_query=%s"
)

func (c *Core) SearchSong() ([]*meta.SearchSongItem, error) {
	var searchSongItems []*meta.SearchSongItem
	searchUrl := fmt.Sprintf(SearchPage, url.QueryEscape(c.Opts.Search.Keyword))

	htmlInfo, err := utils.HttpGet(consts.SourceNameYoutube, searchUrl, map[string]string{
		"user-agent": consts.UAMac,
		"referer":    "https://www.youtube.com/",
	})
	if err != nil {
		return nil, err
	}

	jsonStr, err := utils.RegexSingleMatch(htmlInfo, `var ytInitialData =(.+?);</script>`)
	if err != nil {
		return nil, err
	}

	var searchSongResponse SearchSongResponse
	err = json.Unmarshal([]byte(jsonStr), &searchSongResponse)
	if err != nil {
		return nil, err
	}

	searchList := searchSongResponse.Contents.TwoColumnSearchResultsRenderer.PrimaryContents.SectionListRenderer.Contents[0].ItemSectionRenderer.Contents

	for _, item := range searchList {
		if len(item.VideoRenderer.Title.Runs) == 0 || len(item.VideoRenderer.OwnerText.Runs) == 0 {
			continue
		}
		searchSongItems = append(searchSongItems, &meta.SearchSongItem{
			Name:     item.VideoRenderer.Title.Runs[0].Text,
			Artist:   item.VideoRenderer.OwnerText.Runs[0].Text,
			Duration: utils.DurationStr2Second(item.VideoRenderer.LengthText.SimpleText),
			Url:      fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.VideoRenderer.VideoId),
			Source:   consts.SourceNameYoutube,
		})
	}

	return searchSongItems, nil
}
