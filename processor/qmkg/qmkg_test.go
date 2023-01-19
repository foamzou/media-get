package qmkg

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
			name:   "test share url",
			fields: fields{Opts: &args.Options{Url: "https://kg3.qq.com/node/a6MPpvAot8/play_v2?s=bcY2OWbCGDcTPbjz&shareuid=679f9582262e328e&topsource=a0_pn201001006_z11_u52873355_l1_t1660396943__&pageId=details_of_creations"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:    "我以为",
				Duration: 308,
				CoverUrl: "jpg",
				Artist:   "锋哥",
				Album:    "全民K歌",
				Audios: []meta.Audio{{
					Url: ".m4a",
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
