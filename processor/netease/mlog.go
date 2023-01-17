package netease

import (
	"encoding/json"
	"fmt"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/utils"
)

const ApiMlogVideoConvertId = "https://music.163.com/weapi/mlog/video/convert/id"

type ConvertResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (c *Core) fetchFromMlog() (mediaMeta *meta.MediaMeta, err error) {
	id, err := utils.RegexSingleMatch(c.Opts.Url, `id=(.+?)&`)
	if err != nil {
		return
	}
	params, encSecKey, err := encrypt([]byte(fmt.Sprintf(`{"mlogId": "%s"}`, id)))
	if err != nil {
		return
	}
	videoIdJson, err := utils.PostForm(consts.SourceNameNetease, ApiMlogVideoConvertId,
		map[string]string{
			"params":    params,
			"encSecKey": encSecKey,
		},
		map[string]string{
			"user-agent": consts.UAMac,
			"referer":    c.Opts.Url,
		})

	if err != nil {
		return
	}

	resp := &ConvertResponse{}
	if err = json.Unmarshal([]byte(videoIdJson), &resp); err != nil {
		return
	}

	// use video flow
	c.Opts.Url = "https://music.163.com/#/video?id=" + resp.Data
	return c.fetchFromMV()
}
