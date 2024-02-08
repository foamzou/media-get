package meta

type MediaMeta struct {
	Title             string `json:"title"`
	Description       string `json:"description"`
	Artist            string `json:"artist"`
	Album             string `json:"album"`
	Duration          int    `json:"duration"` // second
	CoverUrl          string `json:"cover_url"`
	PublicTime        int64  `json:"public_time"`
	IsTrial           bool   `json:"is_trial"`
	ResourceForbidden bool   `json:"resource_forbidden"`
	FromMusicPlatform bool   `json:"from_music_platform"`
	Source            string `json:"source"`

	Headers map[string]string `json:"-"`

	ResourceType string `json:"resource_type"`

	Audios []Audio `json:"audios"`
	Videos []Video `json:"videos"`
}

type Audio struct {
	Url          string `json:"url"`
	BitRate      int    `json:"bit_rate"` // kbps
	NotAvailable bool   `json:"not_available"`
}

type Video struct {
	Url            string `json:"url"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	Ratio          string `json:"ratio"`
	NeedExtraAudio bool   `json:"need_extra_audio"`
	NotAvailable   bool   `json:"not_available"`
}
