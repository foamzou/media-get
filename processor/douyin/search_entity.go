package douyin

type SearchSongResponse struct {
	StatusCode int `json:"status_code"`
	Data       []struct {
		Type      int `json:"type"`
		AwemeInfo struct {
			AwemeId    string `json:"aweme_id"`
			Desc       string `json:"desc"`
			CreateTime int    `json:"create_time"`
			Author     struct {
				Uid       string `json:"uid"`
				ShortId   string `json:"short_id"`
				Nickname  string `json:"nickname"`
				Signature string `json:"signature"`
			} `json:"author"`
			Video struct {
				Duration int `json:"duration"`
			} `json:"video"`
		} `json:"aweme_info"`
	} `json:"data"`
}
