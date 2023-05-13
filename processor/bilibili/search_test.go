package bilibili

import (
	"testing"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/test_helper"
	"github.com/foamzou/audio-get/utils"
)

func TestCore_SearchSong(t *testing.T) {
	utils.InitConfig()
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
			fields: fields{Opts: &args.Options{Search: args.Search{Keyword: "珊瑚海 周杰伦", Type: "song"}}},
			wantSongItem: &meta.SearchSongItem{
				Name:   "【1080P修复】周杰伦 梁心颐 - 珊瑚海MV 修复版",
				Artist: "zyl2012",
				Url:    "bilibili",
				Source: consts.SourceNameBilibili,
			},
		},
		{
			name:   "Test search video",
			fields: fields{Opts: &args.Options{Search: args.Search{Keyword: "热带雨林", Type: "song"}}},
			wantSongItem: &meta.SearchSongItem{
				Name:   "【S.H.E】热带雨林-720P修复版 周杰伦作曲",
				Artist: "zyl2012",
				Url:    "bilibili",
				Source: consts.SourceNameBilibili,
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
