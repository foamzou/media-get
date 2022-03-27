package qqmusic

type SongMeta struct {
	Code int `json:"code"`
	Data []struct {
		Action struct {
			Alert    int `json:"alert"`
			Icons    int `json:"icons"`
			Msgdown  int `json:"msgdown"`
			Msgfav   int `json:"msgfav"`
			Msgid    int `json:"msgid"`
			Msgpay   int `json:"msgpay"`
			Msgshare int `json:"msgshare"`
			Switch   int `json:"switch"`
		} `json:"action"`
		Aid   int `json:"aid"`
		Album struct {
			Id         int    `json:"id"`
			Mid        string `json:"mid"`
			Name       string `json:"name"`
			Pmid       string `json:"pmid"`
			Subtitle   string `json:"subtitle"`
			TimePublic string `json:"time_public"`
			Title      string `json:"title"`
		} `json:"album"`
		Bpm      int    `json:"bpm"`
		DataType int    `json:"data_type"`
		Es       string `json:"es"`
		File     struct {
			B30S          int    `json:"b_30s"`
			E30S          int    `json:"e_30s"`
			HiresBitdepth int    `json:"hires_bitdepth"`
			HiresSample   int    `json:"hires_sample"`
			MediaMid      string `json:"media_mid"`
			Size128Mp3    int    `json:"size_128mp3"`
			Size192Aac    int    `json:"size_192aac"`
			Size192Ogg    int    `json:"size_192ogg"`
			Size24Aac     int    `json:"size_24aac"`
			Size320Mp3    int    `json:"size_320mp3"`
			Size48Aac     int    `json:"size_48aac"`
			Size96Aac     int    `json:"size_96aac"`
			Size96Ogg     int    `json:"size_96ogg"`
			SizeApe       int    `json:"size_ape"`
			SizeDts       int    `json:"size_dts"`
			SizeFlac      int    `json:"size_flac"`
			SizeHires     int    `json:"size_hires"`
			SizeTry       int    `json:"size_try"`
			TryBegin      int    `json:"try_begin"`
			TryEnd        int    `json:"try_end"`
			Url           string `json:"url"`
		} `json:"file"`
		Fnote      int `json:"fnote"`
		Genre      int `json:"genre"`
		Id         int `json:"id"`
		IndexAlbum int `json:"index_album"`
		IndexCd    int `json:"index_cd"`
		Interval   int `json:"interval"`
		Isonly     int `json:"isonly"`
		Ksong      struct {
			Id  int    `json:"id"`
			Mid string `json:"mid"`
		} `json:"ksong"`
		Label       string `json:"label"`
		Language    int    `json:"language"`
		Mid         string `json:"mid"`
		ModifyStamp int    `json:"modify_stamp"`
		Mv          struct {
			Id    int    `json:"id"`
			Name  string `json:"name"`
			Title string `json:"title"`
			Vid   string `json:"vid"`
			Vt    int    `json:"vt"`
		} `json:"mv"`
		Name string `json:"name"`
		Ov   int    `json:"ov"`
		Pay  struct {
			PayDown    int `json:"pay_down"`
			PayMonth   int `json:"pay_month"`
			PayPlay    int `json:"pay_play"`
			PayStatus  int `json:"pay_status"`
			PriceAlbum int `json:"price_album"`
			PriceTrack int `json:"price_track"`
			TimeFree   int `json:"time_free"`
		} `json:"pay"`
		Sa     int `json:"sa"`
		Singer []struct {
			Id    int    `json:"id"`
			Mid   string `json:"mid"`
			Name  string `json:"name"`
			Pmid  string `json:"pmid"`
			Title string `json:"title"`
			Type  int    `json:"type"`
			Uin   int    `json:"uin"`
		} `json:"singer"`
		Status     int    `json:"status"`
		Subtitle   string `json:"subtitle"`
		Tid        int    `json:"tid"`
		TimePublic string `json:"time_public"`
		Title      string `json:"title"`
		Trace      string `json:"trace"`
		Type       int    `json:"type"`
		Url        string `json:"url"`
		Version    int    `json:"version"`
		Volume     struct {
			Gain float64 `json:"gain"`
			Lra  float64 `json:"lra"`
			Peak float64 `json:"peak"`
		} `json:"volume"`
	} `json:"data"`
}
