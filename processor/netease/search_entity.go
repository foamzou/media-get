package netease

type SearchSongResponse struct {
	NeedLogin bool `json:"needLogin"`
	Result    struct {
		Songs []struct {
			Name string `json:"name"`
			Id   int    `json:"id"`
			Ar   []struct {
				Id    int      `json:"id"`
				Name  string   `json:"name"`
				Alias []string `json:"alias"`
				Alia  []string `json:"alia,omitempty"`
			} `json:"ar"`
			Alia []interface{} `json:"alia"`
			Pop  float64       `json:"pop"`
			St   int           `json:"st"`
			Rt   *string       `json:"rt"`
			Fee  int           `json:"fee"`
			V    int           `json:"v"`
			Crbt interface{}   `json:"crbt"`
			Cf   string        `json:"cf"`
			Al   struct {
				Id     int           `json:"id"`
				Name   string        `json:"name"`
				PicUrl string        `json:"picUrl"`
				Tns    []interface{} `json:"tns"`
				PicStr string        `json:"pic_str,omitempty"`
				Pic    int64         `json:"pic"`
			} `json:"al"`
			Dt              int           `json:"dt"`
			A               interface{}   `json:"a"`
			Cd              string        `json:"cd"`
			No              int           `json:"no"`
			RtUrl           interface{}   `json:"rtUrl"`
			Ftype           int           `json:"ftype"`
			RtUrls          []interface{} `json:"rtUrls"`
			DjId            int           `json:"djId"`
			Copyright       int           `json:"copyright"`
			SId             int           `json:"s_id"`
			Mark            int           `json:"mark"`
			OriginCoverType int           `json:"originCoverType"`
			ResourceState   bool          `json:"resourceState"`
			Version         int           `json:"version"`
			Single          int           `json:"single"`
			NoCopyrightRcmd interface{}   `json:"noCopyrightRcmd"`
			Rtype           int           `json:"rtype"`
			Rurl            interface{}   `json:"rurl"`
			Mst             int           `json:"mst"`
			Cp              int           `json:"cp"`
			Mv              int           `json:"mv"`
			PublishTime     int64         `json:"publishTime"`
			Privilege       struct {
				Id                 int         `json:"id"`
				Fee                int         `json:"fee"`
				Paid               int         `json:"paid"`
				St                 int         `json:"st"`
				Pl                 int         `json:"pl"`
				Dl                 int         `json:"dl"`
				Sp                 int         `json:"sp"`
				Cp                 int         `json:"cp"`
				Subp               int         `json:"subp"`
				Cs                 bool        `json:"cs"`
				Maxbr              int         `json:"maxbr"`
				Fl                 int         `json:"fl"`
				Toast              bool        `json:"toast"`
				Flag               int         `json:"flag"`
				PreSell            bool        `json:"preSell"`
				PlayMaxbr          int         `json:"playMaxbr"`
				DownloadMaxbr      int         `json:"downloadMaxbr"`
				Rscl               interface{} `json:"rscl"`
				FreeTrialPrivilege struct {
					ResConsumable  bool `json:"resConsumable"`
					UserConsumable bool `json:"userConsumable"`
				} `json:"freeTrialPrivilege"`
				ChargeInfoList []struct {
					Rate          int         `json:"rate"`
					ChargeUrl     interface{} `json:"chargeUrl"`
					ChargeMessage interface{} `json:"chargeMessage"`
					ChargeType    int         `json:"chargeType"`
				} `json:"chargeInfoList"`
			} `json:"privilege"`
			Lyrics []string `json:"lyrics,omitempty"`
		} `json:"songs"`
	} `json:"result"`
	Code int `json:"code"`
}
