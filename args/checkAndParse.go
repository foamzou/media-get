package args

import (
	"os"
	"path"
	"strings"

	"github.com/jessevdk/go-flags"

	"github.com/foamzou/audio-get/consts"
	"github.com/foamzou/audio-get/version"
)

func CheckAndParse() (opt *Options, err error) {
	opt = &Options{}
	_, err = flags.Parse(opt)
	if err != nil {
		if flags.WroteHelp(err) {
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

	return nil
}

func setDefault(opt *Options) {
	if opt.MetaFormat == "" {
		opt.MetaFormat = consts.MetaFormatPlain
	}

	if opt.LogLevel == "" {
		opt.LogLevel = "info"
	}
}
