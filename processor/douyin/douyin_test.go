package douyin

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
			fields: fields{Opts: &args.Options{Url: "https://v.douyin.com/8BxvWBm/"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "【致我们单纯的小美好】主题曲：我多喜欢你，你会知道",
				Description: "【致我们单纯的小美好】主题曲：我多喜欢你，你会知道",
				Duration:    221,
				CoverUrl:    "jpeg",
				Artist:      "环球音乐宅急送",
				Album:       "抖音Video",
				Audios: []meta.Audio{{
					Url: ".mp3",
				}},
				Videos: []meta.Video{{
					Url: ".mp4",
				}},
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
