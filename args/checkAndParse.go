package args

import (
	"os"
	"path"
	"strings"

	"github.com/jessevdk/go-flags"
)

func CheckAndParse() (opt *Options, err error) {
	opt = &Options{}
	_, err = flags.Parse(opt)
	if err != nil {
		//fmt.Println(err)
		return
	}

	if err = checkAndAdjustOpts(opt); err != nil {
		return
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
