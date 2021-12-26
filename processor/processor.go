package processor

import (
	"github.com/foamzou/audio-get/args"
)

type Processor struct {
	Opts *args.Options
}

func (p *Processor) Process() (err error) {
	processor := p.getProcessor()
	mediaMeta, audios, err := processor.FetchMetaAndResourceInfo()
	if err != nil {
		return
	}

	if !p.Opts.MetaOnly {
		err := p.download(mediaMeta, audios)
		if err != nil {
			return err
		}
	}

	p.outputMeta(mediaMeta, audios)
	return
}
