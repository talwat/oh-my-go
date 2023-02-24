package main

import (
	"os"

	"github.com/talwat/oh-my-go/internal/global"
	"github.com/talwat/oh-my-go/internal/log"
	"github.com/talwat/oh-my-go/internal/prompt"
	"github.com/talwat/oh-my-go/internal/utils"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.RawError("you must specify a command")
	}

	pwd := ""
	// Exit is the exit code of $?
	exit := ""

	switch args[0] {
	case "version":
		log.Log("oh-my-go version %s", global.Version)
	case "prompt":
		if len(args) < 4 {
			log.RawError("not enough parameters, need shell, exit code, and pwd")
		}

		global.Shell = args[1]
		shells := []string{"sh", "fish", "zsh", "bash", "cmd", "powershell"}

		if !utils.Contains(shells, global.Shell) {
			log.RawError("invalid shell: %s", global.Shell)
		}

		exit = args[2]
		pwd = args[3]

		log.OutputLog(prompt.GetPrompt(exit, pwd))
	}
}
