package douyin

type MetaResponse struct {
	StatusCode int      `json:"status_code"`
	Detail     MetaItem `json:"aweme_detail"`
}

type MetaItem struct {
	Author struct {
		Nickname string `json:"nickname"`
	} `json:"author"`
	Duration int    `json:"duration"`
	Desc     string `json:"desc"`
	Music    struct {
		PlayUrl struct {
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
		} `json:"play_url"`
	} `json:"music"`
	CreateTime int `json:"create_time"`
	Video      struct {
		Width       int `json:"width"`
		Height      int `json:"height"`
		OriginCover struct {
			UrlList []string `json:"url_list"`
			Uri     string   `json:"uri"`
		} `json:"origin_cover"`
		Duration int    `json:"duration"`
		Vid      string `json:"vid"`
		PlayAddr struct {
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
		} `json:"play_addr"`
		Ratio        string `json:"ratio"`
		HasWatermark bool   `json:"has_watermark"`
	} `json:"video"`
}
