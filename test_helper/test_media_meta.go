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
	checkList := []struct {
		name      string
		got       interface{}
		want      interface{}
		condition string
	}{
		{"Title", gotMediaMeta.Title, wantMediaMeta.Title, "="},
		{"Description", gotMediaMeta.Description, wantMediaMeta.Description, "="},
		{"Artist", gotMediaMeta.Artist, wantMediaMeta.Artist, "="},
		{"Album", gotMediaMeta.Album, wantMediaMeta.Album, "="},
		{"Audio.Url", gotMediaMeta.Audio.Url, wantMediaMeta.Audio.Url, "contain"},
	}
	for _, checkItem := range checkList {
		if checkItem.want != nil {
			if checkItem.condition == "=" && checkItem.got != checkItem.want {
				t.Errorf("check %s failed. Should give: %v, but got: %v", checkItem.name, checkItem.want, checkItem.got)
			}
			if checkItem.condition == "contain" && !strings.Contains(fmt.Sprintf("%v", checkItem.got), fmt.Sprintf("%v", checkItem.want)) {
				t.Errorf("check %s failed. Should contain: %v, but got: %v", checkItem.name, checkItem.want, checkItem.got)
			}
		}
	}
}
