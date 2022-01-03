package youtube

import (
	"encoding/json"
	"strings"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const ApiGetItemInfo = "https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids="

type Core struct {
	Opts *args.Options
}

func (c *Core) FetchMetaAndResourceInfo() (mediaMeta *meta.MediaMeta, err error) {
	htmlInfo, err := utils.HttpGet(c.Opts.Url, map[string]string{
		"user-agent": consts.UAMac,
		"referer":    "https://www.youtube.com/",
	})
	if err != nil {
		return
	}

	jsonInfo, err := utils.RegexSingleMatch(htmlInfo, `var ytInitialPlayerResponse =(.+?);var meta`)
	if err != nil {
		return
	}
	youtubeResponse := &Response{}
	if err = json.Unmarshal([]byte(jsonInfo), youtubeResponse); err != nil {
		return
	}

	audio, videos := getAudioAndVideo(youtubeResponse)

	mediaMeta = &meta.MediaMeta{
		Title:        youtubeResponse.VideoDetails.Title,
		Description:  youtubeResponse.VideoDetails.ShortDescription,
		Duration:     utils.ConvertString2Int(youtubeResponse.VideoDetails.LengthSeconds, 0),
		Album:        "Youtube",
		Artist:       youtubeResponse.VideoDetails.Author,
		CoverUrl:     getCoverUrl(youtubeResponse),
		Audios:       []meta.Audio{audio},
		Videos:       videos,
		ResourceType: consts.ResourceTypeVideo,
		Headers: map[string]string{
			"user-agent": consts.UAMac,
			"referer":    c.Opts.Url,
		},
	}

	return
}

func getCoverUrl(youtubeResponse *Response) string {
	return youtubeResponse.VideoDetails.Thumbnail.Thumbnails[len(youtubeResponse.VideoDetails.Thumbnail.Thumbnails)-1].Url
}

func getAudioAndVideo(youtubeResponse *Response) (audio meta.Audio, videos []meta.Video) {
	for _, media := range youtubeResponse.StreamingData.AdaptiveFormats {
		if !strings.Contains(media.MimeType, "mp4") {
			continue
		}
		if strings.Contains(media.MimeType, "audio") {
			audio = meta.Audio{
				Url:     media.Url,
				BitRate: media.Bitrate / 1024,
			}
		}
		if strings.Contains(media.MimeType, "video") {
			videos = append(videos, meta.Video{
				Url:            media.Url,
				Width:          media.Width,
				Height:         media.Height,
				Ratio:          media.QualityLabel,
				NeedExtraAudio: true,
			})
		}
	}
	return
}
