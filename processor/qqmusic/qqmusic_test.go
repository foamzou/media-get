package qqmusic

import (
	"testing"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/test_helper"
)

func TestCore_FetchMetaAndResourceInfo(t *testing.T) {
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
			fields: fields{Opts: &args.Options{Url: "https://y.qq.com/n/ryqq/songDetail/003OIguR4bcxAf"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "失物招领",
				Description: "《以年为单位的恋爱》电影情感主题曲",
				Duration:    246,
				Artist:      "蔡健雅",
				Album:       "失物招领",
				Audios:      []meta.Audio{{Url: ".m4a"}},
			},
		},
		{
			name:   "Test on mobile",
			fields: fields{Opts: &args.Options{Url: "https://i.y.qq.com/v8/playsong.html?ADTAG=ryqq.songDetail&songmid=003OIguR4bcxAf&songid=0&songtype=0#webchat_redirect"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "失物招领",
				Description: "《以年为单位的恋爱》电影情感主题曲",
				Duration:    246,
				Artist:      "蔡健雅",
				Album:       "失物招领",
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
