package prompt

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/talwat/oh-my-go/config"
	"github.com/talwat/oh-my-go/internal/global"
	"github.com/talwat/oh-my-go/internal/log"
	"github.com/talwat/oh-my-go/internal/log/color"
	"github.com/talwat/oh-my-go/internal/prompt/plugins"
)

func segment(segments *[]string, segmentVal string, segmentColor string) {
	*segments = append(*segments, fmt.Sprintf(
		"%s%s%s",
		color.PromptColor(segmentColor),
		segmentVal,
		color.PromptColor(color.Reset),
	))
}

func loadPlugin(segments *[]string, plugin func() plugins.PluginOutput) {
	out := plugin()

	if !out.Display {
		return
	}

	*segments = append(*segments, fmt.Sprintf(
		"%s%s:(%s%s%s)%s",
		color.PromptColor(out.NameColor),
		out.Name,
		color.PromptColor(out.ValueColor),
		out.Value,
		color.PromptColor(out.NameColor),
		color.PromptColor(color.Reset),
	))
}

func pwd() string {
	home, err := os.UserHomeDir()
	log.Error(err, "an error occurred while getting user home directory")

	pwd := strings.ReplaceAll(global.PWD, home, "~")

	return path.Base(pwd)
}

func loadPlugins(segments *[]string) {
	for _, name := range config.Plugins {
		val, ok := pluginList[name]

		if !ok {
			continue
		}

		loadPlugin(segments, val)
	}
}
