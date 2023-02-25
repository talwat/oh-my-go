package prompt

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/talwat/oh-my-go/config"
	"github.com/talwat/oh-my-go/internal/exec"
	"github.com/talwat/oh-my-go/internal/global"
	"github.com/talwat/oh-my-go/internal/log"
	"github.com/talwat/oh-my-go/internal/log/color"

	"github.com/talwat/oh-my-go/internal/prompt/plugins"
	"github.com/talwat/oh-my-go/internal/prompt/plugins/git"
	"github.com/talwat/oh-my-go/internal/prompt/plugins/node"
	"github.com/talwat/oh-my-go/internal/prompt/plugins/python"
	"github.com/talwat/oh-my-go/internal/prompt/plugins/ruby"
	"github.com/talwat/oh-my-go/internal/prompt/plugins/rust"
)

var pluginList = map[string]func() plugins.PluginOutput{
	"git":    git.Plugin,
	"node":   node.Plugin,
	"rust":   rust.Plugin,
	"python": python.Plugin,
	"ruby":   ruby.Plugin,
}

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

func formatPWD(raw string) string {
	home, err := os.UserHomeDir()
	log.Error(err, "an error occurred while getting user home directory")

	pwd := strings.ReplaceAll(raw, home, "~")

	return pwd
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

func showUserHostname(segments *[]string) {
	hostname, exitcode, err := exec.Run("hostname", "-s")

	if err != nil || exitcode != 0 || hostname == "" {
		return
	}

	segment(segments, fmt.Sprintf("(%s@%s)", global.User, hostname), color.White)
}

func GetPrompt() string {
	segments := []string{}

	if config.ShowUserHostname {
		showUserHostname(&segments)
	}

	// Segments
	if global.ExitCode == "0" {
		segment(&segments, "➜ ", color.Green)
	} else {
		segment(&segments, "➜ ", color.Red)
	}

	segment(&segments, path.Base(formatPWD(global.PWD)), color.Cyan)

	loadPlugins(&segments)

	if git.IsDirty() {
		segment(&segments, "✗", color.Yellow)
	}

	output := fmt.Sprintf("%s ", strings.Join(segments, " "))

	return output
}
