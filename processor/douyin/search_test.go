package douyin

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
			fields: fields{Opts: &args.Options{Search: args.Search{Keyword: "徐子骞被天瑜坚强乐观的性格吸引，金枝妈妈也有意撮合两人", Type: "song"}}},
			wantSongItem: &meta.SearchSongItem{
				Name:   "子骞在和天瑜相处的时间里，渐渐被她乐观开朗的性格吸引， #怀旧经典影视 #陈乔恩 #明道 #精彩片段",
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
				t.Errorf("SearchSong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			test_helper.TestSearchSongResult(t, gotSongItems, tt.wantSongItem)
		})
	}
}
