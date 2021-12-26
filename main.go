package main

import (
	"github.com/foamzou/audio-get/args"
	"github.com/foamzou/audio-get/processor"
)

func main() {
	opt, err := args.CheckAndParse()
	if err != nil {
		//fmt.Println(err)
		return
	}

	//ex, err := os.Executable()
	//if err != nil {
	//	panic(err)
	//}
	//exPath := filepath.Dir(ex)
	//fmt.Println(exPath)
	//return

	//fmt.Println(opt.Url)

	proc := &processor.Processor{Opts: opt}
	err = proc.Process()
	if err != nil {
		panic(err)
	}
}
