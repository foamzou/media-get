package netease

import (
	"encoding/json"
	"strings"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

type ProgramResource struct {
	MainSong struct {
		Name    string `json:"name"`
		Id      int    `json:"id"`
		Artists []struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
		} `json:"artists"`
		Duration int `json:"duration"`
	} `json:"mainSong"`
	Dj struct {
		UserId   int    `json:"userId"`
		Nickname string `json:"nickname"`
		Brand    string `json:"brand"`
	} `json:"dj"`
	Description string `json:"description"`
}

func (c *Core) fetchFromProgram() (mediaMeta *meta.MediaMeta, err error) {
	url := strings.Replace(c.Opts.Url, "#", "/", 1)
	html, err := fetchHtml(url)
	if err != nil {
		return
	}

	matchStr, err := utils.RegexSingleMatch(html, `program\-data.+?({[\s\S]+?)</textarea>`)
	if err != nil {
		return
	}
	resource := &ProgramResource{}
	err = json.Unmarshal([]byte(matchStr), resource)
	if err != nil {
		return
	}

	var artists []string
	for _, artist := range resource.MainSong.Artists {
		artists = append(artists, artist.Name)
	}

	mediaMeta = &meta.MediaMeta{
		Title:       resource.MainSong.Name,
		Description: resource.Description,
		Album:       resource.Dj.Brand,
		Artist:      strings.Join(artists, ", "),
		Audio:       meta.Audio{Url: getSongUrl(resource.MainSong.Id)},
		Headers: map[string]string{
			"user-agent": consts.UAMac,
			"referer":    c.Opts.Url,
		},
	}

	return
}
