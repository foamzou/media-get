package main

import (
	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/debugger"
	"github.com/foamzou/audio-get/logger"
	"github.com/foamzou/audio-get/processor"
	"github.com/foamzou/audio-get/utils"
)

func main() {
	defer func() {
		debugger.AppEnd()
	}()
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

func init() {
	debugger.AppStart()
}

func initApp(opt *args.Options) {
	utils.InitConfig()

	logger.SetLogLevel(opt.LogLevel)

	if !opt.MetaOnly {
		debugger.SelfCheck()
	}

}
