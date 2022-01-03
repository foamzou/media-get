package youtube

type Response struct {
	StreamingData struct {
		AdaptiveFormats []struct {
			Itag      int    `json:"itag"`
			Url       string `json:"url"`
			MimeType  string `json:"mimeType"`
			Bitrate   int    `json:"bitrate"`
			Width     int    `json:"width,omitempty"`
			Height    int    `json:"height,omitempty"`
			InitRange struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"initRange"`
			IndexRange struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"indexRange"`
			LastModified     string `json:"lastModified"`
			ContentLength    string `json:"contentLength"`
			Quality          string `json:"quality"`
			Fps              int    `json:"fps,omitempty"`
			QualityLabel     string `json:"qualityLabel,omitempty"`
			ProjectionType   string `json:"projectionType"`
			AverageBitrate   int    `json:"averageBitrate"`
			ApproxDurationMs string `json:"approxDurationMs"`
			ColorInfo        struct {
				Primaries               string `json:"primaries"`
				TransferCharacteristics string `json:"transferCharacteristics"`
				MatrixCoefficients      string `json:"matrixCoefficients"`
			} `json:"colorInfo,omitempty"`
			HighReplication bool    `json:"highReplication,omitempty"`
			AudioQuality    string  `json:"audioQuality,omitempty"`
			AudioSampleRate string  `json:"audioSampleRate,omitempty"`
			AudioChannels   int     `json:"audioChannels,omitempty"`
			LoudnessDb      float64 `json:"loudnessDb,omitempty"`
		} `json:"adaptiveFormats"`
	} `json:"streamingData"`
	VideoDetails struct {
		VideoId          string   `json:"videoId"`
		Title            string   `json:"title"`
		LengthSeconds    string   `json:"lengthSeconds"`
		Keywords         []string `json:"keywords"`
		ChannelId        string   `json:"channelId"`
		ShortDescription string   `json:"shortDescription"`
		Thumbnail        struct {
			Thumbnails []struct {
				Url    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"thumbnails"`
		} `json:"thumbnail"`
		Author string `json:"author"`
	} `json:"videoDetails"`
}
