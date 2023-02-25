package config

import (
	"os"
	"strings"
)

var (
	Plugins          = strings.Split(os.Getenv("OMGO_PLUGINS"), " ")
	ShowUserHostname = os.Getenv("OMGO_SHOW_USER_HOSTNAME") == "1"
)
