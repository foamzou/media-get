package meta

type IProcessor interface {
	FetchMetaAndResourceInfo() (mediaMeta *MediaMeta, audios []*Audio, err error)
}
