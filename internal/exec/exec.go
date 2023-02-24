package exec

import (
	"os/exec"
	"strings"
)

func Run(cmd string, args ...string) (string, int, error) {
	cmdObj := exec.Command(cmd, args...)

	out, err := cmdObj.Output()
	if err != nil {
		return "", -1, err
	}

	exitCode := cmdObj.ProcessState.ExitCode()

	return strings.TrimSpace(string(out)), exitCode, err
}
