package consts

const (
	SourceNameBilibili = "bilibili"
	SourceNameDouyin   = "douyin"
	SourceNameKugou    = "kugou"
	SourceNameKuwo     = "kuwo"
	SourceNameMigu     = "migu"
	SourceNameNetease  = "netease"
	SourceNameQq       = "qq"
	SourceNameYoutube  = "youtube"
	SourceNameQMKG     = "qmkg"
)

func GetAllSourceName() []string {
	return []string{
		SourceNameBilibili,
		SourceNameDouyin,
		SourceNameKugou,
		SourceNameKuwo,
		SourceNameMigu,
		SourceNameNetease,
		SourceNameQq,
		SourceNameYoutube,
		SourceNameQMKG,
	}
}
