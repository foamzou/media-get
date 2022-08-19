package douyin

import (
	"fmt"

	"github.com/dop251/goja"

	"github.com/foamzou/audio-get/logger"
	gen_xb_js "github.com/foamzou/audio-get/processor/douyin/gen-xb-js"
)

func genXB(urlParams string, ua string) (string, error) {
	vm := goja.New()

	_, err := vm.RunString(gen_xb_js.CryptoMinJS) // executes a script on the global context
	if err != nil {
		logger.Warn("genXB failed", err)
		return "", err
	}
	_, err = vm.RunString(gen_xb_js.AsciiBase64) // executes a script on the global context
	if err != nil {
		logger.Warn("genXB failed", err)
		return "", err
	}
	_, err = vm.RunString(gen_xb_js.XB_JS) // executes a script on the global context
	if err != nil {
		logger.Warn("genXB failed", err)
		return "", err
	}
	val, err := vm.RunString(fmt.Sprintf("xb_main('%s', '%s')", urlParams, ua))
	if err != nil {
		logger.Warn("genXB failed", err)
		return "", err
	}
	// nolint
	v := val.Export().(string)
	return v, nil
}
