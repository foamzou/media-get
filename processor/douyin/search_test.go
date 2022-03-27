package douyin

import (
	"testing"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/test_helper"
)

func TestCore_SearchSong(t *testing.T) {
	type fields struct {
		Opts *args.Options
	}
	tests := []struct {
		name         string
		fields       fields
		wantSongItem *meta.SearchSongItem
		wantErr      bool
	}{
		{
			name:   "Test search video",
			fields: fields{Opts: &args.Options{Search: args.Search{Keyword: "七朵花《我只想要》是多少人的回忆", Type: "song"}}},
			wantSongItem: &meta.SearchSongItem{
				Name:   "七朵花《我只想要》是多少人的回忆#音乐推荐 #经典音乐#情感音乐#热门歌曲#抖音热歌#音乐mv#音乐 ",
				Artist: "音乐而动(激光炮)",
				Url:    "douyin",
				Source: consts.SourceNameDouyin,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Core{
				Opts: tt.fields.Opts,
			}
			gotSongItems, err := c.SearchSong()
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchMetaAndResourceInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			test_helper.TestSearchSongResult(t, gotSongItems, tt.wantSongItem)
		})
	}
}
