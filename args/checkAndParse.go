package args

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/jessevdk/go-flags"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/debugger"
	"github.com/foamzou/audio-get/utils"
	"github.com/foamzou/audio-get/version"
)

func CheckAndParse() (opt *Options, err error) {
	opt = &Options{}
	_, err = flags.Parse(opt)
	if err != nil {
		if flags.WroteHelp(err) {
			if debugger.HasInstalledFFmpeg() {
				fmt.Println("FFmpeg,FFprobe: installed")
			} else {
				fmt.Println("FFmpeg,FFprobe: one of them is not installed")
			}

			version.DisplayVersionInfo()
			return
		}
		return
	}

	setDefault(opt)

	if err = checkAndAdjustOpts(opt); err != nil {
		panic(err)
	}

	return
}

func checkAndAdjustOpts(opt *Options) error {
	if opt.Out != "" {
		fileInfo, err := os.Stat(opt.Out)
		// Return error if the user intent pass a dir but not create
		if err != nil {
			if strings.HasSuffix(opt.Out, string(os.PathSeparator)) {
				return err
			}
		}

		// Assuming that the user intent pass a filepath, but has not been created yet
		// Return error if the parent path is not created
		if fileInfo == nil {
			_, err := os.Stat(path.Dir(opt.Out))
			if err != nil {
				return err
			}
		}

		if fileInfo != nil && fileInfo.IsDir() {
			opt.IsDir = true
		}
	}

	parseSearchSource(opt)

	return nil
}

func setDefault(opt *Options) {
	if opt.InfoFormat == "" {
		opt.InfoFormat = consts.InfoFormatPlain
	}

	if opt.LogLevel == "" {
		opt.LogLevel = "info"
	}
}

func parseSearchSource(opt *Options) {
	allSourceName := consts.GetAllSourceName()

	var sourcesWillBeSearch, includes, excludes []string
	if opt.Search.Sources == "" {
		includes = allSourceName
	} else {
		includes = strings.Split(opt.Search.Sources, ",")
	}
	if opt.Search.ExcludeSources != "" {
		excludes = strings.Split(opt.Search.ExcludeSources, ",")
	}
	includes = trimValueInSlice(includes)
	excludes = trimValueInSlice(excludes)

	for _, sourceName := range includes {
		if !utils.InArray(excludes, sourceName) && utils.InArray(allSourceName, sourceName) {
			sourcesWillBeSearch = append(sourcesWillBeSearch, sourceName)
		}
	}

	opt.Search.SourcesWillBeSearch = sourcesWillBeSearch
}

func trimValueInSlice(slice []string) []string {
	var result []string
	for _, v := range slice {
		result = append(result, strings.TrimSpace(v))
	}
	return result
}
