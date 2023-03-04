package main

import (
	"flag"
	"os"
	"runtime"
	"strings"

	"github.com/talwat/oh-my-go/internal/global"
	"github.com/talwat/oh-my-go/internal/log"
	"github.com/talwat/oh-my-go/internal/prompt"
	"github.com/talwat/oh-my-go/internal/utils"
)

func validateShell(shell string) {
	shells := []string{"sh", "fish", "zsh", "bash", "cmd", "powershell"}

	if !utils.Contains(shells, shell) {
		log.RawError("invalid shell: %s", global.Shell)
	}
}

func parseHostname(hostname string) string {
	if runtime.GOOS == "darwin" {
		return strings.TrimSuffix(hostname, ".local")
	}

	return hostname
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.RawError("you must specify a command")
	}

	switch args[0] {
	case "version":
		log.OutputLog("oh-my-go version %s\n", global.Version)
	case "prompt":
		flag.StringVar(&global.Hostname, "hostname", "unknown", "")
		flag.StringVar(&global.Shell, "shell", "unknown", "")
		flag.StringVar(&global.PWD, "pwd", "unknown", "")
		flag.StringVar(&global.User, "user", "unknown", "")
		flag.StringVar(&global.ExitCode, "exitcode", "0", "")

		err := flag.CommandLine.Parse(os.Args[2:])
		log.Error(err, "an error occurred while parsing CLI flags")

		validateShell(global.Shell)
		global.Hostname = parseHostname(global.Hostname)

		log.OutputLog(prompt.GetPrompt())
	}
}
