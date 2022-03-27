package youtube

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
			fields: fields{Opts: &args.Options{Search: args.Search{Keyword: "123木头人 黑girl", Type: "song"}}},
			wantSongItem: &meta.SearchSongItem{
				Name:   "黑澀會美眉 - 123木頭人 (官方版MV)",
				Artist: "福茂唱片",
				Url:    "youtube",
				Source: consts.SourceNameYoutube,
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
