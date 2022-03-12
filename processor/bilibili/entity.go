package bilibili

type AudioResource struct {
	Data struct {
		Dash struct {
			Duration int     `json:"duration"`
			Videos   []Video `json:"video"`
			Audios   []struct {
				BaseUrl string `json:"baseUrl"`
			} `json:"audio"`
		} `json:"dash"`
	} `json:"data"`
}

type Video struct {
	Id      int    `json:"id"`
	BaseUrl string `json:"baseUrl"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

type AudioMeta struct {
	VideoData struct {
		Title       string `json:"title"`
		Duration    int    `json:"duration"`
		Description string `json:"desc"`
		Pic         string `json:"pic"`
		Owner       struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Staff []struct {
			Mid   int    `json:"mid"`
			Title string `json:"title"`
			Name  string `json:"name"`
		} `json:"staff"`
	} `json:"videoData"`
}
