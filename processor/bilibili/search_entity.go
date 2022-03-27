package bilibili

type SearchSongResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Seid           string `json:"seid"`
		Page           int    `json:"page"`
		Pagesize       int    `json:"pagesize"`
		NumResults     int    `json:"numResults"`
		NumPages       int    `json:"numPages"`
		SuggestKeyword string `json:"suggest_keyword"`
		RqtType        string `json:"rqt_type"`
		Result         []struct {
			ResultType string `json:"result_type"`
			Data       []struct {
				Type         string        `json:"type"`
				Id           int           `json:"id"`
				Author       string        `json:"author"`
				Mid          int           `json:"mid"`
				Typeid       string        `json:"typeid"`
				Typename     string        `json:"typename"`
				Arcurl       string        `json:"arcurl"`
				Aid          int           `json:"aid"`
				Bvid         string        `json:"bvid"`
				Title        string        `json:"title"`
				Description  string        `json:"description"`
				Arcrank      string        `json:"arcrank"`
				Pic          string        `json:"pic"`
				Play         int           `json:"play"`
				VideoReview  int           `json:"video_review"`
				Favorites    int           `json:"favorites"`
				Tag          string        `json:"tag"`
				Review       int           `json:"review"`
				Pubdate      int           `json:"pubdate"`
				Senddate     int           `json:"senddate"`
				Duration     string        `json:"duration"`
				Badgepay     bool          `json:"badgepay"`
				HitColumns   []interface{} `json:"hit_columns"`
				ViewType     string        `json:"view_type"`
				IsPay        int           `json:"is_pay"`
				IsUnionVideo int           `json:"is_union_video"`
				RecTags      interface{}   `json:"rec_tags"`
				NewRecTags   []interface{} `json:"new_rec_tags"`
				RankScore    int           `json:"rank_score"`
				Like         int           `json:"like"`
				Upic         string        `json:"upic"`
				Corner       string        `json:"corner"`
				Cover        string        `json:"cover"`
				Desc         string        `json:"desc"`
				Url          string        `json:"url"`
				RecReason    string        `json:"rec_reason"`
			} `json:"data"`
		} `json:"result"`
	} `json:"data"`
}
