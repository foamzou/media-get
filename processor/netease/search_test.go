package netease

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
			name:   "Test search song",
			fields: fields{Opts: &args.Options{Search: args.Search{Keyword: "我们 很美味", Type: "song"}}},
			wantSongItem: &meta.SearchSongItem{
				Name:   "我们",
				Artist: "很美味",
				Album:  "我们",
				Url:    "music.163.com",
				Source: consts.SourceNameNetease,
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
