package kuwo

type SongMeta struct {
	Data struct {
		Songinfo struct {
			Album           string      `json:"album"`
			AlbumId         string      `json:"albumId"`
			Artist          string      `json:"artist"`
			ArtistId        string      `json:"artistId"`
			ContentType     string      `json:"contentType"`
			CoopFormats     []string    `json:"coopFormats"`
			CopyRight       string      `json:"copyRight"`
			Duration        string      `json:"duration"`
			Formats         string      `json:"formats"`
			HasEcho         interface{} `json:"hasEcho"`
			HasMv           string      `json:"hasMv"`
			Id              string      `json:"id"`
			IsExt           interface{} `json:"isExt"`
			IsNew           interface{} `json:"isNew"`
			IsPoint         string      `json:"isPoint"`
			Isbatch         interface{} `json:"isbatch"`
			Isdownload      string      `json:"isdownload"`
			Isstar          string      `json:"isstar"`
			MkvNsig1        string      `json:"mkvNsig1"`
			MkvNsig2        string      `json:"mkvNsig2"`
			MkvRid          string      `json:"mkvRid"`
			Mp3Nsig1        string      `json:"mp3Nsig1"`
			Mp3Nsig2        string      `json:"mp3Nsig2"`
			Mp3Rid          string      `json:"mp3Rid"`
			Mp3Size         string      `json:"mp3Size"`
			Mp4Sig1         string      `json:"mp4sig1"`
			Mp4Sig2         string      `json:"mp4sig2"`
			MusicrId        string      `json:"musicrId"`
			MutiVer         string      `json:"mutiVer"`
			Mvpayinfo       interface{} `json:"mvpayinfo"`
			Mvpic           interface{} `json:"mvpic"`
			Nsig1           string      `json:"nsig1"`
			Nsig2           string      `json:"nsig2"`
			Online          string      `json:"online"`
			Params          interface{} `json:"params"`
			Pay             string      `json:"pay"`
			Pic             string      `json:"pic"`
			PlayCnt         string      `json:"playCnt"`
			RankChange      interface{} `json:"rankChange"`
			Reason          interface{} `json:"reason"`
			Score           interface{} `json:"score"`
			Score100        string      `json:"score100"`
			SongName        string      `json:"songName"`
			SongTimeMinutes string      `json:"songTimeMinutes"`
			Tpay            interface{} `json:"tpay"`
			Trend           interface{} `json:"trend"`
			UpTime          string      `json:"upTime"`
			Uploader        string      `json:"uploader"`
		} `json:"songinfo"`
	} `json:"data"`
	Msg       string      `json:"msg"`
	Msgs      interface{} `json:"msgs"`
	Profileid string      `json:"profileid"`
	Reqid     string      `json:"reqid"`
	Status    int         `json:"status"`
}
