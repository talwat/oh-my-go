package node

import (
	"github.com/talwat/oh-my-go/internal/exec"
	"github.com/talwat/oh-my-go/internal/log/color"
	"github.com/talwat/oh-my-go/internal/prompt/plugins"
)

func Plugin() plugins.PluginOutput {
	version, code, err := exec.Run("node", "--version")

	if code != 0 || err != nil || version == "" {
		return plugins.PluginOutput{Display: false}
	}

	return plugins.PluginOutput{
		Display:    true,
		Value:      version,
		ValueColor: color.Green,
		Name:       "node",
		NameColor:  color.Blue,
	}
}
