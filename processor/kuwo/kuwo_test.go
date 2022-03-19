package kuwo

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
			name:   "Test song",
			fields: fields{Opts: &args.Options{Url: "https://www.kuwo.cn/play_detail/289457"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "再生花",
				Description: "最爱的主题曲",
				Duration:    242,
				Artist:      "陈慧琳",
				Album:       "最爱的主题曲",
				Audios:      []meta.Audio{{Url: ".mp3"}},
			},
		},
		{
			name:   "Test song",
			fields: fields{Opts: &args.Options{Url: "https://m.kuwo.cn/newh5app/play_detail/289457"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "再生花",
				Description: "最爱的主题曲",
				Duration:    242,
				Artist:      "陈慧琳",
				Album:       "最爱的主题曲",
				Audios:      []meta.Audio{{Url: ".mp3"}},
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
