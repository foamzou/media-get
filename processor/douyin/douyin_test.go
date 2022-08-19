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
			name:   "test share url",
			fields: fields{Opts: &args.Options{Url: "https://v.douyin.com/8BxvWBm/"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "【致我们单纯的小美好】主题曲：我多喜欢你，你会知道",
				Description: "【致我们单纯的小美好】主题曲：我多喜欢你，你会知道",
				Duration:    223,
				CoverUrl:    "jpeg",
				Artist:      "环球音乐宅急送",
				Album:       "抖音Video",
				Audios: []meta.Audio{{
					Url: ".mp3",
				}},
				Videos: []meta.Video{{
					Url: "video",
				}},
			},
		},
		{
			name:   "test pc url",
			fields: fields{Opts: &args.Options{Url: "https://www.douyin.com/video/7035903951505542437"}},
			wantMediaMeta: &meta.MediaMeta{
				Title:       "七朵花《我只想要》是多少人的回忆#音乐推荐 #经典音乐#情感音乐#热门歌曲#抖音热歌#音乐mv#音乐 ",
				Description: "七朵花《我只想要》是多少人的回忆#音乐推荐 #经典音乐#情感音乐#热门歌曲#抖音热歌#音乐mv#音乐 ",
				Duration:    260,
				CoverUrl:    "jpeg",
				Album:       "抖音Video",
				Audios: []meta.Audio{{
					Url: ".mp3",
				}},
				Videos: []meta.Video{{
					Url: "video",
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
