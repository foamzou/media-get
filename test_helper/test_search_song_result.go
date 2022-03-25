package test_helper

import (
	"fmt"
	"strings"
	"testing"

	"github.com/foamzou/audio-get/meta"
)

func TestSearchSongResult(t *testing.T, gotSongItems []*meta.SearchSongItem, wantSongItems *meta.SearchSongItem) {
	if len(gotSongItems) == 0 {
		t.Errorf("gotSongItems is empty")
		return
	}
	firstGotSongItem := gotSongItems[0]
	checkList := []CheckStruct{
		{"Name", firstGotSongItem.Name, wantSongItems.Name, "="},
		{"Artist", firstGotSongItem.Artist, wantSongItems.Artist, "="},
		{"Album", firstGotSongItem.Album, wantSongItems.Album, "="},
		{"Duration", firstGotSongItem.Duration, wantSongItems.Duration, "="},
		{"Url", firstGotSongItem.Url, wantSongItems.Url, "contain"},
		{"Source", firstGotSongItem.Source, wantSongItems.Source, "="},
	}

	for _, checkItem := range checkList {
		if checkItem.want != nil && checkItem.want != 0 && checkItem.want != "" {
			if checkItem.condition == "=" && checkItem.got != checkItem.want {
				t.Errorf("check %s failed. Should give: %v, but got: %v", checkItem.name, checkItem.want, checkItem.got)
			}
			if checkItem.condition == "contain" && !strings.Contains(fmt.Sprintf("%v", checkItem.got), fmt.Sprintf("%v", checkItem.want)) {
				t.Errorf("check %s failed. Should contain: %v, but got: %v", checkItem.name, checkItem.want, checkItem.got)
			}
		}
	}
}
