package kugou

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
			fields: fields{Opts: &args.Options{Url: "https://www.kugou.com/mixsong/456b2ae3.html"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "雨下的瞬間",
				Description: "李梦尹 - 雨下的瞬間",
				Duration:    272,
				Artist:      "李梦尹",
				Album:       "雨下的瞬間",
				Audios:      []meta.Audio{{Url: ".mp3"}},
			},
		},
		{
			name:   "Test song",
			fields: fields{Opts: &args.Options{Url: "https://www.kugou.com/song/#hash=E5D69492BACF314B28519884B6D339BD&album_id=36463673"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "雨下的瞬間",
				Description: "李梦尹 - 雨下的瞬間",
				Duration:    272,
				Artist:      "李梦尹",
				Album:       "雨下的瞬間",
				Audios:      []meta.Audio{{Url: ".mp3"}},
			},
		},
		{
			name:   "Test song",
			fields: fields{Opts: &args.Options{Url: "https://www.kugou.com/mixsong/218xbq2c.html#hash=85E3519452F6B74790775061D882A4C2&album_id=14340507"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "残缺的彩虹",
				Description: "陈绮贞 - 残缺的彩虹",
				Duration:    208,
				Artist:      "陈绮贞",
				Album:       "沙发海",
				Audios:      []meta.Audio{{Url: ".mp3"}},
			},
		},
		{
			name:   "Test song",
			fields: fields{Opts: &args.Options{Url: "https://www.kugou.com/song/#hash=85E3519452F6B74790775061D882A4C2&album_id=14340507"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "残缺的彩虹",
				Description: "陈绮贞 - 残缺的彩虹",
				Duration:    208,
				Artist:      "陈绮贞",
				Album:       "沙发海",
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
