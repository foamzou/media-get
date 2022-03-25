package processor

import (
	"errors"

	"github.com/foamzou/audio-get/args"
)

type Processor struct {
	Opts *args.Options
}

func (p *Processor) Process() (err error) {
	if p.Opts.Url != "" {
		return p.FetchMetaAndResourceInfo()
	}
	if p.Opts.Search.Keyword != "" {
		return p.Search()
	}
	return errors.New("url or keyword is required")
}
