package migu

import (
	"strings"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/meta"
)

type Core struct {
	Opts *args.Options
}

func (c *Core) FetchMetaAndResourceInfo() (mediaMeta *meta.MediaMeta, err error) {
	if strings.Contains(c.Opts.Url, "/song/") {
		return c.fetchFromSong()
	}
	return nil, err
}
