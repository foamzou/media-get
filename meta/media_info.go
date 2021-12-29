package meta

type MediaMeta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	Duration    int    `json:"duration"`
	CoverUrl    string `json:"cover_url"`
	PublicTime  int64  `json:"public_time"`

	Headers map[string]string `json:"-"`

	Audio Audio `json:"audio"`
	Video Video `json:"video"`
}

type Audio struct {
	Url string `json:"url"`
}

type Video struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Ratio  string `json:"ratio"`
}
