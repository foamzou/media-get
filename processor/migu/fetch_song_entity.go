package migu

type RateFormat struct {
	ResourceType         string `json:"resourceType"`
	FormatType           string `json:"formatType"`
	Url                  string `json:"url,omitempty"`
	Format               string `json:"format"`
	Size                 string `json:"size"`
	FileType             string `json:"fileType,omitempty"`
	Price                string `json:"price"`
	IosUrl               string `json:"iosUrl,omitempty"`
	AndroidUrl           string `json:"androidUrl,omitempty"`
	AndroidFileType      string `json:"androidFileType,omitempty"`
	IosFileType          string `json:"iosFileType,omitempty"`
	IosSize              string `json:"iosSize,omitempty"`
	AndroidSize          string `json:"androidSize,omitempty"`
	IosFormat            string `json:"iosFormat,omitempty"`
	AndroidFormat        string `json:"androidFormat,omitempty"`
	IosAccuracyLevel     string `json:"iosAccuracyLevel,omitempty"`
	AndroidAccuracyLevel string `json:"androidAccuracyLevel,omitempty"`
}

type SongInfo struct {
	Code     string `json:"code"`
	Info     string `json:"info"`
	Resource []struct {
		ResourceType string `json:"resourceType"`
		RefId        string `json:"refId"`
		CopyrightId  string `json:"copyrightId"`
		ContentId    string `json:"contentId"`
		SongId       string `json:"songId"`
		SongName     string `json:"songName"`
		SingerId     string `json:"singerId"`
		Singer       string `json:"singer"`
		AlbumId      string `json:"albumId"`
		Album        string `json:"album"`
		AlbumImgs    []struct {
			ImgSizeType string `json:"imgSizeType"`
			Img         string `json:"img"`
		} `json:"albumImgs"`
		OpNumItem struct {
			PlayNum             int    `json:"playNum"`
			PlayNumDesc         string `json:"playNumDesc"`
			KeepNum             int    `json:"keepNum"`
			KeepNumDesc         string `json:"keepNumDesc"`
			CommentNum          int    `json:"commentNum"`
			CommentNumDesc      string `json:"commentNumDesc"`
			ShareNum            int    `json:"shareNum"`
			ShareNumDesc        string `json:"shareNumDesc"`
			OrderNumByWeek      int    `json:"orderNumByWeek"`
			OrderNumByWeekDesc  string `json:"orderNumByWeekDesc"`
			OrderNumByTotal     int    `json:"orderNumByTotal"`
			OrderNumByTotalDesc string `json:"orderNumByTotalDesc"`
			ThumbNum            int    `json:"thumbNum"`
			ThumbNumDesc        string `json:"thumbNumDesc"`
			FollowNum           int    `json:"followNum"`
			FollowNumDesc       string `json:"followNumDesc"`
			SubscribeNum        int    `json:"subscribeNum"`
			SubscribeNumDesc    string `json:"subscribeNumDesc"`
			LivePlayNum         int    `json:"livePlayNum"`
			LivePlayNumDesc     string `json:"livePlayNumDesc"`
			PopularNum          int    `json:"popularNum"`
			PopularNumDesc      string `json:"popularNumDesc"`
			BookingNum          int    `json:"bookingNum"`
			BookingNumDesc      string `json:"bookingNumDesc"`
		} `json:"opNumItem"`
		ToneControl  string `json:"toneControl"`
		RelatedSongs []struct {
			ResourceType     string `json:"resourceType"`
			ResourceTypeName string `json:"resourceTypeName"`
			CopyrightId      string `json:"copyrightId"`
			ProductId        string `json:"productId"`
		} `json:"relatedSongs"`
		RateFormats    []RateFormat `json:"rateFormats"`
		NewRateFormats []struct {
			ResourceType         string `json:"resourceType"`
			FormatType           string `json:"formatType"`
			Url                  string `json:"url,omitempty"`
			Format               string `json:"format,omitempty"`
			Size                 string `json:"size,omitempty"`
			FileType             string `json:"fileType,omitempty"`
			Price                string `json:"price"`
			IosUrl               string `json:"iosUrl,omitempty"`
			AndroidUrl           string `json:"androidUrl,omitempty"`
			AndroidFileType      string `json:"androidFileType,omitempty"`
			IosFileType          string `json:"iosFileType,omitempty"`
			IosSize              string `json:"iosSize,omitempty"`
			AndroidSize          string `json:"androidSize,omitempty"`
			IosFormat            string `json:"iosFormat,omitempty"`
			AndroidFormat        string `json:"androidFormat,omitempty"`
			IosAccuracyLevel     string `json:"iosAccuracyLevel,omitempty"`
			AndroidAccuracyLevel string `json:"androidAccuracyLevel,omitempty"`
			AndroidNewFormat     string `json:"androidNewFormat,omitempty"`
			IosBit               int    `json:"iosBit,omitempty"`
			AndroidBit           int    `json:"androidBit,omitempty"`
		} `json:"newRateFormats"`
		Z3DCode struct {
			ResourceType    string `json:"resourceType"`
			FormatType      string `json:"formatType"`
			Price           string `json:"price"`
			IosUrl          string `json:"iosUrl"`
			AndroidUrl      string `json:"androidUrl"`
			AndroidFileType string `json:"androidFileType"`
			IosFileType     string `json:"iosFileType"`
			IosSize         string `json:"iosSize"`
			AndroidSize     string `json:"androidSize"`
			IosFormat       string `json:"iosFormat"`
			AndroidFormat   string `json:"androidFormat"`
			AndroidFileKey  string `json:"androidFileKey"`
			IosFileKey      string `json:"iosFileKey"`
			H5Url           string `json:"h5Url"`
			H5Size          string `json:"h5Size"`
			H5Format        string `json:"h5Format"`
		} `json:"z3dCode"`
		LrcUrl             string `json:"lrcUrl"`
		DigitalColumnId    string `json:"digitalColumnId"`
		Copyright          string `json:"copyright"`
		ValidStatus        bool   `json:"validStatus"`
		SongDescs          string `json:"songDescs"`
		SongAliasName      string `json:"songAliasName"`
		IsInDAlbum         string `json:"isInDAlbum"`
		IsInSideDalbum     string `json:"isInSideDalbum"`
		IsInSalesPeriod    string `json:"isInSalesPeriod"`
		SongType           string `json:"songType"`
		MrcUrl             string `json:"mrcUrl"`
		InvalidateDate     string `json:"invalidateDate"`
		DalbumId           string `json:"dalbumId"`
		TrcUrl             string `json:"trcUrl"`
		VipType            string `json:"vipType"`
		ScopeOfcopyright   string `json:"scopeOfcopyright"`
		AuditionsType      string `json:"auditionsType"`
		FirstIcon          string `json:"firstIcon"`
		TranslateName      string `json:"translateName"`
		ChargeAuditions    string `json:"chargeAuditions"`
		OldChargeAuditions string `json:"oldChargeAuditions"`
		SongIcon           string `json:"songIcon"`
		CodeRate           struct {
			PQ struct {
				CodeRateChargeAuditions string `json:"codeRateChargeAuditions"`
				IsCodeRateDownload      string `json:"isCodeRateDownload"`
				CodeRateFileSize        string `json:"codeRateFileSize"`
			} `json:"PQ"`
			HQ struct {
				CodeRateChargeAuditions string `json:"codeRateChargeAuditions"`
				IsCodeRateDownload      string `json:"isCodeRateDownload"`
			} `json:"HQ"`
			SQ struct {
				CodeRateChargeAuditions string `json:"codeRateChargeAuditions"`
				IsCodeRateDownload      string `json:"isCodeRateDownload"`
				ContentIdSQ             string `json:"contentIdSQ"`
			} `json:"SQ"`
			ZQ struct {
				CodeRateChargeAuditions string `json:"codeRateChargeAuditions"`
				IsCodeRateDownload      string `json:"isCodeRateDownload"`
			} `json:"ZQ"`
			Z3D struct {
				CodeRateChargeAuditions string `json:"codeRateChargeAuditions"`
				IsCodeRateDownload      string `json:"isCodeRateDownload"`
			} `json:"Z3D"`
		} `json:"codeRate"`
		IsDownload    string `json:"isDownload"`
		HasMv         string `json:"hasMv"`
		MvCopyright   string `json:"mvCopyright"`
		TopQuality    string `json:"topQuality"`
		PreSale       string `json:"preSale"`
		IsShare       string `json:"isShare"`
		IsCollection  string `json:"isCollection"`
		Length        string `json:"length"`
		AuditionsFlag string `json:"auditionsFlag"`
		ListenFlag    string `json:"listenFlag"`
		SingerImg     struct {
			Field1 struct {
				SingerName   string `json:"singerName"`
				MiguImgItems []struct {
					ImgSizeType string `json:"imgSizeType"`
					Img         string `json:"img"`
					FileId      string `json:"fileId"`
				} `json:"miguImgItems"`
			} `json:"112"`
		} `json:"singerImg"`
		SongNamePinyin  string `json:"songNamePinyin"`
		AlbumNamePinyin string `json:"albumNamePinyin"`
		Artists         []struct {
			Id           string `json:"id"`
			Name         string `json:"name"`
			NameSpelling string `json:"nameSpelling"`
		} `json:"artists"`
		LandscapImg         string   `json:"landscapImg"`
		VipStartTime        string   `json:"vipStartTime"`
		VipEndTime          string   `json:"vipEndTime"`
		VipLogo             string   `json:"vipLogo"`
		VipDownload         string   `json:"vipDownload"`
		FirstPublish        string   `json:"firstPublish"`
		ShowTag             []string `json:"showTag"`
		MaterialValidStatus bool     `json:"materialValidStatus"`
	} `json:"resource"`
}
