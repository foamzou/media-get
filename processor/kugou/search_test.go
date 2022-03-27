package kugou

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
			fields: fields{Opts: &args.Options{Search: args.Search{Keyword: "雨下的瞬間 李梦尹", Type: "song"}}},
			wantSongItem: &meta.SearchSongItem{
				Name:   "雨下的瞬間",
				Artist: "李梦尹",
				Album:  "雨下的瞬間",
				Url:    "kugou.com",
				Source: consts.SourceNameKugou,
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
