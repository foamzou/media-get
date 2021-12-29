package bilibili

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

type Core struct {
	Opts *args.Options
}

const Album = "Bilibili"

func (c *Core) FetchMetaAndResourceInfo() (mediaMeta *meta.MediaMeta, err error) {
	html, err := fetchHtml(c.Opts.Url)

	// audio resource
	matchStr, err := utils.RegexSingleMatch(html, `window.__playinfo__=(.+?)<\/script`)
	if err != nil {
		return
	}
	resource := &AudioResource{}
	err = json.Unmarshal([]byte(matchStr), resource)
	if err != nil {
		return
	}

	// audio meta
	mediaMeta = &meta.MediaMeta{}
	metaJson := utils.RegexSingleMatchIgnoreError(html, `__INITIAL_STATE__=(.+?);\(function`, "{}")
	audioMeta := &AudioMeta{}
	err = json.Unmarshal([]byte(metaJson), audioMeta)
	if err != nil {
		fmt.Println("fetch meta json failed", err)
		mediaMeta.Title = utils.RegexSingleMatchIgnoreError(html, `<h1 title="(.+?)"`, utils.Md5(c.Opts.Url))
	} else {
		mediaMeta.Title = audioMeta.VideoData.Title
		mediaMeta.Description = audioMeta.VideoData.Description
	}

	audio := resource.Data.Dash.Audio[0]
	mediaMeta.Artist = getSinger(audioMeta)
	mediaMeta.Album = Album
	mediaMeta.Headers = map[string]string{
		"user-agent": consts.UAMac,
		"referer":    c.Opts.Url,
	}
	mediaMeta.Audio.Url = audio.BaseUrl

	return mediaMeta, nil
}

func getSinger(audioMeta *AudioMeta) string {
	var name string
	if len(audioMeta.VideoData.Staff) == 0 {
		name = audioMeta.VideoData.Owner.Name
	} else {
		var names []string
		for _, staff := range audioMeta.VideoData.Staff {
			names = append(names, staff.Name)
		}
		name = strings.Join(names, ", ")
	}
	if strings.TrimSpace(name) == "" {
		return "unknown"
	}

	return name
}

func fetchHtml(url string) (string, error) {
	return utils.HttpGet(url, map[string]string{
		"user-agent": consts.UAMac,
	})
}
