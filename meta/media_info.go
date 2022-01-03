package meta

type MediaMeta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	Duration    int    `json:"duration"` // second
	CoverUrl    string `json:"cover_url"`
	PublicTime  int64  `json:"public_time"`

	Headers map[string]string `json:"-"`

	ResourceType string `json:"resource_type"`

	Audios []Audio `json:"audios"`
	Videos []Video `json:"videos"`
}

type Audio struct {
	Url     string `json:"url"`
	BitRate int    `json:"bit_rate"` // kbps
}

type Video struct {
	Url            string `json:"url"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Ratio          string `json:"ratio"`
	NeedExtraAudio bool   `json:"need_extra_audio"`
}
