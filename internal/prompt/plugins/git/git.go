package git

import (
	"github.com/talwat/oh-my-go/internal/exec"
	"github.com/talwat/oh-my-go/internal/log/color"
	"github.com/talwat/oh-my-go/internal/prompt/plugins"
)

func Plugin() plugins.PluginOutput {
	branch, code, err := exec.Run("git", "branch", "--show-current")

	if code != 0 || err != nil || branch == "" {
		return plugins.PluginOutput{Display: false}
	}

	return plugins.PluginOutput{
		Display:    true,
		Value:      branch,
		ValueColor: color.Red,
		Name:       "git",
		NameColor:  color.Blue,
	}
}

func IsDirty() bool {
	out, code, err := exec.Run("git", "status", "--porcelain")

	return code == 0 && err == nil && out != ""
}
