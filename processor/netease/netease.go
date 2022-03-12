package netease

import (
	"strings"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/meta"
)

type Core struct {
	Opts *args.Options
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
