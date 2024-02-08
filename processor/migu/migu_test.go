package migu

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
			name:   "Test song",
			fields: fields{Opts: &args.Options{Url: "https://music.migu.cn/v3/music/song/60054701952"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "枫",
				Description: "枫",
				Duration:    275,
				Artist:      "周杰伦",
				Album:       "十一月的萧邦",
				Audios:      []meta.Audio{{Url: ".mp3", NotAvailable: true}},
			},
		},
		{
			name:   "Test song",
			fields: fields{Opts: &args.Options{Url: "https://music.migu.cn/v3/music/song/6005751LTG9"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "I Remember",
				Description: "I Remember",
				Duration:    306,
				Artist:      "飞儿乐团",
				Album:       "Better Life",
				Audios:      []meta.Audio{{Url: ".mp3", NotAvailable: false}},
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
