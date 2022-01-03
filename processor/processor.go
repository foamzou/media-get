package processor

import (
	"github.com/foamzou/audio-get/args"
)

type Processor struct {
	Opts *args.Options
}

func (p *Processor) Process() (err error) {
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
