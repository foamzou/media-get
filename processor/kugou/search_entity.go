package kugou

type SearchSongResponse struct {
	Data struct {
		Info []struct {
			Hash              string        `json:"hash"`
			Songname          string        `json:"songname"`
			AlbumName         string        `json:"album_name"`
			SongnameOriginal  string        `json:"songname_original"`
			Accompany         int           `json:"Accompany"`
			Sqhash            string        `json:"sqhash"`
			FailProcess       int           `json:"fail_process"`
			PayType           int           `json:"pay_type"`
			RpType            string        `json:"rp_type"`
			AlbumId           string        `json:"album_id"`
			OthernameOriginal string        `json:"othername_original"`
			Mvhash            string        `json:"mvhash"`
			Extname           string        `json:"extname"`
			Group             []interface{} `json:"group"`
			Price320          int           `json:"price_320"`
			Hash1             string        `json:"320hash"`
			Topic             string        `json:"topic"`
			Othername         string        `json:"othername"`
			Isnew             int           `json:"isnew"`
			FoldType          int           `json:"fold_type"`
			OldCpy            int           `json:"old_cpy"`
			Srctype           int           `json:"srctype"`
			Singername        string        `json:"singername"`
			AlbumAudioId      int           `json:"album_audio_id"`
			Duration          int           `json:"duration"`
			Filesize          int           `json:"320filesize"`
			PkgPrice320       int           `json:"pkg_price_320"`
			AudioId           int           `json:"audio_id"`
			Feetype           int           `json:"feetype"`
			Price             int           `json:"price"`
			Filename          string        `json:"filename"`
			Source            string        `json:"source"`
			PriceSq           int           `json:"price_sq"`
			FailProcess320    int           `json:"fail_process_320"`
			PkgPrice          int           `json:"pkg_price"`
			PayType320        int           `json:"pay_type_320"`
			TopicUrl          string        `json:"topic_url"`
			M4Afilesize       int           `json:"m4afilesize"`
			RpPublish         int           `json:"rp_publish"`
			Privilege         int           `json:"privilege"`
			Filesize1         int           `json:"filesize"`
			Isoriginal        int           `json:"isoriginal"`
			Privilege1        int           `json:"320privilege"`
			Sqprivilege       int           `json:"sqprivilege"`
			FailProcessSq     int           `json:"fail_process_sq"`
		} `json:"info"`
	} `json:"data"`
	Error string `json:"error"`
}
