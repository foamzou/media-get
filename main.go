package main

import (
	"fmt"

	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/processor"
)

func main() {
	opt, err := args.CheckAndParse()
	if err != nil {
		fmt.Println(err)
		return
	}

	proc := &processor.Processor{Opts: opt}
	err = proc.Process()
	if err != nil {
		panic(err)
	}
}
