package migu

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
			name:   "Test search song",
			fields: fields{Opts: &args.Options{Search: args.Search{Keyword: "淡水海边 周杰伦", Type: "song"}}},
			wantSongItem: &meta.SearchSongItem{
				Name:   "淡水海边(电影《不能说的秘密》背景音乐)",
				Artist: "周杰伦",
				Album:  "《不能说的秘密》电影原声",
				Url:    "music.migu.cn",
				Source: consts.SourceNameMigu,
			},
		},
		{
			name:   "Test search song",
			fields: fields{Opts: &args.Options{Search: args.Search{Keyword: "I Remember F.I.R", Type: "song"}}},
			wantSongItem: &meta.SearchSongItem{
				Name:   "淡水海边(电影《不能说的秘密》背景音乐)",
				Artist: "周杰伦",
				Album:  "《不能说的秘密》电影原声",
				Url:    "music.migu.cn",
				Source: consts.SourceNameMigu,
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
