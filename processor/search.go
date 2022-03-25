package processor

import (
	"encoding/json"
	"fmt"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/logger"
	"github.com/foamzou/audio-get/meta"
)

func (p *Processor) Search() error {
	return p.SearchSong()
}

func (p *Processor) SearchSong() error {
	var searchItems []*meta.SearchSongItem
	processors := p.getProcessors()
	for _, processor := range processors {
		items, err := processor.SearchSong()
		if err != nil {
			logger.Error(err)
			continue
		}
		searchItems = append(searchItems, items...)
	}
	p.outputSearchResult(searchItems)
	return nil
}

func (p *Processor) outputSearchResult(searchItems []*meta.SearchSongItem) {
	// json output
	if p.Opts.InfoFormat != consts.InfoFormatPlain {
		jsonByte, err := json.MarshalIndent(searchItems, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonByte))
		return
	}

	fmt.Println("search result:")
	for _, item := range searchItems {
		fmt.Printf("%s, %s, %s, %s, %d, %s\n", item.Name, item.Album, item.Artist, item.Url, item.Duration, item.Source)
	}
}
