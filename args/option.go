package args

type Options struct {
	Url          string `short:"u" long:"url" description:"Url from website. Start with http[s]://" required:"true"`
	Out          string `short:"o" long:"out" description:"The output file will be place to here(dir/filename). Default: current dir" required:"false"`
	ResourceType string `short:"t" long:"type" description:"The resource type you want to get [auto/audio/video/all]" required:"false" default:"auto"`
	AddMediaTag  bool   `long:"addMediaTag" description:"Add media tag into file from meta information. Default: false" required:"false"`
	MetaOnly     bool   `short:"m" long:"metaOnly" description:"Output meta info only. Would not download audio" required:"false"`
	MetaFormat   string `long:"metaFormat" description:"Support plain/json. Default: plain" required:"false"`
	IsDir        bool
}
