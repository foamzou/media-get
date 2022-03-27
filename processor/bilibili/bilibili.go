package bilibili

import (
	"encoding/json"
	"strings"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/logger"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

type Core struct {
	Opts *args.Options
}

const Album = "Bilibili"

func (c *Core) IsMusicPlatform() bool {
	return false
}

func (c *Core) Domain() string {
	return "bilibili.com"
}

func (c *Core) GetSourceName() string {
	return consts.SourceNameBilibili
}

func (c *Core) FetchMetaAndResourceInfo() (mediaMeta *meta.MediaMeta, err error) {
	html, err := fetchHtml(c.Opts.Url)
	if err != nil {
		return
	}

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
	mediaMeta = &meta.MediaMeta{
		Duration:     resource.Data.Dash.Duration,
		ResourceType: consts.ResourceTypeVideo,
	}
	metaJson := utils.RegexSingleMatchIgnoreError(html, `__INITIAL_STATE__=(.+?);\(function`, "{}")
	audioMeta := &AudioMeta{}
	err = json.Unmarshal([]byte(metaJson), audioMeta)
	if err != nil {
		logger.Warn("fetch meta json failed", err)
		mediaMeta.Title = utils.RegexSingleMatchIgnoreError(html, `<h1 title="(.+?)"`, utils.Md5(c.Opts.Url))
	} else {
		mediaMeta.Title = audioMeta.VideoData.Title
		mediaMeta.Description = audioMeta.VideoData.Description
	}

	audio := resource.Data.Dash.Audios[0]
	mediaMeta.Artist = getSinger(audioMeta)
	mediaMeta.Album = Album
	mediaMeta.CoverUrl = audioMeta.VideoData.Pic
	mediaMeta.Headers = map[string]string{
		"user-agent": consts.UAMac,
		"referer":    c.Opts.Url,
	}
	mediaMeta.Audios = append(mediaMeta.Audios, meta.Audio{
		Url:     audio.BaseUrl,
		BitRate: consts.BitRate128,
	})

	for _, bilibiliVideo := range resource.Data.Dash.Videos {
		mediaMeta.Videos = append(mediaMeta.Videos, meta.Video{
			Url:            bilibiliVideo.BaseUrl,
			Width:          bilibiliVideo.Width,
			Height:         bilibiliVideo.Height,
			Ratio:          getRatioById(bilibiliVideo.Id),
			NeedExtraAudio: true,
		})
	}

	return mediaMeta, nil
}

func getRatioById(id int) string {
	switch id {
	case 16:
		return consts.Ratio360
	case 32:
		return consts.Ratio480
	case 64:
		return consts.Ratio720
	case 80:
		return consts.Ratio1080
	case 112:
		return consts.Ratio1080Plus
	default:
		return consts.RatioUnknown
	}
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
