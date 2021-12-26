package processor

import (
	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/processor/bilibili"
)

type Processor struct {
	Opts *args.Options
}

func (p *Processor) Process() (err error) {
	b := bilibili.Core{
		Opts: p.Opts,
	}
	mediaMeta, audios, err := b.FetchMetaAndResourceInfo()
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
