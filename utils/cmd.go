package utils

import (
	"fmt"
	"os/exec"
)

func ExecCmd(program string, arg ...string) (string, error) {
	out, err := exec.Command(program, arg...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error: %w, output: %s", err, out)
	}
	return string(out), nil
}
