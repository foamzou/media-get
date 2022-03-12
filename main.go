package main

import (
	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/init_app"
	"github.com/foamzou/audio-get/logger"
	"github.com/foamzou/audio-get/processor"
)

func main() {
	opt, err := args.CheckAndParse()
	if err != nil {
		return
	}

	initApp(opt)

	proc := &processor.Processor{Opts: opt}
	err = proc.Process()
	if err != nil {
		panic(err)
	}
}

func initApp(opt *args.Options) {
	logger.SetLogLevel(opt.LogLevel)

	if !opt.MetaOnly {
		init_app.SelfCheck()
	}

}
