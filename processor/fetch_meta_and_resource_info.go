package processor

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/logger"
	"github.com/foamzou/audio-get/meta"
)

func (p *Processor) FetchMetaAndResourceInfo() (err error) {
	if p.Opts.Url == "" {
		logger.Error("the required flag `-u, --url' was not specified")
		os.Exit(1)
	}
	processor := p.getProcessor()
	if processor == nil {
		panic("do not support the host")
	}
	mediaMeta, err := processor.FetchMetaAndResourceInfo()
	if err != nil {
		return
	}

	if !p.Opts.MetaOnly {
		err := p.download(mediaMeta)
		if err != nil {
			return err
		}
	}

	p.outputMeta(mediaMeta)
	return
}

func (p *Processor) outputMeta(mediaMeta *meta.MediaMeta) {
	// json output
	if p.Opts.InfoFormat != consts.InfoFormatPlain {
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
