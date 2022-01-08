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
	if p.Opts.MetaFormat != MetaFormatPlain {
		jsonByte, err := json.MarshalIndent(mediaMeta, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonByte))
		return
	}

	// todo plain output
	fmt.Printf("[Title] %s\n[Description] %s\n\n\n", mediaMeta.Title, mediaMeta.Description)
	fmt.Printf("Artist: %s\n\tAlbum: %s\n\tResource: %s\n",
		mediaMeta.Artist, mediaMeta.Album, mediaMeta.Audios[0].Url)
}
