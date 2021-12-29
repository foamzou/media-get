package processor

import (
	"encoding/json"
	"fmt"

	"github.com/foamzou/audio-get/meta"
)

const MetaFormatJson = "json"
const MetaFormatPlain = "plain"

func (p *Processor) outputMeta(mediaMeta *meta.MediaMeta) {
	// json output
	if p.Opts.MetaFormat == MetaFormatJson {
		jsonByte, err := json.MarshalIndent(mediaMeta, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonByte))
		return
	}

	// plain output
	fmt.Println(fmt.Sprintf("[Title] %s\n[Description] %s\n\n", mediaMeta.Title, mediaMeta.Description))
	fmt.Println(fmt.Sprintf("Artist: %s\n\tAlbum: %s\n\tResource: %s",
		mediaMeta.Artist, mediaMeta.Album, mediaMeta.Audio.Url))
}
