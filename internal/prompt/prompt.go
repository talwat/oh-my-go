package prompt

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/talwat/oh-my-go/internal/global"
	"github.com/talwat/oh-my-go/internal/log"
	"github.com/talwat/oh-my-go/internal/log/color"
	"github.com/talwat/oh-my-go/internal/prompt/plugins"

	"github.com/talwat/oh-my-go/internal/prompt/plugins/git"
	"github.com/talwat/oh-my-go/internal/prompt/plugins/node"
)

func segment(segments *[]string, segmentVal string, segmentColor string) {
	*segments = append(*segments, fmt.Sprintf(
		"%s%s%s",
		color.PromptColor(segmentColor),
		segmentVal,
		color.PromptColor(color.Reset),
	))
}

func plugin(segments *[]string, plugin func() plugins.PluginOutput) {
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

func GetPWD() string {
	pwd := ""

	if global.Shell == "cmd" {
		pwd = os.Getenv("cd")
	} else {
		pwd = os.Getenv("PWD")
	}

	home, err := os.UserHomeDir()
	log.Error(err, "an error occurred while getting user home directory")

	pwd = strings.ReplaceAll(pwd, home, "~")

	return pwd
}

func GetPrompt(exit string) string {
	segments := []string{}

	// Segments
	if exit == "0" {
		segment(&segments, "➜ ", color.Green)
	} else {
		segment(&segments, "➜ ", color.Red)
	}

	segment(&segments, path.Base(GetPWD()), color.Cyan)

	// Plugins
	plugin(&segments, git.Plugin)
	plugin(&segments, node.Plugin)

	if git.IsDirty() {
		segment(&segments, "✗", color.Yellow)
	}

	output := fmt.Sprintf("%s ", strings.Join(segments, " "))

	return output
}
