package processor

import (
	"os"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/logger"
)

type Processor struct {
	Opts *args.Options
}

func (p *Processor) Process() (err error) {
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
