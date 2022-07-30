package migu

import (
	"strings"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/meta"
)

type Core struct {
	Opts *args.Options
}

func (c *Core) IsMusicPlatform() bool {
	return true
}

func (c *Core) GetSourceName() string {
	return consts.SourceNameMigu
}

func (c *Core) Domains() []string {
	return []string{"migu.cn"}
}

func (c *Core) FetchMetaAndResourceInfo() (mediaMeta *meta.MediaMeta, err error) {
	if strings.Contains(c.Opts.Url, "/song/") {
		return c.fetchFromSong()
	}
	return nil, err
}
