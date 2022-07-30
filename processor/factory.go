package processor

import (
	"strings"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
	"github.com/foamzou/audio-get/processor/bilibili"
	"github.com/foamzou/audio-get/processor/douyin"
	"github.com/foamzou/audio-get/processor/kugou"
	"github.com/foamzou/audio-get/processor/kuwo"
	"github.com/foamzou/audio-get/processor/migu"
	"github.com/foamzou/audio-get/processor/netease"
	"github.com/foamzou/audio-get/processor/qqmusic"
	"github.com/foamzou/audio-get/processor/youtube"
)

func (p *Processor) getProcessor() meta.IProcessor {
	var ProcessorMap = p.getProcessorMap()

	for _, processor := range ProcessorMap {
		domains := processor.Domains()
		for _, domain := range domains {
			if strings.Contains(p.Opts.Url, domain) {
				return processor
			}
		}
	}

	return nil
}

func (p *Processor) getProcessorBySourceName(sourceName string) meta.IProcessor {
	var ProcessorMap = p.getProcessorMap()

	provider, found := ProcessorMap[sourceName]
	if found {
		return provider
	}
	return nil
}

func (p *Processor) getProcessors() []meta.IProcessor {
	var processors = make([]meta.IProcessor, 0)

	for _, processor := range p.getProcessorMap() {
		processors = append(processors, processor)
	}

	return processors
}

func (p *Processor) getProcessorMap() map[string]meta.IProcessor {
	var ProcessorMap = map[string]meta.IProcessor{
		consts.SourceNameNetease:  &netease.Core{Opts: p.Opts},
		consts.SourceNameBilibili: &bilibili.Core{Opts: p.Opts},
		consts.SourceNameDouyin:   &douyin.Core{Opts: p.Opts},
		consts.SourceNameYoutube:  &youtube.Core{Opts: p.Opts},
		consts.SourceNameMigu:     &migu.Core{Opts: p.Opts},
		consts.SourceNameKugou:    &kugou.Core{Opts: p.Opts},
		consts.SourceNameKuwo:     &kuwo.Core{Opts: p.Opts},
		consts.SourceNameQq:       &qqmusic.Core{Opts: p.Opts},
	}

	return ProcessorMap
}
