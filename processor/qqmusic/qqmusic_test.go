package qqmusic

import (
	"testing"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/test_helper"
	"github.com/foamzou/audio-get/utils"
)

func TestCore_FetchMetaAndResourceInfo(t *testing.T) {
	utils.InitConfig()
	type fields struct {
		Opts *args.Options
	}
	tests := []struct {
		name          string
		fields        fields
		wantMediaMeta *meta.MediaMeta
		wantErr       bool
	}{
		{
			name:   "Test on PC",
			fields: fields{Opts: &args.Options{Url: "https://y.qq.com/n/ryqq/songDetail/000Afq8O0JRzXv"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "南山忆",
				Description: "《剑仙》网游主题曲",
				Duration:    201,
				Artist:      "许嵩",
				Album:       "许嵩早期单曲集",
				Audios:      []meta.Audio{{Url: ".m4a"}},
			},
		},
		{
			name:   "Test on mobile",
			fields: fields{Opts: &args.Options{Url: "https://i.y.qq.com/v8/playsong.html?ADTAG=ryqq.songDetail&songmid=000Afq8O0JRzXv&songid=0&songtype=0#webchat_redirect"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "南山忆",
				Description: "《剑仙》网游主题曲",
				Duration:    201,
				Artist:      "许嵩",
				Album:       "许嵩早期单曲集",
				Audios:      []meta.Audio{{Url: ".m4a"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Core{
				Opts: tt.fields.Opts,
			}
			gotMediaMeta, err := c.FetchMetaAndResourceInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchMetaAndResourceInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			test_helper.TestMediaMeta(t, gotMediaMeta, tt.wantMediaMeta)
		})
	}
}
