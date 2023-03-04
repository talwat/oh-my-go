package prompt

import (
	"fmt"
	"strings"

	"github.com/talwat/oh-my-go/config"
	"github.com/talwat/oh-my-go/internal/global"
	"github.com/talwat/oh-my-go/internal/log/color"

	"github.com/talwat/oh-my-go/internal/prompt/plugins"
	"github.com/talwat/oh-my-go/internal/prompt/plugins/git"
	"github.com/talwat/oh-my-go/internal/prompt/plugins/node"
	"github.com/talwat/oh-my-go/internal/prompt/plugins/python"
	"github.com/talwat/oh-my-go/internal/prompt/plugins/ruby"
	"github.com/talwat/oh-my-go/internal/prompt/plugins/rust"
)

//nolint:gochecknoglobals
var pluginList = map[string]func() plugins.PluginOutput{
	"git":    git.Plugin,
	"node":   node.Plugin,
	"rust":   rust.Plugin,
	"python": python.Plugin,
	"ruby":   ruby.Plugin,
}

func GetPrompt() string {
	segments := []string{}

	if config.ShowUserHostname {
		segment(&segments, fmt.Sprintf("(%s@%s)", global.User, global.Hostname), color.White)
	}

	// Segments
	if global.ExitCode == "0" {
		segment(&segments, "➜ ", color.Green)
	} else {
		segment(&segments, "➜ ", color.Red)
	}

	segment(&segments, pwd(), color.Cyan)

	loadPlugins(&segments)

	if git.IsDirty() {
		segment(&segments, "✗", color.Yellow)
	}

	output := fmt.Sprintf("%s ", strings.Join(segments, " "))

	return output
}
