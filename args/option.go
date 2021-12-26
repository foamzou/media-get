package args

type Options struct {
	Url        string `short:"u" long:"url" description:"Url from website. Start with http[s]://" required:"true"`
	Out        string `short:"o" long:"out" description:"The output file will be place to here(dir/filename)" required:"false"`
	MetaOnly   bool   `short:"m" long:"metaonly" description:"Output meta info only. Would not download audio" required:"false"`
	MetaFormat string `long:"metaformat" description:"Support raw(default), json" required:"false"`
	IsDir      bool
}
