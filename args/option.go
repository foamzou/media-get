package args

type Options struct {
	Url          string `short:"u" long:"url" description:"Url from website. Start with http[s]://" required:"false"`
	Out          string `short:"o" long:"out" description:"The output file will be place to here(dir/filename). Default: current dir" required:"false"`
	ResourceType string `short:"t" long:"type" description:"The resource type you want to get [auto/audio/video/all]" required:"false" default:"auto"`
	MetaOnly     bool   `short:"m" long:"metaOnly" description:"Output meta info only. Would not download audio" required:"false"`
	InfoFormat   string `long:"infoFormat" description:"Support plain/json. Default: plain" required:"false"`
	AddMediaTag  bool   `long:"addMediaTag" description:"Add media tag into file from meta information" required:"false"`
	LogLevel     string `short:"l" long:"logLevel" description:"Support: silence/error/warn/info/debug. Default: info" required:"false"`

	Search Search `group:"Search Options"`

	IsDir bool
}

//nolint
type Search struct {
	Keyword        string `short:"k" long:"keyword" description:"If search song, keyword or searchSongName is required. Would use keyword to search if both of them are specify. Such as \"songName singer\"" required:"false"`
	SongName       string `long:"searchSongName" description:"song name" required:"false"`
	Artist         string `long:"searchArtist" description:"artist name" required:"false"`
	Album          string `long:"searchAlbum" description:"album name" required:"false"`
	Type           string `long:"searchType" description:"Support: song, Default: song" required:"false"`
	Sources        string `long:"sources" description:"Search in the specific source, separate with comma. Support: bilibili,douyin,kugou,kuwo,migu,netease,qq,youtube. Default: All" required:"false"`
	ExcludeSources string `long:"excludeSource" description:"Search not in the specific source, separate with comma" required:"false"`

	SourcesWillBeSearch []string
}
