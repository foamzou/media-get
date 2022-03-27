package meta

type SearchSongItem struct {
	Name           string
	Artist         string
	Album          string
	Duration       int
	Url            string
	CannotDownload bool
	Source         string

	Score float64
}
