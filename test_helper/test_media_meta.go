package test_helper

import (
	"fmt"
	"strings"
	"testing"

	"github.com/foamzou/audio-get/meta"
)

func TestMediaMeta(t *testing.T, gotMediaMeta, wantMediaMeta *meta.MediaMeta) {
	if gotMediaMeta == nil {
		t.Errorf("gotMediaMeta is nil")
		return
	}
	checkList := []CheckStruct{
		{"Title", gotMediaMeta.Title, wantMediaMeta.Title, "="},
		{"Description", gotMediaMeta.Description, wantMediaMeta.Description, "="},
		{"Duration", gotMediaMeta.Duration, wantMediaMeta.Duration, "="},
		{"CoverUrl", gotMediaMeta.CoverUrl, wantMediaMeta.CoverUrl, "contain"},
		{"Artist", gotMediaMeta.Artist, wantMediaMeta.Artist, "="},
		{"Album", gotMediaMeta.Album, wantMediaMeta.Album, "="},
		{"Audios.Url", gotMediaMeta.Audios[0].Url, wantMediaMeta.Audios[0].Url, "contain"},
		{"Audios.NotAvailable", gotMediaMeta.Audios[0].NotAvailable, wantMediaMeta.Audios[0].NotAvailable, "="},
	}
	if wantMediaMeta.Videos != nil && len(wantMediaMeta.Videos) >= 1 {
		checkList = append(checkList, CheckStruct{"Videos.Url", gotMediaMeta.Videos[0].Url, wantMediaMeta.Videos[0].Url, "contain"})
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
