package config

import (
	"os"
	"strings"
)

//nolint:gochecknoglobals
var (
	Plugins          = strings.Split(os.Getenv("OMGO_PLUGINS"), " ")
	ShowUserHostname = os.Getenv("OMGO_SHOW_USER_HOSTNAME") == "1"
)
