package youtube

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
			fields: fields{Opts: &args.Options{Url: "https://www.youtube.com/watch?v=_BR8-Qp4j5M"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:    "IU - 夜信 (Through the Night)   (華納official HD 高畫質官方中文版)",
				Duration: 283,
				CoverUrl: "http",
				Artist:   "台灣華納日韓官方頻道",
				Album:    "Youtube",
				Audios: []meta.Audio{{
					Url: "googlevideo",
				}},
				Videos: []meta.Video{{
					Url: "googlevideo",
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
