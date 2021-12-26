package ffmpeg

type MediaFormat struct {
	Format struct {
		Filename string `json:"filename"`
		Tags     struct {
			Album  string `json:"album"`
			Title  string `json:"title"`
			Artist string `json:"artist"`
		} `json:"tags"`
	} `json:"format"`
}

type MetaTag struct {
	Album  string
	Title  string
	Artist string
}
