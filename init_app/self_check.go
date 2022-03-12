package init_app

import (
	"os"

	"github.com/foamzou/audio-get/ffmpeg"
	"github.com/foamzou/audio-get/logger"
)

func SelfCheck() {
	logger.Debug("self check start")

	ffmpegBin := ffmpeg.GetFfmpegBin()
	ffprobeBin := ffmpeg.GetFfprobeBin()
	if ffmpegBin == "" || ffprobeBin == "" {
		logger.Error("ffmpeg or ffprobe not found, please install them first")
		os.Exit(1)
	}

	logger.Debug("[ffmpeg]: use " + ffmpegBin)
	logger.Debug("[ffprobe]: use " + ffprobeBin)

	logger.Debug("self check finish")
}
