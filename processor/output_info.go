package processor

import (
	"encoding/json"
	"fmt"

	"github.com/foamzou/audio-get/meta"
)

const MetaFormatJson = "json"
const MetaFormatPlain = "plain"

type OutputMeta struct {
	MediaMeta *meta.MediaMeta `json:"mediaMeta"`
	Audios    []*meta.Audio   `json:"audios"`
}

func (p *Processor) outputMeta(mediaMeta *meta.MediaMeta, audios []*meta.Audio) {
	// json output
	if p.Opts.MetaFormat == MetaFormatJson {
		o := &OutputMeta{
			MediaMeta: mediaMeta,
			Audios:    audios,
		}
		jsonByte, err := json.MarshalIndent(o, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonByte))
		return
	}

	// plain output
	fmt.Println(fmt.Sprintf("[Title] %s\n[Description] %s\n\n", mediaMeta.Title, mediaMeta.Description))
	audioCount := len(audios)
	fmt.Println(fmt.Sprintf("[Audio Count] %d", audioCount))
	for i, audio := range audios {
		fmt.Println(fmt.Sprintf("[%d/%d] Title: %s\n\tArtist: %s\n\tAlbum: %s\n\tResource: %s",
			i+1, audioCount, audio.Title, audio.Artist, audio.Album, audio.Resource.Url))
	}
}
