package processor

import (
	"strings"

	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/processor/bilibili"
	"github.com/foamzou/audio-get/processor/netease"
)

func (p *Processor) getProcessor() meta.IProcessor {
	var ProcessorMap = map[string]meta.IProcessor{
		"163.com":      &netease.Core{Opts: p.Opts},
		"bilibili.com": &bilibili.Core{Opts: p.Opts},
	}

	for host, processor := range ProcessorMap {
		if strings.Contains(p.Opts.Url, host) {
			return processor
		}
	}

	return nil
}
