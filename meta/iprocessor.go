package meta

type IProcessor interface {
	FetchMetaAndResourceInfo() (mediaMeta *MediaMeta, err error)
	SearchSong() (searchItems []*SearchSongItem, err error)
	Domains() []string
	GetSourceName() string
	IsMusicPlatform() bool
}
