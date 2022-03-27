package netease

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

func (c *Core) Domain() string {
	return "163.com"
}

func (c *Core) GetSourceName() string {
	return consts.SourceNameNetease
}

func (c *Core) FetchMetaAndResourceInfo() (mediaMeta *meta.MediaMeta, err error) {
	if strings.Contains(c.Opts.Url, "/program") || strings.Contains(c.Opts.Url, "/dj") {
		return c.fetchFromProgram()
	}
	if strings.Contains(c.Opts.Url, "/song") {
		return c.fetchFromSong()
	}
	if strings.Contains(c.Opts.Url, "/mlog") {
		return c.fetchFromMlog()
	}
	if strings.Contains(c.Opts.Url, "/mv") || strings.Contains(c.Opts.Url, "/video") {
		return c.fetchFromMV()
	}
	return nil, err
}
