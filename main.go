package main

import (
	"os"

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

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.RawError("you must specify a command")
	}

	switch args[0] {
	case "version":
		log.OutputLog("oh-my-go version %s\n", global.Version)
	case "prompt":
		if len(args) < 3 {
			log.RawError("not enough parameters, need shell, user, and pwd")
		}

		global.Shell = args[1]
		validateShell(global.Shell)

		global.PWD = args[2]
		global.User = args[3]

		if len(args) > 4 {
			global.ExitCode = args[4]
		}

		log.OutputLog(prompt.GetPrompt())
	}
}
