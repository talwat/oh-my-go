package ruby

import (
	"strings"

	"github.com/talwat/oh-my-go/internal/exec"
	"github.com/talwat/oh-my-go/internal/log/color"
	"github.com/talwat/oh-my-go/internal/prompt/plugins"
)

func Plugin() plugins.PluginOutput {
	output, code, err := exec.Run("ruby", "--version")

	if code != 0 || err != nil || output == "" {
		return plugins.PluginOutput{Display: false}
	}

	version := strings.Split(output, " ")[1]

	return plugins.PluginOutput{
		Display:    true,
		Value:      version,
		ValueColor: color.Magenta,
		Name:       "ruby",
		NameColor:  color.Blue,
	}
}
