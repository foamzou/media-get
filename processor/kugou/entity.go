package kugou

type SongMeta struct {
	Status  int `json:"status"`
	ErrCode int `json:"err_code"`
	Data    struct {
		Hash         string `json:"hash"`
		Timelength   int    `json:"timelength"`
		Filesize     int    `json:"filesize"`
		AudioName    string `json:"audio_name"`
		AlbumName    string `json:"album_name"`
		AlbumId      string `json:"album_id"`
		Img          string `json:"img"`
		AuthorName   string `json:"author_name"`
		SongName     string `json:"song_name"`
		Lyrics       string `json:"lyrics"`
		Privilege    int    `json:"privilege"`
		Privilege2   string `json:"privilege2"`
		PlayUrl      string `json:"play_url"`
		IsFreePart   int    `json:"is_free_part"`
		Bitrate      int    `json:"bitrate"`
		HasPrivilege bool   `json:"has_privilege"`
	} `json:"data"`
}
