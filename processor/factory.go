package processor

import (
	"strings"

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
	var ProcessorMap = map[string]meta.IProcessor{
		"163.com":      &netease.Core{Opts: p.Opts},
		"bilibili.com": &bilibili.Core{Opts: p.Opts},
		"douyin.com":   &douyin.Core{Opts: p.Opts},
		"youtube.com":  &youtube.Core{Opts: p.Opts},
		"migu.cn":      &migu.Core{Opts: p.Opts},
		"kugou.com":    &kugou.Core{Opts: p.Opts},
		"kuwo.cn":      &kuwo.Core{Opts: p.Opts},
		"qq.com":       &qqmusic.Core{Opts: p.Opts},
	}

	for host, processor := range ProcessorMap {
		if strings.Contains(p.Opts.Url, host) {
			return processor
		}
	}

	return nil
}
